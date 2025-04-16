package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/mocks"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
)

func TestUpdateItemQuantity(t *testing.T) {
	type args struct {
		userID       string
		productID    string
		restaurantID string
		quantity     int
	}
	tests := []struct {
		name       string
		args       args
		repoMocker func(*mocks.MockCartRepo)
		wantErr    error
	}{
		{
			name: "Success",
			args: args{
				userID:       "user123",
				productID:    "product456",
				restaurantID: "restaurant789",
				quantity:     3,
			},
			repoMocker: func(repo *mocks.MockCartRepo) {
				repo.EXPECT().UpdateItemQuantity(gomock.Any(), "user123", "product456", "restaurant789", 3).Return(nil).Times(1)
			},
			wantErr: nil,
		},
		{
			name: "Update quantity failure",
			args: args{
				userID:       "user123",
				productID:    "product456",
				restaurantID: "restaurant789",
				quantity:     3,
			},
			repoMocker: func(repo *mocks.MockCartRepo) {
				repo.EXPECT().UpdateItemQuantity(gomock.Any(), "user123", "product456", "restaurant789", 3).Return(errors.New("update error")).Times(1)
			},
			wantErr: errors.New("update error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockCartRepo(ctrl)
			uc := NewCartUsecase(repo, nil)

			tt.repoMocker(repo)

			err := uc.UpdateItemQuantity(context.Background(), tt.args.userID, tt.args.productID, tt.args.restaurantID, tt.args.quantity)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("UpdateItemQuantity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClearCart(t *testing.T) {
	tests := []struct {
		name       string
		userID     string
		repoMocker func(*mocks.MockCartRepo)
		wantErr    error
	}{
		{
			name:   "Success",
			userID: "user123",
			repoMocker: func(repo *mocks.MockCartRepo) {
				repo.EXPECT().ClearCart(gomock.Any(), "user123").Return(nil).Times(1)
			},
			wantErr: nil,
		},
		{
			name:   "Clear cart failure",
			userID: "user123",
			repoMocker: func(repo *mocks.MockCartRepo) {
				repo.EXPECT().ClearCart(gomock.Any(), "user123").Return(errors.New("clear error")).Times(1)
			},
			wantErr: errors.New("clear error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockCartRepo(ctrl)
			uc := NewCartUsecase(repo, nil)

			tt.repoMocker(repo)

			err := uc.ClearCart(context.Background(), tt.userID)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("ClearCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateOrder(t *testing.T) {
	type args struct {
		userID string
		req    models.OrderInReq
		cart   models.Cart
	}
	tests := []struct {
		name       string
		args       args
		repoMocker func(*mocks.MockRestaurantRepo)
		wantErr    error
	}{
		{
			name: "Success",
			args: args{
				userID: "user123",
				req: models.OrderInReq{
					Status:     "new",
					Address:    "123 Street",
					FinalPrice: 100.50,
				},
				cart: models.Cart{
					Id:   uuid.NewV4(),
					Name: "Test Cart",
				},
			},
			repoMocker: func(repo *mocks.MockRestaurantRepo) {
				repo.EXPECT().Save(gomock.Any(), gomock.Any(), "user123").Return(nil).Times(1)
			},
			wantErr: nil,
		},
		{
			name: "Save order failure",
			args: args{
				userID: "user123",
				req: models.OrderInReq{
					Status:     "new",
					Address:    "123 Street",
					FinalPrice: 100.50,
				},
				cart: models.Cart{
					Id:   uuid.NewV4(),
					Name: "Test Cart",
				},
			},
			repoMocker: func(repo *mocks.MockRestaurantRepo) {
				repo.EXPECT().Save(gomock.Any(), gomock.Any(), "user123").Return(errors.New("save error")).Times(1)
			},
			wantErr: errors.New("save error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := mocks.NewMockRestaurantRepo(ctrl)
			uc := NewCartUsecase(nil, repo)

			tt.repoMocker(repo)

			_, err := uc.CreateOrder(context.Background(), tt.args.userID, tt.args.req, tt.args.cart)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("CreateOrder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
