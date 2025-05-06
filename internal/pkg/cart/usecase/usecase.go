package usecase

import (
	"context"
	"time"

	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/satori/uuid"
)

type CartUsecase struct {
	cartRepo       cart.CartRepo
	restaurantRepo cart.RestaurantRepo
}

func NewCartUsecase(cartRepo cart.CartRepo, restaurantRepo cart.RestaurantRepo) *CartUsecase {
	return &CartUsecase{
		cartRepo:       cartRepo,
		restaurantRepo: restaurantRepo,
	}
}

func (uc *CartUsecase) GetCart(ctx context.Context, login string) (models.Cart, error, bool) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	cartRaw, restaurantID, err := uc.cartRepo.GetCart(ctx, login)
	if err != nil {
		logger.Error("ошибка получения корзины", slog.String("error", err.Error()))
		return models.Cart{}, err, false
	}

	if restaurantID == "" || cartRaw == nil {
		logger.Info("корзина пуста или нет restaurantID")
		return models.Cart{}, nil, false
	}

	productIDs := make([]string, 0, len(cartRaw))
	for id := range cartRaw {
		productIDs = append(productIDs, id)
	}

	items, err := uc.restaurantRepo.GetCartItem(ctx, productIDs, cartRaw, restaurantID)
	if err != nil {
		logger.Error("ошибка получения данных по товарам", slog.String("restaurantID", restaurantID), slog.String("error", err.Error()))
		return models.Cart{}, err, false
	}

	logger.Info("успешное получение корзины")
	return items, nil, true
}

func (uc *CartUsecase) UpdateItemQuantity(ctx context.Context, login, productID string, restaurantId string, quantity int) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))
	err := uc.cartRepo.UpdateItemQuantity(ctx, login, productID, restaurantId, quantity)
	if err != nil {
		logger.Error("не удалось обновить количество", slog.String("error", err.Error()))
	} else {
		logger.Info("успешно обновлено количество", slog.String("productID", productID), slog.Int("quantity", quantity))
	}
	return err
}

func (uc *CartUsecase) ClearCart(ctx context.Context, login string) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))
	err := uc.cartRepo.ClearCart(ctx, login)
	if err != nil {
		logger.Error("ошибка при очистке корзины", slog.String("error", err.Error()))
	} else {
		logger.Info("корзина успешно очищена")
	}
	return err
}

func (u *CartUsecase) CreateOrder(ctx context.Context, userID string, req models.OrderInReq, cart models.Cart) (models.Order, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	order := models.Order{
		ID:                uuid.NewV4(),
		UserID:            userID,
		Status:            req.Status,
		Address:           req.Address,
		OrderProducts:     cart,
		ApartmentOrOffice: req.ApartmentOrOffice,
		Intercom:          req.Intercom,
		Entrance:          req.Entrance,
		Floor:             req.Floor,
		CourierComment:    req.CourierComment,
		LeaveAtDoor:       req.LeaveAtDoor,
		CreatedAt:         time.Now(),
		FinalPrice:        req.FinalPrice,
	}

	order.Sanitize()

	if err := u.restaurantRepo.Save(ctx, order, userID); err != nil {
		logger.Error("не удалось сохранить заказ", slog.String("error", err.Error()))
		return models.Order{}, err
	}

	logger.Info("заказ успешно создан", slog.String("orderID", order.ID.String()))
	return order, nil
}

func (u *CartUsecase) GetOrders(ctx context.Context, user_id uuid.UUID, count, offset int) ([]models.Order, error) {
	return u.restaurantRepo.GetOrders(ctx, user_id, count, offset)
}

func (u *CartUsecase) GetOrderById(ctx context.Context, order_id, user_id uuid.UUID) (models.Order, error) {
	return u.restaurantRepo.GetOrderById(ctx, order_id, user_id)
}

func (u *CartUsecase) UpdateOrderStatus(ctx context.Context, order_id uuid.UUID) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	status := "paid"
	err := u.restaurantRepo.UpdateOrderStatus(ctx, order_id, status)
	if err != nil{
		logger.Error(err.Error())
		return err
	}

	go u.scheduleDeliveryStatusUpdate(order_id)

	logger.Info("success")
	return nil
}

func (u *CartUsecase) scheduleDeliveryStatusUpdate(orderID uuid.UUID) {
    ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
    defer cancel()
	
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

    time.Sleep(60 * time.Second)

    err := u.restaurantRepo.UpdateOrderStatus(ctx, orderID, "in_delivery")
    if err != nil {
        logger.Error("Failed to update delivery status",
            slog.String("error", err.Error()))
        return
    }

    logger.Info("Order status updated to 'in_delivery'")

	time.Sleep(20 * time.Second)
    err = u.restaurantRepo.UpdateOrderStatus(ctx, orderID, "delivered")
    if err != nil {
        logger.Error("Failed to update to 'delivered'",
            slog.String("error", err.Error()))
        return
    }
    logger.Info("Order status updated to 'delivered'")
}
