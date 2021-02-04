package entity

import (
	"context"
)

type TraceLogRepository interface {
	FindAll(ctx context.Context) ([]TraceLog, error)
	Create(ctx context.Context, tlog TraceLog) error
}
