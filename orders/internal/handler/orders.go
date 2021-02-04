package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"orders/internal/entity"

	"github.com/gorilla/mux"
)

func DecodeFindAllOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return entity.FindAllOrderRequest{}, nil
}

func DecodeGetByIDOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}
	return entity.GetByIDOrderRequest{ID: id}, nil
}

func DecodeCreateOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req entity.CreateOrderRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Order); e != nil {
		return nil, e
	}
	return req, nil
}

func DecodeDeleteOrderRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}
	return entity.DeleteOrderRequest{ID: id}, nil
}

func EncodeOrderResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
