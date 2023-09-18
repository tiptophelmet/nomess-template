package model

import (
	"github.com/tiptophelmet/nomess-core/v3/db/orm/sql"
)

type Item struct {
	sql.Model
	Name        string
	Description string
}
