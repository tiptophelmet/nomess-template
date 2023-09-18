package response

import (
	"github.com/tiptophelmet/nomess-core/v4/intl"
	"github.com/tiptophelmet/nomess-template/model"
)

type itemCreated struct {
	Message string
	Item    *model.Item
}

func ItemCreated(created *model.Item) *itemCreated {
	return &itemCreated{
		Message: intl.Localize("item_created.message"),
		Item:    created,
	}
}
