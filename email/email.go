package email

import (
	"github.com/tiptophelmet/nomess/config"
)

type Email interface {
	Send(mailTo string, template string, data string, mailFrom string) (bool, error)
}

func InitEmail() Email {
	mailDriver := config.Get("mail.driver").Required().Str()

	switch mailDriver {
	case "ses":
		return InitSES()
	default:
		return InitSES()
	}
}
