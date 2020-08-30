package product

import (
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/mocks/repository"
	"github.com/williamchang80/sea-apd/mocks/usecase"
	usecase2 "github.com/williamchang80/sea-apd/usecase"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestNewProductController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := echo.New()
	repo := repository.NewMockRepository(ctrl)
	type args struct {
		ctx *echo.Echo
	}
	tests := []struct {
		name     string
		args     args
		want     domain.ProductController
		initMock func() domain.ProductUsecase
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
			},
			want: &ProductController{
				usecase: usecase2.NewProductUseCase(repo),
			},
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return usecase2.NewProductUseCase(c)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()

			if got := NewProductController(tt.args.ctx, mock); reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("NewProductController() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductController_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx *echo.Echo
	}
	defer ctrl.Finish()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() domain.ProductUsecase
	}{
		{
			name: "success",
			args: args{
				ctx: echo.New(),
			},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return usecase2.NewProductUseCase(c)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			req, err := http.NewRequest(echo.GET, "/products", strings.NewReader(""))
			if err != nil {
				t.Errorf("GetProducts() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewProductController(c, mock)
			if got := controller.GetProducts(ctx); (got != nil) != tt.wantErr {
				t.Errorf("GetProducts() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestProductController_CreateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx *echo.Echo
	}

	defer ctrl.Finish()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() domain.ProductUsecase
	}{
		{
			name: "success",
			args: args{
				ctx: echo.New(),
			},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return usecase2.NewProductUseCase(c)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			req, err := http.NewRequest(echo.POST, "/product", strings.NewReader(""))
			if err != nil {
				t.Errorf("CreateProduct() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewProductController(c, mock)
			if got := controller.CreateProduct(ctx); (got != nil) != tt.wantErr {
				t.Errorf("CreateProduct() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func TestProductController_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx *echo.Echo
	}

	defer ctrl.Finish()
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() domain.ProductUsecase
	}{
		{
			name: "success",
			args: args{
				ctx: echo.New(),
			},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return usecase2.NewProductUseCase(c)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			req, err := http.NewRequest(echo.DELETE, "/product", strings.NewReader(""))
			if err != nil {
				t.Errorf("DeleteProduct() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewProductController(c, mock)
			if got := controller.DeleteProduct(ctx); (got != nil) != tt.wantErr {
				t.Errorf("DeleteProduct() error= %v, want %v", got, tt.wantErr)
			}
		})
	}
}
