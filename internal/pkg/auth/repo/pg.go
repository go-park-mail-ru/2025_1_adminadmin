package repo

import (
	"context"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	insertUser        = "INSERT INTO users (id, login, first_name, last_name, phone_number, description, user_pic, password_hash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
	selectUserByLogin = "SELECT id, first_name, last_name, phone_number, description, user_pic, password_hash FROM users WHERE login = $1"
	updateUser        = "UPDATE users SET phone_number = $1, first_name = $2, last_name = $3, description = $4, password_hash = $5 WHERE id = $6;"
	updateUserPic     = "UPDATE users set user_pic = $1 WHERE login = $2"
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

func (repo *AuthRepo) UpdateUser(ctx context.Context, user models.User) error {

	_, err := repo.db.Exec(ctx, updateUser, user.PhoneNumber, user.FirstName, user.LastName, user.Description, user.PasswordHash, user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *AuthRepo) UpdateUserPic(ctx context.Context, login string, userPic string) error {
	_, err := repo.db.Exec(ctx, updateUserPic, userPic, login)
	if err != nil {
		return err
	}

	return nil
}
