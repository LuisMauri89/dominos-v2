package entity

import (
	"errors"
	"html"
	"strings"

	"github.com/gofrs/uuid"
)

type Delivery struct {
	ID          string  `gorm:"primary_key" json:"id"`
	OrderID     string  `gorm:"size:255;not null" json:"orderid"`
	Status      string  `gorm:"size:255;not null" json:"status"`
	Name        string  `gorm:"size:255;not null" json:"name"`
	FinalPrice  float64 `json:"final_price"`
	Address     string  `gorm:"size:255;not null" json:"address"`
	Description string  `gorm:"size:255" json:"address"`
}

func (d *Delivery) Prepare() {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	d.ID = id

	d.Name = html.EscapeString(strings.TrimSpace(d.Name))
	d.Address = html.EscapeString(strings.TrimSpace(d.Address))
}

func (d *Delivery) Validate() error {
	if d.Name == "" {
		return errors.New("required name")
	}
	if d.Address == "" {
		return errors.New("required address")
	}
	if d.FinalPrice <= 0.0 {
		return errors.New("total price should be other than cero")
	}

	return nil
}

type FindAllRequest struct {
}

type FindAllResponse struct {
	TDely []Delivery `json:"tdely"`
	Err   error      `json:"error,omitempty"`
}

func (r FindAllResponse) error() error { return r.Err }

type CreateRequest struct {
	Delivery Delivery `json:"delivery"`
}

type CreateResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateResponse) error() error { return r.Err }

type GetByStatusRequest struct {
	Status string `json:"status"`
}

type GetByStatusResponse struct {
	Delivery Delivery `json:"delivery"`
	Err      error    `json:"error,omitempty"`
}

func (r GetByStatusResponse) error() error { return r.Err }
