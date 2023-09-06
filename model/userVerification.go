package model

import (
	"github.com/tiptophelmet/nomess/internal/db/orm/doc/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserVerification struct {
	mongo.Model
	UserID primitive.ObjectID `bson:"user_id"`
	Code   string             `bson:"code"`
}
