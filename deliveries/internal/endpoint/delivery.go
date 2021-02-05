package endpoint

import (
	"context"
	"deliveries/internal/service"

	"github.com/KingDiegoA/dominos-v2/deliveries/internal/entity"
	"github.com/go-kit/kit/endpoint"
)

func MakeOrderEndpoints(s service.OrderService, endpoints map[string]endpoint.Endpoint) map[string]endpoint.Endpoint {
	endpoints["FindAllEndpoint"] = makeFindAllEndpoint(s)
	endpoints["GetByStatusEndpoint"] = makeGetByStatusEndpoint(s)
	endpoints["CreateOrderEndpoint"] = makeCreateOrderEndpoint(s)

	return endpoints
}
func makeFindAllEndpoint(s service.DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		deliveries, e := s.FindAll(ctx)
		return entity.FindAllResponse{TDely: deliveries, Err: e}, nil
	}
}

func makeCreateOrderEndpoint(s service.DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.CreateOrderRequest)
		e := s.Create(ctx, req.Delivery)
		return entity.CreateOrderResponse{Err: e}, nil
	}
}

func makeGetByStatusEndpoint(s service.DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.GetByStatusRequest)
		e := s.GetByStatus(ctx, req.Status)
		return entity.GetByStatusResponse{Err: e}, nil
	}
}
