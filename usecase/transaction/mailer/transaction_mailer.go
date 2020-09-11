package mailer

import (
	"fmt"
	"github.com/williamchang80/sea-apd/common/constants/transaction_status"
	"github.com/williamchang80/sea-apd/common/mailer"
	"github.com/williamchang80/sea-apd/domain/transaction"
	"strings"
)

type TransactionMailer struct {
}

func (t *TransactionMailer) CreateMail(tr ...interface{}) []mailer.Mail {
	s, _ := tr[0].(transaction.Transaction)
	customerEmail, _ := tr[1].(string)
	merchantEmail, _ := tr[2].(string)
	switch transaction_status.ParseToEnum(s.Status) {
	case transaction_status.WAITING_CONFIRMATION:
		return CreateInvoiceAndNotificationMailer(s, customerEmail, merchantEmail)
	case transaction_status.WAITING_DELIVERY:
		return CreateRequestMailer(s, customerEmail, merchantEmail)
	case transaction_status.ACCEPTED:
		return CreateArrivalMailer(s, customerEmail, merchantEmail)
	}
	return nil
}

func CreateInvoiceAndNotificationMailer(transaction transaction.Transaction, customerEmail string,
	merchantEmail string) []mailer.Mail {
	invoiceMailer := mailer.Mail{
		Sender:    mailer.MailSender,
		Subject:   fmt.Sprintf("Invoice for transaction id %v", transaction.ID),
		Recipient: customerEmail,
		Body: fmt.Sprintf(`Hello, %v thank you for your purchasing, +
			Your transaction details:
				Id: %v
				Total: %v
				Merchant: %v
			We will notify you soon`, strings.TrimSuffix(customerEmail, "@"),
			transaction.ID, transaction.Amount, strings.TrimSuffix(merchantEmail, "@")),
	}
	notificationMailer := mailer.Mail{
		Sender:    mailer.MailSender,
		Subject:   fmt.Sprintf("New transaction with id %v", transaction.ID),
		Recipient: mailer.AdminEmail,
		Body: fmt.Sprintf(`Hello, admin please confirm transaction with id %v
			From merchant %v with customer %v`, transaction.ID,
			strings.TrimSuffix(merchantEmail, "@"),
			strings.TrimSuffix(customerEmail, "@")),
	}
	mailers := []mailer.Mail{
		invoiceMailer,
		notificationMailer,
	}
	return mailers
}

func CreateRequestMailer(transaction transaction.Transaction, customerEmail string,
	merchantEmail string) []mailer.Mail {
	createRequestMailer := mailer.Mail{
		Sender:    mailer.MailSender,
		Subject:   fmt.Sprintf("New Request item from transaction id %v", transaction.ID),
		Recipient: merchantEmail,
		Body: fmt.Sprintf(`Hello, %v please confirm request with transaction id %v
			From customer %v, please check your store`,
			strings.TrimSuffix(merchantEmail, "@"), transaction.ID,
			strings.TrimSuffix(customerEmail, "@"),
		),
	}
	mailers := []mailer.Mail{
		createRequestMailer,
	}
	return mailers
}

func CreateArrivalMailer(transaction transaction.Transaction, customerEmail string,
	merchantEmail string) []mailer.Mail {
	createRequestMailer := mailer.Mail{
		Sender:    mailer.MailSender,
		Subject:   "Item confirmed!",
		Recipient: customerEmail,
		Body: fmt.Sprintf(`Hello, %v your transaction with id %v
			has been confimed by %v and delivered! Please wait for item to arrived`,
			strings.TrimSuffix(customerEmail, "@"),
			transaction.ID, strings.TrimSuffix(merchantEmail, "@"),
		),
	}
	mailers := []mailer.Mail{
		createRequestMailer,
	}
	return mailers
}
