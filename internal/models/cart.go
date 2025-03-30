package models

// easyjson:json
type Cart struct {
	UserID int        `json:"user_id"`
	Items  []CartItem `json:"items"`
}
