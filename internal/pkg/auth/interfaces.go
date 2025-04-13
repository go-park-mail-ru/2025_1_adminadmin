package auth

import (
	"context"
	"errors"
	"io"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
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
	ErrSamePassword       = errors.New("Новый пароль совпадает со старым")
	ErrSameName           = errors.New("Новые имя и фамилия совпадают со старыми")
	ErrSamePhone          = errors.New("Новый телефон совпадает со старым")
	ErrBasePath           = errors.New("Базовый путь для картинок не установлен")
	ErrFileCreation       = errors.New("Ошибка при создании файла")
	ErrFileSaving         = errors.New("Ошибка при сохранении файла")
	ErrFileDeletion       = errors.New("Ошибка при удалении файла")
	ErrDBError            = errors.New("Ошибка БД")
	ErrAddressNotFound    = errors.New("Ошибка поиска адреса")
)

type AuthRepo interface {
	InsertUser(ctx context.Context, user models.User) error
	SelectUserByLogin(ctx context.Context, login string) (models.User, error)
	UpdateUser(ctx context.Context, user models.User) error
	UpdateUserPic(ctx context.Context, login string, userPic string) error
	InsertAddress(ctx context.Context, address models.Address) error
	DeleteAddress(ctx context.Context, addressId uuid.UUID) error
	SelectUserAddresses(ctx context.Context, login string) ([]models.Address, error)
	AddressExists(ctx context.Context, address string, userID uuid.UUID) (bool, error)
}

type AuthUsecase interface {
	SignIn(ctx context.Context, data models.SignInReq) (models.User, string, string, error)
	SignUp(ctx context.Context, data models.SignUpReq) (models.User, string, string, error)
	Check(ctx context.Context, login string) (models.User, error)
	UpdateUser(ctx context.Context, login string, updateData models.UpdateUserReq) (models.User, error)
	UpdateUserPic(ctx context.Context, login string, picture io.ReadSeeker, extension string) (models.User, error)
	GetUserAddresses(ctx context.Context, login string) ([]models.Address, error)
	DeleteAddress(ctx context.Context, addressId uuid.UUID) error
	AddAddress(ctx context.Context, address models.Address) error
}
