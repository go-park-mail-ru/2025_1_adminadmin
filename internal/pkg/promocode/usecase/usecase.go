package usecase

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	interfaces "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/promocode"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/satori/uuid"
)

type PromocodeUsecase struct {
	repo interfaces.PromocodeRepo
}

func NewPromocodeUsecase(r interfaces.PromocodeRepo) *PromocodeUsecase {
	return &PromocodeUsecase{repo: r}
}

func (u *PromocodeUsecase) GetPromocodes(ctx context.Context, user_id uuid.UUID, count int, offset int) ([]models.Promocode, error) {
	return u.repo.GetPromocodes(ctx, user_id, count, offset)
}

func (u *PromocodeUsecase) CheckPromocode(ctx context.Context, user_id uuid.UUID, promocode string) (float64, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	discount, err := u.repo.GetDiscount(ctx, user_id, promocode)
	if err != nil {
		logger.Error(err.Error())
		return 0.0, err
	}

	logger.Info("success")
	
	return discount, nil
}
