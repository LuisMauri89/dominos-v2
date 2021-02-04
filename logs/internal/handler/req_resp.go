package handler

import (
	"logs/internal/entity"
)

type FindAllRequest struct {
}

type CreateRequest struct {
	Tlog entity.TraceLog `json:"tlog"`
}

type FindAllResponse struct {
	Tlogs []entity.TraceLog `json:"tlogs"`
	Err   error             `json:"error,omitempty"`
}

func (r FindAllResponse) error() error { return r.Err }

type CreateResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateResponse) error() error { return r.Err }
