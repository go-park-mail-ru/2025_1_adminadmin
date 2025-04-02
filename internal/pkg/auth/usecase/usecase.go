package usecase

import (
	"bytes"
	"context"
	"crypto/rand"
	"io"
	"log/slog"
	"os"
	"path"
	"strings"
	"time"
	"unicode"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/log"
	"github.com/golang-jwt/jwt"
	"github.com/satori/uuid"
	"golang.org/x/crypto/argon2"
)

func HashPassword(salt []byte, plainPassword string) []byte {
	hashedPass := argon2.IDKey([]byte(plainPassword), salt, 1, 64*1024, 4, 32)
	return append(salt, hashedPass...)
}

func checkPassword(passHash []byte, plainPassword string) bool {
	salt := make([]byte, 8)
	copy(salt, passHash[:8])
	userPassHash := HashPassword(salt, plainPassword)
	return bytes.Equal(userPassHash, passHash)
}

const (
	minNameLength  = 2
	maxNameLength  = 25
	minPhoneLength = 7
	maxPhoneLength = 15
	maxLoginLength = 20
	minLoginLength = 3
	minPassLength  = 8
	maxPassLength  = 25
)

const allowedRunes = "абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"
const allowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-"

func isValidName(name string) bool {
	if len(name) < minNameLength || len(name) > maxNameLength {
		return false
	}
	for _, r := range name {
		if !strings.ContainsRune(allowedRunes, r) {
			return false
		}
	}
	return true
}

