// consignment-service/repository.go
package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	pb "github.com/loogo/wms/consignment-service/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName                = "shippy"
	consignmentCollection = "consignments"
)

type Repository interface {
	Create(*pb.Consignment) error
	GetAll() ([]*pb.Consignment, error)
	Close()
}

type ConsignmentRepository struct {
	client *mongo.Client
}

// Create a new consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) error {
	res, err := repo.collection().InsertOne(context.Background(), consignment)
	return err
}

// GetAll consignments
func (repo *ConsignmentRepository) GetAll() ([]*pb.Consignment, error) {
	var consignments []*pb.Consignment
	// Find normally takes a query, but as we want everything, we can nil this.
	// We then bind our consignments variable by passing it as an argument to .All().
	// That sets consignments to the result of the find query.
	// There's also a `One()` function for single results.
	cur, err := repo.collection().Find(nil, bson.D{})
	if err != nil {
		return nil, err
	}
	for cur.Next(nil) {
		con := &pb.Consignment{}
		err = cur.Decode(con)
		if err != nil {
			return nil, err
		}
		consignments = append(consignments, con)
	}
	return consignments, err
}

// Close closes the database session after each query has ran.
// Mgo creates a 'master' session on start-up, it's then good practice
// to copy a new session for each request that's made. This means that
// each request has its own database session. This is safer and more efficient,
// as under the hood each session has its own database socket and error handling.
// Using one main database socket means requests having to wait for that session.
// I.e this approach avoids locking and allows for requests to be processed concurrently. Nice!
// But... it does mean we need to ensure each session is closed on completion. Otherwise
// you'll likely build up loads of dud connections and hit a connection limit. Not nice!
func (repo *ConsignmentRepository) Close() {
	repo.client.Disconnect(nil)
}

func (repo *ConsignmentRepository) collection() *mongo.Collection {
	return repo.client.Database(dbName).Collection(consignmentCollection)
}
