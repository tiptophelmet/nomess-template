package model

import (
	"github.com/tiptophelmet/nomess/db/orm/doc/mongo"
)

type User struct {
	mongo.Model
	Email        string `bson:"email"`
	PasswordHash string `bson:"password_hash"`
	AuthProvider string `bson:"auth_provider"`
	Verified     bool   `bson:"verified"`
}
