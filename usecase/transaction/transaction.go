package transaction

import (
	"github.com/williamchang80/sea-apd/domain/transaction"
	transaction2 "github.com/williamchang80/sea-apd/dto/request/transaction"
)

type TransactionUsecase struct {
	tr transaction.TransactionRepository
}

func NewTransactionUsecase(repo transaction.TransactionRepository) transaction.TransactionUsecase {
	return &TransactionUsecase{tr: repo}
}

func (t TransactionUsecase) CreateTransaction(request transaction2.TransactionRequest) error {
	panic("implement me")
}

func (t TransactionUsecase) UpdateTransactionStatus(request transaction2.UpdateTransactionRequest) error {
	err := t.tr.UpdateTransactionStatus(request.Status, request.Id)
	return err
}


