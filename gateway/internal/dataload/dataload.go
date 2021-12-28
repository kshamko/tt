package dataload

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/kshamko/tt/gateway/internal/models"
	"github.com/kshamko/tt/grpc/pkg/grpcapi"
	log "github.com/sirupsen/logrus"
)

type Dataload struct {
	src        io.Reader
	grpcClient grpcapi.GRPCServiceClient
	logger     *log.Entry
}

func New(src io.Reader, grpcClient grpcapi.GRPCServiceClient, logger *log.Entry) *Dataload {
	return &Dataload{
		src:        src,
		grpcClient: grpcClient,
		logger:     logger,
	}
}

func (d *Dataload) Start(ctx context.Context) error {
	dec := json.NewDecoder(d.src)
	_, err := dec.Token()
	if err != nil {
		return err
	}

	for dec.More() {
		t, err := dec.Token()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		var data models.Data
		err = dec.Decode(&data)
		if err != nil {
			d.logger.Error(err)
			return err
		}

		if err := d.saveData(ctx, t, &data); err != nil {
			return err
		}
	}
	return nil
}

//
func (d *Dataload) saveData(ctx context.Context, token json.Token, data *models.Data) error {

	coordinates := []*grpcapi.Coordinate{}
	if len(data.Coordinates) > 0 {
		lonInt, lonFract := splitCoordinate(data.Coordinates[0])
		latInt, latFract := splitCoordinate(data.Coordinates[1])

		coordinates = []*grpcapi.Coordinate{
			{
				Deg: lonInt,
				Min: lonFract,
			},
			{
				Deg: latInt,
				Min: latFract,
			},
		}
	}

	in := grpcapi.AddReq{
		ID: fmt.Sprintf("%s", token),
		Data: &grpcapi.Entity{
			Name:        data.Name,
			City:        data.City,
			Country:     data.Country,
			Alias:       data.Alias,
			Regions:     data.Regions,
			Coordinates: coordinates,
			Province:    data.Province,
			Timezone:    data.Timezone,
			Unlocs:      data.Unlocs,
			Code:        data.Code,
		},
	}
	d.logger.Info(token)
	_, err := d.grpcClient.AddEntity(ctx, &in)
	return err
}

func splitCoordinate(coord float64) (int32, int32) {
	ipart := int32(coord)
	frpart := fmt.Sprintf("%.7f", coord-float64(ipart))[2:]
	frint, _ := strconv.ParseInt(frpart, 10, 64)

	return ipart, int32(frint)
}
