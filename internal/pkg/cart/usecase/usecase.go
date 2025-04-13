package usecase

import (
	"context"
	"log"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	pgRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/pg"
	redisRepo "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/repo/redis"
	"github.com/satori/uuid"
)

type CartUsecase struct {
	cartRepo       *redisRepo.CartRepository
	restaurantRepo *pgRepo.RestaurantRepository
}

func NewCartUsecase(cartRepo *redisRepo.CartRepository, restaurantRepo *pgRepo.RestaurantRepository) *CartUsecase {
	return &CartUsecase{
		cartRepo:       cartRepo,
		restaurantRepo: restaurantRepo,
	}
}

func (uc *CartUsecase) GetCart(ctx context.Context, userID string) (models.Cart, error) {
	cartRaw, restaurantID, err := uc.cartRepo.GetCart(ctx, userID)
	log.Printf("[GetCart] usecase1 %v", err)
	if err != nil {
		return models.Cart{}, err
	}

	if restaurantID == "" {
		return models.Cart{}, nil
	}

	productIDs := make([]string, 0, len(cartRaw))
	for id := range cartRaw {
		productIDs = append(productIDs, id)
	}

	items, err := uc.restaurantRepo.GetCartItem(ctx, productIDs, cartRaw, restaurantID)
	log.Printf("[GetCart] usecase2 %v", err)
	if err != nil {
		return models.Cart{}, err
	}

	return items, nil
}

func (uc *CartUsecase) UpdateItemQuantity(ctx context.Context, userID, productID string, restaurantId string, quantity int) error {
	return uc.cartRepo.UpdateItemQuantity(ctx, userID, productID, restaurantId, quantity)
}

func (uc *CartUsecase) ClearCart(ctx context.Context, userID string) error {
	return uc.cartRepo.ClearCart(ctx, userID)
}

func (u *CartUsecase) CreateOrder(ctx context.Context, userID string, req models.OrderInReq, cart models.Cart) (*models.Order, error) {
	order := &models.Order{
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
	}

	if err := u.restaurantRepo.Save(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}
