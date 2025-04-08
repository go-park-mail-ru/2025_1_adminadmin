package models

// easyjson:json
type CartItem struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
	Weight   int     `json:"weight"`
	Amount   int     `json:"amount"`
}
