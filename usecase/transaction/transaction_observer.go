package transaction

import (
	"fmt"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/domain/transaction"
	merchant2 "github.com/williamchang80/sea-apd/dto/request/merchant"
)

type UpdateMerchantBalanceObserver struct {
}

func (u *UpdateMerchantBalanceObserver) Update(transaction transaction.Transaction,
	t TransactionUsecase) error {
	var err error
	if transaction_status.ParseToEnum(transaction.Status) == transaction_status.WAITING_DELIVERY {
		err = t.merchantUseCase.UpdateMerchantBalance(merchant2.UpdateMerchantBalanceRequest{
			Amount:     transaction.Amount,
			MerchantId: transaction.UserId,
		})
	}
	return err
}

type NotifyAdminObserver struct {
}

func (n *NotifyAdminObserver) Update(transaction transaction.Transaction) error {
	fmt.Println("UPDATED")
}
