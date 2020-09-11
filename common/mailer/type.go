package mailer

type MailType int

const (
	TRANSACTION MailType = iota
	AUTH
)
