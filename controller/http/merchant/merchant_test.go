package merchant

import (
	"encoding/json"
	"github.com/williamchang80/sea-apd/common/constants/merchant_status"
	merchant3 "github.com/williamchang80/sea-apd/dto/request/merchant"
	"github.com/williamchang80/sea-apd/mocks/usecase/user"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	domain "github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/mocks/repository/merchant"
	merchant_mock_usecase "github.com/williamchang80/sea-apd/mocks/usecase/merchant"
	merchant2 "github.com/williamchang80/sea-apd/usecase/merchant"
)

var (
	mockId              = "1"
	mockMerchantRequest = merchant3.MerchantRequest{
		Name:    "name",
		UserId:  "2",
		Brand:   "teest",
		Address: "address",
	}
	mockUpdateMerchantApprovalStatusRequest = merchant3.UpdateMerchantApprovalStatusRequest{
		Status:     merchant_status.ACCEPTED,
		MerchantId: "1",
	}
	mockUpdateMerchantRequest = merchant3.UpdateMerchantRequest{
		MerchantId: "1",
		Name:       "name",
		Brand:      "test",
		Address:    "address",
	}
)

func TestNewMerchantController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := echo.New()
	repo := merchant.NewMockRepository(ctrl)
	userUsecase := user.NewMockUsecase(ctrl)
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
				usecase: merchant2.NewMerchantUsecase(repo, userUsecase),
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

func TestMerchantController_RegisterMerchant(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx     *echo.Echo
		request merchant3.MerchantRequest
	}
	defer ctrl.Finish()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() domain.MerchantUsecase
	}{
		{
			name: "success",
			args: args{
				ctx:     echo.New(),
				request: mockMerchantRequest,
			},
			wantErr: false,
			initMock: func() domain.MerchantUsecase {
				c := merchant_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
		{
			name: "fail with empty request",
			args: args{
				ctx:     echo.New(),
				request: merchant3.MerchantRequest{},
			},
			wantErr: false,
			initMock: func() domain.MerchantUsecase {
				c := merchant_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			data, _ := json.Marshal(tt.args.request)
			req, err := http.NewRequest(echo.POST, "api/merchant", strings.NewReader(string(data)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			if err != nil {
				t.Errorf("RegisterMerchant() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewMerchantController(c, mock)
			if got := controller.RegisterMerchant(ctx); (got != nil) != tt.wantErr {
				t.Errorf("RegisterMerchant() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestMerchantController_GetMerchantById(t *testing.T) {
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
			req, err := http.NewRequest(echo.GET, "api/merchant"+"?"+params.Encode(), nil)
			if err != nil && !tt.wantErr {
				t.Errorf("GetMerchantById() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewMerchantController(c, mock)
			if controller.GetMerchantById(ctx); (rec.Code != tt.wantStatus) || tt.wantErr {
				t.Errorf("GetMerchantById() error= %v, want %v", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestMerchantController_GetMerchants(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx *echo.Echo
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
			name: "success",
			args: args{
				ctx: echo.New(),
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
			c := echo.New()
			req, err := http.NewRequest(echo.GET, "api/merchants", nil)
			if err != nil && !tt.wantErr {
				t.Errorf("GetMerchants() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewMerchantController(c, mock)
			if controller.GetMerchants(ctx); (rec.Code != tt.wantStatus) || tt.wantErr {
				t.Errorf("GetMerchants() error= %v, want %v", rec.Code, tt.wantStatus)
			}
		})
	}
}

func TestMerchantController_UpdateMerchantApprovalStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx     *echo.Echo
		request merchant3.UpdateMerchantApprovalStatusRequest
	}
	defer ctrl.Finish()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() domain.MerchantUsecase
	}{
		{
			name: "success",
			args: args{
				ctx:     echo.New(),
				request: mockUpdateMerchantApprovalStatusRequest,
			},
			wantErr: false,
			initMock: func() domain.MerchantUsecase {
				c := merchant_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
		{
			name: "fail with empty request",
			args: args{
				ctx:     echo.New(),
				request: merchant3.UpdateMerchantApprovalStatusRequest{},
			},
			wantErr: false,
			initMock: func() domain.MerchantUsecase {
				c := merchant_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			data, _ := json.Marshal(tt.args.request)
			req, err := http.NewRequest(echo.POST, "api/merchant/status",
				strings.NewReader(string(data)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			if err != nil {
				t.Errorf("UpdateMerchantApprovalStatus() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewMerchantController(c, mock)
			if got := controller.UpdateMerchantApprovalStatus(ctx); (got != nil) != tt.wantErr {
				t.Errorf("UpdateMerchantApprovalStatus() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestMerchantController_UpdateMerchant(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx     *echo.Echo
		request merchant3.UpdateMerchantRequest
	}
	defer ctrl.Finish()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() domain.MerchantUsecase
	}{
		{
			name: "success",
			args: args{
				ctx:     echo.New(),
				request: mockUpdateMerchantRequest,
			},
			wantErr: false,
			initMock: func() domain.MerchantUsecase {
				c := merchant_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
		{
			name: "fail with empty request",
			args: args{
				ctx:     echo.New(),
				request: merchant3.UpdateMerchantRequest{},
			},
			wantErr: false,
			initMock: func() domain.MerchantUsecase {
				c := merchant_mock_usecase.NewMockUsecase(ctrl)
				return c
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			data, _ := json.Marshal(tt.args.request)
			req, err := http.NewRequest(echo.PUT, "api/merchant",
				strings.NewReader(string(data)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			if err != nil {
				t.Errorf("UpdateMerchant() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewMerchantController(c, mock)
			if got := controller.UpdateMerchant(ctx); (got != nil) != tt.wantErr {
				t.Errorf("UpdateMerchant() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}