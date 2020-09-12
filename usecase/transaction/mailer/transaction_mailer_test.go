package mailer

import (
	"github.com/golang/mock/gomock"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/common/mailer"
	"github.com/williamchang80/sea-apd/domain/transaction"
	"reflect"
	"testing"
)

var (
	mockConfirmationTransactionEntity = transaction.Transaction{
		Status:     transaction_status.ToString(transaction_status.WAITING_CONFIRMATION),
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		CustomerId: "1",
		MerchantId: "1",
	}
	mockAcceptedTransactionEntity = transaction.Transaction{
		Status:     transaction_status.ToString(transaction_status.ACCEPTED),
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		CustomerId: "1",
		MerchantId: "1",
	}
	mockDeliveryTransactionEntity = transaction.Transaction{
		Status:     transaction_status.ToString(transaction_status.WAITING_DELIVERY),
		BankNumber: "123456789",
		BankName:   "Mock Bank",
		Amount:     10000,
		CustomerId: "1",
		MerchantId: "1",
	}
	mockCustomerEmail = "test1@test.com"
	mockMerchantEmail = "test@test.com"
)

func TestTransactionMailer_CreateMail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		transaction   transaction.Transaction
		customerEmail string
		merchantEmail string
	}
	tests := []struct {
		name string
		args args
		want []mailer.Mail
	}{
		{
			name: "success with accepted transaction",
			args: args{
				transaction:   mockAcceptedTransactionEntity,
				customerEmail: mockCustomerEmail,
				merchantEmail: mockMerchantEmail,
			},
			want: []mailer.Mail{},
		},
		{
			name: "success with on delivery transaction",
			args: args{
				transaction:   mockDeliveryTransactionEntity,
				customerEmail: mockCustomerEmail,
				merchantEmail: mockMerchantEmail,
			},
			want: []mailer.Mail{},
		},
		{
			name: "success with waiting confirmation transaction",
			args: args{
				transaction:   mockConfirmationTransactionEntity,
				customerEmail: mockCustomerEmail,
				merchantEmail: mockMerchantEmail,
			},
			want: []mailer.Mail{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := TransactionMailer{}
			if got := c.CreateMail(tt.args.transaction,
				tt.args.customerEmail, tt.args.merchantEmail);
				reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Errorf("CreateMail() = %v, want %v", got, tt.want)
			}
		})
	}
}
