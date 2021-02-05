package entity

/*type DeliveryItemRepository interface {
	FindAll(ctx context.Context, deliveryID string) ([]DeliveryItem, error)
}

type deliveryItemRepository struct {
	conn Connection
}

func NewDeliveryItemRepository(conn Connection) DeliveryItemRepository {
	return &deliveryItemRepository{
		conn: conn,
	}
}

func (r *deliveryItemRepository) FindAll(ctx context.Context, deliveryID string) ([]DeliveryItem, error) {
	deliveryitems := []DeliveryItem{}
	err := r.conn.DB.Where("delivery_id = ?", deliveryID).Find(&deliveryitems).Error
	if err != nil {
		return []DeliveryItem{}, err
	}
	return deliveryitems, nil
}*/
