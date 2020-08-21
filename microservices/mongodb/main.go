package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Customer struct {
	FirstName  string
	LastName   string
	CreditCard CreditCard
}

type CreditCard struct {
	Number string
}

func main() {
	client, err := mongo.NewClient(
		options.Client().
			ApplyURI("mongodb://mongorootuser:mongorootpw@localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	collection := client.Database("shop").Collection("customer")
	customer := Customer{FirstName: "Hans", LastName: "wurst"}
	customer.CreditCard = CreditCard{Number: "123-123-123"}
	collection.InsertOne(ctx, customer)

	cursor, err := collection.Find(ctx, bson.D{{"firstname", "Hans"}})
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			panic(err)
		}
		fmt.Printf("result is %v\n", result)
	}
	if err := cursor.Err(); err != nil {
		panic(err)
	}

}
