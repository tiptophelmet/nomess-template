package model

import (
	"github.com/tiptophelmet/nomess-core/v5/db/orm/sql"
)

type Item struct {
	sql.Model
	Name        string
	Description string
}
