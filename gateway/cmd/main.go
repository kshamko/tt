package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-openapi/loads"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/InVisionApp/go-health"
	flags "github.com/jessevdk/go-flags"
	"github.com/kshamko/tt/gateway/internal/dataload"
	"github.com/kshamko/tt/gateway/internal/debug"
	"github.com/kshamko/tt/gateway/internal/handler"
	"github.com/kshamko/tt/gateway/internal/restapi"
	"github.com/kshamko/tt/gateway/internal/restapi/operations"
	"github.com/kshamko/tt/grpc/pkg/grpcapi"
	"golang.org/x/sync/errgroup"
)

const APP_ID = "gateway"

func main() { //nolint: funlen
	var opts = struct {
		HTTPListenHost string `long:"http.listen.host" env:"HTTP_LISTEN_HOST" default:"" description:"http server interface host"`
		HTTPListenPort int    `long:"http.listen.port" env:"HTTP_LISTEN_PORT" default:"8081" description:"http server interface port"`
		DebugListen    string `long:"debug.listen" env:"DEBUG_LISTEN" default:":2111" description:"Interface for serve debug information(metrics/health/pprof)"`
		Verbose        bool   `long:"v" env:"VERBOSE" description:"Enable Verbose log output"`
		GRPCEnpoint    string `long:"service.grpc" env:"SERVICE_GRPC" default:"localhost:6060"`
	}{}

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	logger := logrus.WithField("app_id", APP_ID)
	logger.Logger.SetOutput(os.Stdout)
	logger.Logger.SetLevel(logrus.InfoLevel)

	if opts.Verbose {
		logger.Logger.SetLevel(logrus.DebugLevel)
	}

	logger.Infof("Launching Application with: %+v", opts)

	grpcClient, err := initGRPCClient(opts.GRPCEnpoint, grpc.WithInsecure())
	if err != nil {
		logger.Fatal(err)
	}

	gr, appctx := errgroup.WithContext(context.Background())

	// init healthchecks & metrics
	gr.Go(func() error {
		healthd := health.New()
		d := debug.New(healthd)

		return d.Serve(appctx, opts.DebugListen)
	})

	// init http server with API served
	gr.Go(func() error {
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			logger.Error(err)
			return err
		}
		api := operations.NewDataGWSwaggerAPI(swaggerSpec)
		api.DataDataHandler = handler.NewData(
			grpcClient,
		)

		server := restapi.NewServer(api)
		// nolint: errcheck
		defer server.Shutdown()

		server.Host = opts.HTTPListenHost
		server.Port = opts.HTTPListenPort

		go func() {
			<-appctx.Done()
			server.Shutdown()
		}()

		return server.Serve()
	})

	//load data from file (would move that part to grpc service)
	gr.Go(func() error {
		f, err := os.Open("assets/data/ports.json")
		if err != nil {
			return nil
		}
		defer f.Close()

		dl := dataload.New(f, grpcClient, logger)
		return dl.Start(appctx)
	})

	errCanceled := errors.New("Canceled")
	// listen to os signal to stop app gracefully/ change logs level in runtime
	gr.Go(func() error {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		cusr := make(chan os.Signal, 1)
		signal.Notify(cusr, syscall.SIGUSR1)
		for {
			select {
			case <-appctx.Done():
				return appctx.Err()
			case <-sigs:
				logger.Info("Caught stop signal. Exiting ...")
				return errCanceled
			case <-cusr:
				if logger.Level == logrus.DebugLevel {
					logger.Logger.SetLevel(logrus.InfoLevel)
					logger.Info("Caught SIGUSR1 signal. Log level changed to INFO")
					continue
				}
				logger.Info("Caught SIGUSR1 signal. Log level changed to DEBUG")
				logger.Logger.SetLevel(logrus.DebugLevel)
			}
		}
	})

	err = gr.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

//
func initGRPCClient(target string, opts ...grpc.DialOption) (grpcapi.GRPCServiceClient, error) {

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, err
	}
	return grpcapi.NewGRPCServiceClient(conn), nil
}
