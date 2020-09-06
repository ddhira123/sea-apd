package merchant

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	domain "github.com/williamchang80/sea-apd/domain/merchant"
	mock_psql "github.com/williamchang80/sea-apd/mocks/postgres"
	"reflect"
	"regexp"
	"testing"
)

func TestNewMerchantRepository(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want domain.MerchantRepository
	}{
		{
			name: "success with null value on db",
			args: args{
				db: nil,
			},
			want: &MerchantRepository{
				db: nil,
			},
		},
		{
			name: "success with value on db",
			args: args{
				db: db,
			},
			want: &MerchantRepository{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMerchantRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMerchantRepository_GetMerchantBalance(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		merchantId string
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "failed with not matched query",
			args: args{
				merchantId: "",
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
		{
			name: "success",
			args: args{
				merchantId: "1",
			},
			want:    10000,
			wantErr: false,
			initMock: func() *gorm.DB {
				rows := sqlmock.NewRows([]string{
					"balance",
					"name",
					"user_id",
				}).AddRow(10000, "name", "1")
				mocks.ExpectQuery(regexp.QuoteMeta(`
					SELECT
						*
					FROM
						"merchants"
					WHERE
						"merchants"."deleted_at" IS NULL
					AND (("merchants"."id" = $1))
				`)).WillReturnRows(rows)
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := MerchantRepository{
				db: tt.initMock(),
			}
			balance, err := pr.GetMerchantBalance(tt.args.merchantId)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantRepository.GetMerchantBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(balance, tt.want) {
				t.Errorf("MerchantRepository.Find() = %v, want %v", balance, tt.want)
			}
		})
	}
}

func TestMerchantRepository_UpdateMerchantBalance(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		merchantId string
		amount     int
	}
	tests := []struct {
		name     string
		args     args
		want     int
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "failed with not matched query",
			args: args{
				merchantId: "",
				amount:     0,
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
		{
			name: "failed",
			args: args{
				merchantId: "1",
				amount:     10000,
			},
			want:    10000,
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := MerchantRepository{
				db: tt.initMock(),
			}
			err := pr.UpdateMerchantBalance(tt.args.amount, tt.args.merchantId)
			if err != nil && !tt.wantErr {
				t.Errorf("MerchantRepository.UpdateMerchantBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
