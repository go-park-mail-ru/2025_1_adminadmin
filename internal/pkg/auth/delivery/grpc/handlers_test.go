package grpc

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/models"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2025_1_adminadmin/internal/pkg/auth/mocks"
	"github.com/satori/uuid"
	"github.com/golang/mock/gomock"
	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func TestAuthHandler_SignIn(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    type fields struct {
        uc *mocks.MockAuthUsecase
    }
    type args struct {
        ctx context.Context
        in  *gen.SignInRequest
    }

    testUserID := uuid.NewV4()
    testCases := []struct {
        name        string
        setup       func(f *fields)
        args        args
        want        *gen.UserResponse
        wantErr     bool
        wantErrCode codes.Code
    }{
        {
            name: "Success",
            setup: func(f *fields) {
                // Исправляем возвращаемый тип - убираем указатель
                f.uc.EXPECT().SignIn(gomock.Any(), models.SignInReq{
                    Login:    "test@example.com",
                    Password: "password123",
                }).Return(
                    models.User{ // Убрали & - возвращаем значение, а не указатель
                        Id:          testUserID,
                        Login:       "test@example.com",
                        PhoneNumber: "+1234567890",
                        FirstName:   "John",
                        LastName:    "Doe",
                        Description: "Test user",
                        UserPic:     "avatar.jpg",
                    },
                    "access_token",
                    "csrf_token",
                    nil,
                )
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignInRequest{
                    Login:    "test@example.com",
                    Password: "password123",
                },
            },
            want: &gen.UserResponse{
                Login:       "test@example.com",
                PhoneNumber: "+1234567890",
                Id:          testUserID.String(),
                FirstName:   "John",
                LastName:    "Doe",
                Description: "Test user",
                UserPic:     "avatar.jpg",
                Token:       "access_token",
                CsrfToken:   "csrf_token",
            },
            wantErr: false,
        },
        // Остальные тестовые случаи остаются без изменений
        {
            name: "Invalid login format",
            setup: func(f *fields) {
                f.uc.EXPECT().SignIn(gomock.Any(), models.SignInReq{
                    Login:    "invalid",
                    Password: "password123",
                }).Return(models.User{}, "", "", auth.ErrInvalidLogin) // Возвращаем пустую структуру вместо nil
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignInRequest{
                    Login:    "invalid",
                    Password: "password123",
                },
            },
            want:        nil,
            wantErr:     true,
            wantErrCode: codes.InvalidArgument,
        },
        {
            name: "User not found",
            setup: func(f *fields) {
                f.uc.EXPECT().SignIn(gomock.Any(), models.SignInReq{
                    Login:    "notfound@example.com",
                    Password: "password123",
                }).Return(models.User{}, "", "", auth.ErrUserNotFound)
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignInRequest{
                    Login:    "notfound@example.com",
                    Password: "password123",
                },
            },
            want:        nil,
            wantErr:     true,
            wantErrCode: codes.InvalidArgument,
        },
        {
            name: "Invalid credentials",
            setup: func(f *fields) {
                f.uc.EXPECT().SignIn(gomock.Any(), models.SignInReq{
                    Login:    "test@example.com",
                    Password: "wrongpassword",
                }).Return(models.User{}, "", "", auth.ErrInvalidCredentials)
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignInRequest{
                    Login:    "test@example.com",
                    Password: "wrongpassword",
                },
            },
            want:        nil,
            wantErr:     true,
            wantErrCode: codes.Unauthenticated,
        },
        {
            name: "Internal server error",
            setup: func(f *fields) {
                f.uc.EXPECT().SignIn(gomock.Any(), models.SignInReq{
                    Login:    "test@example.com",
                    Password: "password123",
                }).Return(models.User{}, "", "", errors.New("database error"))
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignInRequest{
                    Login:    "test@example.com",
                    Password: "password123",
                },
            },
            want:        nil,
            wantErr:     true,
            wantErrCode: codes.Internal,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            fields := fields{
                uc: mocks.NewMockAuthUsecase(ctrl),
            }
            if tc.setup != nil {
                tc.setup(&fields)
            }

            h := &AuthHandler{
                uc: fields.uc,
            }

            got, err := h.SignIn(tc.args.ctx, tc.args.in)
            if (err != nil) != tc.wantErr {
                t.Errorf("SignIn() error = %v, wantErr %v", err, tc.wantErr)
                return
            }

            if tc.wantErr {
                st, ok := status.FromError(err)
                if !ok {
                    t.Errorf("SignIn() expected gRPC status error")
                    return
                }
                if st.Code() != tc.wantErrCode {
                    t.Errorf("SignIn() error code = %v, want %v", st.Code(), tc.wantErrCode)
                }
                return
            }

            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("SignIn() = %v, want %v", got, tc.want)
            }
        })
    }
}

