package handler

/* import (
	"context"
	"encoding/json"
	"net/http"

	"deliveries/internal/entity"

	"github.com/gorilla/mux"
)

func DecodeFindAllDeliveryItemRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrMissingRequiredArguments
	}

	return entity.FindAllDeliveryItemRequest{DeliveryrID: id}, nil
}

func EncodeDeliveryItemResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		EncodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
} */
