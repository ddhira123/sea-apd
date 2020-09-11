package merchant

import (
	"github.com/williamchang80/sea-apd/domain/user"
	user2 "github.com/williamchang80/sea-apd/mocks/usecase/user"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/merchant"
	request "github.com/williamchang80/sea-apd/dto/request/merchant"
	merchant2 "github.com/williamchang80/sea-apd/mocks/repository/merchant"
)

func TestNewMerchantUsecase(t *testing.T) {
	type args struct {
		repository merchant.MerchantRepository
		uc    user.UserUsecase
	}
	tests := []struct {
		name string
		args args
		want merchant.MerchantUsecase
	}{
		{
			name: "success",
			args: args{
				repository: nil,
				uc:    nil,
			},
			want: MerchantUsecase{
				mc: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMerchantUsecase(tt.args.repository, tt.args.uc);
				!reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMerchantUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerchantUsecase_GetMerchantBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		merchantId string
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantErr  bool
		initMock func() merchant.MerchantUsecase
	}{
		{
			name: "success",
			args: args{
				merchantId: "1",
			},
			want:    100,
			wantErr: false,
			initMock: func() merchant.MerchantUsecase {
				r := merchant2.NewMockRepository(ctrl)
				u := user2.NewMockUsecase(ctrl)
				return NewMerchantUsecase(r, u)
			},
		},
		{
			name: "failed with empty id args",
			args: args{
				merchantId: "",
			},
			wantErr: true,
			initMock: func() merchant.MerchantUsecase {
				r := merchant2.NewMockRepository(ctrl)
				u := user2.NewMockUsecase(ctrl)
				return NewMerchantUsecase(r, u)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			p, err := c.GetMerchantBalance(tt.args.merchantId)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantUsecase.GetMerchantBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("MerchantUsecase.GetMerchantBalance() = %v, got %v", tt.want, p)
			}
		})
	}
}

func TestMerchantUsecase_UpdateMerchantBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		request request.UpdateMerchantBalanceRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() merchant.MerchantUsecase
	}{
		{
			name: "success",
			args: args{
				request: request.UpdateMerchantBalanceRequest{
					Amount:     10000,
					MerchantId: "1",
				},
			},
			wantErr: false,
			initMock: func() merchant.MerchantUsecase {
				r := merchant2.NewMockRepository(ctrl)
				u := user2.NewMockUsecase(ctrl)
				return NewMerchantUsecase(r, u)
			},
		},
		{
			name: "failed with empty object request",
			args: args{
				request: request.UpdateMerchantBalanceRequest{},
			},
			wantErr: true,
			initMock: func() merchant.MerchantUsecase {
				r := merchant2.NewMockRepository(ctrl)
				u := user2.NewMockUsecase(ctrl)
				return NewMerchantUsecase(r, u)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.UpdateMerchantBalance(tt.args.request)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantUsecase.UpdateMerchantBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
