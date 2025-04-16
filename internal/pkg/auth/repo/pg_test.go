package repo

import (
	"context"
	"errors"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/usecase"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertUser(t *testing.T) {
	salt := make([]byte, 8)
	userId := uuid.NewV4()
	testUser := models.User{
		Login:        "test_user",
		PasswordHash: usecase.HashPassword(salt, "password123"),
		Id:           userId,
		FirstName:    "Тайлер",
		LastName:     "Дерден",
		PhoneNumber:  "88005553535",
		Description:  "Some User",
		UserPic:      "default.png",
	}

	tests := []struct {
		name           string
		repoMocker     func(*pgxpoolmock.MockPgxPool)
		expectedLogger string
		err            error
	}{
		{
			name: "Success",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(gomock.Any(), insertUser,
					testUser.Id,
					testUser.Login,
					testUser.FirstName,
					testUser.LastName,
					testUser.PhoneNumber,
					testUser.Description,
					testUser.UserPic,
					testUser.PasswordHash,
				).Return(nil, nil)
			},
			expectedLogger: "Success",
			err:            nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			defer ctrl.Finish()

			test.repoMocker(mockPool)

			repo := CreateAuthRepo(mockPool)
			err := repo.InsertUser(context.Background(), models.User{
				Login:        "test_user",
				PasswordHash: usecase.HashPassword(salt, "password123"),
				Id:           userId,
				FirstName:    "Тайлер",
				LastName:     "Дерден",
				PhoneNumber:  "88005553535",
				Description:  "Some User",
				UserPic:      "default.png",
			})

			assert.Equal(t, test.err, err)
		})
	}

}

func TestSelectUserByLogin(t *testing.T) {
	columns := []string{"id", "first_name", "last_name", "phone_number", "description", "user_pic", "password_hash"}

	salt := make([]byte, 8)
	userId := uuid.NewV4()
	testUser := models.User{
		Login:        "test_user",
		PasswordHash: usecase.HashPassword(salt, "password123"),
		Id:           userId,
		FirstName:    "Тайлер",
		LastName:     "Дерден",
		PhoneNumber:  "88005553535",
		Description:  "Some User",
		UserPic:      "default.png",
	}

	tests := []struct {
		name         string
		repoMocker   func(*pgxpoolmock.MockPgxPool, pgx.Rows, string)
		login        string
		expectedUser models.User
		expectedErr  error
	}{
		{
			name: "Success",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows, login string) {
				mockPool.EXPECT().QueryRow(gomock.Any(), selectUserByLogin, login).Return(pgxRows)
			},
			login:        testUser.Login,
			expectedUser: testUser,
			expectedErr:  nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			defer ctrl.Finish()

			pgxRows := pgxpoolmock.NewRows(columns).
				AddRow(
					testUser.Id,
					testUser.FirstName,
					testUser.LastName,
					testUser.PhoneNumber,
					testUser.Description,
					testUser.UserPic,
					testUser.PasswordHash,
				).ToPgxRows()
			pgxRows.Next()
			test.repoMocker(mockPool, pgxRows, test.login)

			repo := CreateAuthRepo(mockPool)
			_, err := repo.SelectUserByLogin(context.Background(), test.login)

			assert.Equal(t, test.expectedErr, err)

		})
	}
}

