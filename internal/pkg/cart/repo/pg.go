package repo

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type CartRepository struct {
	redisClient *redis.Client
}

func NewCartRepository(redisClient *redis.Client) *CartRepository {
	return &CartRepository{redisClient: redisClient}
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
