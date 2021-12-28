package server

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/kshamko/tt/grpc/internal/datasource"
	api "github.com/kshamko/tt/grpc/pkg/grpcapi"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// nolint
var (
	// ErrNotFound returns if no info was found
	ErrNotFound = status.Error(codes.NotFound, "not found")
	// ErrBadRequest returns if malformed input was provided
	ErrBadRequest = status.Error(codes.InvalidArgument, "invalid input")
)

//Datasource for invites database
type Datasource interface {
	GetEntry(ctx context.Context, id string) (datasource.Data, error)
	AddEntry(ctx context.Context, id string, data datasource.Data) error
}

//Server for grpc server
type Server struct {
	db     Datasource
	logger grpclog.LoggerV2
}

//Serve creates and starts gRPC server
func Serve(ctx context.Context, listen string, l grpclog.LoggerV2, db Datasource) error {
	s := &Server{
		db:     db,
		logger: l,
	}

	lis, err := net.Listen("tcp", listen)
	if err != nil {
		return errors.Wrap(err, "Coudn't start API server")
	}

	grpc_prometheus.EnableHandlingTimeHistogram(func(h *prometheus.HistogramOpts) {})

	mw := []grpc.UnaryServerInterceptor{
		grpc_prometheus.UnaryServerInterceptor,
	}
	srv := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(mw...)))
	if l != nil {
		grpclog.SetLoggerV2(l)
	}
	api.RegisterGRPCServiceServer(srv, s)

	e := make(chan error, 1)
	go func() {
		e <- srv.Serve(lis)
	}()
	select {
	case <-ctx.Done():
		srv.GracefulStop()
		return nil
	case err := <-e:
		return err
	}
}

//GetUser return user by condition passed with user var
func (s *Server) GetEntity(ctx context.Context, in *api.GetReq) (*api.Entity, error) {
	if in == nil {
		return nil, ErrBadRequest
	}

	data, err := s.db.GetEntry(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	return &api.Entity{
		ID:      in.ID,
		Name:    data.Name,
		City:    data.City,
		Country: data.Country,
		Alias:   data.Alias,
		Regions: data.Regions,
		Coordinates: []*api.Coordinate{
			{
				Deg: data.Coordinates[0].Deg,
				Min: data.Coordinates[0].Min,
			},
			{
				Deg: data.Coordinates[1].Deg,
				Min: data.Coordinates[1].Min,
			},
		},
		Province: data.Province,
		Timezone: data.Timezone,
		Unlocs:   data.Unlocs,
		Code:     data.Code,
	}, nil
}

//AddUser insert user into storage
func (s *Server) AddEntity(ctx context.Context, in *api.AddReq) (*emptypb.Empty, error) {

	coordinates := []datasource.Coordinate{}
	if len(in.Data.Coordinates) != 0 {
		coordinates = []datasource.Coordinate{
			{
				Deg: in.Data.Coordinates[0].Deg,
				Min: in.Data.Coordinates[0].Min,
			},
			{
				Deg: in.Data.Coordinates[1].Deg,
				Min: in.Data.Coordinates[1].Min,
			},
		}
	}

	err := s.db.AddEntry(
		ctx,
		in.ID,
		datasource.Data{
			Name:        in.Data.Name,
			City:        in.Data.City,
			Country:     in.Data.Country,
			Alias:       in.Data.Alias,
			Regions:     in.Data.Regions,
			Coordinates: coordinates,
			Province:    in.Data.Province,
			Timezone:    in.Data.Timezone,
			Unlocs:      in.Data.Unlocs,
			Code:        in.Data.Code,
		},
	)
	return &emptypb.Empty{}, err
}
