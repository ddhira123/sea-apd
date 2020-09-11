package mailer

import (
	"context"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/mailgun/mailgun-go/v4"
	"os"
	"time"
)

const (
	MailSender = "william_chang80@rocketmail.com"
	AdminEmail = "admin@admin.com"
)

type Mail struct {
	Sender    string
	Subject   string
	Recipient string
	Body      string
}

var Mailer *mailgun.MailgunImpl

func InitMail() {
	API_KEY := os.Getenv("API_KEY")
	DOMAIN_NAME := os.Getenv("DOMAIN_NAME")
	if Mailer == nil {
		Mailer = mailgun.NewMailgun(DOMAIN_NAME, API_KEY)
	}
}

func CreateMailer(mail Mail) *mailgun.Message {
	m := Mailer.NewMessage(mail.Sender, mail.Subject, mail.Body, mail.Recipient)
	return m
}

func SendEmail(mails []Mail) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	for _, mail := range mails {
		message := CreateMailer(mail)
		Mailer.Send(ctx, message)
		log.Info(fmt.Sprintf("Mail sent from %v to %v !", mail.Sender, mail.Recipient))
	}
	return nil
}
