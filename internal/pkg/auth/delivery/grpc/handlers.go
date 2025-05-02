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
	//logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

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
			//log.LogHandlerError(logger, fmt.Errorf("Неправильный логин или пароль: %w", err), http.StatusBadRequest)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		case auth.ErrInvalidName, auth.ErrInvalidPhone, auth.ErrCreatingUser:
			//log.LogHandlerError(logger, err, http.StatusBadRequest)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		default:
			//log.LogHandlerError(logger, fmt.Errorf("Неизвестная ошибка: %w", err), http.StatusInternalServerError)
			return nil, status.Errorf(codes.Internal, err.Error())
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
	//logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	user, err := h.uc.Check(ctx, in.Login)

	if err != nil {
		if errors.Is(err, auth.ErrUserNotFound) {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
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
	//logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

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
			//log.LogHandlerError(logger, err, http.StatusBadRequest)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		default:
			//log.LogHandlerError(logger, fmt.Errorf("Ошибка обновления данных пользователя: %w", err), http.StatusInternalServerError)
			return nil, status.Errorf(codes.Internal, err.Error())
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
	//logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	reader := bytes.NewReader(in.UserPic)
	user, err := h.uc.UpdateUserPic(ctx, in.Login, reader, in.FileExtension)
	if err != nil {
		switch err {
		case auth.ErrUserNotFound:
			//log.LogHandlerError(logger, err, http.StatusNotFound)
			return nil, status.Errorf(codes.NotFound, err.Error())
		case auth.ErrBasePath:
			//log.LogHandlerError(logger, err, http.StatusInternalServerError)
			return nil, status.Errorf(codes.Internal, err.Error())
		case auth.ErrFileCreation, auth.ErrFileSaving, auth.ErrFileDeletion:
			//log.LogHandlerError(logger, fmt.Errorf("Ошибка при работе с файлом: %w", err), http.StatusInternalServerError)
			return nil, status.Errorf(codes.Internal, err.Error())
		default:
			//log.LogHandlerError(logger, fmt.Errorf("Ошибка при обновлении аватарки: %w", err), http.StatusInternalServerError)
			return nil, status.Errorf(codes.Internal, err.Error())
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
	//logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))

	addresses, err := h.uc.GetUserAddresses(ctx, in.Login)
	if err != nil {
		//log.LogHandlerError(logger, fmt.Errorf("Ошибка на уровне ниже (usecase): %w", err), http.StatusInternalServerError)
		return nil, status.Errorf(codes.Internal, err.Error())
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
	//logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))
	parsedUUID, err := uuid.FromString(in.Id)
	if err != nil {
		//log.LogHandlerError(logger, fmt.Errorf("некорректный id адреса: %w", err), http.StatusUnauthorized)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	address := models.Address{
		Id: parsedUUID,
	}

	err = h.uc.DeleteAddress(ctx, address.Id)
	if err != nil {
		//log.LogHandlerError(logger, fmt.Errorf("Ошибка на уровне ниже (usecase): %w", err), http.StatusInternalServerError)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}

func (h *AuthHandler) AddAddress(ctx context.Context, in *gen.Address) (*emptypb.Empty, error) {
	//logger := log.GetLoggerFromContext(r.Context()).With(slog.String("func", log.GetFuncName()))
	parsedUUIDa, err := uuid.FromString(in.Id)
	if err != nil {
		//log.LogHandlerError(logger, fmt.Errorf("некорректный id адреса: %w", err), http.StatusUnauthorized)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	parsedUUIDu, err := uuid.FromString(in.UserId)
	if err != nil {
		//log.LogHandlerError(logger, fmt.Errorf("некорректный id адреса: %w", err), http.StatusUnauthorized)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	address := models.Address{
		Id:      parsedUUIDa,
		Address: in.Address,
		UserId:  parsedUUIDu,
	}

	err = h.uc.AddAddress(ctx, address)
	if err != nil {
		//log.LogHandlerError(logger, fmt.Errorf("Ошибка на уровне ниже (usecase): %w", err), http.StatusInternalServerError)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
