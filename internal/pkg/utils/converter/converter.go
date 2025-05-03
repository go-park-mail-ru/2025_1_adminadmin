package converter

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc/gen"
	"github.com/satori/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CartToProto(cart models.Cart) *gen.CartResponse {
	return &gen.CartResponse{
		RestaurantId:   cart.Id.String(),
		RestaurantName: cart.Name,
		Products:       CartItemsToProto(cart.CartItems),
		FullCart:       len(cart.CartItems) > 0,
	}
}

func ProtoToCart(protoCart *gen.CartResponse) (models.Cart, error) {
	if protoCart == nil {
		return models.Cart{}, errors.New("nil cart response")
	}

	restaurantID, err := uuid.FromString(protoCart.RestaurantId)
	if err != nil {
		return models.Cart{}, err
	}

	items, err := ProtoToCartItems(protoCart.Products)
	if err != nil {
		return models.Cart{}, err
	}

	return models.Cart{
		Id:        restaurantID,
		Name:      protoCart.RestaurantName,
		CartItems: items,
	}, nil
}

func CartItemsToProto(items []models.CartItem) []*gen.CartItem {
	if items == nil {
		return nil
	}

	protoItems := make([]*gen.CartItem, len(items))
	for i, item := range items {
		protoItems[i] = &gen.CartItem{
			Id:       item.Id.String(),
			Name:     item.Name,
			Price:    item.Price,
			ImageUrl: item.ImageURL,
			Weight:   int32(item.Weight),
			Amount:   int32(item.Amount),
		}
	}
	return protoItems
}

func ProtoToCartItems(protoItems []*gen.CartItem) ([]models.CartItem, error) {
	if protoItems == nil {
		return nil, nil
	}

	items := make([]models.CartItem, 0, len(protoItems))
	for _, protoItem := range protoItems {
		if protoItem == nil {
			continue
		}

		id, err := uuid.FromString(protoItem.Id)
		if err != nil {
			return nil, err
		}

		items = append(items, models.CartItem{
			Id:       id,
			Name:     protoItem.Name,
			Price:    protoItem.Price,
			ImageURL: protoItem.ImageUrl,
			Weight:   int(protoItem.Weight),
			Amount:   int(protoItem.Amount),
		})
	}
	return items, nil
}

func OrderInReqToProto(req models.OrderInReq, cart models.Cart, login string) *gen.CreateOrderRequest {
	return &gen.CreateOrderRequest{
		Status:            req.Status,
		Address:           req.Address,
		ApartmentOrOffice: req.ApartmentOrOffice,
		Intercom:          req.Intercom,
		Entrance:          req.Entrance,
		Floor:             req.Floor,
		CourierComment:    req.CourierComment,
		LeaveAtDoor:       req.LeaveAtDoor,
		FinalPrice:        req.FinalPrice,
		Cart:              CartToProto(cart),
		Login:             login,
	}
}

func OrderToProto(order models.Order, userId string) (*gen.OrderResponse, error) {
	createdAtProto := timestamppb.New(order.CreatedAt)
	err := createdAtProto.CheckValid(); 
	if err != nil {
		return nil, err
	}

	grpcCart := CartToProto(order.OrderProducts)
	if grpcCart == nil {
		return nil, fmt.Errorf("invalid cart data")
	}

	return &gen.OrderResponse{
		Id:                order.ID.String(),
		UserId:            userId,
		Status:            order.Status,
		Address:           order.Address,
		OrderProducts:     grpcCart,
		ApartmentOrOffice: order.ApartmentOrOffice,
		Intercom:          order.Intercom,
		Entrance:          order.Entrance,
		Floor:             order.Floor,
		CourierComment:    order.CourierComment,
		LeaveAtDoor:       order.LeaveAtDoor,
		CreatedAt:         createdAtProto,
		FinalPrice:        order.FinalPrice,
	}, nil
}

func ProtoToOrder(grpcOrder *gen.OrderResponse) (models.Order, error) {
	if grpcOrder == nil {
		return models.Order{}, fmt.Errorf("nil order response")
	}

	orderID, err := uuid.FromString(grpcOrder.Id)
	if err != nil {
		return models.Order{}, fmt.Errorf("invalid order ID: %v", err)
	}

	userID, err := uuid.FromString(grpcOrder.UserId)
	if err != nil {
		return models.Order{}, fmt.Errorf("invalid user ID: %v", err)
	}

	var createdAt time.Time
	if grpcOrder.CreatedAt != nil {
		if err := grpcOrder.CreatedAt.CheckValid(); err != nil {
			return models.Order{}, fmt.Errorf("invalid timestamp: %v", err)
		}
		createdAt = grpcOrder.CreatedAt.AsTime()
	}

	cart, err := ProtoToCart(grpcOrder.OrderProducts)
	if err != nil {
		return models.Order{}, fmt.Errorf("failed to convert cart: %v", err)
	}

	return models.Order{
		ID:                orderID,
		UserID:            userID.String(),
		Status:            grpcOrder.Status,
		Address:           grpcOrder.Address,
		OrderProducts:     cart,
		ApartmentOrOffice: grpcOrder.ApartmentOrOffice,
		Intercom:          grpcOrder.Intercom,
		Entrance:          grpcOrder.Entrance,
		Floor:             grpcOrder.Floor,
		CourierComment:    grpcOrder.CourierComment,
		LeaveAtDoor:       grpcOrder.LeaveAtDoor,
		CreatedAt:         createdAt,
		FinalPrice:        grpcOrder.FinalPrice,
	}, nil
}