func TestUpdateUser(t *testing.T) {
	salt := make([]byte, 8)
	userId := uuid.NewV4()

	testUser := models.User{
		Login:        "test_user",
		PasswordHash: usecase.HashPassword(salt, "password123"),
		Id:           userId,
		FirstName:    "Тайлер",
		LastName:     "Дерден",
		PhoneNumber:  "88005553535",
		Description:  "Some User",
		UserPic:      "default.png",
	}

	tests := []struct {
		name        string
		repoMocker  func(*pgxpoolmock.MockPgxPool)
		user        models.User
		expectedErr error
	}{
		{
			name: "Success",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(
					gomock.Any(),
					updateUser,
					testUser.PhoneNumber,
					testUser.FirstName,
					testUser.LastName,
					testUser.Description,
					testUser.PasswordHash,
					testUser.Id,
				).Return(nil, nil)
			},
			user:        testUser,
			expectedErr: nil,
		},
		{
			name: "DB Error",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(
					gomock.Any(),
					updateUser,
					testUser.PhoneNumber,
					testUser.FirstName,
					testUser.LastName,
					testUser.Description,
					testUser.PasswordHash,
					testUser.Id,
				).Return(nil, errors.New("db error"))
			},
			user:        testUser,
			expectedErr: errors.New("db error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			defer ctrl.Finish()

			test.repoMocker(mockPool)

			repo := CreateAuthRepo(mockPool)
			err := repo.UpdateUser(context.Background(), test.user)

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateUserPic(t *testing.T) {
	testLogin := "test_user"
	testPic := "new_pic.png"

	tests := []struct {
		name        string
		repoMocker  func(*pgxpoolmock.MockPgxPool)
		login       string
		userPic     string
		expectedErr error
	}{
		{
			name: "Success",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(
					gomock.Any(),
					updateUserPic,
					testPic,
					testLogin,
				).Return(nil, nil)
			},
			login:       testLogin,
			userPic:     testPic,
			expectedErr: nil,
		},
		{
			name: "DB Error",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(
					gomock.Any(),
					updateUserPic,
					testPic,
					testLogin,
				).Return(nil, errors.New("update failed"))
			},
			login:       testLogin,
			userPic:     testPic,
			expectedErr: errors.New("update failed"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			defer ctrl.Finish()

			test.repoMocker(mockPool)

			repo := CreateAuthRepo(mockPool)
			err := repo.UpdateUserPic(context.Background(), test.login, test.userPic)

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestSelectUserAddresses(t *testing.T) {
	testLogin := "test_user"
	testAddress := models.Address{
		Id:      uuid.NewV4(),
		Address: "123 Test Street",
		UserId:  uuid.NewV4(),
	}

	columns := []string{"id", "address", "user_id"}

	tests := []struct {
		name           string
		repoMocker     func(*pgxpoolmock.MockPgxPool, pgx.Rows, string)
		login          string
		expectedResult []models.Address
		expectedErr    error
	}{
		{
			name: "Success",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows, login string) {
				mockPool.EXPECT().Query(gomock.Any(), selectUserAddresses, login).Return(pgxRows, nil)
			},
			login: testLogin,
			expectedResult: []models.Address{
				{
					Id:      testAddress.Id,
					Address: testAddress.Address,
					UserId:  testAddress.UserId,
				},
			},
			expectedErr: nil,
		},
		{
			name: "Query error",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool, _ pgx.Rows, login string) {
				mockPool.EXPECT().Query(gomock.Any(), selectUserAddresses, login).Return(nil, errors.New("query failed"))
			},
			login:          testLogin,
			expectedResult: nil,
			expectedErr:    errors.New("query failed"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			defer ctrl.Finish()

			pgxRows := pgxpoolmock.NewRows(columns).
				AddRow(
					testAddress.Id,
					testAddress.Address,
					testAddress.UserId,
				).ToPgxRows()

			test.repoMocker(mockPool, pgxRows, test.login)

			repo := CreateAuthRepo(mockPool)
			result, err := repo.SelectUserAddresses(context.Background(), test.login)

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedResult, result)
			}
		})
	}
}

func TestDeleteAddress(t *testing.T) {
	testAddressID := uuid.NewV4()

	tests := []struct {
		name        string
		setupMock   func(mockPool *pgxpoolmock.MockPgxPool)
		addressID   uuid.UUID
		expectedErr error
	}{
		{
			name: "Success",
			setupMock: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().
					Exec(gomock.Any(), deleteAddress, testAddressID).
					Return(pgconn.CommandTag("DELETE 1"), nil)
			},
			addressID:   testAddressID,
			expectedErr: nil,
		},
		{
			name: "Address not found",
			setupMock: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().
					Exec(gomock.Any(), deleteAddress, testAddressID).
					Return(pgconn.CommandTag("DELETE 0"), nil)
			},
			addressID:   testAddressID,
			expectedErr: errors.New("Адрес не найден"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			defer ctrl.Finish()

			test.setupMock(mockPool)

			repo := CreateAuthRepo(mockPool)
			err := repo.DeleteAddress(context.Background(), test.addressID)

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestInsertAddress(t *testing.T) {
	testAddress := models.Address{
		Id:      uuid.NewV4(),
		Address: "123 Test Street",
		UserId:  uuid.NewV4(),
	}

	tests := []struct {
		name        string
		repoMocker  func(*pgxpoolmock.MockPgxPool)
		address     models.Address
		expectedErr error
	}{
		{
			name: "Success",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(
					gomock.Any(),
					insertAddress,
					testAddress.Id,
					testAddress.Address,
					testAddress.UserId,
				).Return(nil, nil)
			},
			address:     testAddress,
			expectedErr: nil,
		},
		{
			name: "DB Error",
			repoMocker: func(mockPool *pgxpoolmock.MockPgxPool) {
				mockPool.EXPECT().Exec(
					gomock.Any(),
					insertAddress,
					testAddress.Id,
					testAddress.Address,
					testAddress.UserId,
				).Return(nil, errors.New("db error"))
			},
			address:     testAddress,
			expectedErr: errors.New("db error"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockPool := pgxpoolmock.NewMockPgxPool(ctrl)
			defer ctrl.Finish()

			test.repoMocker(mockPool)

			repo := CreateAuthRepo(mockPool)
			err := repo.InsertAddress(context.Background(), test.address)

			if test.expectedErr != nil {
				assert.EqualError(t, err, test.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
