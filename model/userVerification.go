package model

import (
	"github.com/tiptophelmet/nomess-template/internal/db/orm/sql"
)

type UserVerification struct {
	sql.Model
	User User
	Code string
}
