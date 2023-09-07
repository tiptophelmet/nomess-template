package model

import (
	"github.com/tiptophelmet/nomess-core/db/orm/sql"
)

type UserVerification struct {
	sql.Model
	User User
	Code string
}
