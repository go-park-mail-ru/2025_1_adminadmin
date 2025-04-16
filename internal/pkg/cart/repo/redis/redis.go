package repo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/redis/go-redis/v9"
)

type CartRepository struct {
	redisClient *redis.Client
}

func NewCartRepository(redisClient *redis.Client) *CartRepository {
	return &CartRepository{redisClient: redisClient}
}

func (r *CartRepository) GetCart(ctx context.Context, userID string) (map[string]int, string, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()), slog.String("user_id", userID))

	key := "cart:" + userID
	logger.Info("Получение корзины", slog.String("cart_key", key))

	items, err := r.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		logger.Error("Ошибка при HGetAll Redis", slog.String("error", err.Error()))
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
			logger.Error("Ошибка при конвертации количества товара", slog.String("product_id", productID), slog.String("quantity", quantity), slog.String("error", err.Error()))
			continue
		}

		cart[productID] = qty
	}

	logger.Info("Итоговая корзина", slog.Any("cart", cart), slog.String("restaurant_id", restaurantID))

	return cart, restaurantID, nil
}

func (r *CartRepository) UpdateItemQuantity(ctx context.Context, userID, productID, restaurantID string, quantity int) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()), slog.String("user_id", userID), slog.String("product_id", productID), slog.String("restaurant_id", restaurantID), slog.Int("quantity", quantity))

	key := "cart:" + userID
	logger.Info("Обновление количества товара")

	currentRestaurantID, err := r.redisClient.HGet(ctx, key, "restaurant_id").Result()
	if err != nil && err != redis.Nil {
		logger.Error("Ошибка при получении restaurant_id из Redis", slog.String("error", err.Error()))
		return err
	}

	if currentRestaurantID != "" && currentRestaurantID != restaurantID {
		logger.Info("Рестораны не совпадают, очищаем корзину", slog.String("current_restaurant_id", currentRestaurantID), slog.String("new_restaurant_id", restaurantID))
		if err := r.redisClient.Del(ctx, key).Err(); err != nil {
			logger.Error("Ошибка при удалении ключа Redis", slog.String("error", err.Error()))
			return err
		}
	}

	if quantity <= 0 {
		logger.Info("Количество товара <= 0, удаляем товар", slog.String("product_id", productID))
		err := r.redisClient.HDel(ctx, key, productID).Err()
		if err != nil {
			logger.Error("Ошибка при удалении товара из корзины", slog.String("error", err.Error()))
			return err
		}

		fields, err := r.redisClient.HKeys(ctx, key).Result()
		if err == nil {
			onlyRestaurantID := len(fields) == 1 && fields[0] == "restaurant_id"
			if onlyRestaurantID || len(fields) == 0 {
				logger.Info("Корзина пуста, удаляем restaurant_id")
				_ = r.redisClient.HDel(ctx, key, "restaurant_id").Err()
			}
		} else {
			logger.Error("Ошибка при получении ключей из Redis", slog.String("error", err.Error()))
		}

		return nil
	}

	if quantity > 999 {
		logger.Warn("Превышен лимит количества товара", slog.Int("quantity", quantity))
		return fmt.Errorf("товар уже в корзине")
	}

	pipe := r.redisClient.TxPipeline()
	pipe.HSet(ctx, key, productID, quantity)
	pipe.HSet(ctx, key, "restaurant_id", restaurantID)

	_, err = pipe.Exec(ctx)
	if err != nil {
		logger.Error("Ошибка при выполнении транзакции Redis", slog.String("error", err.Error()))
	} else {
		logger.Info("Успешно обновлено", slog.String("product_id", productID), slog.Int("quantity", quantity))
	}
	return err
}

func (r *CartRepository) ClearCart(ctx context.Context, userID string) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()), slog.String("user_id", userID))

	key := "cart:" + userID
	logger.Info("Очистка корзины")

	err := r.redisClient.Del(ctx, key).Err()
	if err != nil {
		logger.Error("Ошибка при удалении корзины из Redis", slog.String("error", err.Error()))
		return err
	}

	logger.Info("Корзина успешно очищена")
	return nil
}
