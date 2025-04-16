package usecase

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/mocks"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
)

func TestSignIn(t *testing.T) {
	salt := make([]byte, 8)
	os.Setenv("JWT_SECRET", "testsecret")

	type args struct {
		data models.SignInReq
	}
	tests := []struct {
		name       string
		repoMocker func(*mocks.MockAuthRepo, string, string)
		args       args
		wantErr    error
	}{
		{
			name: "Success",
			repoMocker: func(repo *mocks.MockAuthRepo, login, password string) {
				repo.EXPECT().SelectUserByLogin(gomock.Any(), login).Return(models.User{
					Id:           uuid.NewV4(),
					Login:        login,
					PasswordHash: HashPassword(salt, password),
				}, nil).Times(1)
			},
			args: args{
				data: models.SignInReq{
					Login:    "testuser",
					Password: "Pass@123",
				},
			},
			wantErr: nil,
		},
		{
			name:       "Invalid login format",
			repoMocker: func(repo *mocks.MockAuthRepo, _, _ string) {},
			args: args{
				data: models.SignInReq{
					Login:    "inv@lid!",
					Password: "Pass@123",
				},
			},
			wantErr: auth.ErrInvalidLogin,
		},
		{
			name: "User not found",
			repoMocker: func(repo *mocks.MockAuthRepo, login, _ string) {
				repo.EXPECT().SelectUserByLogin(gomock.Any(), login).Return(models.User{}, auth.ErrUserNotFound).Times(1)
			},
			args: args{
				data: models.SignInReq{
					Login:    "nouser",
					Password: "Pass@123",
				},
			},
			wantErr: auth.ErrUserNotFound,
		},
		{
			name: "Wrong password",
			repoMocker: func(repo *mocks.MockAuthRepo, login, _ string) {
				repo.EXPECT().SelectUserByLogin(gomock.Any(), login).Return(models.User{
					Id:           uuid.NewV4(),
					Login:        login,
					PasswordHash: HashPassword(salt, "Correct@123"),
				}, nil).Times(1)
			},
			args: args{
				data: models.SignInReq{
					Login:    "testuser",
					Password: "WrongPassword1!",
				},
			},
			wantErr: auth.ErrInvalidCredentials,
		},
		{
			name: "Token generation error (no secret)",
			repoMocker: func(repo *mocks.MockAuthRepo, login, password string) {
				repo.EXPECT().SelectUserByLogin(gomock.Any(), login).Return(models.User{
					Id:           uuid.NewV4(),
					Login:        login,
					PasswordHash: HashPassword(salt, password),
				}, nil).Times(1)

				os.Setenv("JWT_SECRET", "")
			},
			args: args{
				data: models.SignInReq{
					Login:    "testuser",
					Password: "Pass@123",
				},
			},
			wantErr: auth.ErrGeneratingToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockAuthRepo(ctrl)
			uc := CreateAuthUsecase(repo)

			tt.repoMocker(repo, tt.args.data.Login, tt.args.data.Password)

			_, token, csrf, err := uc.SignIn(context.Background(), tt.args.data)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("SignIn() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil {
				if token == "" || csrf == "" {
					t.Errorf("Expected non-empty token and csrf, got token: '%s', csrf: '%s'", token, csrf)
				}
			}
		})
	}
}

