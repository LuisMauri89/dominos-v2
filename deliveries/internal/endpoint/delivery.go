package endpoint

import (
	"context"

	"deliveries/internal/entity"
	"deliveries/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeDeliveriesEndpoints(s service.DeliveryService, endpoints map[string]endpoint.Endpoint) map[string]endpoint.Endpoint {
	endpoints["FindAllDeliveryEndpoint"] = makeFindAllEndpoint(s)
	endpoints["GetByStatusDeliveryEndpoint"] = makeGetByStatusEndpoint(s)
	endpoints["CreateDeliveryEndpoint"] = makeCreateEndpoint(s)

	return endpoints
}
func makeFindAllEndpoint(s service.DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		deliveries, e := s.FindAll(ctx)
		return entity.FindAllResponse{TDely: deliveries, Err: e}, nil
	}
}

func makeCreateEndpoint(s service.DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.CreateRequest)
		e := s.Create(ctx, req.Delivery)
		return entity.CreateResponse{Err: e}, nil
	}
}

func makeGetByStatusEndpoint(s service.DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.GetByStatusRequest)
		delivery, e := s.GetByStatus(ctx, req.Status)
		return entity.GetByStatusResponse{Delivery: delivery, Err: e}, nil
	}
}
