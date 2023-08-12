package email

import (
	"github.com/tiptophelmet/nomess/config"
	"github.com/tiptophelmet/nomess/logger"
)

type Email interface {
	Send(mailTo string, template string, data string, mailFrom string) (bool, error)
}

func InitEmail() Email {
	mailer, err := config.Str("mailer")
	if err != nil {
		logger.Alert("could not resolve mailer")
	}

	switch mailer {
	case "ses":
		return InitSES()
	default:
		return InitSES()
	}
}
