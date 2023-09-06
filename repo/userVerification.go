package repo

import (
	"github.com/tiptophelmet/nomess/internal/db"
	"github.com/tiptophelmet/nomess/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserVerification struct {
	db *db.MongoDB
}

func InitUserVerificationRepo() *UserVerification {
	return &UserVerification{db: db.InitMongoDB()}
}

func (repo *UserVerification) Save(model *model.UserVerification) (*mongo.InsertOneResult, error) {
	defer repo.db.CtxCancel()
	defer repo.db.Client.Disconnect(repo.db.Ctx)
	collection := repo.db.Client.Database("mydb").Collection("userVerification")

	result, err := collection.InsertOne(repo.db.Ctx, model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
