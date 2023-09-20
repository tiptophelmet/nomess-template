package response

import (
	"github.com/tiptophelmet/nomess-core/v5/intl"
)

type itemNotFound struct {
	Message string
	Text    string
}

func ItemNotFound() *itemNotFound {
	return &itemNotFound{
		Message: intl.Localize("item_not_found.message"),
		Text:    intl.Localize("item_not_found.text"),
	}
}
