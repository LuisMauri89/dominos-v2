package entity

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/gofrs/uuid"
)

type Order struct {
	ID         string      `gorm:"primary_key" json:"id"`
	Name       string      `gorm:"size:255;not null" json:"name"`
	Email      string      `gorm:"size:100;not null;unique" json:"email"`
	Address    string      `gorm:"size:255;not null;" json:"address"`
	TotalPrice float64     `json:"total_price"`
	CreatedAt  time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Items      []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
}

func (o *Order) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	o.ID = id

	o.Name = html.EscapeString(strings.TrimSpace(o.Name))
	o.Email = html.EscapeString(strings.TrimSpace(o.Email))
	o.Address = html.EscapeString(strings.TrimSpace(o.Address))
	o.CreatedAt = time.Now()
	o.UpdatedAt = time.Now()
}

func (o *Order) Validate() error {
	if o.Name == "" {
		return errors.New("required name")
	}
	if o.Email == "" {
		return errors.New("required email")
	}
	if err := checkmail.ValidateFormat(o.Email); err != nil {
		return errors.New("invalid email")
	}
	if o.Address == "" {
		return errors.New("required address")
	}
	if o.TotalPrice <= 0.0 {
		return errors.New("total price should be other than cero")
	}
	if o.Items == nil || len(o.Items) == 0 {
		return errors.New("items can not be empty")
	}

	return nil
}

type FindAllOrderRequest struct {
}

type GetByIDOrderRequest struct {
	ID string `json:"id"`
}

type CreateOrderRequest struct {
	Order Order `json:"order"`
}

type DeleteOrderRequest struct {
	ID string `json:"id"`
}

type FindAllOrderResponse struct {
	Orders []Order `json:"orders"`
	Err    error   `json:"error,omitempty"`
}

func (r FindAllOrderResponse) error() error { return r.Err }

type GetByIDOrderResponse struct {
	Order Order `json:"order"`
	Err   error `json:"error,omitempty"`
}

func (r GetByIDOrderResponse) error() error { return r.Err }

type CreateOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateOrderResponse) error() error { return r.Err }

type DeleteOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r DeleteOrderResponse) error() error { return r.Err }
