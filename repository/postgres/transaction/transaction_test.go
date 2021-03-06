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
		CustomerId: "1",
	}
	mockTransactionEntity = domain.Transaction{
		Status:     transaction_status.ToString(transaction_status.WAITING_PAYMENT),
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		CustomerId: "1",
		MerchantId: "1",
	}
	mockUpdateTransaction = request.UpdateTransactionRequest{
		TransactionId: "1",
		Status:        transaction_status.WAITING_PAYMENT,
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

func TestTransactionRepository_CreateCart(t *testing.T) {
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
					transaction_status.WAITING_PAYMENT,
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
			err := pr.CreateCart(tt.args.transaction)
			if err != nil && !tt.wantErr {
				t.Errorf("ProductRepository.CreateCart() error = %v, wantErr %v", err, tt.wantErr)
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
				status:    transaction_status.ToString(transaction_status.ACCEPTED),
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectExec(regexp.QuoteMeta(`
					UPDATE "transactions" 
					SET "status" = $1
					WHERE "transactions"."deleted_at" 
						IS NULL AND ((id = 0))
				`)).WithArgs(transaction_status.ToString(transaction_status.ACCEPTED)).
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
			_, err := pr.UpdateTransactionStatus(tt.args.productId, tt.args.status)
			if err != nil && !tt.wantErr {
				t.Errorf("TransactionRepository.UpdateTransactionStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTransactionRepository_GetTransactionByRequiredStatus(t *testing.T) {
	db, mocks := mock_psql.Connection()
	defer db.Close()
	type args struct {
		productId      string
		requiredStatus []string
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
				requiredStatus: []string{
					transaction_status.ToString(transaction_status.OTHER),
				},
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				mocks.ExpectExec(regexp.QuoteMeta(`
					UPDATE "transactions" 
					SET "status" = $1
					WHERE "transactions"."deleted_at" 
						IS NULL AND ((id = 0))
				`)).WithArgs(transaction_status.ToString(transaction_status.ACCEPTED)).
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
			_, err := pr.GetTransactionByRequiredStatus(tt.args.requiredStatus, tt.args.productId)
			if err != nil && !tt.wantErr {
				t.Errorf("TransactionRepository.GetTransactionByRequiredStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTransactionRepository_GetMerchantRequestItem(t *testing.T) {
	db, _ := mock_psql.Connection()
	defer db.Close()
	type args struct {
		merchantId     string
		requiredStatus []string
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
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := TransactionRepository{
				db: tt.initMock(),
			}
			_, err := pr.GetMerchantRequestItem(tt.args.merchantId)
			if err != nil && !tt.wantErr {
				t.Errorf("TransactionRepository.GetMerchantRequestItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTransactionRepository_UpdateTransaction(t *testing.T) {
	db, _ := mock_psql.Connection()
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
			name: "fail with invalid id",
			args: args{
				transaction: domain.Transaction{},
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
		{
			name: "success",
			args: args{
				transaction: domain.Transaction{
					Amount: 123,
				},
			},
			wantErr: true,
			initMock: func() *gorm.DB {
				return db
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := TransactionRepository{
				db: tt.initMock(),
			}
			err := pr.UpdateTransaction(tt.args.transaction)
			if err != nil && !tt.wantErr {
				t.Errorf("TransactionRepository.UpdateTransaction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