func TestAuthHandler_SignUp(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    type fields struct {
        uc *mocks.MockAuthUsecase
    }
    type args struct {
        ctx context.Context
        in  *gen.SignUpRequest
    }

    testUserID := uuid.NewV4()
    testCases := []struct {
        name        string
        setup       func(f *fields)
        args        args
        want        *gen.UserResponse
        wantErr     bool
        wantErrCode codes.Code
    }{
        {
            name: "Success",
            setup: func(f *fields) {
                f.uc.EXPECT().SignUp(gomock.Any(), models.SignUpReq{
                    Login:       "test@example.com",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "password123",
                }).Return(
                    models.User{
                        Id:          testUserID,
                        Login:       "test@example.com",
                        PhoneNumber: "+1234567890",
                        FirstName:   "John",
                        LastName:    "Doe",
                        Description: "Test user",
                        UserPic:     "avatar.jpg",
                    },
                    "access_token",
                    "csrf_token",
                    nil,
                )
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignUpRequest{
                    Login:       "test@example.com",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "password123",
                },
            },
            want: &gen.UserResponse{
                Login:       "test@example.com",
                PhoneNumber: "+1234567890",
                Id:          testUserID.String(),
                FirstName:   "John",
                LastName:    "Doe",
                Description: "Test user",
                UserPic:     "avatar.jpg",
                Token:       "access_token",
                CsrfToken:   "csrf_token",
            },
            wantErr: false,
        },
        {
            name: "Invalid login",
            setup: func(f *fields) {
                f.uc.EXPECT().SignUp(gomock.Any(), models.SignUpReq{
                    Login:       "invalid",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "password123",
                }).Return(models.User{}, "", "", auth.ErrInvalidLogin)
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignUpRequest{
                    Login:       "invalid",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "password123",
                },
            },
            want:        nil,
            wantErr:     true,
            wantErrCode: codes.InvalidArgument,
        },
        {
            name: "Invalid password",
            setup: func(f *fields) {
                f.uc.EXPECT().SignUp(gomock.Any(), models.SignUpReq{
                    Login:       "test@example.com",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "short",
                }).Return(models.User{}, "", "", auth.ErrInvalidPassword)
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignUpRequest{
                    Login:       "test@example.com",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "short",
                },
            },
            want:        nil,
            wantErr:     true,
            wantErrCode: codes.InvalidArgument,
        },
        {
            name: "Invalid name",
            setup: func(f *fields) {
                f.uc.EXPECT().SignUp(gomock.Any(), models.SignUpReq{
                    Login:       "test@example.com",
                    FirstName:   "",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "password123",
                }).Return(models.User{}, "", "", auth.ErrInvalidName)
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignUpRequest{
                    Login:       "test@example.com",
                    FirstName:   "",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "password123",
                },
            },
            want:        nil,
            wantErr:     true,
            wantErrCode: codes.InvalidArgument,
        },
        {
            name: "Invalid phone",
            setup: func(f *fields) {
                f.uc.EXPECT().SignUp(gomock.Any(), models.SignUpReq{
                    Login:       "test@example.com",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "invalid",
                    Password:    "password123",
                }).Return(models.User{}, "", "", auth.ErrInvalidPhone)
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignUpRequest{
                    Login:       "test@example.com",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "invalid",
                    Password:    "password123",
                },
            },
            want:        nil,
            wantErr:     true,
            wantErrCode: codes.InvalidArgument,
        },
        {
            name: "User creation error",
            setup: func(f *fields) {
                f.uc.EXPECT().SignUp(gomock.Any(), models.SignUpReq{
                    Login:       "test@example.com",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "password123",
                }).Return(models.User{}, "", "", auth.ErrCreatingUser)
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignUpRequest{
                    Login:       "test@example.com",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "password123",
                },
            },
            want:        nil,
            wantErr:     true,
            wantErrCode: codes.InvalidArgument,
        },
        {
            name: "Internal server error",
            setup: func(f *fields) {
                f.uc.EXPECT().SignUp(gomock.Any(), models.SignUpReq{
                    Login:       "test@example.com",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "password123",
                }).Return(models.User{}, "", "", errors.New("database error"))
            },
            args: args{
                ctx: context.Background(),
                in: &gen.SignUpRequest{
                    Login:       "test@example.com",
                    FirstName:   "John",
                    LastName:    "Doe",
                    PhoneNumber: "+1234567890",
                    Password:    "password123",
                },
            },
            want:        nil,
            wantErr:     true,
            wantErrCode: codes.Internal,
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            fields := fields{
                uc: mocks.NewMockAuthUsecase(ctrl),
            }
            if tc.setup != nil {
                tc.setup(&fields)
            }

            h := &AuthHandler{
                uc: fields.uc,
            }

            got, err := h.SignUp(tc.args.ctx, tc.args.in)
            if (err != nil) != tc.wantErr {
                t.Errorf("SignUp() error = %v, wantErr %v", err, tc.wantErr)
                return
            }

            if tc.wantErr {
                st, ok := status.FromError(err)
                if !ok {
                    t.Errorf("SignUp() expected gRPC status error")
                    return
                }
                if st.Code() != tc.wantErrCode {
                    t.Errorf("SignUp() error code = %v, want %v", st.Code(), tc.wantErrCode)
                }
                return
            }

            if !reflect.DeepEqual(got, tc.want) {
                t.Errorf("SignUp() = %v, want %v", got, tc.want)
            }
        })
    }
}

