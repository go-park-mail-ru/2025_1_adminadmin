package interfaces

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

type PromocodeRepo interface {
	GetPromocodes(ctx context.Context, user_id uuid.UUID, count, offset int) ([]models.Promocode, error)
	GetDiscount(ctx context.Context, user_id uuid.UUID, promocode string) (float64, error)
}

type PromocodeUsecase interface {
	GetPromocodes(ctx context.Context, user_id uuid.UUID, count int, offset int) ([]models.Promocode, error)
	CheckPromocode(ctx context.Context, user_id uuid.UUID, promocode string) (float64, error)
}
