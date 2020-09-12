package mailer

import (
	"fmt"
	"github.com/williamchang80/sea-apd/common/constants/user_role"
	"github.com/williamchang80/sea-apd/common/mailer"
	"github.com/williamchang80/sea-apd/domain/user"
)

type AuthMailer struct {
}

func (a AuthMailer) CreateMail(i ...interface{}) []mailer.Mail {
	u, _ := i[0].(user.User)
	role := user_role.ParseToEnum(i[1].(string))
	switch role {
	case user_role.MERCHANT:
		return CreateMerchantProposalMailer(u)
	}
	return nil
}

func CreateMerchantProposalMailer(u user.User) []mailer.Mail {
	merchantProposalMailer := mailer.Mail{
		Sender:    mailer.MailSender,
		Subject:   fmt.Sprintf("New Merchant Proposal Request with user id %v", u.ID),
		Recipient: mailer.AdminEmail,
		Body: fmt.Sprintf(`Hello, admin there are new Merchant Proposal Request with 
		user id %v and name %v, please check the application for details`, u.ID, u.Name),
	}
	mailers := []mailer.Mail{
		merchantProposalMailer,
	}
	return mailers
}
