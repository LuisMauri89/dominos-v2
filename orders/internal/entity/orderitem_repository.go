package entity

import (
	"context"
)

type OrderItemRepository interface {
	FindAll(ctx context.Context, orderID string) ([]OrderItem, error)
}

type orderItemRepository struct {
	conn Connection
}

func NewOrderItemRepository(conn Connection) OrderItemRepository {
	return &orderItemRepository{
		conn: conn,
	}
}

func (r *orderItemRepository) FindAll(ctx context.Context, orderID string) ([]OrderItem, error) {
	orderitems := []OrderItem{}
	err := r.conn.DB.Where("order_id = ?", orderID).Find(&orderitems).Error
	if err != nil {
		return []OrderItem{}, err
	}
	return orderitems, nil
}
