package handler

import (
	"context"
	"encoding/json"
	"time"

	"orders/internal/entity"
	"orders/internal/service"

	"github.com/go-kit/kit/log"
)

type LoggingOrderServiceMiddleware func(s service.OrderService) service.OrderService

type loggingOrderServiceMiddleware struct {
	service.OrderService
	tlogger service.LogsService
	logger  log.Logger
}

func NewLoggingOrderServiceMiddleware(logger log.Logger, tlogger service.LogsService) LoggingOrderServiceMiddleware {
	return func(next service.OrderService) service.OrderService {
		return &loggingOrderServiceMiddleware{next, tlogger, logger}
	}
}

func (mw *loggingOrderServiceMiddleware) FindAll(ctx context.Context) ([]entity.Order, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "FindAllOrders", "took", time.Since(begin))
	}(time.Now())

	collection, err := mw.OrderService.FindAll(ctx)
	tlog := service.TLog{
		ServiceName: "ORDERS",
		Caller:      "Order->FindAll",
		Event:       "GET",
		Extra:       "Find all orders. " + time.Now().String(),
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return collection, err
}

func (mw *loggingOrderServiceMiddleware) GetByID(ctx context.Context, id string) (entity.Order, error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetByID", "took", time.Since(begin))
	}(time.Now())

	order, err := mw.OrderService.GetByID(ctx, id)
	tlog := service.TLog{
		ServiceName: "ORDERS",
		Caller:      "Order->GetByID",
		Event:       "GET",
		Extra:       "Get order by ID. " + id,
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return order, err
}

func (mw *loggingOrderServiceMiddleware) Create(ctx context.Context, order entity.Order) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Create", "took", time.Since(begin))
	}(time.Now())

	err := mw.OrderService.Create(ctx, order)
	tlog := service.TLog{
		ServiceName: "ORDERS",
		Caller:      "Order->Create",
		Event:       "POST",
		Extra:       "Create new order. " + order.Name,
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return err
}

func (mw *loggingOrderServiceMiddleware) Delete(ctx context.Context, id string) error {
	defer func(begin time.Time) {
		mw.logger.Log("method", "Delete", "took", time.Since(begin))
		mw.tlogger.SaveLog(service.TLog{
			ServiceName: "ORDERS",
			Caller:      "Order->Delete",
			Event:       "DELETE",
			Extra:       "Delete order by ID.",
		})
	}(time.Now())

	err := mw.OrderService.Delete(ctx, id)
	tlog := service.TLog{
		ServiceName: "ORDERS",
		Caller:      "Order->Delete",
		Event:       "DELETE",
		Extra:       "Delete order by ID. " + id,
	}

	if err != nil {
		extra, _ := json.Marshal(err)
		tlog.Extra = string(extra)
	}
	go mw.tlogger.SaveLog(tlog)

	return err
}