func TestSignUp(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")

	type args struct {
		data models.SignUpReq
	}

	tests := []struct {
		name       string
		args       args
		repoMocker func(*mocks.MockAuthRepo, models.User)
		wantErr    error
	}{
		{
			name: "Success",
			args: args{
				data: models.SignUpReq{
					Login:       "newuser",
					Password:    "Password@123",
					FirstName:   "Иван",
					LastName:    "Петров",
					PhoneNumber: "1234567890",
				},
			},
			repoMocker: func(repo *mocks.MockAuthRepo, user models.User) {
				repo.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: nil,
		},
		{
			name: "Invalid login",
			args: args{
				data: models.SignUpReq{
					Login:       "inv@lid",
					Password:    "Password@123",
					FirstName:   "Иван",
					LastName:    "Петров",
					PhoneNumber: "1234567890",
				},
			},
			repoMocker: func(repo *mocks.MockAuthRepo, user models.User) {},
			wantErr:    auth.ErrInvalidLogin,
		},
		{
			name: "Invalid password",
			args: args{
				data: models.SignUpReq{
					Login:       "validlogin",
					Password:    "nopunct",
					FirstName:   "Иван",
					LastName:    "Петров",
					PhoneNumber: "1234567890",
				},
			},
			repoMocker: func(repo *mocks.MockAuthRepo, user models.User) {},
			wantErr:    auth.ErrInvalidPassword,
		},
		{
			name: "Invalid name",
			args: args{
				data: models.SignUpReq{
					Login:       "validlogin",
					Password:    "Password@123",
					FirstName:   "John1",
					LastName:    "Петров",
					PhoneNumber: "1234567890",
				},
			},
			repoMocker: func(repo *mocks.MockAuthRepo, user models.User) {},
			wantErr:    auth.ErrInvalidName,
		},
		{
			name: "Invalid phone",
			args: args{
				data: models.SignUpReq{
					Login:       "validlogin",
					Password:    "Password@123",
					FirstName:   "Иван",
					LastName:    "Петров",
					PhoneNumber: "123-456",
				},
			},
			repoMocker: func(repo *mocks.MockAuthRepo, user models.User) {},
			wantErr:    auth.ErrInvalidPhone,
		},
		{
			name: "Insert user error",
			args: args{
				data: models.SignUpReq{
					Login:       "newuser",
					Password:    "Password@123",
					FirstName:   "Иван",
					LastName:    "Петров",
					PhoneNumber: "1234567890",
				},
			},
			repoMocker: func(repo *mocks.MockAuthRepo, user models.User) {
				repo.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(auth.ErrCreatingUser).Times(1)
			},
			wantErr: auth.ErrCreatingUser,
		},
		{
			name: "Token generation failure (no secret)",
			args: args{
				data: models.SignUpReq{
					Login:       "newuser",
					Password:    "Password@123",
					FirstName:   "Иван",
					LastName:    "Петров",
					PhoneNumber: "1234567890",
				},
			},
			repoMocker: func(repo *mocks.MockAuthRepo, user models.User) {
				os.Setenv("JWT_SECRET", "")
				repo.EXPECT().InsertUser(gomock.Any(), gomock.Any()).Return(nil).Times(1)
			},
			wantErr: auth.ErrGeneratingToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockAuthRepo(ctrl)
			uc := CreateAuthUsecase(repo)

			testUser := models.User{
				Login:       tt.args.data.Login,
				PhoneNumber: tt.args.data.PhoneNumber,
				FirstName:   tt.args.data.FirstName,
				LastName:    tt.args.data.LastName,
			}

			tt.repoMocker(repo, testUser)

			_, token, csrf, err := uc.SignUp(context.Background(), tt.args.data)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err == nil && (token == "" || csrf == "") {
				t.Errorf("Expected non-empty token and csrf, got token: '%s', csrf: '%s'", token, csrf)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	tests := []struct {
		name       string
		login      string
		repoMocker func(*mocks.MockAuthRepo, string)
		wantErr    error
	}{
		{
			name:  "Success",
			login: "validuser",
			repoMocker: func(repo *mocks.MockAuthRepo, login string) {
				repo.EXPECT().SelectUserByLogin(gomock.Any(), login).Return(models.User{
					Id:    uuid.NewV4(),
					Login: login,
				}, nil).Times(1)
			},
			wantErr: nil,
		},
		{
			name:  "User not found",
			login: "missinguser",
			repoMocker: func(repo *mocks.MockAuthRepo, login string) {
				repo.EXPECT().SelectUserByLogin(gomock.Any(), login).Return(models.User{}, auth.ErrUserNotFound).Times(1)
			},
			wantErr: auth.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockAuthRepo(ctrl)
			uc := CreateAuthUsecase(repo)

			tt.repoMocker(repo, tt.login)

			_, err := uc.Check(context.Background(), tt.login)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	salt := make([]byte, 8)

	oldPass := "OldPass@123"
	oldHash := HashPassword(salt, oldPass)

	oldUser := models.User{
		Id:           uuid.NewV4(),
		Login:        "testuser",
		FirstName:    "Иван",
		LastName:     "Петров",
		PhoneNumber:  "89101112233",
		PasswordHash: oldHash,
		Description:  "Старое описание",
	}

	tests := []struct {
		name        string
		login       string
		updateData  models.UpdateUserReq
		repoMocker  func(*mocks.MockAuthRepo)
		expectedErr error
	}{
		{
			name:  "Successful update",
			login: oldUser.Login,
			updateData: models.UpdateUserReq{
				Password:    "NewPass@123",
				FirstName:   "Сергей",
				LastName:    "Сидоров",
				PhoneNumber: "89223334455",
				Description: "Новое описание",
			},
			repoMocker: func(repo *mocks.MockAuthRepo) {
				repo.EXPECT().
					SelectUserByLogin(gomock.Any(), oldUser.Login).
					Return(oldUser, nil).Times(1)

				repo.EXPECT().
					UpdateUser(gomock.Any(), gomock.Any()).
					Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name:  "Invalid password format",
			login: oldUser.Login,
			updateData: models.UpdateUserReq{
				Password: "short",
			},
			repoMocker:  func(repo *mocks.MockAuthRepo) {},
			expectedErr: auth.ErrInvalidPassword,
		},
		{
			name:  "User not found",
			login: "nouser",
			updateData: models.UpdateUserReq{
				FirstName: "Игорь",
			},
			repoMocker: func(repo *mocks.MockAuthRepo) {
				repo.EXPECT().
					SelectUserByLogin(gomock.Any(), "nouser").
					Return(models.User{}, auth.ErrUserNotFound).Times(1)
			},
			expectedErr: auth.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockAuthRepo(ctrl)
			uc := CreateAuthUsecase(repo)

			tt.repoMocker(repo)

			_, err := uc.UpdateUser(context.Background(), tt.login, tt.updateData)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("UpdateUser() error = %v, wantErr = %v", err, tt.expectedErr)
			}
		})
	}
}

func TestGetUserAddresses(t *testing.T) {
	login := "testuser"
	addressList := []models.Address{
		{Id: uuid.NewV4(), Address: "г. Москва, ул. Ленина, д. 1", UserId: uuid.NewV4()},
	}

	tests := []struct {
		name        string
		login       string
		repoMocker  func(*mocks.MockAuthRepo)
		expectedErr error
	}{
		{
			name:  "Successful",
			login: login,
			repoMocker: func(repo *mocks.MockAuthRepo) {
				repo.EXPECT().
					SelectUserAddresses(gomock.Any(), login).
					Return(addressList, nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name:  "Error while fetching addresses",
			login: login,
			repoMocker: func(repo *mocks.MockAuthRepo) {
				repo.EXPECT().
					SelectUserAddresses(gomock.Any(), login).
					Return(nil, auth.ErrDBError).Times(1)
			},
			expectedErr: auth.ErrDBError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockAuthRepo(ctrl)
			uc := CreateAuthUsecase(mockRepo)
			tt.repoMocker(mockRepo)

			_, err := uc.GetUserAddresses(context.Background(), tt.login)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("GetUserAddresses() error = %v, wantErr = %v", err, tt.expectedErr)
			}
		})
	}
}

func TestDeleteAddress(t *testing.T) {
	addressId := uuid.NewV4()

	tests := []struct {
		name        string
		addressId   uuid.UUID
		repoMocker  func(*mocks.MockAuthRepo)
		expectedErr error
	}{
		{
			name:      "Successful",
			addressId: addressId,
			repoMocker: func(repo *mocks.MockAuthRepo) {
				repo.EXPECT().
					DeleteAddress(gomock.Any(), addressId).
					Return(nil).Times(1)
			},
			expectedErr: nil,
		},
		{
			name:      "Error deleting address",
			addressId: addressId,
			repoMocker: func(repo *mocks.MockAuthRepo) {
				repo.EXPECT().
					DeleteAddress(gomock.Any(), addressId).
					Return(auth.ErrDBError).Times(1)
			},
			expectedErr: auth.ErrDBError,
		},
		{
			name:      "Address not found",
			addressId: uuid.NewV4(),
			repoMocker: func(repo *mocks.MockAuthRepo) {
				repo.EXPECT().
					DeleteAddress(gomock.Any(), gomock.Any()).
					Return(auth.ErrAddressNotFound).Times(1)
			},
			expectedErr: auth.ErrAddressNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockAuthRepo(ctrl)
			uc := CreateAuthUsecase(mockRepo)
			tt.repoMocker(mockRepo)

			err := uc.DeleteAddress(context.Background(), tt.addressId)

			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("DeleteAddress() error = %v, wantErr = %v", err, tt.expectedErr)
			}
		})
	}
}
