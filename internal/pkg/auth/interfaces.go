package auth

import (
	"context"
	"errors"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
)

var (
	ErrCreatingUser       = errors.New("Ошибка в создании пользователя")
	ErrUserNotFound       = errors.New("Пользователь не найден")
	ErrInvalidLogin       = errors.New("Неверный формат логина")
	ErrInvalidPassword    = errors.New("Неверный формат пароля")
	ErrInvalidCredentials = errors.New("Неверный логин или пароль")
	ErrGeneratingToken    = errors.New("Ошибка генерации токена")
	ErrInvalidName        = errors.New("Имя и фамилия должны содержать только русские буквы и быть от 2 до 25 символов")
	ErrInvalidPhone       = errors.New("Некорректный номер телефона")
	ErrUUID               = errors.New("Ошибка создания UUID")
)

type AuthRepo interface {
	InsertUser(ctx context.Context, user models.User) error
	SelectUserByLogin(ctx context.Context, login string) (models.User, error)
}

type AuthUsecase interface {
	SignIn(ctx context.Context, data models.SignInReq) (models.User, string, string, error)
	SignUp(ctx context.Context, data models.SignUpReq) (models.User, string, string, error)
	Check(ctx context.Context, login string) (models.User, error)
}
