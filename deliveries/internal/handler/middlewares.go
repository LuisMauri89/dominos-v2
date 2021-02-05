package handler

import (
	"context"
	"deliveries/internal/service"
	"time"

	"deliveries/internal/entity"

	"github.com/go-kit/kit/log"
)

type LoggingDeliveryServiceMiddleware func(s service.DeliveryService) service.DeliveryService

type loggingDeliveryServiceMiddleware struct {
	service.DeliveryService
	logger log.Logger
}

func NewLoggingDeliveryMiddleware(logger log.Logger) LoggingDeliveryServiceMiddleware {
	return func(next service.DeliveryService) service.DeliveryService {
		return &loggingDeliveryServiceMiddleware{next, logger}
	}
}

func (mw *loggingDeliveryServiceMiddleware) FindAll(ctx context.Context) ([]entity.Delivery, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "FindAll", "took", time.Since(begin))
	}(time.Now())
	return mw.DeliveryService.FindAll(ctx)
}

func (mw *loggingDeliveryServiceMiddleware) Create(ctx context.Context, td entity.Delivery) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Create", "took", time.Since(begin))
	}(time.Now())
	return mw.DeliveryService.Create(ctx, td)
}

func (mw *loggingDeliveryServiceMiddleware) GetByStatus(ctx context.Context, status string) (entity.Delivery, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetByStatus", "took", time.Since(begin))
	}(time.Now())
	return mw.DeliveryService.GetByStatus(ctx, status)

}
