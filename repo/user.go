package repo

import (
	"github.com/tiptophelmet/nomess/internal/db"
	"github.com/tiptophelmet/nomess/model"

	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	db *db.MongoDB
}

func InitUserRepo() *User {
	return &User{db: db.InitMongoDB()}
}

func (repo *User) Save(model *model.User) (*mongo.InsertOneResult, error) {
	defer repo.db.CtxCancel()
	defer repo.db.Client.Disconnect(repo.db.Ctx)
	collection := repo.db.Client.Database("mydb").Collection("user")

	result, err := collection.InsertOne(repo.db.Ctx, model)
	if err != nil {
		return nil, err
	}

	return result, nil
}
