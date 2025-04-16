package repo

import (
	"context"
	"errors"
	"log/slog"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/satori/uuid"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	insertUser          = "INSERT INTO users (id, login, first_name, last_name, phone_number, description, user_pic, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	selectUserByLogin   = "SELECT id, first_name, last_name, phone_number, description, user_pic, password_hash FROM users WHERE login = $1"
	updateUser          = "UPDATE users SET phone_number = $1, first_name = $2, last_name = $3, description = $4, password_hash = $5 WHERE id = $6;"
	updateUserPic       = "UPDATE users SET user_pic = $1 WHERE login = $2"
	selectUserAddresses = `
		SELECT a.id, a.address, a.user_id 
		FROM addresses a
		JOIN users u ON a.user_id = u.id
		WHERE u.login = $1
	`
	deleteAddress = "DELETE FROM addresses WHERE id = $1;"
	insertAddress = "INSERT INTO addresses (id, address, user_id) VALUES ($1, $2, $3)"
	addressExists = "SELECT EXISTS(SELECT 1 FROM addresses WHERE address = $1 AND user_id = $2)"
)

type AuthRepo struct {
	db pgxtype.Querier
}

func CreateAuthRepo(db pgxtype.Querier) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (repo *AuthRepo) InsertUser(ctx context.Context, user models.User) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	_, err := repo.db.Exec(ctx, insertUser, user.Id, user.Login, user.FirstName, user.LastName, user.PhoneNumber, user.Description, user.UserPic, user.PasswordHash)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	logger.Info("Successful")
	return nil
}

func (repo *AuthRepo) SelectUserByLogin(ctx context.Context, login string) (models.User, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	resultUser := models.User{Login: login}
	err := repo.db.QueryRow(ctx, selectUserByLogin, login).Scan(
		&resultUser.Id,
		&resultUser.FirstName,
		&resultUser.LastName,
		&resultUser.PhoneNumber,
		&resultUser.Description,
		&resultUser.UserPic,
		&resultUser.PasswordHash,
	)

	if err != nil {
		logger.Error(err.Error())
		return models.User{}, err
	}
	resultUser.Sanitize()

	logger.Info("Successful")
	return resultUser, nil
}

func (repo *AuthRepo) UpdateUser(ctx context.Context, user models.User) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	_, err := repo.db.Exec(ctx, updateUser, user.PhoneNumber, user.FirstName, user.LastName, user.Description, user.PasswordHash, user.Id)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("Successful")
	return nil
}

func (repo *AuthRepo) UpdateUserPic(ctx context.Context, login string, userPic string) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	_, err := repo.db.Exec(ctx, updateUserPic, userPic, login)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("Successful")
	return nil
}

func (repo *AuthRepo) SelectUserAddresses(ctx context.Context, login string) ([]models.Address, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	rows, err := repo.db.Query(ctx, selectUserAddresses, login)
	if err != nil {
		logger.Error(err.Error())
		return []models.Address{}, err
	}
	defer rows.Close()

	var addresses []models.Address
	for rows.Next() {
		var addr models.Address
		if err := rows.Scan(&addr.Id, &addr.Address, &addr.UserId); err != nil {
			logger.Error(err.Error())
			return []models.Address{}, err
		}
		addresses = append(addresses, addr)
		addr.Sanitize()
	}

	logger.Info("Successful")
	return addresses, nil
}

func (repo *AuthRepo) DeleteAddress(ctx context.Context, addressId uuid.UUID) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	result, err := repo.db.Exec(ctx, deleteAddress, addressId)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		logger.Error("Адрес не найден")
		return errors.New("Адрес не найден")
	}

	logger.Info("Successful")
	return nil
}

func (repo *AuthRepo) InsertAddress(ctx context.Context, address models.Address) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	_, err := repo.db.Exec(ctx, insertAddress, address.Id, address.Address, address.UserId)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("Successful")
	return nil
}

func (repo *AuthRepo) AddressExists(ctx context.Context, address string, userID uuid.UUID) (bool, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	var exists bool
	err := repo.db.QueryRow(ctx, addressExists, address, userID).Scan(&exists)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	return exists, nil
}