func TestAuthHandler_Check(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	validLogin := "johndoe"
	user := models.User{
		Login:       validLogin,
		PhoneNumber: "1234567890",
		Id:          uuid.NewV4(),
		FirstName:   "John",
		LastName:    "Doe",
		Description: "Test user",
		UserPic:     "pic.jpg",
	}

	type fields struct {
		uc *mocks.MockAuthUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.CheckRequest
	}
	tests := []struct {
		name        string
		setup       func(f *fields)
		args        args
		want        *gen.UserResponse
		wantErr     bool
		wantErrCode codes.Code
	}{
		{
			name: "Success",
			setup: func(f *fields) {
				f.uc.EXPECT().
					Check(gomock.Any(), validLogin).
					Return(user, nil)
			},
			args: args{
				ctx: context.Background(),
				in:  &gen.CheckRequest{Login: validLogin},
			},
			want: &gen.UserResponse{
				Login:       user.Login,
				PhoneNumber: user.PhoneNumber,
				Id:          user.Id.String(),
				FirstName:   user.FirstName,
				LastName:    user.LastName,
				Description: user.Description,
				UserPic:     user.UserPic,
			},
			wantErr: false,
		},
		{
			name: "User not found",
			setup: func(f *fields) {
				f.uc.EXPECT().
					Check(gomock.Any(), validLogin).
					Return(models.User{}, auth.ErrUserNotFound)
			},
			args: args{
				ctx: context.Background(),
				in:  &gen.CheckRequest{Login: validLogin},
			},
			wantErr:     true,
			wantErrCode: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fields{
				uc: mocks.NewMockAuthUsecase(ctrl),
			}
			if tt.setup != nil {
				tt.setup(f)
			}
			h := &AuthHandler{uc: f.uc}
			got, err := h.Check(tt.args.ctx, tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Fatalf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				st, ok := status.FromError(err)
				if !ok {
					t.Fatalf("expected gRPC status error, got: %v", err)
				}
				if st.Code() != tt.wantErrCode {
					t.Errorf("expected gRPC error code %v, got %v", tt.wantErrCode, st.Code())
				}
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Check() got = %v, want = %v", got, tt.want)
			}
		})
	}
}


