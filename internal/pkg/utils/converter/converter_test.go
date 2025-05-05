package converter

import (
	"testing"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc/gen"
	"github.com/satori/uuid"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCartConversion(t *testing.T) {
	cartID := uuid.NewV4()
	itemID := uuid.NewV4()

	tests := []struct {
		name      string
		input     models.Cart
		expectErr bool
	}{
		{
			name: "Normal cart",
			input: models.Cart{
				Id:   cartID,
				Name: "Test Rest",
				CartItems: []models.CartItem{
					{
						Id:       itemID,
						Name:     "Burger",
						Price:    500,
						ImageURL: "http://img",
						Weight:   300,
						Amount:   1,
					},
				},
			},
			expectErr: false,
		},
		{
			name: "Empty cart",
			input: models.Cart{
				Id:        cartID,
				Name:      "Empty Rest",
				CartItems: []models.CartItem{},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proto := CartToProto(tt.input)
			result, err := ProtoToCart(proto)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.input.Id, result.Id)
				assert.Equal(t, tt.input.Name, result.Name)
				assert.Equal(t, len(tt.input.CartItems), len(result.CartItems))
			}
		})
	}
}

func TestOrderConversion(t *testing.T) {
	orderID := uuid.NewV4()
	userID := "testuser"
	now := time.Now().Truncate(time.Second)
	itemID := uuid.NewV4()

	tests := []struct {
		name      string
		order     models.Order
		expectErr bool
	}{
		{
			name: "Valid order",
			order: models.Order{
				ID:                orderID,
				UserID:            userID,
				Status:            "created",
				Address:           "Main St",
				OrderProducts:     models.Cart{
					Id:   uuid.NewV4(),
					Name: "Place",
					CartItems: []models.CartItem{
						{
							Id:       itemID,
							Name:     "Item",
							Price:    123,
							ImageURL: "url",
							Weight:   123,
							Amount:   1,
						},
					},
				},
				ApartmentOrOffice: "12",
				Intercom:          "123",
				Entrance:          "1",
				Floor:             "2",
				CourierComment:    "comment",
				LeaveAtDoor:       true,
				CreatedAt:         now,
				FinalPrice:        1000,
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proto, err := OrderToProto(tt.order, tt.order.UserID)
			assert.NoError(t, err)

			orderResult, err := ProtoToOrder(proto)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.order.ID, orderResult.ID)
				assert.Equal(t, tt.order.Address, orderResult.Address)
				assert.Equal(t, tt.order.CreatedAt.Unix(), orderResult.CreatedAt.Unix())
			}
		})
	}
}

func TestProtoToCartItems_Errors(t *testing.T) {
	tests := []struct {
		name      string
		input     []*gen.CartItem
		expectErr bool
	}{
		{
			name: "Nil input",
			input: nil,
			expectErr: false,
		},
		{
			name: "Invalid UUID",
			input: []*gen.CartItem{
				{Id: "invalid-uuid"},
			},
			expectErr: true,
		},
		{
			name: "Valid item",
			input: []*gen.CartItem{
				{Id: uuid.NewV4().String(), Name: "Test", Price: 100, Weight: 100, Amount: 1},
			},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ProtoToCartItems(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestProtoToOrder_Errors(t *testing.T) {
	tests := []struct {
		name      string
		input     *gen.OrderResponse
		expectErr bool
	}{
		{
			name:      "Nil input",
			input:     nil,
			expectErr: true,
		},
		{
			name: "Invalid UUID",
			input: &gen.OrderResponse{
				Id: "bad-uuid",
			},
			expectErr: true,
		},
		{
			name: "Invalid timestamp",
			input: &gen.OrderResponse{
				Id:        uuid.NewV4().String(),
				UserId:    "user",
				CreatedAt: &timestamppb.Timestamp{Seconds: -999999999999},
			},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ProtoToOrder(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
