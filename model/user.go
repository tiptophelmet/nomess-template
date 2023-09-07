package model

import (
	"github.com/tiptophelmet/nomess-template/internal/db/orm/sql"
)

type User struct {
	sql.Model
	Email        string
	PasswordHash string
	AuthProvider string
	Verified     bool
}
