package response

import (
	"github.com/tiptophelmet/nomess-core/v5/intl"
	"github.com/tiptophelmet/nomess-template/model"
)

type itemUpdated struct {
	Message string
	Item    *model.Item
}

func ItemUpdated(updated *model.Item) *itemDeleted {
	return &itemDeleted{
		Message: intl.Localize("item_updated.message"),
		Item:    updated,
	}
}
