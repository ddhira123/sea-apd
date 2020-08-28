package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/domain"
	mock_psql "github.com/williamchang80/sea-apd/mocks/postgres"
	"reflect"
	"regexp"
	"testing"
)

func TestNewProductRepository(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want domain.ProductRepository
	}{
		{
			name: "success with null value on db",
			args: args{
				db: nil,
			},
			want: &ProductRepository{
				db: nil,
			},
		},
		{
			name: "success with value on db",
			args: args{
				db: db,
			},
			want: &ProductRepository{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProductRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductRepository_GetProducts(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()

	type args struct {
		product domain.Product
	}
	tests := []struct {
		name     string
		args     args
		want     []domain.Product
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "failed with not matched query",
			args: args{
				product: domain.Product{
					Name:        "mock name",
					Description: "mock description",
					Price:       1000,
					Image:       "https://mock.image.com/mock-image",
					Stock:       10,
				},
			},
			want:    []domain.Product(nil),
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
		{
			name: "success",
			args: args{
				product: domain.Product{
					Name:        "mock name",
					Description: "mock description",
					Price:       1000,
					Image:       "https://mock.image.com/mock-image",
					Stock:       10,
				},
			},
			want:    []domain.Product{},
			wantErr: false,
			initMock: func() *gorm.DB {
				mocks.ExpectQuery(regexp.QuoteMeta(`
					SELECT
						*
					FROM
						"products"
					WHERE
						"products"."deleted_at" IS NULL
				`)).WillReturnRows(sqlmock.NewRows([]string{
					"name",
					"description",
					"price",
					"image",
					"stock",
				}))
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := ProductRepository{
				db: tt.initMock(),
			}
			products, err := pr.GetProducts()
			if err != nil && !tt.wantErr {
				t.Errorf("ProductRepository.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(products, tt.want) {
				t.Errorf("ProductRepository.Find() = %v, want %v", products, tt.want)
			}
		})
	}
}
