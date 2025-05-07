package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/mocks"
	"github.com/golang/mock/gomock"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
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

func TestGetCart(t *testing.T) {
	restaurantId := uuid.NewV4()
	product1Id := uuid.NewV4()
	product2Id := uuid.NewV4()
	type args struct {
		login string
	}
	tests := []struct {
		name               string
		args               args
		cartRepoMock       func(*mocks.MockCartRepo)
		restaurantRepoMock func(*mocks.MockRestaurantRepo)
		want               models.Cart
		wantErr            error
		wantBool           bool
	}{
		{
			name: "Success",
			args: args{login: "user123"},
			cartRepoMock: func(repo *mocks.MockCartRepo) {
				repo.EXPECT().
					GetCart(gomock.Any(), "user123").
					Return(
						map[string]int{"product1": 2, "product2": 1},
						"restaurant123",
						nil,
					).Times(1)
			},
			restaurantRepoMock: func(repo *mocks.MockRestaurantRepo) {
				repo.EXPECT().
					GetCartItem(
						gomock.Any(),
						[]string{"product1", "product2"},
						map[string]int{"product1": 2, "product2": 1},
						"restaurant123",
					).
					Return(models.Cart{
						Id:   restaurantId,
						Name: "Test Restaurant",
						CartItems: []models.CartItem{
							{
								Id:       product1Id,
								Name:     "Product 1",
								Price:    100,
								ImageURL: "image1.jpg",
								Weight:   200,
								Amount:   2,
							},
							{
								Id:       product2Id,
								Name:     "Product 2",
								Price:    200,
								ImageURL: "image2.jpg",
								Weight:   300,
								Amount:   1,
							},
						},
					}, nil).Times(1)
			},
			want: models.Cart{
				Id:   restaurantId,
				Name: "Test Restaurant",
				CartItems: []models.CartItem{
					{
						Id:       product1Id,
						Name:     "Product 1",
						Price:    100,
						ImageURL: "image1.jpg",
						Weight:   200,
						Amount:   2,
					},
					{
						Id:       product2Id,
						Name:     "Product 2",
						Price:    200,
						ImageURL: "image2.jpg",
						Weight:   300,
						Amount:   1,
					},
				},
			},
			wantErr:  nil,
			wantBool: true,
		},
		{
			name: "Empty cart",
			args: args{login: "user123"},
			cartRepoMock: func(repo *mocks.MockCartRepo) {
				repo.EXPECT().
					GetCart(gomock.Any(), "user123").
					Return(nil, "", nil).
					Times(1)
			},
			restaurantRepoMock: func(repo *mocks.MockRestaurantRepo) {},
			want:               models.Cart{},
			wantErr:            nil,
			wantBool:           false,
		},
		{
			name: "Get cart error",
			args: args{login: "user123"},
			cartRepoMock: func(repo *mocks.MockCartRepo) {
				repo.EXPECT().
					GetCart(gomock.Any(), "user123").
					Return(nil, "", errors.New("cart error")).
					Times(1)
			},
			restaurantRepoMock: func(repo *mocks.MockRestaurantRepo) {},
			want:               models.Cart{},
			wantErr:            errors.New("cart error"),
			wantBool:           false,
		},
		{
			name: "Get cart items error",
			args: args{login: "user123"},
			cartRepoMock: func(repo *mocks.MockCartRepo) {
				repo.EXPECT().
					GetCart(gomock.Any(), "user123").
					Return(
						map[string]int{"product1": 2},
						"restaurant123",
						nil,
					).Times(1)
			},
			restaurantRepoMock: func(repo *mocks.MockRestaurantRepo) {
				repo.EXPECT().
					GetCartItem(
						gomock.Any(),
						[]string{"product1"},
						map[string]int{"product1": 2},
						"restaurant123",
					).
					Return(models.Cart{}, errors.New("items error")).Times(1)
			},
			want:     models.Cart{},
			wantErr:  errors.New("items error"),
			wantBool: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			cartRepo := mocks.NewMockCartRepo(ctrl)
			restaurantRepo := mocks.NewMockRestaurantRepo(ctrl)
			uc := NewCartUsecase(cartRepo, restaurantRepo)

			tt.cartRepoMock(cartRepo)
			tt.restaurantRepoMock(restaurantRepo)

			got, err, ok := uc.GetCart(context.Background(), tt.args.login)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCart() got = %v, want %v", got, tt.want)
			}
			if (err != nil && tt.wantErr == nil) ||
				(err == nil && tt.wantErr != nil) ||
				(err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("GetCart() error = %v, wantErr %v", err, tt.wantErr)
			}
			if ok != tt.wantBool {
				t.Errorf("GetCart() ok = %v, want %v", ok, tt.wantBool)
			}
		})
	}
}

func TestUpdateOrderStatus(t *testing.T) {
	testOrderID := uuid.NewV4()

	tests := []struct {
		name          string
		orderID       uuid.UUID
		repoMocker    func(*mocks.MockRestaurantRepo)
		expectedError error
	}{
		{
			name:    "Success",
			orderID: testOrderID,
			repoMocker: func(repo *mocks.MockRestaurantRepo) {
				repo.EXPECT().
					UpdateOrderStatus(gomock.Any(), testOrderID, "paid").
					Return(nil).
					Times(1)
			},
			expectedError: nil,
		},
		{
			name:    "Update status error",
			orderID: testOrderID,
			repoMocker: func(repo *mocks.MockRestaurantRepo) {
				repo.EXPECT().
					UpdateOrderStatus(gomock.Any(), testOrderID, "paid").
					Return(errors.New("database error")).
					Times(1)
			},
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			restaurantRepo := mocks.NewMockRestaurantRepo(ctrl)
			uc := &CartUsecase{
				restaurantRepo: restaurantRepo,
			}

			tt.repoMocker(restaurantRepo)

			err := uc.UpdateOrderStatus(context.Background(), tt.orderID)

			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
