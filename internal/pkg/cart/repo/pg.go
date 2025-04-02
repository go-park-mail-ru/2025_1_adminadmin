package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgtype/pgxtype"
	"github.com/redis/go-redis/v9"
)

const (
	insertUser = "INSERT INTO users (id, login, first_name, last_name, phone_number, description, user_pic, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
)

type CartRepository struct {
	redisClient *redis.Client
	db          pgxtype.Querier
}

func NewCartRepository(redisClient *redis.Client, db pgxtype.Querier) *CartRepository {
	return &CartRepository{redisClient: redisClient, db: db}
}

func (r *CartRepository) GetCart(ctx context.Context, userID string) (map[string]int, error) {

	key := "cart:" + userID
	items, err := r.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	cart := make(map[string]int)
	for productID, quantity := range items {
		var qty int
		fmt.Sscanf(quantity, "%d", &qty)
		cart[productID] = qty
	}

	return cart, nil
}

func (r *CartRepository) AddItem(ctx context.Context, userID, productID string) error {
	key := "cart:" + userID

	return r.redisClient.HIncrBy(ctx, key, productID, 1).Err()
}

func (r *CartRepository) RemoveItem(ctx context.Context, userID, productID string) error {
	key := "cart:" + userID
	quantity, err := r.redisClient.HGet(ctx, key, productID).Int()
	if err != nil {
		return err
	}

	if quantity > 1 {
		return r.redisClient.HIncrBy(ctx, key, productID, -1).Err()
	}
	return r.redisClient.HDel(ctx, key, productID).Err()
}
