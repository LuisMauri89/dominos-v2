package service

/*
type DeliveryItemService interface {
	FindAll(ctx context.Context, deliveryID string) ([]entity.DeliveryItem, error)
}

type deliveryItemService struct {
	repository entity.DeliveryItemRepository
	logger     log.Logger
}

func NewDeliveryItemService(repository entity.DeliveryItemRepository, logger log.Logger) DeliveryItemService {
	return &deliveryItemService{
		repository: repository,
		logger:     logger,
	}
}

func (s *deliveryItemService) FindAll(ctx context.Context, deliveryID string) ([]entity.DeliveryItem, error) {
	deliveryitem, err := s.repository.FindAll(ctx, deliveryID)
	if err != nil {
		s.logger.Log("error:", err)
		return nil, err
	}
	s.logger.Log("findall:", "success")
	return deliveryitems, nil
}
*/
