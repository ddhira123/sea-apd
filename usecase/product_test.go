package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/domain"
	"github.com/williamchang80/sea-apd/dto/request"
	repository2 "github.com/williamchang80/sea-apd/mocks/repository"
	"reflect"
	"testing"
)

func TestNewProductUseCase(t *testing.T) {
	type args struct {
		repository domain.ProductRepository
	}
	tests := []struct {
		name string
		args args
		want domain.ProductUsecase
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
		want     []domain.Product
		wantErr  bool
		initMock func() domain.ProductUsecase
	}{
		{
			name:    "success",
			want:    []domain.Product{},
			wantErr: false,
			initMock: func() domain.ProductUsecase {
				r := repository2.NewMockRepository(ctrl)
				return NewProductUseCase(r)
			},
		},
		{
			name:    "failed with error as return type",
			want:    []domain.Product{},
			wantErr: true,
			initMock: func() domain.ProductUsecase {
				r := repository2.NewMockRepository(ctrl)
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
		productRequest request.Product
	}
	tests := []struct {
		name string
		args args
		want domain.Product
	}{
		{
			name: "success",
			args: args{
				productRequest: request.Product{
					Name:        "Mock name",
					Stock:       10,
					Description: "Mock desc",
					Price:       30,
					Image:       nil,
				},
			},
			want: domain.Product{
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
