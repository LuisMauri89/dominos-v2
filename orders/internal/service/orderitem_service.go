package service

import (
	"context"

	"orders/internal/entity"

	"github.com/go-kit/kit/log"
)

type OrderItemService interface {
	FindAll(ctx context.Context, orderID string) ([]entity.OrderItem, error)
}

type orderItemService struct {
	repository entity.OrderItemRepository
	logger     log.Logger
}

func NewOrderItemService(repository entity.OrderItemRepository, logger log.Logger) OrderItemService {
	return &orderItemService{
		repository: repository,
		logger:     logger,
	}
}

func (s *orderItemService) FindAll(ctx context.Context, orderID string) ([]entity.OrderItem, error) {
	orderitems, err := s.repository.FindAll(ctx, orderID)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return orderitems, nil
}
