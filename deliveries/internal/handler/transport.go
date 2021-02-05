package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"deliveries/internal/service"

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
func MakeHTTPHandler(os service.DeliveryService, endpoints map[string]endpoint.Endpoint, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(EncodeError),
	}

	router.Methods("GET").Path("/dely/").Handler(httptransport.NewServer(
		endpoints["FindAllDeliveryEndpoint"],
		DecodeFindAllDeliveryRequest,
		EncodeDeliveryResponse,
		options...,
	))
	router.Methods("GET").Path("/dely/{status}").Handler(httptransport.NewServer(
		endpoints["GetByStatusDeliveryEndpoint"],
		DecodeGetByStatusDeliveryRequest,
		EncodeDeliveryResponse,
		options...,
	))
	router.Methods("POST").Path("/dely/").Handler(httptransport.NewServer(
		endpoints["CreateDeliveryEndpoint"],
		DecodeCreateDeliveryRequest,
		EncodeDeliveryResponse,
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
