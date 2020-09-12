package auth

import (
	"github.com/golang/mock/gomock"
	auth_domain "github.com/williamchang80/sea-apd/domain/auth"
	"github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/dto/request/auth"
	user2 "github.com/williamchang80/sea-apd/mocks/repository/user"
	"reflect"
	"testing"
)

var (
	mockLoginRequest = auth.LoginRequest{
		Email:    "test@test.com",
		Password: "123",
	}
)

func TestNewAuthUsecase(t *testing.T) {
	type args struct {
		repository user.UserRepository
	}
	tests := []struct {
		name string
		args args
		want auth_domain.AuthUsecase
	}{
		{
			name: "success",
			args: args{
				repository: nil,
			},
			want: AuthUsecase{
				repo: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAuthUsecase(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAuthUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthUsecase_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request auth.LoginRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() auth_domain.AuthUsecase
	}{

		{
			name: "failed with empty object request",
			args: args{
				request: auth.LoginRequest{},
			},
			wantErr: true,
			initMock: func() auth_domain.AuthUsecase {
				r := user2.NewMockRepository(ctrl)
				return NewAuthUsecase(r)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			_, err := c.Login(tt.args.request)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductUsecase.Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

