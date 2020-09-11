package mailer

import (
	"github.com/williamchang80/sea-apd/usecase/auth/mailer"
	mailer2 "github.com/williamchang80/sea-apd/usecase/transaction/mailer"
)

type MailFactory interface {
	CreateMail(...interface{}) []Mail
}

func CreateMailerFactory(mailType MailType) MailFactory {
	switch mailType {
	case AUTH:
		return &mailer.AuthMailer{}
	case TRANSACTION:
		return &mailer2.TransactionMailer{}
	}
	return nil
}
