package response

import (
	"github.com/tiptophelmet/nomess-core/v5/intl"
	"github.com/tiptophelmet/nomess-template/model"
)

type itemDeleted struct {
	Message string
	Item    *model.Item
}

func ItemDeleted() *itemDeleted {
	return &itemDeleted{
		Message: intl.Localize("item_deleted.message"),
	}
}
