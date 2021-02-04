package endpoint

import (
	"context"
	"orders/internal/entity"
	"orders/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeOrderItemEndpoints(ois service.OrderItemService, endpoints map[string]endpoint.Endpoint) map[string]endpoint.Endpoint {
	endpoints["FindAllOrderItemEndpoint"] = makeFindAllOrderItemEndpoint(ois)
	return endpoints
}

func makeFindAllOrderItemEndpoint(ois service.OrderItemService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(entity.FindAllOrderItemRequest)
		orderitems, e := ois.FindAll(ctx, req.OrderID)
		return entity.FindAllOrderItemResponse{OrderItems: orderitems, Err: e}, nil
	}
}
