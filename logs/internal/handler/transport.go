package handler

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type errorer interface {
	error() error
}

func MakeHTTPHandler(endpoints map[string]endpoint.Endpoint, logger log.Logger) http.Handler {
	router := mux.NewRouter()
	decodeEncoders := GetDecodersEncoders()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(decodeEncoders.ErrorEncoder),
	}

	router.Methods("GET").Path("/tlogs/").Handler(httptransport.NewServer(
		endpoints["FindAllEndpoint"],
		decodeEncoders.FindAllDecoder,
		decodeEncoders.Encoder,
		options...,
	))
	router.Methods("POST").Path("/tlogs/").Handler(httptransport.NewServer(
		endpoints["CreateEndpoint"],
		decodeEncoders.CreateDecoder,
		decodeEncoders.Encoder,
		options...,
	))
	return router
}
