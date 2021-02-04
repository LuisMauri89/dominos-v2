package endpoint

import (
	"context"

	"logs/internal/handler"
	"logs/internal/service"

	"github.com/go-kit/kit/endpoint"
)

func MakeEndpoints(s service.TraceLogService) map[string]endpoint.Endpoint {
	endpoints := make(map[string]endpoint.Endpoint)
	endpoints["FindAllEndpoint"] = makeFindAllEndpoint(s)
	endpoints["CreateEndpoint"] = makeCreateEndpoint(s)
	return endpoints
}

func makeFindAllEndpoint(s service.TraceLogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		tlogs, e := s.FindAll(ctx)
		return handler.FindAllResponse{Tlogs: tlogs, Err: e}, nil
	}
}

func makeCreateEndpoint(s service.TraceLogService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(handler.CreateRequest)
		e := s.Create(ctx, req.Tlog)
		return handler.CreateResponse{Err: e}, nil
	}
}
