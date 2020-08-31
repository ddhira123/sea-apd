package product

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	domain "github.com/williamchang80/sea-apd/domain/product"
	request "github.com/williamchang80/sea-apd/dto/request/product"
	"github.com/williamchang80/sea-apd/mocks/repository"
	"github.com/williamchang80/sea-apd/mocks/usecase"
	"github.com/williamchang80/sea-apd/usecase/product"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
)

var (
	mockData = request.ProductRequest{
		Name:        "Mock name",
		Stock:       20,
		Description: "Mock desc",
		Price:       10,
		Image:       nil,
	}
	mockId = "1"
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
				usecase: product.NewProductUseCase(repo),
			},
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return product.NewProductUseCase(c)
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
				return product.NewProductUseCase(c)
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
		ctx     *echo.Echo
		request request.ProductRequest
	}
	const FILE_PATH = "../mocks/file/mock_image.jpg"
	image, _ := os.Open(FILE_PATH)
	defer image.Close()
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
				request: request.ProductRequest{
					Name:        "Mock name",
					Stock:       10,
					Description: "Mock desc",
					Price:       20,
					Image:       image,
				},
			},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return product.NewProductUseCase(c)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			data, _ := json.Marshal(mockData)
			req, err := http.NewRequest(echo.POST, "/product", strings.NewReader(string(data)))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
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
		ctx       *echo.Echo
		getParams func() url.Values
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
				getParams: func() url.Values {
					q := make(url.Values)
					q.Set("productId", mockId)
					return q
				},
			},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return product.NewProductUseCase(c)
			},
		},
		{
			name: "failed with no params",
			args: args{
				ctx: echo.New(),
				getParams: func() url.Values {
					q := make(url.Values)
					return q
				},
			},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return product.NewProductUseCase(c)
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

func TestProductController_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	type UpdateRequest struct {
		request.ProductRequest
		ProductId string
	}

	type args struct {
		ctx     *echo.Echo
		request UpdateRequest
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
				request: UpdateRequest{
					ProductRequest: mockData,
					ProductId:      mockId,
				},
			},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return product.NewProductUseCase(c)
			},
		},
		{
			name: "failed with invalid request",
			args: args{
				ctx: echo.New(),
			},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return product.NewProductUseCase(c)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			data, _ := json.Marshal(tt.args.request)
			req, err := http.NewRequest(echo.PUT, "/product", strings.NewReader(string(data)))
			if err != nil {
				t.Errorf("DeleteProduct() request error= %v", tt.wantErr)
			}
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewProductController(c, mock)
			if err := controller.UpdateProduct(ctx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateProduct() error= %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductController_GetProductById(t *testing.T) {
	ctrl := gomock.NewController(t)
	type args struct {
		ctx       *echo.Echo
		getParams func() url.Values
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
				getParams: func() url.Values {
					q := make(url.Values)
					q.Set("productId", mockId)
					return q
				},
			},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return product.NewProductUseCase(c)
			},
		},
		{
			name: "failed with no params",
			args: args{
				ctx: echo.New(),
				getParams: func() url.Values {
					q := make(url.Values)
					return q
				},
			},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				c := usecase.NewMockUsecase(ctrl)
				return product.NewProductUseCase(c)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := tt.initMock()
			c := echo.New()
			params := tt.args.getParams()
			req := httptest.NewRequest(echo.GET, "/product"+"/?"+params.Encode(), nil)
			rec := httptest.NewRecorder()
			ctx := c.NewContext(req, rec)
			controller := NewProductController(c, mock)
			if err := controller.GetProductById(ctx); (err != nil) != tt.wantErr {
				t.Errorf("GetProductById() error= %v, want %v", err, tt.wantErr)
			}
		})
	}
}
