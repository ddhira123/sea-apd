package transaction

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	domain "github.com/williamchang80/sea-apd/domain/transaction"
	request "github.com/williamchang80/sea-apd/dto/request/transaction"
	mock_psql "github.com/williamchang80/sea-apd/mocks/postgres"
	"reflect"
	"regexp"
	"testing"
)

var (
	mockCreateTransactionRequest = request.TransactionRequest{
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		UserId:     "1",
	}
	transactionStatus     = transaction_status.GetTransactionStatus()
	mockTransactionEntity = domain.Transaction{
		Status:     transactionStatus["ONPROGRESS"],
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		UserId:     "1",
	}
	mockUpdateTransaction = request.UpdateTransactionRequest{
		Id:     "1",
		Status: "accepted",
	}
	mockTransactionId = "1"
)

func TestNewTransactionRepository(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want domain.TransactionRepository
	}{
		{
			name: "success",
			args: args{
				db: db,
			},
			want: &TransactionRepository{
				db: db,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewTransactionRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProductRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTransactionRepository_CreateTransaction(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		transaction domain.Transaction
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
				transaction: mockTransactionEntity,
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectBegin()
				mocks.ExpectQuery(regexp.QuoteMeta(`
					INSERT INTO "transactions" 
                   ("created_at",
					"updated_at",
					"deleted_at",
					"status",
					"bank_number",
					"bank_name",
					"amount",
					"user_id")
					VALUES (?,?,?,?,?,?,?,?) RETURNING "transactions"."id"
					`)).WithArgs(
					sqlmock.AnyArg(),
					transactionStatus["onprogress"],
					"123456789",
					"Mock name",
					10,
					"1",
				).WillReturnError(sqlmock.ErrCancelled)
				mocks.ExpectCommit()
				return db
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := TransactionRepository{
				db: tt.initMock(),
			}
			err := pr.CreateTransaction(tt.args.transaction)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductRepository.CreateProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTransactionRepository_GetTransactionById(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		transactionId string
	}
	tests := []struct {
		name     string
		args     args
		want     *domain.Transaction
		wantErr  bool
		initMock func() *gorm.DB
	}{
		{
			name: "failed with record not found",
			args: args{
				transactionId: "0",
			},
			want:    nil,
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectQuery(`
					SELECT
						*
					FROM 
						"transactions"
					WHERE
						"transactions"."deleted_at" IS NULL AND
						(("transactions"."id" = $1))
				`)
				return db
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := TransactionRepository{
				db: tt.initMock(),
			}
			products, err := pr.GetTransactionById(tt.args.transactionId)
			if err != nil && !tt.wantErr {
				t.Errorf("TransactionRepository.GetTransactionById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(products, tt.want) {
				t.Errorf("TransactionRepository.Find() = %v, want %v", products, tt.want)
			}
		})
	}
}

func TestTransactionRepository_UpdateTransactionStatus(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		productId string
		status    string
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
				status:    transactionStatus["Accepted"],
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectExec(regexp.QuoteMeta(`
					UPDATE "transactions" 
					SET "status" = $1
					WHERE "transactions"."deleted_at" 
						IS NULL AND ((id = 0))
				`)).WithArgs(transactionStatus["Accepted"]).
					WillReturnError(errors.New("No data found"))
				return db
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := TransactionRepository{
				db: tt.initMock(),
			}
			err := pr.UpdateTransactionStatus(tt.args.productId, tt.args.status)
			if err != nil && !tt.wantErr {
				t.Errorf("TransactionRepository.UpdateTransactionStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
