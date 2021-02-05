package entity

type Delivery struct {
	ID          string  `gorm:"primary_key" json:"id"`
	OrderID     string  `gorm:"size:255;not null" json:"orderid"`
	Status      string  `gorm:"size:255;not null" json:"status"`
	Name        string  `gorm:"size:255;not null" json:"name"`
	FinalPrice  float64 `json:"final_price"`
	Address     string  `gorm:"size:255;not null" json:"address"`
	Description string  `gorm:"size:255" json:"address"`
}
type FindAllRequest struct {
}

type FindAllResponse struct {
	TDely []Delivery `json:"tdely"`
	Err   error      `json:"error,omitempty"`
}

func (r FindAllResponse) error() error { return r.Err }

type CreateOrderRequest struct {
	Delivery Delivery `json:"delivery"`
}

type CreateOrderResponse struct {
	Err error `json:"error,omitempty"`
}

func (r CreateOrderResponse) error() error { return r.Err }

type GetByStatusRequest struct {
	Status string `json:"status"`
}

type GetByStatusResponse struct {
	Delivery Delivery `json:"delivery"`
	Err      error    `json:"error,omitempty"`
}

func (r GetByStatusResponse) error() error { return r.Err }
