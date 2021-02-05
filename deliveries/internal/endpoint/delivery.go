package endpoint

import (
	"context"
	"deliveries/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeOrderEndpoints(s service.OrderService, endpoints map[string]endpoint.Endpoint) map[string]endpoint.Endpoint {
	endpoints["FindAllEndpoint"] = makeFindAllEndpoint(s)
	endpoints["GetByStatusEndpoint"] = makeGetByStatusEndpoint(s)
	endpoints["CreateOrderEndpoint"] = makeCreateOrderEndpoint(s)

	return endpoints
}
func makeFindAllEndpoint(s DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		deliveries, e := s.FindAll(ctx)
		return FindAllResponse{TDely: deliveries, Err: e}, nil
	}
}

func makeCreateOrderEndpoint(s DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateOrderRequest)
		e := s.Create(ctx, req.Delivery)
		return CreateOrderResponse{Err: e}, nil
	}
}

func makeGetByStatusEndpoint(s DeliveryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetByStatusRequest)
		deliveries, e := os.GetByStatus(ctx, req.Status)
		return GetByStatusResponse{Order: order, Err: e}, nil
	}
}
