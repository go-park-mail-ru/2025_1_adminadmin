package models

import (
	"html"

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
	Id        uuid.UUID  `json:"restaurant_id"`
	Name      string     `json:"restaurant_name"`
	CartItems []CartItem `json:"products"`
}

// easyjson:json
type CartInReq struct {
	Quantity     int    `json:"quantity"`
	RestaurantId string `json:"restaurant_id"`
}

// easyjson:json
type Order struct {
    ID                uuid.UUID `json:"id"`
    UserID            uuid.UUID `json:"user_id"`
    Status            string    `json:"status"`
    AddressID         uuid.UUID `json:"address_id"`
    OrderProducts     Cart      `json:"order_products"` 

    ApartmentOrOffice string    `json:"apartment_or_office"`
    Intercom          string    `json:"intercom"`
    Entrance          string    `json:"entrance"`
    Floor             string    `json:"floor"`
    CourierComment    string    `json:"courier_comment"`
    LeaveAtDoor       bool      `json:"leave_at_door"`
}

func (a *CartItem) Sanitize() {
	a.Name = html.EscapeString(a.Name)
	a.ImageURL = html.EscapeString(a.ImageURL)
}

func (a *Cart) Sanitize() {
	a.Name = html.EscapeString(a.Name)
}

func (a *CartInReq) Sanitize() {
	a.RestaurantId = html.EscapeString(a.RestaurantId)
}

func (a *Order) Sanitize() {
	a.Status = html.EscapeString(a.Status)
	a.ApartmentOrOffice = html.EscapeString(a.ApartmentOrOffice)
	a.Intercom = html.EscapeString(a.Intercom)
	a.Entrance = html.EscapeString(a.Entrance)
	a.Floor = html.EscapeString(a.Floor)
	a.CourierComment = html.EscapeString(a.CourierComment)
}