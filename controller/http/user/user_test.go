package user

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	user2 "github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/auth"
	"github.com/williamchang80/sea-apd/dto/request/user"
	user5 "github.com/williamchang80/sea-apd/mocks/repository/user"
	user3 "github.com/williamchang80/sea-apd/mocks/usecase/user"
	user4 "github.com/williamchang80/sea-apd/usecase/user"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

var (
	mockRegisterUserRequest = auth.RegisterUserRequest{
		Email:                "email@email.com",
		Password:             "123",
		Name:                 "test",
	}
	mockUpdateUserRequest = user.UpdateUserRequest{
		OldPassword: "pass",
		NewPassword: "new",
		NewEmail:    "new@old.com",
		OldEmail:    "old@old.com",
		UserId:      "123",
	}
)

func TestNewUserController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := echo.New()
	repo := user5.NewMockRepository(ctrl)
	type args struct {
		ctx *echo.Echo
	}
	tests := []struct {
		name     string
		args     args
		want     user2.UserController
		initMock func() user2.UserUsecase
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: &UserController{
				usecase: user4.NewUserUsecase(repo, nil),
			},
			initMock: func() user2.UserUsecase {
				return user3.NewMockUsecase(ctrl)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()

			if got := NewUserController(tt.args.ctx, mock); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("NewUserController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserController_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx     *echo.Echo
		request auth.RegisterUserRequest
	}
	defer ctrl.Finish()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() user2.UserUsecase
	}{
		{
			name: "success",
			args: args{
				ctx:     echo.New(),
				request: mockRegisterUserRequest,
			},
			wantErr: false,
			initMock: func() user2.UserUsecase {
				c := user3.NewMockUsecase(ctrl)
				return c
			},
		},
		{
			name: "fail with empty request",
			args: args{
				ctx:     echo.New(),
				request: auth.RegisterUserRequest{},
			},
			wantErr: false,
			initMock: func() user2.UserUsecase {
				c := user3.NewMockUsecase(ctrl)
				return c
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			data, _ := json.Marshal(tt.args.request)
			req, err := http.NewRequest(echo.POST, "api/auth/register",
				strings.NewReader(string(data)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			if err != nil {
				t.Errorf("CreateUser() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewUserController(c, mock)
			if got := controller.CreateUser(ctx); (got != nil) != tt.wantErr {
				t.Errorf("CreateUser() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestUserController_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx     *echo.Echo
		request user.UpdateUserRequest
	}
	defer ctrl.Finish()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() user2.UserUsecase
	}{
		{
			name: "success",
			args: args{
				ctx:     echo.New(),
				request: mockUpdateUserRequest,
			},
			wantErr: false,
			initMock: func() user2.UserUsecase {
				c := user3.NewMockUsecase(ctrl)
				return c
			},
		},
		{
			name: "fail with empty request",
			args: args{
				ctx:     echo.New(),
				request: user.UpdateUserRequest{},
			},
			wantErr: false,
			initMock: func() user2.UserUsecase {
				c := user3.NewMockUsecase(ctrl)
				return c
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			data, _ := json.Marshal(tt.args.request)
			req, err := http.NewRequest(echo.PUT, "api/user", strings.NewReader(string(data)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			if err != nil {
				t.Errorf("UpdateUser() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewUserController(c, mock)
			if got := controller.UpdateUser(ctx); (got != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}
