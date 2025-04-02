package models

// easyjson:json
type CartItem struct {
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    ImageURL string  `json:"image_url"`
    Weight   float64 `json:"weight"`
    Amount   int     `json:"amount"`
}
