package repo

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/driftprogramming/pgxpoolmock"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/usecase"
	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
)

var testLogger *slog.Logger

func init() {
	testLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}

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
		name               string
		expectedRepoAction func(*pgxpoolmock.MockPgxPool)
		expectedLogger     string
		err                error
	}{
		{
			name: "Success",
			expectedRepoAction: func(mockPool *pgxpoolmock.MockPgxPool) {
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

			test.expectedRepoAction(mockPool)

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
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    ctx := context.WithValue(context.Background(), "logger", logger)
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
		name               string
		expectedRepoAction func(*pgxpoolmock.MockPgxPool, pgx.Rows, string)
		login              string
		expectedUser       models.User
		expectedErr        error
	}{
		{
			name: "Success",
			expectedRepoAction: func(mockPool *pgxpoolmock.MockPgxPool, pgxRows pgx.Rows, login string) {
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

			test.expectedRepoAction(mockPool, pgxRows, test.login)

			repo := CreateAuthRepo(mockPool)
			_, err := repo.SelectUserByLogin(ctx, test.login)

			assert.Equal(t, test.expectedErr, err)
			
		})
	}
}
