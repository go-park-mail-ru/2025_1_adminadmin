package grpc

import (
	"context"
	"os"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/cart/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/utils/converter"
	"github.com/satori/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CartHandler struct {
	uc     cart.CartUsecase
	secret string
	gen.CartServiceServer
}

func CreateCartHandler(uc cart.CartUsecase) *CartHandler {
	return &CartHandler{uc: uc, secret: os.Getenv("JWT_SECRET")}
}

func (h *CartHandler) GetCart(ctx context.Context, in *gen.GetCartRequest) (*gen.CartResponse, error) {
	cart, err, full_cart := h.uc.GetCart(ctx, in.Login)

	if err != nil {
		return &gen.CartResponse{}, status.Errorf(codes.Internal, "ошибка получения корзины")
	}

	return &gen.CartResponse{
		RestaurantId:   cart.Id.String(),
		RestaurantName: cart.Name,
		Products:       converter.CartItemsToProto(cart.CartItems),
		FullCart:       full_cart,
	}, nil
}

func (h *CartHandler) UpdateItemQuantity(ctx context.Context, in *gen.UpdateQuantityRequest) (*emptypb.Empty, error) {
	err := h.uc.UpdateItemQuantity(ctx, in.Login, in.ProductId, in.RestaurantId, int(in.Quantity))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (h *CartHandler) ClearCart(ctx context.Context, in *gen.ClearCartRequest) (*emptypb.Empty, error) {
	err := h.uc.ClearCart(ctx, in.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func (h *CartHandler) CreateOrder(ctx context.Context, in *gen.CreateOrderRequest) (*gen.OrderResponse, error) {
	req := models.OrderInReq{
		Status:            in.Status,
		Address:           in.Address,
		ApartmentOrOffice: in.ApartmentOrOffice,
		Intercom:          in.Intercom,
		Entrance:          in.Entrance,
		Floor:             in.Floor,
		CourierComment:    in.CourierComment,
		LeaveAtDoor:       in.LeaveAtDoor,
		FinalPrice:        in.FinalPrice,
	}

	restId, err := uuid.FromString(in.Cart.RestaurantId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid restaurant ID: %v", err)
	}

	cartItems, err := converter.ProtoToCartItems(in.Cart.Products)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to convert cart items: %v", err)
	}

	cart := models.Cart{
		Id:        restId,
		Name:      in.Cart.RestaurantName,
		CartItems: cartItems,
	}

	order, err := h.uc.CreateOrder(ctx, in.Login, req, cart)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return converter.OrderToProto(order, in.Login)
}

func (h *CartHandler) GetOrders(ctx context.Context, in *gen.GetOrdersRequest) (*gen.OrderListResponse, error) {
	userId, err := uuid.FromString(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid restaurant ID: %v", err)
	}
	orders, err := h.uc.GetOrders(ctx, userId, int(in.Count), int(in.Offset))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	protoOrders := make([]*gen.OrderResponse, 0, len(orders))
	for _, order := range orders {
		protoOrder, err := converter.OrderToProto(order, in.UserId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "order conversion failed: %v", err)
		}
		protoOrders = append(protoOrders, protoOrder)
	}

	return &gen.OrderListResponse{
		Orders: protoOrders,
	}, nil
}

func (h *CartHandler) GetOrderById(ctx context.Context, in *gen.GetOrderByIdRequest) (*gen.OrderResponse, error) {
	userId, err := uuid.FromString(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user ID: %v", err)
	}
	orderId, err := uuid.FromString(in.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid order ID: %v", err)
	}
	order, err := h.uc.GetOrderById(ctx, orderId, userId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get order: %v", err)
	}

	return converter.OrderToProto(order, in.UserId)
}

func (h *CartHandler) UpdateOrderStatus(ctx context.Context, in *gen.UpdateOrderStatusRequest) (*emptypb.Empty, error) {
	orderId, err := uuid.FromString(in.OrderId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid order ID: %v", err)
	}
	err = h.uc.UpdateOrderStatus(ctx, orderId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update order status: %v", err)
	}
	return &emptypb.Empty{}, nil
}
