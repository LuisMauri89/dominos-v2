package service

import (
	"context"

	"deliveries/internal/entity"

	"github.com/go-kit/kit/log"
)

type DeliveryService interface {
	FindAll(ctx context.Context) ([]entity.Delivery, error)
	GetByStatus(ctx context.Context, status string) (entity.Delivery, error)
	Create(ctx context.Context, delivery entity.Delivery) error
}

type deliveryService struct {
	repository entity.DeliveryRepository
	logger     log.Logger
	tlogger    LogsService
}

func NewDeliveryService(repository entity.DeliveryRepository, logger log.Logger, tlogger LogsService) DeliveryService {
	return &deliveryService{
		repository: repository,
		logger:     logger,
		tlogger:    tlogger,
	}
}

func (s *deliveryService) FindAll(ctx context.Context) ([]entity.Delivery, error) {
	deliveries, err := s.repository.FindAll(ctx)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return deliveries, nil
}

func (s *deliveryService) Create(ctx context.Context, delivery entity.Delivery) error {
	delivery.Prepare()
	err := delivery.Validate()
	if err != nil {
		s.logger.Log("error:", err)
		return err
	}

	/* items := []entity.DeliveryItem{}
	for _, deliveryitem := range delivery.Items {
		deliveryitem.Prepare()
		err = deliveryitem.Validate()
		if err != nil {
			s.logger.Log("error:", err)
			return err
		}
		items = append(items, deliveryitem)
	}
	delivery.Items = items */

	if err := s.repository.Create(ctx, delivery); err != nil {
		s.logger.Log("error:", err)
		return err
	}
	s.logger.Log("create:", "success")

	go s.tlogger.SaveLog(TLog{
		ServiceName: "DELIVERIES",
		Caller:      "Delivery->Create",
		Event:       "POST",
		Extra:       "Create new delivery.",
	})
	return nil
}

func (s *deliveryService) GetByStatus(ctx context.Context, status string) (entity.Delivery, error) {
	return entity.Delivery{}, nil
}
