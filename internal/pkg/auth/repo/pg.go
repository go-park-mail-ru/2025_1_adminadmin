package repo

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	insertUser        = "INSERT INTO users (id, login, first_name, last_name, phone_number, description, user_pic, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	selectUserByLogin = "SELECT id, first_name, last_name, phone_number, description, user_pic, password_hash FROM users WHERE login = $1"
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
	_, err := repo.db.Exec(ctx, insertUser, user.Id, user.FirstName, user.LastName, user.PhoneNumber, user.Description, user.UserPic, user.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

func (repo *AuthRepo) SelectUserByLogin(ctx context.Context, login string) (models.User, error) {
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
		return models.User{}, err
	}
	resultUser.Login = login

	return resultUser, nil
}
