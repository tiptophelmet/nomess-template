package response

import (
	"github.com/tiptophelmet/nomess-core/v5/intl"
)

type itemsNotFound struct {
	Message string
	Text    string
}

func ItemsNotFound() *itemsNotFound {
	return &itemsNotFound{
		Message: intl.Localize("items_not_found.message"),
		Text:    intl.Localize("items_not_found.text"),
	}
}
