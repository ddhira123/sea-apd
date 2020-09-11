package mailer_type

type MailType int

const (
	TRANSACTION MailType = iota
	AUTH
)
