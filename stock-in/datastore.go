// consignment-service/datastore.go
package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "gopkg.in/mgo.v2"
)

// CreateSession creates the main session to our mongodb instance
func CreateSession() (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:27017", "root", "root", "127.0.0.1")
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.NewClient(clientOptions)
	// if err != nil {
	// 	return nil, err
	// }
	// fmt.Println(client)

	// session, err := mgo.Dial(url)

	if err != nil {
		return nil, err
	}

	// session.SetMode(mgo.Monotonic, true)

	return client, nil
}
