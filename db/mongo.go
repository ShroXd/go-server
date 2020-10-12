package db

import (
	"context"
	"github.com/qiniu/qmgo"
	"log"

	"go-server/setting"
)

var (
	Client     *qmgo.Client
	Mongo      *qmgo.Database
	MongoDBUrl = setting.LoadMongoDB()
)

func Connect() {
	ctx := context.Background()
	config := &qmgo.Config{Uri: MongoDBUrl, Database: "dev", Coll: "books"}
	client, err := qmgo.NewClient(ctx, config)

	if err != nil {
		log.Fatal(err)
	}
	Client = client
	Mongo = client.Database("dev")
}