func isValidPhone(phone string) bool {
	if len(phone) < minPhoneLength || len(phone) > maxPhoneLength {
		return false
	}
	for _, r := range phone {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func validLogin(login string) bool {
	if len(login) < minLoginLength || len(login) > maxLoginLength {
		return false
	}
	for _, char := range login {
		if !strings.Contains(allowedChars, string(char)) {
			return false
		}
	}
	return true
}

func validPassword(password string) bool {
	var up, low, digit, special bool

	if len(password) < minPassLength || len(password) > maxPassLength {
		return false
	}

	for _, char := range password {

		switch {
		case unicode.IsUpper(char):
			up = true
		case unicode.IsLower(char):
			low = true
		case unicode.IsDigit(char):
			digit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			special = true
		default:
			return false
		}
	}

	return up && low && digit && special
}

func generateToken(login string, id uuid.UUID) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "", nil
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"id": id,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString([]byte(secret))
}

type AuthUsecase struct {
	repo auth.AuthRepo
}

func CreateAuthUsecase(repo auth.AuthRepo) *AuthUsecase {
	return &AuthUsecase{repo: repo}
}

func (uc *AuthUsecase) SignIn(ctx context.Context, data models.SignInReq) (models.User, string, string, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	if !validLogin(data.Login) {
		logger.Error(auth.ErrInvalidLogin.Error())
		return models.User{}, "", "", auth.ErrInvalidLogin
	}

	user, err := uc.repo.SelectUserByLogin(ctx, data.Login)
	if err != nil {
		logger.Error(auth.ErrUserNotFound.Error())
		return models.User{}, "", "", auth.ErrUserNotFound
	}

	if !checkPassword(user.PasswordHash, data.Password) {
		logger.Error(auth.ErrInvalidCredentials.Error())
		return models.User{}, "", "", auth.ErrInvalidCredentials
	}

	token, err := generateToken(user.Login, user.Id)
	if err != nil {
		logger.Error(auth.ErrGeneratingToken.Error())
		return models.User{}, "", "", auth.ErrGeneratingToken
	}

	csrfToken := uuid.NewV4().String()

	logger.Info("Successful")
	return user, token, csrfToken, nil
}

func (uc *AuthUsecase) SignUp(ctx context.Context, data models.SignUpReq) (models.User, string, string, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	if !validLogin(data.Login) {
		logger.Error(auth.ErrInvalidLogin.Error())
		return models.User{}, "", "", auth.ErrInvalidLogin
	}

	if !validPassword(data.Password) {
		logger.Error(auth.ErrInvalidPassword.Error())
		return models.User{}, "", "", auth.ErrInvalidPassword
	}

	if !isValidName(data.FirstName) || !isValidName(data.LastName) {
		logger.Error(auth.ErrInvalidName.Error())
		return models.User{}, "", "", auth.ErrInvalidName
	}

	if !isValidPhone(data.PhoneNumber) {
		logger.Error(auth.ErrInvalidPhone.Error())
		return models.User{}, "", "", auth.ErrInvalidPhone
	}

	salt := make([]byte, 8)
	rand.Read(salt)
	hashedPassword := HashPassword(salt, data.Password)

	newUser := models.User{
		Login:        data.Login,
		PhoneNumber:  data.PhoneNumber,
		Id:           uuid.NewV4(),
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Description:  "",
		UserPic:      "default.png",
		PasswordHash: hashedPassword,
	}

	err := uc.repo.InsertUser(ctx, newUser)
	if err != nil {
		logger.Error(err.Error())
		return models.User{}, "", "", auth.ErrCreatingUser
	}

	token, err := generateToken(newUser.Login, newUser.Id)
	if err != nil {
		logger.Error(err.Error())
		return models.User{}, "", "", auth.ErrGeneratingToken
	}

	csrfToken := uuid.NewV4().String()

	logger.Info("Successful")
	return newUser, token, csrfToken, nil
}

func (uc *AuthUsecase) Check(ctx context.Context, login string) (models.User, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	user, err := uc.repo.SelectUserByLogin(ctx, login)
	if err != nil {
		logger.Error(err.Error())
		return models.User{}, auth.ErrUserNotFound
	}

	return user, nil
}

func (uc *AuthUsecase) UpdateUser(ctx context.Context, login string, updateData models.UpdateUserReq) (models.User, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	if updateData.Password != "" && !validPassword(updateData.Password) {
		logger.Error(auth.ErrInvalidPassword.Error())
		return models.User{}, auth.ErrInvalidPassword
	}

	if (updateData.FirstName != "" && !isValidName(updateData.FirstName)) || (updateData.LastName != "" && !isValidName(updateData.LastName)) {
		logger.Error(auth.ErrInvalidName.Error())
		return models.User{}, auth.ErrInvalidName
	}

	if updateData.PhoneNumber != "" && !isValidPhone(updateData.PhoneNumber) {
		logger.Error(auth.ErrInvalidPhone.Error())
		return models.User{}, auth.ErrInvalidPhone
	}

	user, err := uc.repo.SelectUserByLogin(ctx, login)
	if err != nil {
		logger.Error(err.Error())
		return models.User{}, auth.ErrUserNotFound
	}

	if updateData.Password != "" {
		salt := make([]byte, 8)
		rand.Read(salt)
		hashedPassword := HashPassword(salt, updateData.Password)

		if bytes.Equal(hashedPassword, user.PasswordHash) {
			return models.User{}, auth.ErrSamePassword
		}

		user.PasswordHash = hashedPassword
	}

	if updateData.FirstName != "" && updateData.FirstName != user.FirstName {
		user.FirstName = updateData.FirstName
	} else if updateData.FirstName == user.FirstName {
		logger.Error(auth.ErrSameName.Error())
		return models.User{}, auth.ErrSameName
	}

	if updateData.LastName != "" && updateData.LastName != user.LastName {
		user.LastName = updateData.LastName
	} else if updateData.LastName == user.LastName {
		logger.Error(auth.ErrSameName.Error())
		return models.User{}, auth.ErrSameName
	}

	if updateData.PhoneNumber != "" && updateData.PhoneNumber != user.PhoneNumber {
		user.PhoneNumber = updateData.PhoneNumber
	} else if updateData.PhoneNumber == user.PhoneNumber {
		logger.Error(auth.ErrSamePhone.Error())
		return models.User{}, auth.ErrSamePhone
	}

	user.Description = updateData.Description

	err = uc.repo.UpdateUser(ctx, user)
	if err != nil {
		logger.Error(err.Error())
		return models.User{}, err
	}

	logger.Info("Successful")
	return user, nil
}

func (uc *AuthUsecase) UpdateUserPic(ctx context.Context, login string, picture io.ReadSeeker, extension string) (models.User, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	user, err := uc.repo.SelectUserByLogin(ctx, login)
	if err != nil {
		logger.Error(err.Error())
		return models.User{}, auth.ErrUserNotFound
	}

	pictureBasePath := os.Getenv("PICTURE_BASE_PATH")
	if pictureBasePath == "" {
		logger.Error(auth.ErrBasePath.Error())
		return models.User{}, auth.ErrBasePath
	}

	imageName := uuid.NewV4().String()
	newImagePath := path.Join(pictureBasePath, imageName+extension)

	dst, err := os.Create(newImagePath)
	if err != nil {
		logger.Error(err.Error())
		return models.User{}, auth.ErrFileCreation
	}
	defer dst.Close()

	if _, err := io.Copy(dst, picture); err != nil {
		logger.Error(err.Error())
		return models.User{}, auth.ErrFileSaving
	}

	if err := uc.repo.UpdateUserPic(ctx, login, imageName+extension); err != nil {
		logger.Error(err.Error())
		return models.User{}, err
	}

	if user.UserPic != "default.png" {
		oldImagePath := path.Join(pictureBasePath, user.UserPic)
		if err := os.Remove(oldImagePath); err != nil && !os.IsNotExist(err) {
			logger.Error(err.Error())
			return models.User{}, auth.ErrFileDeletion
		}
	}

	user.UserPic = imageName + extension

	logger.Info("Successful")
	return user, nil
}

func (uc *AuthUsecase) GetUserAddresses(ctx context.Context, login string) ([]models.Address, error) {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	addresses, err := uc.repo.SelectUserAddresses(ctx, login)
	if err != nil {
		logger.Error(err.Error())
		return []models.Address{}, err
	}

	logger.Info("Successful")
	return addresses, nil
}

func (uc *AuthUsecase) DeleteAddress(ctx context.Context, addressId uuid.UUID) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	err := uc.repo.DeleteAddress(ctx, addressId)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("Successful")
	return nil
}

func (uc *AuthUsecase) AddAddress(ctx context.Context, address models.Address) error {
	logger := log.GetLoggerFromContext(ctx).With(slog.String("func", log.GetFuncName()))

	address.Id = uuid.NewV4()
	err := uc.repo.InsertAddress(ctx, address)
	if err != nil {
		logger.Error(err.Error())
		return err
	}

	logger.Info("Successful")
	return nil
}
