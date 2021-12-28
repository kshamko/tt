package handler

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/kshamko/tt/gateway/internal/models"
	"github.com/kshamko/tt/gateway/internal/restapi/operations/data"
	"github.com/kshamko/tt/grpc/pkg/grpcapi"
)

// Data struct to define Handle func on.
type Data struct {
	apiClient grpcapi.GRPCServiceClient
}

func NewData(apiClient grpcapi.GRPCServiceClient) *Data {
	return &Data{
		apiClient: apiClient,
	}
}

// Handle function processes http request, needed by swagger generated code.
//
func (dt *Data) Handle(in data.DataParams) middleware.Responder {

	ctx := context.Background()
	result, err := dt.apiClient.GetEntity(ctx, &grpcapi.GetReq{ID: in.ID})

	if err != nil {
		return data.NewDataNotFound().WithPayload(
			&models.APIInvalidResponse{
				Code:    404,
				Message: err.Error(),
			},
		)
	}

	coordinates := []float64{}
	for _, coord := range result.Coordinates {
		coordinates = append(coordinates, float64(coord.Deg)+float64(coord.Min)/float64(10000000))
	}

	return data.NewDataOK().WithPayload(
		&models.Data{
			Alias:       result.Alias,
			City:        result.City,
			Code:        result.Code,
			Coordinates: coordinates,
			Country:     result.Country,
			Name:        result.Name,
			Province:    result.Province,
			Regions:     result.Regions,
			Timezone:    result.Timezone,
			Unlocs:      result.Unlocs,
		},
	)
}
