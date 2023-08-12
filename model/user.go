package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	AuthProvider string             `bson:"auth_provider"`
	Verified     bool               `bson:"verified"`
}
