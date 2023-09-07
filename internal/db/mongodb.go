package db

import (
	"context"
	"time"

	"github.com/tiptophelmet/nomess/internal/config"
	"github.com/tiptophelmet/nomess/internal/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client    *mongo.Client
	Ctx       context.Context
	CtxCancel context.CancelFunc
}

func InitMongoDB() *MongoDB {
	uri := config.Get("mongodb-uri").Required().Str()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal(err.Error())
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		logger.Fatal(err.Error())
	}

	return &MongoDB{client, ctx, cancel}
}