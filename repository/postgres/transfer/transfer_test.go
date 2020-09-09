package transfer

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	domain "github.com/williamchang80/sea-apd/domain/transfer"
	mock_psql "github.com/williamchang80/sea-apd/mocks/postgres"
	"reflect"
	"regexp"
	"testing"
)

func TestNewTransferRepository(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want domain.TransferRepository
	}{
		{
			name: "success with null value on db",
			args: args{
				db: nil,
			},
			want: &TransferRepository{
				db: nil,
			},
		},
		{
			name: "success with value on db",
			args: args{
				db: db,
			},
			want: &TransferRepository{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransferRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewTransferRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransferRepository_GetTransferHistory(t *testing.T) {
	db, _ := mock_psql.Connection()
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
			name: "failed",
			args: args{
				merchantId: "1",
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
			pr := TransferRepository{
				db: tt.initMock(),
			}
			_, err := pr.GetTransferHistory(tt.args.merchantId)
			if err != nil && !tt.wantErr {
				t.Errorf("TransferRepository.GetTransferHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTransferRepository_CreateTransferHistory(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		transfer domain.Transfer
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
				transfer: domain.Transfer{
					Amount:     1000,
					BankName:   "name",
					BankNumber: "23",
					MerchantId: "1",
				},
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectBegin()
				mocks.ExpectQuery(regexp.QuoteMeta(`
				INSERT INTO "transfers" 
					("id","created_at","updated_at",
					 "deleted_at","amount","bank_name",
					 "bank_number","merchant_id") 
                VALUES (?,?,?,?,?,?,?,?)
				`)).WithArgs(sqlmock.AnyArg(), 3000, "name", "123", "1").WillReturnError(sqlmock.ErrCancelled)
				mocks.ExpectCommit()
				return db
			},
		},
		{
			name: "failed",
			args: args{
				transfer: domain.Transfer{},
			},
			want:    10000,
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectBegin()
				mocks.ExpectCommit()
				return db
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := TransferRepository{
				db: tt.initMock(),
			}
			err := pr.CreateTransferHistory(tt.args.transfer)
			if err != nil && !tt.wantErr {
				t.Errorf("TransferRepository.CreateTransferHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
