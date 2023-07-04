package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID   string
	Name string
	Age  uint
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("test").Collection("users")

	user := User{ID: "1", Name: "John Doe", Age: 25}
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	var queriedUser User
	err = collection.FindOne(context.Background(), bson.M{}).Decode(&queriedUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Queried User:", queriedUser)

	update := bson.M{"$set": bson.M{"Name": "Jane Smith"}}
	_, err = collection.UpdateOne(context.Background(), bson.M{"_id": queriedUser.ID}, update)
	if err != nil {
		log.Fatal(err)
	}

	err = collection.FindOne(context.Background(), bson.M{"_id": queriedUser.ID}).Decode(&queriedUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated User:", queriedUser)
}
