package mailer

import (
	"fmt"
	"github.com/williamchang80/sea-apd/common/constants/user_role"
	"github.com/williamchang80/sea-apd/common/mailer"
)

type AuthMailer struct {
}

func (a AuthMailer) CreateMail(i ...interface{}) []mailer.Mail {
	id, _ := i[0].(string)
	senderStatus, _ := i[1].(user_role.UserRole)
	switch senderStatus {
	case user_role.CUSTOMER:
		return CreateMerchantProposalMailer(id)
	}
	return nil
}

func CreateMerchantProposalMailer(id string) []mailer.Mail {
	merchantProposalMailer := mailer.Mail{
		Sender:    mailer.MailSender,
		Subject:   fmt.Sprintf("New Merchant Proposal Request with user id %v", id),
		Recipient: mailer.AdminEmail,
		Body: fmt.Sprintf(`Hello, admin there are new Merchant Proposal Request with 
		user id %v, please check the application for details`, id),
	}
	mailers := []mailer.Mail{
		merchantProposalMailer,
	}
	return mailers
}
