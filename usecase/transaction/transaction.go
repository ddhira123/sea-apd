package transaction

import (
	"errors"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/domain/transaction"
	merchant2 "github.com/williamchang80/sea-apd/dto/request/merchant"
	transaction2 "github.com/williamchang80/sea-apd/dto/request/transaction"
	"strings"
)

var transactionStatus = transaction_status.GetTransactionStatus()

type TransactionUsecase struct {
	tr              transaction.TransactionRepository
	merchantUseCase merchant.MerchantUsecase
}

func ConvertToDomain(t transaction2.TransactionRequest) transaction.Transaction {
	return transaction.Transaction{
		Status:     transactionStatus["ONPROGRESS"],
		BankNumber: t.BankNumber,
		BankName:   t.BankName,
		Amount:     t.Amount,
		UserId:     t.UserId,
	}
}

func NewTransactionUsecase(repo transaction.TransactionRepository, merchantUseCase merchant.MerchantUsecase) transaction.TransactionUsecase {
	return &TransactionUsecase{tr: repo, merchantUseCase: merchantUseCase}
}

func (t TransactionUsecase) CreateTransaction(request transaction2.TransactionRequest) error {
	tran := ConvertToDomain(request)
	err := t.tr.CreateTransaction(tran)
	return err
}

func (t TransactionUsecase) UpdateTransactionStatus(request transaction2.UpdateTransactionRequest) error {
	status := transactionStatus[strings.ToUpper(request.Status)]
	if len(status) == 0 {
		return errors.New("cannot find status")
	}
	tran, err := t.tr.UpdateTransactionStatus(status, request.TransactionId)
	if err != nil{
		return err
	}
	if status == transactionStatus["ACCEPTED"] {
		t.merchantUseCase.UpdateMerchantBalance(merchant2.UpdateMerchantBalanceRequest{
			Amount: tran.Amount,
			MerchantId: tran.UserId,
		})
	}
	return err
}

func (t TransactionUsecase) GetTransactionById(id string) (*transaction.Transaction, error) {
	tr, err := t.tr.GetTransactionById(id)
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func (t TransactionUsecase) GetTransactionHistory(userId string) ([]transaction.Transaction, error) {
	requiredStatusForTransactionHistory := transaction_status.GetRequiredStatus()
	tr, err := t.tr.GetTransactionByRequiredStatus(requiredStatusForTransactionHistory, userId)
	if err != nil {
		return nil, err
	}
	return tr, nil
}
