package handler

import (
	"context"
	"time"

	"logs/internal/entity"
	"logs/internal/service"

	"github.com/go-kit/kit/log"
)

type LoggingTraceLogServiceMiddleware func(s service.TraceLogService) service.TraceLogService

type loggingTraceLogServiceMiddleware struct {
	service.TraceLogService
	logger log.Logger
}

func NewLoggingTraceLogServiceMiddleware(logger log.Logger) LoggingTraceLogServiceMiddleware {
	return func(next service.TraceLogService) service.TraceLogService {
		return &loggingTraceLogServiceMiddleware{next, logger}
	}
}

func (mw *loggingTraceLogServiceMiddleware) FindAll(ctx context.Context) ([]entity.TraceLog, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "FindAllTlogs", "took", time.Since(begin))
	}(time.Now())
	return mw.TraceLogService.FindAll(ctx)
}

func (mw *loggingTraceLogServiceMiddleware) Create(ctx context.Context, tlog entity.TraceLog) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "CreateTlog", "took", time.Since(begin))
	}(time.Now())
	return mw.TraceLogService.Create(ctx, tlog)
}
