package factory

import (
	"github.com/williamchang80/sea-apd/common/constants/mailer_type"
	mailer3 "github.com/williamchang80/sea-apd/common/mailer"
	"github.com/williamchang80/sea-apd/usecase/auth/mailer"
	mailer2 "github.com/williamchang80/sea-apd/usecase/transaction/mailer"
)

type MailFactory interface {
	CreateMail(...interface{}) []mailer3.Mail
}

func CreateMailerFactory(mailType mailer_type.MailType) MailFactory {
	switch mailType {
	case mailer_type.AUTH:
		return &mailer.AuthMailer{}
	case mailer_type.TRANSACTION:
		return &mailer2.TransactionMailer{}
	}
	return nil
}
