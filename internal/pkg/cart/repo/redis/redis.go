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

func (r *CartRepository) GetCart(ctx context.Context, userID string) (map[string]int, string, error) {
	key := "cart:" + userID
	items, err := r.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, "", err
	}

	cart := make(map[string]int)
	var restaurantID string

	for productID, quantity := range items {
		if productID == "restaurant_id" {
			restaurantID = quantity
			continue
		}

		var qty int
		if _, err := fmt.Sscanf(quantity, "%d", &qty); err == nil {
			cart[productID] = qty
		}
	}

	return cart, restaurantID, nil
}



func (r *CartRepository) AddItem(ctx context.Context, userID, productID string) error {
    key := "cart:" + userID
    quantity, err := r.redisClient.HGet(ctx, key, productID).Int()
    if err == nil && quantity > 0 {
        return fmt.Errorf("товар уже в корзине")
    }
    return r.redisClient.HSet(ctx, key, productID, 1).Err()
}

func (r *CartRepository) UpdateItemQuantity(ctx context.Context, userID, productID, restaurantID string, quantity int) error {
	key := "cart:" + userID

	currentRestaurantID, err := r.redisClient.HGet(ctx, key, "restaurant_id").Result()
	if err != nil && err != redis.Nil {
		return err
	}

	if currentRestaurantID != "" && currentRestaurantID != restaurantID {
		if err := r.redisClient.Del(ctx, key).Err(); err != nil {
			return err
		}
	}

	if quantity <= 0 {
		err := r.redisClient.HDel(ctx, key, productID).Err()
		if err != nil {
			return err
		}

		fields, err := r.redisClient.HKeys(ctx, key).Result()
		if err == nil {
			onlyRestaurantID := len(fields) == 1 && fields[0] == "restaurant_id"
			if onlyRestaurantID || len(fields) == 0 {
				_ = r.redisClient.HDel(ctx, key, "restaurant_id").Err()
			}
		}

		return nil
	}

	if quantity > 999 {
		return fmt.Errorf("товар уже в корзине")
	}

	pipe := r.redisClient.TxPipeline()
	pipe.HSet(ctx, key, productID, quantity)
	pipe.HSet(ctx, key, "restaurant_id", restaurantID)

	_, err = pipe.Exec(ctx)
	return err
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
