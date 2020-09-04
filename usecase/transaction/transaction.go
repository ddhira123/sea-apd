package transaction

import (
	"errors"
	"github.com/williamchang80/sea-apd/common/constants/payment_status"
	"github.com/williamchang80/sea-apd/domain/transaction"
	transaction2 "github.com/williamchang80/sea-apd/dto/request/transaction"
	"strings"
)

var paymentStatus = payment_status.GetPaymentStatus()

type TransactionUsecase struct {
	tr transaction.TransactionRepository
}

func ConvertToDomain(t transaction2.TransactionRequest) transaction.Transaction {
	return transaction.Transaction{
		Status:     paymentStatus["ONPROGRESS"],
		BankNumber: t.BankNumber,
		BankName:   t.BankName,
		Amount:     t.Amount,
		UserId:     t.UserId,
	}
}

func NewTransactionUsecase(repo transaction.TransactionRepository) transaction.TransactionUsecase {
	return &TransactionUsecase{tr: repo}
}

func (t TransactionUsecase) CreateTransaction(request transaction2.TransactionRequest) error {
	tran := ConvertToDomain(request)
	err := t.tr.CreateTransaction(tran)
	return err
}

func (t TransactionUsecase) UpdateTransactionStatus(request transaction2.UpdateTransactionRequest) error {
	status := paymentStatus[strings.ToUpper(request.Status)]
	if len(status) == 0 {
		return errors.New("Cannot find status")
	}
	err := t.tr.UpdateTransactionStatus(status, request.Id)
	return err
}
