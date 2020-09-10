package transaction

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/domain/transaction"
	merchant2 "github.com/williamchang80/sea-apd/dto/request/merchant"
	"os"
	"strings"
	"time"
)

type UpdateMerchantBalanceObserver struct {
}

func (u *UpdateMerchantBalanceObserver) Update(transaction transaction.Transaction,
	t merchant.MerchantUsecase) error {
	var err error
	if transaction_status.ParseToEnum(transaction.Status) == transaction_status.WAITING_DELIVERY {
		err = t.UpdateMerchantBalance(merchant2.UpdateMerchantBalanceRequest{
			Amount:     transaction.Amount,
			MerchantId: transaction.MerchantId,
		})
	}
	return err
}

type NotifyAdminObserver struct {
}

func (n *NotifyAdminObserver) Update(transaction transaction.Transaction,
	t merchant.MerchantUsecase) error {
	if transaction_status.ParseToEnum(transaction.Status) == transaction_status.WAITING_CONFIRMATION {
		API_KEY := os.Getenv("API_KEY")
		DOMAIN_NAME := os.Getenv("DOMAIN_NAME")
		mg := mailgun.NewMailgun(DOMAIN_NAME, API_KEY)
		sender := "test@test.com"
		subject := "Payment Confirmation"
		recipient := "test123@test.com"
		body := fmt.Sprintf("Hello, %v please confirm payment with id %v",
			strings.TrimSuffix(recipient, "@"), transaction.ID)
		message := mg.NewMessage(sender, subject, body, recipient)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		mg.Send(ctx, message)
	}
	return nil
}
