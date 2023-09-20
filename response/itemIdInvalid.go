package response

import (
	"github.com/tiptophelmet/nomess-core/v5/intl"
)

type itemIdInvalid struct {
	Message string
	Text    string
}

func ItemIdInvalid() *itemIdInvalid {
	return &itemIdInvalid{
		Message: intl.Localize("item_id_invalid.message"),
		Text:    intl.Localize("item_id_invalid.text"),
	}
}
