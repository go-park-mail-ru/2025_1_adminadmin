package grpc

import (
	"context"
	"os"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen"
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
	//logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	req := models.SignInReq{
		Login:    in.Login,
		Password: in.Password,
	}
	req.Sanitize()

	user, token, csrfToken, err := h.uc.SignIn(ctx, req)

	if err != nil {
		switch err {
		case auth.ErrInvalidLogin, auth.ErrUserNotFound:
			//log.LogHandlerError(logger, err, http.StatusBadRequest)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		case auth.ErrInvalidCredentials:
			//log.LogHandlerError(logger, err, http.StatusUnauthorized)
			return nil, status.Errorf(codes.Unauthenticated, err.Error())
		default:
			//log.LogHandlerError(logger, fmt.Errorf("Неизвестная ошибка: %w", err), http.StatusInternalServerError)
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	return &gen.UserResponse{
		Login:  user.Login,
		Token:  token,
		Token2: csrfToken,
	}, nil
}

func (h *AuthHandler) SignUp(context.Context, *gen.SignUpRequest) (*gen.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (h *AuthHandler) Check(context.Context, *emptypb.Empty) (*gen.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (h *AuthHandler) UpdateUser(context.Context, *gen.UpdateUserRequest) (*gen.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (h *AuthHandler) UpdateUserPic(context.Context, *gen.UpdateUserPicRequest) (*gen.UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserPic not implemented")
}
func (h *AuthHandler) GetUserAddresses(context.Context, *emptypb.Empty) (*gen.AddressListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAddresses not implemented")
}
func (h *AuthHandler) DeleteAddress(context.Context, *gen.DeleteAddressRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAddress not implemented")
}
func (h *AuthHandler) AddAddress(context.Context, *gen.Address) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAddress not implemented")
}
