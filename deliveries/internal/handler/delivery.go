package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"deliveries/internal/entity"

	"github.com/gorilla/mux"
)

func DecodeFindAllDeliveryRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return entity.FindAllRequest{}, nil
}

func DecodeGetByStatusDeliveryRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["status"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}
	return entity.GetByStatusRequest{id}, nil
}

func DecodeCreateDeliveryRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req entity.CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Delivery); e != nil {
		return nil, e
	}
	return req, nil
}

func EncodeDeliveryResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
