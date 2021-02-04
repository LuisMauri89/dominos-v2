package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"orders/internal/entity"

	"github.com/gorilla/mux"
)

func DecodeFindAllOrderItemRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}

	return entity.FindAllOrderItemRequest{OrderID: id}, nil
}

func EncodeOrderItemResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
