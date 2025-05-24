package models

import (
	"html"
	"time"

	"github.com/satori/uuid"
)

// easyjson:json
type CartItem struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	ImageURL string    `json:"image_url"`
	Weight   int       `json:"weight"`
	Amount   int       `json:"amount"`
}

// easyjson:json
type Cart struct {
	Id               uuid.UUID  `json:"restaurant_id"`
	Name             string     `json:"restaurant_name"`
	CartItems        []CartItem `json:"products"`
	TotalSum         float64    `json:"total_sum"`
	RecommendedItems []CartItem `json:"recommended_products"`
}

// easyjson:json
type CartInReq struct {
	Quantity     int    `json:"quantity"`
	RestaurantId string `json:"restaurant_id"`
}

// easyjson:json
type Order struct {
	ID            uuid.UUID `json:"id"`
	UserID        string    `json:"user"`
	Status        string    `json:"status"`
	Address       string    `json:"address"`
	OrderProducts Cart      `json:"order_products"`

	ApartmentOrOffice string    `json:"apartment_or_office"`
	Intercom          string    `json:"intercom"`
	Entrance          string    `json:"entrance"`
	Floor             string    `json:"floor"`
	CourierComment    string    `json:"courier_comment"`
	LeaveAtDoor       bool      `json:"leave_at_door"`
	CreatedAt         time.Time `json:"created_at"`
	FinalPrice        float64   `json:"final_price"`
}

// easyjson:json
type OrderInReq struct {
	Status  string `json:"status"`
	Address string `json:"address"`

	ApartmentOrOffice string  `json:"apartment_or_office"`
	Intercom          string  `json:"intercom"`
	Entrance          string  `json:"entrance"`
	Floor             string  `json:"floor"`
	CourierComment    string  `json:"courier_comment"`
	LeaveAtDoor       bool    `json:"leave_at_door"`
	FinalPrice        float64 `json:"final_price"`
	Promocode         string  `json:"promocode"`
}

// easyjson:json
type OrderResp struct {
	Orders []Order `json:"orders"`
	Total  int     `json:"total"`
}

// easyjson:json
type Promocode struct {
	Id        uuid.UUID `json:"id"`
	Promocode string    `json:"promocode"`
	Discount  float64   `json:"discount"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type PromocodeResp struct {
	Discount float64 `json:"discount"`
}

func (p *Promocode) Sanitize() {
	p.Promocode = html.EscapeString(p.Promocode)
}

func (c *CartItem) Sanitize() {
	c.Name = html.EscapeString(c.Name)
	c.ImageURL = html.EscapeString(c.ImageURL)
}

func (c *Cart) Sanitize() {
	c.Name = html.EscapeString(c.Name)
	for i := range c.CartItems {
		c.CartItems[i].Sanitize()
	}
}

func (c *CartInReq) Sanitize() {
	c.RestaurantId = html.EscapeString(c.RestaurantId)
}

func (o *Order) Sanitize() {
	o.Status = html.EscapeString(o.Status)
	o.Address = html.EscapeString(o.Address)
	o.ApartmentOrOffice = html.EscapeString(o.ApartmentOrOffice)
	o.Intercom = html.EscapeString(o.Intercom)
	o.Entrance = html.EscapeString(o.Entrance)
	o.Floor = html.EscapeString(o.Floor)
	o.CourierComment = html.EscapeString(o.CourierComment)
	o.OrderProducts.Sanitize()
}

func (o *OrderInReq) Sanitize() {
	o.Status = html.EscapeString(o.Status)
	o.Address = html.EscapeString(o.Address)
	o.ApartmentOrOffice = html.EscapeString(o.ApartmentOrOffice)
	o.Intercom = html.EscapeString(o.Intercom)
	o.Entrance = html.EscapeString(o.Entrance)
	o.Floor = html.EscapeString(o.Floor)
	o.CourierComment = html.EscapeString(o.CourierComment)
}
