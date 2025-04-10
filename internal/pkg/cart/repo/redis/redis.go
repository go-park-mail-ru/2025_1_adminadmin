package repo

import (
	"context"
	"fmt"
	"log"

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
	log.Printf("[GetCart] Получение корзины для ключа: %s", key)

	items, err := r.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		log.Printf("[GetCart] Ошибка при HGetAll Redis: %v", err)
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
		if _, err := fmt.Sscanf(quantity, "%d", &qty); err != nil {
			log.Printf("[GetCart] Ошибка при конвертации количества товара (productID: %s, value: %s): %v", productID, quantity, err)
			continue
		}

		cart[productID] = qty
	}

	log.Printf("[GetCart] Итоговая корзина: %+v, restaurantID: %s", cart, restaurantID)

	return cart, restaurantID, nil
}

func (r *CartRepository) UpdateItemQuantity(ctx context.Context, userID, productID, restaurantID string, quantity int) error {
	key := "cart:" + userID
	log.Printf("[UpdateItemQuantity] Обновление товара %s для пользователя %s, ресторан: %s, количество: %d", productID, userID, restaurantID, quantity)

	currentRestaurantID, err := r.redisClient.HGet(ctx, key, "restaurant_id").Result()
	if err != nil && err != redis.Nil {
		log.Printf("[UpdateItemQuantity] Ошибка при получении restaurant_id из Redis: %v", err)
		return err
	}

	if currentRestaurantID != "" && currentRestaurantID != restaurantID {
		log.Printf("[UpdateItemQuantity] Рестораны не совпадают (текущий: %s, новый: %s), очищаем корзину", currentRestaurantID, restaurantID)
		if err := r.redisClient.Del(ctx, key).Err(); err != nil {
			log.Printf("[UpdateItemQuantity] Ошибка при удалении ключа Redis: %v", err)
			return err
		}
	}

	if quantity <= 0 {
		log.Printf("[UpdateItemQuantity] Количество <= 0, удаляем товар %s", productID)
		err := r.redisClient.HDel(ctx, key, productID).Err()
		if err != nil {
			log.Printf("[UpdateItemQuantity] Ошибка при удалении товара из корзины: %v", err)
			return err
		}

		fields, err := r.redisClient.HKeys(ctx, key).Result()
		if err == nil {
			onlyRestaurantID := len(fields) == 1 && fields[0] == "restaurant_id"
			if onlyRestaurantID || len(fields) == 0 {
				log.Printf("[UpdateItemQuantity] Корзина пуста, удаляем restaurant_id")
				_ = r.redisClient.HDel(ctx, key, "restaurant_id").Err()
			}
		} else {
			log.Printf("[UpdateItemQuantity] Ошибка при получении ключей из Redis: %v", err)
		}

		return nil
	}

	if quantity > 999 {
		log.Printf("[UpdateItemQuantity] Превышен лимит количества товара (%d)", quantity)
		return fmt.Errorf("товар уже в корзине")
	}

	pipe := r.redisClient.TxPipeline()
	pipe.HSet(ctx, key, productID, quantity)
	pipe.HSet(ctx, key, "restaurant_id", restaurantID)

	_, err = pipe.Exec(ctx)
	if err != nil {
		log.Printf("[UpdateItemQuantity] Ошибка при выполнении транзакции Redis: %v", err)
	} else {
		log.Printf("[UpdateItemQuantity] Успешно обновлено: %s -> %d", productID, quantity)
	}
	return err
}
