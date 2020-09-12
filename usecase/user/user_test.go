package user

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/common/constants/user_role"
	auth2 "github.com/williamchang80/sea-apd/domain/auth"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/auth"
	user2 "github.com/williamchang80/sea-apd/dto/request/user"
	user3 "github.com/williamchang80/sea-apd/mocks/repository/user"
	auth4 "github.com/williamchang80/sea-apd/mocks/usecase/auth"
	"reflect"
	"testing"
)

var (
	mockRegisterUserRequest = auth.RegisterUserRequest{
		Email:                "test@test.com",
		Password:             "123",
		Name:                 "name",
	}
	mockUpdateUserRoleRequest = user2.UpdateUserRoleRequest{
		Role:   user_role.MERCHANT,
		UserId: "1",
	}
	mockUpdateUserRequest = user2.UpdateUserRequest{
		OldPassword: "123",
		NewPassword: "123",
		NewEmail:    "123",
		OldEmail:    "",
		UserId:      "1123",
	}
	mockId = "1"
)

func TestNewUserUsecase(t *testing.T) {
	type args struct {
		repository user.UserRepository
		usecase    auth2.AuthUsecase
	}
	tests := []struct {
		name string
		args args
		want user.UserUsecase
	}{
		{
			name: "success",
			args: args{
				repository: nil,
			},
			want: UserUsecase{
				repo:    nil,
				usecase: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUsecase(tt.args.repository, tt.args.usecase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserUsecase_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request auth.RegisterUserRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() user.UserUsecase
	}{
		{
			name: "success",
			args: args{
				request: mockRegisterUserRequest,
			},
			wantErr: false,
			initMock: func() user.UserUsecase {
				r := user3.NewMockRepository(ctrl)
				a := auth4.NewMockUsecase(ctrl)
				return NewUserUsecase(r, a)
			},
		},
		{
			name: "failed with empty object args",
			args: args{
				request: auth.RegisterUserRequest{},
			},
			wantErr: true,
			initMock: func() user.UserUsecase {
				r := user3.NewMockRepository(ctrl)
				a := auth4.NewMockUsecase(ctrl)
				return NewUserUsecase(r, a)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.CreateUser(tt.args.request)
			if err != nil && !tt.wantErr {
				t.Errorf("UserUsecase.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserUsecase_GetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		userId string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() user.UserUsecase
	}{
		{
			name: "success",
			args: args{
				userId: mockId,
			},
			wantErr: false,
			initMock: func() user.UserUsecase {
				r := user3.NewMockRepository(ctrl)
				a := auth4.NewMockUsecase(ctrl)
				return NewUserUsecase(r, a)
			},
		},
		{
			name: "failed with empty object args",
			args: args{
				userId: "",
			},
			wantErr: true,
			initMock: func() user.UserUsecase {
				r := user3.NewMockRepository(ctrl)
				a := auth4.NewMockUsecase(ctrl)
				return NewUserUsecase(r, a)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			_, err := c.GetUserById(tt.args.userId)
			if err != nil && !tt.wantErr {
				t.Errorf("UserUsecase.GetUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request user2.UpdateUserRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() user.UserUsecase
	}{
		{
			name: "success",
			args: args{
				request: mockUpdateUserRequest,
			},
			wantErr: true,
			initMock: func() user.UserUsecase {
				r := user3.NewMockRepository(ctrl)
				a := auth4.NewMockUsecase(ctrl)
				return NewUserUsecase(r, a)
			},
		},
		{
			name: "failed with not matched request",
			args: args{
				request: user2.UpdateUserRequest{
					OldPassword: "123",
					NewPassword: "123",
					NewEmail:    "123",
					OldEmail:    "123",
					UserId:      "1",
				},
			},
			wantErr: true,
			initMock: func() user.UserUsecase {
				r := user3.NewMockRepository(ctrl)
				a := auth4.NewMockUsecase(ctrl)
				return NewUserUsecase(r, a)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.UpdateUser(tt.args.request)
			if err != nil && !tt.wantErr {
				t.Errorf("UserUsecase.UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestUserUsecase_UpdateUserRole(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request user2.UpdateUserRoleRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() user.UserUsecase
	}{
		{
			name: "success",
			args: args{
				request: mockUpdateUserRoleRequest,
			},
			wantErr: false,
			initMock: func() user.UserUsecase {
				r := user3.NewMockRepository(ctrl)
				a := auth4.NewMockUsecase(ctrl)
				return NewUserUsecase(r, a)
			},
		},
		{
			name: "failed with empty object args",
			args: args{
				request: user2.UpdateUserRoleRequest{},
			},
			wantErr: true,
			initMock: func() user.UserUsecase {
				r := user3.NewMockRepository(ctrl)
				a := auth4.NewMockUsecase(ctrl)
				return NewUserUsecase(r, a)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.UpdateUserRole(tt.args.request)
			if err != nil && !tt.wantErr {
				t.Errorf("UserUsecase.UpdateUserRole() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
