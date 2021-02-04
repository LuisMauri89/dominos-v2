package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"orders/internal/service"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var (
	ErrMissingRequiredArguments = errors.New("missing required argument")
)

type errorer interface {
	error() error
}

// MakeHTTPHandler -
func MakeHTTPHandler(os service.OrderService, ois service.OrderItemService, endpoints map[string]endpoint.Endpoint, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(EncodeError),
	}

	router.Methods("GET").Path("/orders/").Handler(httptransport.NewServer(
		endpoints["FindAllOrderEndpoint"],
		DecodeFindAllOrderRequest,
		EncodeOrderResponse,
		options...,
	))
	router.Methods("GET").Path("/orders/{id}").Handler(httptransport.NewServer(
		endpoints["GetByIDOrderEndpoint"],
		DecodeGetByIDOrderRequest,
		EncodeOrderResponse,
		options...,
	))
	router.Methods("POST").Path("/orders/").Handler(httptransport.NewServer(
		endpoints["CreateOrderEndpoint"],
		DecodeCreateOrderRequest,
		EncodeOrderResponse,
		options...,
	))
	router.Methods("DELETE").Path("/orders/{id}").Handler(httptransport.NewServer(
		endpoints["DeleteOrderEndpoint"],
		DecodeDeleteOrderRequest,
		EncodeOrderResponse,
		options...,
	))
	router.Methods("GET").Path("/order/{id}/orderitems/").Handler(httptransport.NewServer(
		endpoints["FindAllOrderItemEndpoint"],
		DecodeFindAllOrderItemRequest,
		EncodeOrderItemResponse,
		options...,
	))
	return router
}

func EncodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("nil error - can not encode nil error.")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	default:
		return http.StatusInternalServerError
	}
}
