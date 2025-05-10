package repo

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	dbUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/db"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/redis/go-redis/v9"
)

type CartRepository struct {
	redisClient *redis.Client
}

func NewCartRepository() (*CartRepository, error) {
	redisClient, err := dbUtils.InitRedis()
	return &CartRepository{redisClient: redisClient}, err
}

func (r *CartRepository) GetCart(ctx context.Context, userID string) (map[string]int, string, float64, error) {
	logger := log.GetLoggerFromContext(ctx).With(
		slog.String("func", log.GetFuncName()),
		slog.String("user_id", userID),
	)

	key := "cart:" + userID
	items, err := r.redisClient.HGetAll(ctx, key).Result()
	if err != nil {
		logger.Error("Ошибка при HGetAll Redis", slog.String("error", err.Error()))
		return nil, "", 0, err
	}

	cart := make(map[string]int)
	var restaurantID string
	var totalSum float64

	for field, value := range items {
		switch field {
		case "restaurant_id":
			restaurantID = value
		case "total_sum":
			if sum, err := strconv.ParseFloat(value, 64); err == nil {
				totalSum = sum
			} else {
				logger.Error("Ошибка при конвертации total_sum", slog.String("value", value), slog.String("error", err.Error()))
			}
		default:
			if qty, err := strconv.Atoi(value); err == nil {
				cart[field] = qty
			} else {
				logger.Error("Ошибка при конвертации количества товара", slog.String("product_id", field), slog.String("quantity", value), slog.String("error", err.Error()))
			}
		}
	}

	logger.Info("Итоговая корзина",
		slog.Any("cart", cart),
		slog.String("restaurant_id", restaurantID),
		slog.Float64("total_sum", totalSum),
	)

	return cart, restaurantID, totalSum, nil
}

func (r *CartRepository) UpdateItemQuantity(ctx context.Context, userID, productID, restaurantID string, quantity int, price float64) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()),
		slog.String("user_id", userID),
		slog.String("product_id", productID),
		slog.String("restaurant_id", restaurantID),
		slog.Int("quantity", quantity))

	key := "cart:" + userID

	currentRestaurantID, err := r.redisClient.HGet(ctx, key, "restaurant_id").Result()
	if err != nil && err != redis.Nil {
		logger.Error("Ошибка при получении restaurant_id из Redis", slog.String("error", err.Error()))
		return err
	}

	if currentRestaurantID != "" && currentRestaurantID != restaurantID {
		if err := r.redisClient.Del(ctx, key).Err(); err != nil {
			logger.Error("Ошибка при удалении ключа Redis", slog.String("error", err.Error()))
			return err
		}
	}

	oldQtyStr, _ := r.redisClient.HGet(ctx, key, productID).Result()
	oldQty, _ := strconv.Atoi(oldQtyStr)

	if quantity <= 0 {
		err := r.redisClient.HDel(ctx, key, productID).Err()
		if err != nil {
			logger.Error("Ошибка при удалении товара из корзины", slog.String("error", err.Error()))
			return err
		}

		totalStr, _ := r.redisClient.HGet(ctx, key, "total_sum").Result()
		totalSum, _ := strconv.ParseFloat(totalStr, 64)

		totalSum -= float64(oldQty) * price

		_, err = r.redisClient.HSet(ctx, key, "total_sum", totalSum).Result()
		if err != nil {
			logger.Error("Ошибка при обновлении суммы корзины", slog.String("error", err.Error()))
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

	delta := quantity - oldQty
	deltaSum := float64(delta) * price

	totalStr, _ := r.redisClient.HGet(ctx, key, "total_sum").Result()
	totalSum, _ := strconv.ParseFloat(totalStr, 64)

	newTotal := totalSum + deltaSum

	pipe := r.redisClient.TxPipeline()
	pipe.HSet(ctx, key, productID, quantity)
	pipe.HSet(ctx, key, "restaurant_id", restaurantID)
	pipe.HSet(ctx, key, "total_sum", newTotal) 

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

	err := r.redisClient.Del(ctx, key).Err()
	if err != nil {
		logger.Error("Ошибка при удалении корзины из Redis", slog.String("error", err.Error()))
		return err
	}

	logger.Info("Корзина успешно очищена")
	return nil
}
