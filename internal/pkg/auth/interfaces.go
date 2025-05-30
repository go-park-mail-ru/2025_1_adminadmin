package auth

import (
	"context"
	"errors"
	"io"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/satori/uuid"
)

var (
	ErrCreatingUser       = errors.New("ошибка в создании пользователя")
	ErrUserNotFound       = errors.New("пользователь не найден")
	ErrInvalidLogin       = errors.New("неверный формат логина")
	ErrInvalidPassword    = errors.New("неверный формат пароля")
	ErrInvalidCredentials = errors.New("неверный логин или пароль")
	ErrAlreadyExists      = errors.New("пользователь с таким логином уже существует")
	ErrGeneratingToken    = errors.New("ошибка генерации токена")
	ErrInvalidName        = errors.New("имя и фамилия должны содержать только русские буквы и быть от 2 до 25 символов")
	ErrInvalidPhone       = errors.New("некорректный номер телефона")
	ErrUUID               = errors.New("ошибка создания UUID")
	ErrSamePassword       = errors.New("новый пароль совпадает со старым")
	ErrBasePath           = errors.New("базовый путь для картинок не установлен")
	ErrFileCreation       = errors.New("ошибка при создании файла")
	ErrFileSaving         = errors.New("ошибка при сохранении файла")
	ErrFileDeletion       = errors.New("ошибка при удалении файла")
	ErrDBError            = errors.New("ошибка БД")
	ErrAddressNotFound    = errors.New("ошибка поиска адреса")
)

type AuthRepo interface {
	InsertUser(ctx context.Context, user models.User) error
	AddPromocode(ctx context.Context, userId uuid.UUID) error
	SelectUserByLogin(ctx context.Context, login string) (models.User, error)
	SetSecret2fa(ctx context.Context, secret2fa []byte, login string) error
	GetSecret2fa(ctx context.Context, login string) ([]byte, error)
	Disable2fa(ctx context.Context, login string) error
	UpdateUser(ctx context.Context, user models.User) error
	UpdateUserPic(ctx context.Context, login string, userPic string) error
	InsertAddress(ctx context.Context, address models.Address) error
	DeleteAddress(ctx context.Context, addressId uuid.UUID) error
	SelectUserAddresses(ctx context.Context, login string) ([]models.Address, error)
	AddressExists(ctx context.Context, address string, userID uuid.UUID) (bool, error)
	GetActiveAddress(ctx context.Context, userId uuid.UUID) (models.Address, error)
	ActiveAddressExists(ctx context.Context, userID uuid.UUID) (bool, error)
}

type AuthUsecase interface {
	SignIn(ctx context.Context, data models.SignInReq) (models.User, string, string, error)
	GetQRCode(ctx context.Context, secret2fa []byte, login string) error
	GetSecret2fa(ctx context.Context, login string) ([]byte, error)
	Disable2fa(ctx context.Context, login string) error
	SignUp(ctx context.Context, data models.SignUpReq) (models.User, string, string, error)
	Check(ctx context.Context, login string) (models.User, error)
	UpdateUser(ctx context.Context, login string, updateData models.UpdateUserReq) (models.User, error)
	UpdateUserPic(ctx context.Context, login string, picture io.ReadSeeker, extension string) (models.User, error)
	GetUserAddresses(ctx context.Context, login string) ([]models.Address, error)
	DeleteAddress(ctx context.Context, addressId uuid.UUID) error
	AddAddress(ctx context.Context, address models.Address) error
}
