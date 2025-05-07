package grpc

import (
	"bytes"
	"context"
	"errors"
	"os"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen"
	"github.com/satori/uuid"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type AuthHandler struct {
	uc     auth.AuthUsecase
	secret string
	gen.AuthServiceServer
}

func CreateAuthHandler(uc auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{uc: uc, secret: os.Getenv("JWT_SECRET")}
}

func (h *AuthHandler) SignIn(ctx context.Context, in *gen.SignInRequest) (*gen.UserResponse, error) {
	req := models.SignInReq{
		Login:    in.Login,
		Password: in.Password,
	}
	req.Sanitize()

	user, token, csrfToken, err := h.uc.SignIn(ctx, req)

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin, auth.ErrUserNotFound:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		case auth.ErrInvalidCredentials:
			return nil, status.Errorf(codes.Unauthenticated, "%v", err)
		default:
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
	}

	return &gen.UserResponse{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          user.Id.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
		Token:       token,
		CsrfToken:   csrfToken,
	}, nil
}

func (h *AuthHandler) SignUp(ctx context.Context, in *gen.SignUpRequest) (*gen.UserResponse, error) {
	req := models.SignUpReq{
		Login:       in.Login,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		PhoneNumber: in.PhoneNumber,
		Password:    in.Password,
	}
	req.Sanitize()

	user, token, csrfToken, err := h.uc.SignUp(ctx, req)

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin, auth.ErrInvalidPassword:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		case auth.ErrInvalidName, auth.ErrInvalidPhone, auth.ErrCreatingUser:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		default:
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
	}
	return &gen.UserResponse{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          user.Id.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
		Token:       token,
		CsrfToken:   csrfToken,
	}, nil

}

func (h *AuthHandler) Check(ctx context.Context, in *gen.CheckRequest) (*gen.UserResponse, error) {
	user, err := h.uc.Check(ctx, in.Login)

	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		}
	}

	return &gen.UserResponse{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          user.Id.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}, nil
}

func (h *AuthHandler) UpdateUser(ctx context.Context, in *gen.UpdateUserRequest) (*gen.UserResponse, error) {
	req := models.UpdateUserReq{
		Description: in.Description,
		FirstName:   in.FirstName,
		LastName:    in.LastName,
		PhoneNumber: in.PhoneNumber,
		Password:    in.Password,
	}
	req.Sanitize()

	user, err := h.uc.UpdateUser(ctx, in.Login, req)
	if err != nil {
		switch err {
		case auth.ErrInvalidPassword, auth.ErrInvalidName, auth.ErrInvalidPhone, auth.ErrSamePassword:
			return nil, status.Errorf(codes.InvalidArgument, "%v", err)
		default:
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
	}
	return &gen.UserResponse{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          user.Id.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}, nil

}

func (h *AuthHandler) UpdateUserPic(ctx context.Context, in *gen.UpdateUserPicRequest) (*gen.UserResponse, error) {
	reader := bytes.NewReader(in.UserPic)
	user, err := h.uc.UpdateUserPic(ctx, in.Login, reader, in.FileExtension)
	if err != nil {
		switch err {
		case auth.ErrUserNotFound:
			return nil, status.Errorf(codes.NotFound, "%v", err)
		case auth.ErrBasePath:
			return nil, status.Errorf(codes.Internal, "%v", err)
		case auth.ErrFileCreation, auth.ErrFileSaving, auth.ErrFileDeletion:
			return nil, status.Errorf(codes.Internal, "%v", err)
		default:
			return nil, status.Errorf(codes.Internal, "%v", err)
		}
	}
	return &gen.UserResponse{
		Login:       user.Login,
		PhoneNumber: user.PhoneNumber,
		Id:          user.Id.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Description: user.Description,
		UserPic:     user.UserPic,
	}, nil

}

func (h *AuthHandler) GetUserAddresses(ctx context.Context, in *gen.AddressRequest) (*gen.AddressListResponse, error) {
	addresses, err := h.uc.GetUserAddresses(ctx, in.Login)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}

	var grpcAddresses []*gen.Address
	for _, addr := range addresses {
		grpcAddresses = append(grpcAddresses, &gen.Address{
			Id:      addr.Id.String(),
			Address: addr.Address,
			UserId:  addr.UserId.String(),
		})
	}

	return &gen.AddressListResponse{
		Addresses: grpcAddresses,
	}, nil
}

func (h *AuthHandler) DeleteAddress(ctx context.Context, in *gen.DeleteAddressRequest) (*emptypb.Empty, error) {
	parsedUUID, err := uuid.FromString(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	address := models.Address{
		Id: parsedUUID,
	}

	err = h.uc.DeleteAddress(ctx, address.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return &emptypb.Empty{}, nil
}

func (h *AuthHandler) AddAddress(ctx context.Context, in *gen.Address) (*emptypb.Empty, error) {
	parsedUUIDa, err := uuid.FromString(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	parsedUUIDu, err := uuid.FromString(in.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	address := models.Address{
		Id:      parsedUUIDa,
		Address: in.Address,
		UserId:  parsedUUIDu,
	}

	err = h.uc.AddAddress(ctx, address)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "%v", err)
	}
	return &emptypb.Empty{}, nil
}
