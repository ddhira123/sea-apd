package transaction

import (
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/domain/transaction"
	merchant2 "github.com/williamchang80/sea-apd/dto/request/merchant"
	transaction2 "github.com/williamchang80/sea-apd/dto/request/transaction"
)

type TransactionUsecase struct {
	tr              transaction.TransactionRepository
	merchantUseCase merchant.MerchantUsecase
}

func convertTransactionRequestToDomain(t transaction2.TransactionRequest) transaction.Transaction {
	return transaction.Transaction{
		Status:     transaction_status.ToString(transaction_status.WAITING_CONFIRMATION),
		BankNumber: t.BankNumber,
		BankName:   t.BankName,
		Amount:     t.Amount,
		UserId:     t.UserId,
	}
}

func NewTransactionUsecase(repo transaction.TransactionRepository,
	merchantUseCase merchant.MerchantUsecase) transaction.TransactionUsecase {
	return &TransactionUsecase{tr: repo, merchantUseCase: merchantUseCase}
}

func (t TransactionUsecase) CreateTransaction(request transaction2.TransactionRequest) error {
	tran := convertTransactionRequestToDomain(request)
	err := t.tr.CreateTransaction(tran)
	return err
}

func (t TransactionUsecase) UpdateTransactionStatus(request transaction2.
UpdateTransactionRequest) error {
	status := transaction_status.ToString(request.Status)
	tran, err := t.tr.UpdateTransactionStatus(status, request.TransactionId)
	if err != nil {
		return err
	}
	if request.Status == transaction_status.WAITING_DELIVERY {
		t.merchantUseCase.UpdateMerchantBalance(merchant2.UpdateMerchantBalanceRequest{
			Amount:     tran.Amount,
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

func (t TransactionUsecase) GetTransactionHistory(userId string) ([]transaction.
Transaction, error) {
	requiredStatusForTransactionHistory := transaction_status.GetStatusListForTransactionHistory()
	tr, err := t.tr.GetTransactionByRequiredStatus(requiredStatusForTransactionHistory, userId)
	if err != nil {
		return nil, err
	}
	return tr, nil
}
