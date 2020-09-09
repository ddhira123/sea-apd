package user

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	domain "github.com/williamchang80/sea-apd/domain/user"
	request "github.com/williamchang80/sea-apd/dto/request/user"
	user_repository "github.com/williamchang80/sea-apd/mocks/repository/user"
	user_mock_usecase "github.com/williamchang80/sea-apd/mocks/usecase/user"
	user_usecase "github.com/williamchang80/sea-apd/usecase/user"
)

var (
	mockData = request.UserRequest{
		Name:     "Mock name",
		Email:    "Mock email",
		Password: "Mock desc",
		Role:     "4",
		Phone:    "Mock Phone",
	}
	mockId = "1"
)

func TestNewUserController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := echo.New()
	repo := user_repository.NewMockRepository(ctrl)
	type args struct {
		ctx *echo.Echo
	}
	tests := []struct {
		name     string
		args     args
		want     domain.UserController
		initMock func() domain.UserUsecase
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: &UserController{
				usecase: user_usecase.NewUserUsecase(repo),
			},
			initMock: func() domain.UserUsecase {
				return user_mock_usecase.NewMockUsecase(ctrl)
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

// func TestUserController_GetUsers(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	type args struct {
// 		ctx *echo.Echo
// 	}
// 	tests := []struct {
// 		name     string
// 		args     args
// 		wantErr  bool
// 		initMock func() domain.UserUsecase
// 	}{
// 		{
// 			name: "success",
// 			args: args{
// 				ctx: echo.New(),
// 			},
// 			wantErr: false,
// 			initMock: func() domain.UserUsecase {
// 				c := user_mock_usecase.NewMockUsecase(ctrl)
// 				return c
// 			},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mock := tt.initMock()
// 			c := echo.New()
// 			req, err := http.NewRequest(echo.GET, "api/list-users", strings.NewReader(""))
// 			if err != nil {
// 				t.Errorf("GetUsers() request error= %v", tt.wantErr)
// 			}
// 			rec := httptest.NewRecorder()
// 			ctx := c.NewContext(req, rec)
// 			controller := NewUserController(c, mock)
// 			if got := controller.GetUsers(ctx); (got != nil) != tt.wantErr {
// 				t.Errorf("GetUsers() error= %v, want %v", got, tt.wantErr)
// 			}
// 		})
// 	}
// }
