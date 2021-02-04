package entity

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/gofrs/uuid"
)

type OrderItem struct {
	ID        string    `gorm:"primary_key" json:"id"`
	Product   string    `gorm:"size:255;not null" json:"product"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Order     Order     `json:"order"`
	OrderID   string    `gorm:"not null" json:"order_id"`
}

func (oi *OrderItem) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	oi.ID = id

	oi.Product = html.EscapeString(strings.TrimSpace(oi.Product))
	oi.CreatedAt = time.Now()
	oi.UpdatedAt = time.Now()
	oi.Order = Order{}
}

func (oi *OrderItem) Validate() error {
	if oi.Product == "" {
		return errors.New("required product")
	}
	if oi.Quantity == 0 {
		return errors.New("quantity should be other than cero")
	}

	return nil
}

type FindAllOrderItemRequest struct {
	OrderID string
}

type FindAllOrderItemResponse struct {
	OrderItems []OrderItem `json:"orderitems"`
	Err        error       `json:"error,omitempty"`
}

func (r FindAllOrderItemResponse) error() error { return r.Err }
