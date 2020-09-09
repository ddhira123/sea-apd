package merchant

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	domain "github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/mocks/repository/merchant"
	merchant_mock_usecase "github.com/williamchang80/sea-apd/mocks/usecase/merchant"
	merchant2 "github.com/williamchang80/sea-apd/usecase/merchant"
)

var (
	mockId = "1"
)

func TestNewMerchantController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := echo.New()
	repo := merchant.NewMockRepository(ctrl)
	type args struct {
		ctx *echo.Echo
	}
	tests := []struct {
		name     string
		args     args
		want     domain.MerchantController
		initMock func() domain.MerchantUsecase
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: &MerchantController{
				usecase: merchant2.NewMerchantUsecase(repo),
			},
			initMock: func() domain.MerchantUsecase {
				c := merchant_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()

			if got := NewMerchantController(tt.args.ctx, mock); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("NewProductController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerchantController_GetMerchantBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx       *echo.Echo
		getParams func() url.Values
	}
	defer ctrl.Finish()
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantStatus int
		initMock   func() domain.MerchantUsecase
	}{
		{
			name: "failed with empty request and invalid status",
			args: args{
				ctx: echo.New(),
				getParams: func() url.Values {
					q := make(url.Values)
					return q
				},
			},
			wantErr:    false,
			wantStatus: http.StatusNotFound,
			initMock: func() domain.MerchantUsecase {
				return merchant_mock_usecase.NewMockUsecase(ctrl)
			},
		},
		{
			name: "success",
			args: args{
				ctx: echo.New(),
				getParams: func() url.Values {
					q := make(url.Values)
					q.Set("merchantId", mockId)
					return q
				},
			},
			wantErr:    false,
			wantStatus: http.StatusOK,
			initMock: func() domain.MerchantUsecase {
				return merchant_mock_usecase.NewMockUsecase(ctrl)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			params := tt.args.getParams()
			c := echo.New()
			req, err := http.NewRequest(echo.GET, "api/merchant/balance"+"?"+params.Encode(), nil)
			if err != nil && !tt.wantErr {
				t.Errorf("GetMerchantBalance() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewMerchantController(c, mock)
			if controller.GetMerchantBalance(ctx); (rec.Code != tt.wantStatus) || tt.wantErr {
				t.Errorf("GetMerchantBalance() error= %v, want %v", rec.Code, tt.wantStatus)
			}
		})
	}
}
