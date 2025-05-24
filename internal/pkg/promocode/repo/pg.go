package repo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	dbUtils "github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/db"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/satori/uuid"
)

type PromocodeRepository struct {
	db pgxtype.Querier
}

func NewPromocodeRepository() (*PromocodeRepository, error) {
	db, err := dbUtils.InitDB()
	return &PromocodeRepository{db: db}, err
}

const (
	getAllPromocodes = "SELECT id, promocode, discount, created_at, expires_at FROM promocodes WHERE user_id = $1 AND is_used = FALSE ORDER BY id ASC LIMIT $2 OFFSET $3;"
	getDiscount      = "SELECT discount FROM promocodes WHERE user_id = $1 AND promocode = $2 AND is_used = FALSE"
	deletePromocode = "UPDATE promocodes SET is_used = TRUE WHERE user_id = $1 AND promocode = $2"
)

func (r *PromocodeRepository) GetPromocodes(ctx context.Context, user_id uuid.UUID, count, offset int) ([]models.Promocode, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	rows, err := r.db.Query(ctx, getAllPromocodes, user_id, count, offset)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	var promocodes []models.Promocode
	for rows.Next() {
		var promocode models.Promocode
		if err := rows.Scan(&promocode.Id, &promocode.Promocode, &promocode.Discount, &promocode.CreatedAt, &promocode.ExpiresAt); err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		promocode.UserId = user_id
		promocode.Sanitize()
		promocodes = append(promocodes, promocode)
	}

	logger.Info("Successful")
	return promocodes, rows.Err()
}

func (r *PromocodeRepository) GetDiscount(ctx context.Context, user_id uuid.UUID, promocode string) (float64, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	var discount float64
	err := r.db.QueryRow(ctx, getDiscount, user_id, promocode).Scan(&discount)
	if err != nil {
		logger.Error("ошибка при получении скидки", slog.String("error", err.Error()))
		return 0.0, fmt.Errorf("не удалось получить скидку: %w", err)
	}

	logger.Info("Successful")
	return discount, nil
}
