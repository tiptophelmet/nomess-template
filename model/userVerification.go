package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserVerification struct {
	ID     primitive.ObjectID `bson:"_id"`
	UserID primitive.ObjectID `bson:"user_id"`
	Code   string             `bson:"code"`
}
