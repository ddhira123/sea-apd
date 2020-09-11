package transaction

import (
	"github.com/williamchang80/sea-apd/common/constants/mailer_type"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/common/mailer"
	"github.com/williamchang80/sea-apd/common/mailer/factory"
	"github.com/williamchang80/sea-apd/domain/merchant"
	"github.com/williamchang80/sea-apd/domain/transaction"
	merchant2 "github.com/williamchang80/sea-apd/dto/request/merchant"
)

var mail factory.MailFactory

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
		mail = factory.CreateMailerFactory(mailer_type.TRANSACTION)
		mails := mail.CreateMail(transaction, "customer@customer.com", "merchant@merchant.com")
		err := mailer.SendEmail(mails)
		if err != nil {
			return err
		}
	}
	return nil
}

type SendPaymentInvoiceObserver struct {
}

func (n *SendPaymentInvoiceObserver) Update(transaction transaction.Transaction,
	t merchant.MerchantUsecase) error {
	if transaction_status.ParseToEnum(transaction.Status) == transaction_status.WAITING_CONFIRMATION {
		mail = factory.CreateMailerFactory(mailer_type.TRANSACTION)
		mails := mail.CreateMail(transaction, "customer@customer.com", "merchant@merchant.com")
		err := mailer.SendEmail(mails)
		if err != nil {
			return err
		}
	}
	return nil
}
