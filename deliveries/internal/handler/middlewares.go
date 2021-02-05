package handler

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

type LoggingDeliveryServiceMiddleware func(s DeliveryService) DeliveryService

type loggingDeliveryServiceMiddleware struct {
	DeliveryService
	logger log.Logger
}

func NewLoggingDeliveryMiddleware(logger log.Logger) LoggingDeliveryServiceMiddleware {
	return func(next DeliveryService) DeliveryService {
		return &loggingDeliveryServiceMiddleware{next, logger}
	}
}

func (mw *loggingDeliveryServiceMiddleware) FindAll(ctx context.Context) ([]Delivery, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "FindAll", "took", time.Since(begin))
	}(time.Now())
	return mw.DeliveryService.FindAll(ctx)
}

func (mw *loggingDeliveryServiceMiddleware) Create(ctx context.Context, td Delivery) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Create", "took", time.Since(begin))
	}(time.Now())
	return mw.DeliveryService.Create(ctx, td)
}

func (mw *loggingDeliveryServiceMiddleware) GetByStatus(ctx context.Context) ([]Delivery, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetByStatus", "took", time.Since(begin))
	}(time.Now())
	return mw.DeliveryService.GetByStatus(ctx)

}
