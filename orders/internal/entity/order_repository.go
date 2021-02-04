package entity

import (
	"context"
	"errors"

	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	FindAll(ctx context.Context) ([]Order, error)
	GetByID(ctx context.Context, id string) (Order, error)
	Create(ctx context.Context, order Order) error
	Delete(ctx context.Context, id string) error
}

type orderRepository struct {
	conn Connection
}

func NewOrderRepository(conn Connection) OrderRepository {
	return &orderRepository{
		conn: conn,
	}
}

func (r *orderRepository) FindAll(ctx context.Context) ([]Order, error) {
	orders := []Order{}
	err := r.conn.DB.Find(&orders).Error
	if err != nil {
		return []Order{}, err
	}
	return orders, nil
}

func (r *orderRepository) GetByID(ctx context.Context, id string) (Order, error) {
	order := Order{}
	err := r.conn.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return Order{}, errors.New("user not found")
		}
		return Order{}, err
	}

	return order, nil
}

func (r *orderRepository) Create(ctx context.Context, order Order) error {
	err := r.conn.DB.Create(&order).Error
	if err != nil {
		return err
	}
	r.conn.DB.Save(&order)
	return nil
}

func (r *orderRepository) Delete(ctx context.Context, id string) error {
	err := r.conn.DB.Delete(&Order{}, id).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return errors.New("user not found")
		}
		return err
	}
	return nil
}
