package service

import (
	"context"

	"orders/internal/entity"

	"github.com/go-kit/kit/log"
)

type OrderService interface {
	FindAll(ctx context.Context) ([]entity.Order, error)
	GetByID(ctx context.Context, id string) (entity.Order, error)
	Create(ctx context.Context, order entity.Order) error
	Delete(ctx context.Context, id string) error
}

type orderService struct {
	repository   entity.OrderRepository
	logger       log.Logger
	kafkaService KafkaService
}

func NewOrderService(repository entity.OrderRepository, logger log.Logger, kafkaService KafkaService) OrderService {
	return &orderService{
		repository:   repository,
		logger:       logger,
		kafkaService: kafkaService,
	}
}

func (s *orderService) FindAll(ctx context.Context) ([]entity.Order, error) {
	orders, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return orders, nil
}

func (s *orderService) GetByID(ctx context.Context, id string) (entity.Order, error) {
	order, err := s.repository.GetByID(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return order, err
	}
	s.logger.Log("getbyid:", "success")
	return order, nil
}

func (s *orderService) Create(ctx context.Context, order entity.Order) error {
	order.Prepare()
	err := order.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	items := []entity.OrderItem{}
	for _, orderitem := range order.Items {
		orderitem.Prepare()
		err = orderitem.Validate()
		if err != nil {
			s.logger.Log("error:", err)
			return err
		}
		items = append(items, orderitem)
	}
	order.Items = items

	if err := s.repository.Create(ctx, order); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	payload := Payload{
		ID:         order.ID,
		Name:       order.Name,
		Address:    order.Address,
		TotalPrice: order.TotalPrice,
		Action:     "NEW",
	}
	s.kafkaService.ProduceOrderAction(payload)
	return nil
}

func (s *orderService) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("delete:", "success")
	return nil
}
