package product

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	domain "github.com/williamchang80/sea-apd/domain/product"
	mock_psql "github.com/williamchang80/sea-apd/mocks/postgres"
	"reflect"
	"regexp"
	"testing"
	"time"
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
				t.Errorf("ProductRepository.GetProducts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(products, tt.want) {
				t.Errorf("ProductRepository.Find() = %v, want %v", products, tt.want)
			}
		})
	}
}

func TestProductRepository_GetProductById(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		productId string
	}
	tests := []struct {
		name     string
		args     args
		want     *domain.Product
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "failed with record not found",
			args: args{
				productId: "0",
			},
			want:    nil,
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectQuery(regexp.QuoteMeta(`
					select
						*
					FROM 
						"products"
					WHERE
						"products"."deleted_at" IS NULL AND
						(("products"."id" = $1))
				`))
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := ProductRepository{
				db: tt.initMock(),
			}
			products, err := pr.GetProductById(tt.args.productId)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductRepository.GetProductById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(products, tt.want) {
				t.Errorf("ProductRepository.Find() = %v, want %v", products, tt.want)
			}
		})
	}
}

func TestProductRepository_CreateProduct(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		product domain.Product
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "fail with invalid db query",
			args: args{
				product: domain.Product{
					Name:        "mock name",
					Description: "mock description",
					Price:       1000,
					Image:       "https://mock.image.com/mock-image",
					Stock:       10,
				},
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectBegin()
				mocks.ExpectQuery(`
					"INSERT INTO "products" 
						("created_at", 
						 "updated_at", 
 						 "deleted_at", 
						 "name", 
						 "description", 
						 "price", 
						 "image",
						 "stock") 
                     VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
					 RETURNING "products"."id"
				`).WithArgs(
					time.Date(2020, time.January, 12, 12, 12, 12, 12, time.UTC),
					time.Date(2020, time.January, 12, 12, 12, 12, 12, time.UTC),
					nil,
					nil,
					"mock description",
					1000,
					"https://mock.image.com/mock-image",
					10,
				).WillReturnRows(sqlmock.NewRows([]string{
					"id",
				}))
				mocks.ExpectCommit()
				return db
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := ProductRepository{
				db: tt.initMock(),
			}
			err := pr.CreateProduct(tt.args.product)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductRepository.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProductRepository_DeleteProduct(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		productId string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "fail with invalid id",
			args: args{
				productId: "0",
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectBegin()
				mocks.ExpectExec(regexp.QuoteMeta(`
					UPDATE "products" 
					SET "deleted_at"= NOW() 
					WHERE "products"."deleted_at" 
					IS NULL AND (("products"."id" = $1))
				`)).WithArgs("0").
					WillReturnError(errors.New("No data found"))
				mocks.ExpectCommit()
				return db
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := ProductRepository{
				db: tt.initMock(),
			}
			err := pr.DeleteProduct(tt.args.productId)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductRepository.DeleteProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProductRepository_UpdateProduct(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		productId string
		product   domain.Product
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "fail with invalid id",
			args: args{
				productId: "0",
				product: domain.Product{
					Name:        "mock name",
					Description: "mock description",
					Price:       1000,
					Image:       "https://mock.image.com/mock-image",
					Stock:       10,
				},
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectBegin()
				mocks.ExpectExec(regexp.QuoteMeta(`
					UPDATE "products" 
					SET "description" = $1, 
						"image" = $2, 
						"name" = $3, 
						"price" = $4, 
						"stock" = $5, 
						"updated_at" = $6 
					WHERE "products"."deleted_at" 
						IS NULL AND ((id = 0))
				`)).WithArgs("mock description", "https://mock.image.com/mock-image",
					"mock name", 1000, 10, time.Now()).
					WillReturnError(errors.New("No data found"))
				mocks.ExpectCommit()
				return db
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := ProductRepository{
				db: tt.initMock(),
			}
			err := pr.UpdateProduct(tt.args.productId, tt.args.product)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductRepository.UpdateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProductRepository_GetProductsByMerchant(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		merchantId string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		want     []domain.Product
		initMock func() *gorm.DB
	}{
		{
			name: "fail with invalid id",
			args: args{
				merchantId: "0",
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectQuery(regexp.QuoteMeta(`
					SELECT *
					FROM "products"
					WHERE "products"."deleted_at" 
						IS NULL AND ((merchant_id = $1))
				`)).WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{
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
			prod, err := pr.GetProductsByMerchant(tt.args.merchantId)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductRepository.GetProductsByMerchant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(tt.want, prod) {
				t.Errorf("ProductRepository.GetProductsByMerchant() Find = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
