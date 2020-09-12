package mailer

type Mailer interface {
	InitMail()
	SendEmail(mails []Mail) error
	CreateMailer(mail Mail) interface{}
}
