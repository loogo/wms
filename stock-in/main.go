package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type message struct {
	Name string
	Age  int
}

type MyDoc struct {
	SomeInt    int    `bson:"some_int"`
	SomeString string `bson:"some_string,omitempty"`
	CustomType MyType `bson:"custom_type,omitempty"`
}

type MyType struct {
	Value string `bson:"value,omitempty"`
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	url := fmt.Sprintf("mongodb://%s:%s@%s:27017", "root", "root", "127.0.0.1")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))

	if err != nil {
		fmt.Println(err)
	}
	db := client.Database("demo")
	fmt.Println("switch database to demo")
	// res, err := db.Collection("demo").InsertOne(ctx, bson.D{
	// 	{"task", "1"},
	// 	{"createAt", "2019-04-02"},
	// })
	msg := &message{"王春雷", 22}
	_, err = db.Collection("demo").InsertOne(ctx, msg)
	// collection := db.Collection("demo")
	// var myType = MyType{Value: "ABCD"}

	// docToInsert := MyDoc{42, "The Answer", myType}

	// _, err = collection.InsertOne(nil, docToInsert)
	if err != nil {
		fmt.Println(err)
	}
	client.Disconnect(nil)
}
