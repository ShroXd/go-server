package db

import (
	"context"
	"github.com/qiniu/qmgo"
	"log"
)

var (
	Client *qmgo.Client
	Mongo  *qmgo.Database
)

const (
	MongoDBUrl = "mongodb://49.232.5.176:34541"
)

func Connect() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: MongoDBUrl, Database: "dev", Coll: "books"})

	//defer func() {
	//	if err = client.Close(ctx); err != nil {
	//		panic(err)
	//	}
	//}()

	if err != nil {
		log.Fatal(err)
	}
	Client = client
	Mongo = client.Database("dev")
}
