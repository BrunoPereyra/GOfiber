package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb+srv://bruno:Ej622XsRFaiuh9kg@cluster0.0leyfts.mongodb.net/TDD?retryWrites=true&w=majority")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connect")
	return client
}
