package response

import (
	"github.com/tiptophelmet/nomess-core/v5/intl"
)

type itemIdAbsent struct {
	Message string
	Text    string
}

func ItemIdAbsent() *itemIdAbsent {
	return &itemIdAbsent{
		Message: intl.Localize("item_id_absent.message"),
		Text:    intl.Localize("item_id_absent.text"),
	}
}
