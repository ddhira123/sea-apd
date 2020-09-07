package product

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain/product"
	request "github.com/williamchang80/sea-apd/dto/request/product"
	product2 "github.com/williamchang80/sea-apd/mocks/repository/product"
	"os"
	"reflect"
	"testing"
)

func TestNewProductUseCase(t *testing.T) {
	type args struct {
		repository product.ProductRepository
	}
	tests := []struct {
		name string
		args args
		want product.ProductUsecase
	}{
		{
			name: "success",
			args: args{
				repository: nil,
			},
			want: &ProductUsecase{
				pr: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductUseCase(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductUseCase_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name     string
		want     []product.Product
		wantErr  bool
		initMock func() product.ProductUsecase
	}{
		{
			name:    "success",
			want:    []product.Product{},
			wantErr: false,
			initMock: func() product.ProductUsecase {
				r := product2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
		{
			name:    "failed with error as return type",
			want:    []product.Product{},
			wantErr: true,
			initMock: func() product.ProductUsecase {
				r := product2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			p, err := c.GetProducts()
			if err != nil && !tt.wantErr {
				t.Errorf("ProductUsecase.GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("ProductUsecase.GetProducts() = %v, got %v", tt.want, p)
			}
		})
	}
}

func TestConvertToDomain(t *testing.T) {
	type args struct {
		productRequest request.ProductRequest
	}
	tests := []struct {
		name string
		args args
		want product.Product
	}{
		{
			name: "success",
			args: args{
				productRequest: request.ProductRequest{
					Name:        "Mock name",
					Stock:       10,
					Description: "Mock desc",
					Price:       30,
					Image:       nil,
				},
			},
			want: product.Product{
				Name:        "Mock name",
				Stock:       10,
				Description: "Mock desc",
				Price:       30,
				Image:       "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToDomain(tt.args.productRequest); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvertToDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductUsecase_GetProductById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		id string
	}
	tests := []struct {
		name     string
		args     args
		want     *product.Product
		wantErr  bool
		initMock func() product.ProductUsecase
	}{
		{
			name: "success",
			args: args{
				id: "Mock id",
			},
			want: &product.Product{
				Name:        "Mock Name",
				Description: "Mock Desc",
				Price:       20,
				Image:       "Mock image",
				Stock:       30,
			},
			wantErr: false,
			initMock: func() product.ProductUsecase {
				r := product2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
		{
			name: "failed with empty id args",
			args: args{
				id: "",
			},
			wantErr: true,
			initMock: func() product.ProductUsecase {
				r := product2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			p, err := c.GetProductById(tt.args.id)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductUsecase.GetProductById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(p, tt.want) {
				t.Errorf("ProductUsecase.GetProductById() = %v, got %v", tt.want, p)
			}
		})
	}
}

func TestProductUsecase_DeleteProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		id string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() product.ProductUsecase
	}{
		{
			name: "success",
			args: args{
				id: "Mock id",
			},
			wantErr: false,
			initMock: func() product.ProductUsecase {
				r := product2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
		{
			name: "failed with empty id args",
			args: args{
				id: "",
			},
			wantErr: true,
			initMock: func() product.ProductUsecase {
				r := product2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.DeleteProduct(tt.args.id)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductUsecase.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProductUsecase_CreateProduct(t *testing.T) {
	const FILE_PATH = "../mocks/file/mock_image.jpg"
	ctrl := gomock.NewController(t)
	image, _ := os.Open(FILE_PATH)
	defer image.Close()
	defer ctrl.Finish()
	type args struct {
		request request.ProductRequest
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() product.ProductUsecase
	}{
		{
			name: "success",
			args: args{
				request: request.ProductRequest{
					Name:        "Mock name",
					Stock:       10,
					Description: "Mock desc",
					Price:       20,
					Image:       image,
				},
			},
			wantErr: false,
			initMock: func() product.ProductUsecase {
				r := product2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
		{
			name: "failed with empty object request",
			args: args{
				request: request.ProductRequest{},
			},
			wantErr: true,
			initMock: func() product.ProductUsecase {
				r := product2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.CreateProduct(tt.args.request)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductUsecase.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProductUsecase_UpdateProduct(t *testing.T) {
	const FILE_PATH = "../mocks/file/mock_image.jpg"
	ctrl := gomock.NewController(t)
	image, _ := os.Open(FILE_PATH)
	defer image.Close()
	defer ctrl.Finish()
	type args struct {
		request   request.ProductRequest
		productId string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() product.ProductUsecase
	}{
		{
			name: "success",
			args: args{
				request: request.ProductRequest{
					Name:        "Mock name",
					Stock:       10,
					Description: "Mock desc",
					Price:       20,
					Image:       image,
				},
				productId: "1",
			},
			wantErr: false,
			initMock: func() product.ProductUsecase {
				r := product2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
		{
			name: "failed with empty object request",
			args: args{
				request: request.ProductRequest{},
				productId: "",
			},
			wantErr: true,
			initMock: func() product.ProductUsecase {
				r := product2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := tt.initMock()
			err := c.UpdateProduct(tt.args.productId, tt.args.request)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductUsecase.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
