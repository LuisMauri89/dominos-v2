package endpoint

import (
	"context"
	"orders/internal/entity"
	"orders/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeOrderEndpoints(os service.OrderService, endpoints map[string]endpoint.Endpoint) map[string]endpoint.Endpoint {
	endpoints["FindAllOrderEndpoint"] = makeFindAllOrderEndpoint(os)
	endpoints["GetByIDOrderEndpoint"] = makeGetByIDOrderEndpoint(os)
	endpoints["CreateOrderEndpoint"] = makeCreateOrderEndpoint(os)
	endpoints["DeleteOrderEndpoint"] = makeDeleteOrderEndpoint(os)
	return endpoints
}

func makeFindAllOrderEndpoint(os service.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		orders, e := os.FindAll(ctx)
		return entity.FindAllOrderResponse{Orders: orders, Err: e}, nil
	}
}

func makeGetByIDOrderEndpoint(os service.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.GetByIDOrderRequest)
		order, e := os.GetByID(ctx, req.ID)
		return entity.GetByIDOrderResponse{Order: order, Err: e}, nil
	}
}

func makeCreateOrderEndpoint(os service.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.CreateOrderRequest)
		e := os.Create(ctx, req.Order)
		return entity.CreateOrderResponse{Err: e}, nil
	}
}

func makeDeleteOrderEndpoint(os service.OrderService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.DeleteOrderRequest)
		e := os.Delete(ctx, req.ID)
		return entity.DeleteOrderResponse{Err: e}, nil
	}
}
