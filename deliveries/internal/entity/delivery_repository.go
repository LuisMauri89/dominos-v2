package entity

import (
	"context"
	"errors"
)

type DeliveryRepository interface {
	FindAll(ctx context.Context) ([]Delivery, error)
	GetByStatus(ctx context.Context, status string) (Delivery, error)
	Create(ctx context.Context, td Delivery) error
}
type deliveryRepository struct {
	conn Connection
}

func NewDeliveryRepository(conn Connection) DeliveryRepository {
	return &deliveryRepository{
		conn: conn,
	}
}
func (r *deliveryRepository) FindAll(ctx context.Context) ([]Delivery, error) {
	deliveries := []Delivery{}
	err := r.conn.DB.Find(&delivery).Error
	if err != nil {
		return []Delivery{}, err
	}
	return deliveries, nil
}

func (r *deliveryRepository) Create(ctx context.Context, delivery Delivery) error {
	err := r.conn.DB.Create(&delivery).Error
	if err != nil {
		return err
	}
	r.conn.DB.Save(&delivery)
	return nil
}
func (r *deliveryRepository) GetByStatus(ctx context.Context, status string) (Delivery, error) {
	deliveries := Delivery{}
	err := r.conn.DB.Where("status = ?", status).First(&delivery).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return Delivery{}, errors.New("user not found")
		}
		return Delivery{}, err
	}

	return delivery, nil
}