func TestAuthHandler_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		uc *mocks.MockAuthUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.UpdateUserRequest
	}

	testUserID := uuid.NewV4()
	testCases := []struct {
		name        string
		setup       func(f *fields)
		args        args
		want        *gen.UserResponse
		wantErr     bool
		wantErrCode codes.Code
	}{
		{
			name: "Success",
			setup: func(f *fields) {
				f.uc.EXPECT().UpdateUser(gomock.Any(), "test@example.com", models.UpdateUserReq{
					Description: "Updated description",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "+1234567890",
					Password:    "newPassword123",
				}).Return(models.User{
					Id:          testUserID,
					Login:       "test@example.com",
					PhoneNumber: "+1234567890",
					FirstName:   "John",
					LastName:    "Doe",
					Description: "Updated description",
					UserPic:     "updated.jpg",
				}, nil)
			},
			args: args{
				ctx: context.Background(),
				in: &gen.UpdateUserRequest{
					Login:       "test@example.com",
					Description: "Updated description",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "+1234567890",
					Password:    "newPassword123",
				},
			},
			want: &gen.UserResponse{
				Login:       "test@example.com",
				PhoneNumber: "+1234567890",
				Id:          testUserID.String(),
				FirstName:   "John",
				LastName:    "Doe",
				Description: "Updated description",
				UserPic:     "updated.jpg",
			},
			wantErr: false,
		},
		{
			name: "Invalid password",
			setup: func(f *fields) {
				f.uc.EXPECT().UpdateUser(gomock.Any(), "test@example.com", models.UpdateUserReq{
					Description: "",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "+1234567890",
					Password:    "123",
				}).Return(models.User{}, auth.ErrInvalidPassword)
			},
			args: args{
				ctx: context.Background(),
				in: &gen.UpdateUserRequest{
					Login:       "test@example.com",
					Description: "",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "+1234567890",
					Password:    "123",
				},
			},
			want:        nil,
			wantErr:     true,
			wantErrCode: codes.InvalidArgument,
		},
		{
			name: "Invalid phone",
			setup: func(f *fields) {
				f.uc.EXPECT().UpdateUser(gomock.Any(), "test@example.com", models.UpdateUserReq{
					Description: "",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "invalid-phone",
					Password:    "password123",
				}).Return(models.User{}, auth.ErrInvalidPhone)
			},
			args: args{
				ctx: context.Background(),
				in: &gen.UpdateUserRequest{
					Login:       "test@example.com",
					Description: "",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "invalid-phone",
					Password:    "password123",
				},
			},
			want:        nil,
			wantErr:     true,
			wantErrCode: codes.InvalidArgument,
		},
		{
			name: "Same password error",
			setup: func(f *fields) {
				f.uc.EXPECT().UpdateUser(gomock.Any(), "test@example.com", models.UpdateUserReq{
					Description: "",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "+1234567890",
					Password:    "oldPassword",
				}).Return(models.User{}, auth.ErrSamePassword)
			},
			args: args{
				ctx: context.Background(),
				in: &gen.UpdateUserRequest{
					Login:       "test@example.com",
					Description: "",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "+1234567890",
					Password:    "oldPassword",
				},
			},
			want:        nil,
			wantErr:     true,
			wantErrCode: codes.InvalidArgument,
		},
		{
			name: "Internal error",
			setup: func(f *fields) {
				f.uc.EXPECT().UpdateUser(gomock.Any(), "test@example.com", models.UpdateUserReq{
					Description: "",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "+1234567890",
					Password:    "password123",
				}).Return(models.User{}, errors.New("db connection failed"))
			},
			args: args{
				ctx: context.Background(),
				in: &gen.UpdateUserRequest{
					Login:       "test@example.com",
					Description: "",
					FirstName:   "John",
					LastName:    "Doe",
					PhoneNumber: "+1234567890",
					Password:    "password123",
				},
			},
			want:        nil,
			wantErr:     true,
			wantErrCode: codes.Internal,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fields := fields{
				uc: mocks.NewMockAuthUsecase(ctrl),
			}
			if tc.setup != nil {
				tc.setup(&fields)
			}

			h := &AuthHandler{
				uc: fields.uc,
			}

			got, err := h.UpdateUser(tc.args.ctx, tc.args.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if tc.wantErr {
				st, ok := status.FromError(err)
				if !ok {
					t.Errorf("UpdateUser() expected gRPC status error")
					return
				}
				if st.Code() != tc.wantErrCode {
					t.Errorf("UpdateUser() error code = %v, want %v", st.Code(), tc.wantErrCode)
				}
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("UpdateUser() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAuthHandler_UpdateUserPic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testUserID := uuid.NewV4()
	picData := []byte("fake-image-data")

	type fields struct {
		uc *mocks.MockAuthUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.UpdateUserPicRequest
	}
	tests := []struct {
		name        string
		setup       func(f *fields)
		args        args
		want        *gen.UserResponse
		wantErr     bool
		wantErrCode codes.Code
	}{
		{
			name: "Success",
			setup: func(f *fields) {
				f.uc.EXPECT().UpdateUserPic(gomock.Any(), "testuser", gomock.Any(), ".jpg").
					Return(models.User{
						Id:          testUserID,
						Login:       "testuser",
						PhoneNumber: "+123456789",
						FirstName:   "John",
						LastName:    "Doe",
						Description: "desc",
						UserPic:     "path/to/pic.jpg",
					}, nil)
			},
			args: args{
				ctx: context.Background(),
				in: &gen.UpdateUserPicRequest{
					Login:         "testuser",
					UserPic:       picData,
					FileExtension: ".jpg",
				},
			},
			want: &gen.UserResponse{
				Login:       "testuser",
				PhoneNumber: "+123456789",
				Id:          testUserID.String(),
				FirstName:   "John",
				LastName:    "Doe",
				Description: "desc",
				UserPic:     "path/to/pic.jpg",
			},
			wantErr: false,
		},
		{
			name: "User not found",
			setup: func(f *fields) {
				f.uc.EXPECT().UpdateUserPic(gomock.Any(), "ghost", gomock.Any(), ".jpg").
					Return(models.User{}, auth.ErrUserNotFound)
			},
			args: args{
				ctx: context.Background(),
				in: &gen.UpdateUserPicRequest{
					Login:         "ghost",
					UserPic:       picData,
					FileExtension: ".jpg",
				},
			},
			wantErr:     true,
			wantErrCode: codes.NotFound,
		},
		{
			name: "Base path error",
			setup: func(f *fields) {
				f.uc.EXPECT().UpdateUserPic(gomock.Any(), "testuser", gomock.Any(), ".jpg").
					Return(models.User{}, auth.ErrBasePath)
			},
			args: args{
				ctx: context.Background(),
				in: &gen.UpdateUserPicRequest{
					Login:         "testuser",
					UserPic:       picData,
					FileExtension: ".jpg",
				},
			},
			wantErr:     true,
			wantErrCode: codes.Internal,
		},
		{
			name: "File saving error",
			setup: func(f *fields) {
				f.uc.EXPECT().UpdateUserPic(gomock.Any(), "testuser", gomock.Any(), ".jpg").
					Return(models.User{}, auth.ErrFileSaving)
			},
			args: args{
				ctx: context.Background(),
				in: &gen.UpdateUserPicRequest{
					Login:         "testuser",
					UserPic:       picData,
					FileExtension: ".jpg",
				},
			},
			wantErr:     true,
			wantErrCode: codes.Internal,
		},
		{
			name: "Unexpected error",
			setup: func(f *fields) {
				f.uc.EXPECT().UpdateUserPic(gomock.Any(), "testuser", gomock.Any(), ".jpg").
					Return(models.User{}, errors.New("unknown error"))
			},
			args: args{
				ctx: context.Background(),
				in: &gen.UpdateUserPicRequest{
					Login:         "testuser",
					UserPic:       picData,
					FileExtension: ".jpg",
				},
			},
			wantErr:     true,
			wantErrCode: codes.Internal,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := &fields{
				uc: mocks.NewMockAuthUsecase(ctrl),
			}
			if tc.setup != nil {
				tc.setup(f)
			}

			h := &AuthHandler{
				uc: f.uc,
			}

			got, err := h.UpdateUserPic(tc.args.ctx, tc.args.in)
			if (err != nil) != tc.wantErr {
				t.Errorf("UpdateUserPic() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if tc.wantErr {
				st, ok := status.FromError(err)
				if !ok {
					t.Errorf("Expected gRPC status error, got: %v", err)
					return
				}
				if st.Code() != tc.wantErrCode {
					t.Errorf("Expected error code %v, got %v", tc.wantErrCode, st.Code())
				}
				return
			}

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestAuthHandler_GetUserAddresses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		uc *mocks.MockAuthUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.AddressRequest
	}
	tests := []struct {
		name        string
		setup       func(f *fields)
		args        args
		want        *gen.AddressListResponse
		wantErr     bool
		wantErrCode codes.Code
	}{
		{
			name: "Success",
			setup: func(f *fields) {
				userID := uuid.NewV4()
				addr1 := models.Address{
					Id:      uuid.NewV4(),
					UserId:  userID,
					Address: "123 Main St",
				}
				addr2 := models.Address{
					Id:      uuid.NewV4(),
					UserId:  userID,
					Address: "456 Side Ave",
				}
				f.uc.EXPECT().GetUserAddresses(gomock.Any(), "testuser").
					Return([]models.Address{addr1, addr2}, nil)
			},
			args: args{
				ctx: context.Background(),
				in:  &gen.AddressRequest{Login: "testuser"},
			},
			want: &gen.AddressListResponse{
				Addresses: []*gen.Address{
					{
						Id:      "", 
						Address: "123 Main St",
						UserId:  "",
					},
					{
						Id:      "",
						Address: "456 Side Ave",
						UserId:  "",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Usecase error",
			setup: func(f *fields) {
				f.uc.EXPECT().GetUserAddresses(gomock.Any(), "broken").
					Return(nil, errors.New("db error"))
			},
			args: args{
				ctx: context.Background(),
				in:  &gen.AddressRequest{Login: "broken"},
			},
			wantErr:     true,
			wantErrCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fields{
				uc: mocks.NewMockAuthUsecase(ctrl),
			}
			if tt.setup != nil {
				tt.setup(f)
			}

			h := &AuthHandler{uc: f.uc}
			got, err := h.GetUserAddresses(tt.args.ctx, tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				st, ok := status.FromError(err)
				if !ok {
					t.Fatalf("Expected gRPC status error, got: %v", err)
				}
				if st.Code() != tt.wantErrCode {
					t.Errorf("Expected gRPC error code %v, got %v", tt.wantErrCode, st.Code())
				}
				return
			}

			if len(got.Addresses) != len(tt.want.Addresses) {
				t.Fatalf("Expected %d addresses, got %d", len(tt.want.Addresses), len(got.Addresses))
			}
			for i := range got.Addresses {
				if got.Addresses[i].Address != tt.want.Addresses[i].Address {
					t.Errorf("Address mismatch at index %d: got %v, want %v", i, got.Addresses[i].Address, tt.want.Addresses[i].Address)
				}
			}
		})
	}
}

func TestAuthHandler_DeleteAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	validID := uuid.NewV4()
	invalidID := "invalid-uuid-format"

	type fields struct {
		uc *mocks.MockAuthUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.DeleteAddressRequest
	}
	tests := []struct {
		name        string
		setup       func(f *fields)
		args        args
		wantErr     bool
		wantErrCode codes.Code
	}{
		{
			name: "Success",
			setup: func(f *fields) {
				f.uc.EXPECT().
					DeleteAddress(gomock.Any(), validID).
					Return(nil)
			},
			args: args{
				ctx: context.Background(),
				in:  &gen.DeleteAddressRequest{Id: validID.String()},
			},
			wantErr: false,
		},
		{
			name: "Invalid UUID",
			setup: func(f *fields) {}, 
			args: args{
				ctx: context.Background(),
				in:  &gen.DeleteAddressRequest{Id: invalidID},
			},
			wantErr:     true,
			wantErrCode: codes.InvalidArgument,
		},
		{
			name: "Usecase error",
			setup: func(f *fields) {
				f.uc.EXPECT().
					DeleteAddress(gomock.Any(), validID).
					Return(errors.New("delete failed"))
			},
			args: args{
				ctx: context.Background(),
				in:  &gen.DeleteAddressRequest{Id: validID.String()},
			},
			wantErr:     true,
			wantErrCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fields{
				uc: mocks.NewMockAuthUsecase(ctrl),
			}
			if tt.setup != nil {
				tt.setup(f)
			}

			h := &AuthHandler{uc: f.uc}
			_, err := h.DeleteAddress(tt.args.ctx, tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Fatalf("DeleteAddress() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				st, ok := status.FromError(err)
				if !ok {
					t.Fatalf("expected gRPC status error, got: %v", err)
				}
				if st.Code() != tt.wantErrCode {
					t.Errorf("expected gRPC error code %v, got %v", tt.wantErrCode, st.Code())
				}
			}
		})
	}
}

func TestAuthHandler_AddAddress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	validAddressID := uuid.NewV4()
	validUserID := uuid.NewV4()
	invalidUUID := "invalid-uuid"

	type fields struct {
		uc *mocks.MockAuthUsecase
	}
	type args struct {
		ctx context.Context
		in  *gen.Address
	}
	tests := []struct {
		name        string
		setup       func(f *fields)
		args        args
		wantErr     bool
		wantErrCode codes.Code
	}{
		{
			name: "Success",
			setup: func(f *fields) {
				f.uc.EXPECT().
					AddAddress(gomock.Any(), models.Address{
						Id:      validAddressID,
						Address: "123 Main St",
						UserId:  validUserID,
					}).
					Return(nil)
			},
			args: args{
				ctx: context.Background(),
				in: &gen.Address{
					Id:      validAddressID.String(),
					Address: "123 Main St",
					UserId:  validUserID.String(),
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Address ID",
			setup: func(f *fields) {},
			args: args{
				ctx: context.Background(),
				in: &gen.Address{
					Id:      invalidUUID,
					Address: "Somewhere",
					UserId:  validUserID.String(),
				},
			},
			wantErr:     true,
			wantErrCode: codes.InvalidArgument,
		},
		{
			name: "Invalid User ID",
			setup: func(f *fields) {}, 
			args: args{
				ctx: context.Background(),
				in: &gen.Address{
					Id:      validAddressID.String(),
					Address: "Somewhere",
					UserId:  invalidUUID,
				},
			},
			wantErr:     true,
			wantErrCode: codes.InvalidArgument,
		},
		{
			name: "Usecase error",
			setup: func(f *fields) {
				f.uc.EXPECT().
					AddAddress(gomock.Any(), models.Address{
						Id:      validAddressID,
						Address: "Fail St",
						UserId:  validUserID,
					}).
					Return(errors.New("internal error"))
			},
			args: args{
				ctx: context.Background(),
				in: &gen.Address{
					Id:      validAddressID.String(),
					Address: "Fail St",
					UserId:  validUserID.String(),
				},
			},
			wantErr:     true,
			wantErrCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fields{
				uc: mocks.NewMockAuthUsecase(ctrl),
			}
			if tt.setup != nil {
				tt.setup(f)
			}

			h := &AuthHandler{uc: f.uc}
			_, err := h.AddAddress(tt.args.ctx, tt.args.in)

			if (err != nil) != tt.wantErr {
				t.Fatalf("AddAddress() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				st, ok := status.FromError(err)
				if !ok {
					t.Fatalf("expected gRPC status error, got: %v", err)
				}
				if st.Code() != tt.wantErrCode {
					t.Errorf("expected gRPC error code %v, got %v", tt.wantErrCode, st.Code())
				}
			}
		})
	}
}
