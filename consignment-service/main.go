package main

import (

	// Import the generated protobuf code
	"fmt"
	"log"

	pb "github.com/loogo/wms/consignment-service/proto/consignment"
	vesselProto "github.com/loogo/wms/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"os"
)

const (
	defaultHost = "localhost:27017"
)

func main() {
	// Database host from the environment variables
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = defaultHost
	}
	url := fmt.Sprintf("mongodb://%s:%s@%s", "root", "root", host)
	session, err := CreateSession(url)

	// Mgo creates a 'master' session, we need to end that session
	// before the main function closes.
	defer session.Close()

	if err != nil {

		// We're wrapping the error returned from our CreateSession
		// here to add some context to the error.
		log.Panicf("Could not connect to datastore with host %s - %v", host, err)
	}

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("go.micro.srv.consignment"),
		micro.Version("latest"),
	)

	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", srv.Client())

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
// const (
// 	port = ":50051"
// )

// type IRepository interface {
// 	Create(*pb.Consignment) (*pb.Consignment, error)
// 	GetAll() []*pb.Consignment
// }

// type Repository struct {
// 	consignments []*pb.Consignment
// }

// func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
// 	updated := append(repo.consignments, consignment)
// 	repo.consignments = updated
// 	return consignment, nil
// }
// func (repo *Repository) GetAll() []*pb.Consignment {
// 	return repo.consignments
// }

// type service struct {
// 	repo         Repository
// 	vesselClient vesselProto.VesselService
// }

// func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
// 	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
// 		MaxWeight: req.Weight,
// 		Capacity:  int32(len(req.Containers)),
// 	})

// 	req.VesselId = vesselResponse.Vessel.Id

// 	consignment, err := s.repo.Create(req)
// 	if err != nil {
// 		return err
// 	}

// 	res.Created = true
// 	res.Consignment = consignment
// 	return nil
// }

// func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
// 	consignments := s.repo.GetAll()
// 	res.Consignments = consignments
// 	return nil
// }

//func main() {
//	session, err := CreateSession()
//
//	// Mgo creates a 'master' session, we need to end that session
//	// before the main function closes.
//	defer session.Close()
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// lis, err := net.Listen("tcp", port)
//	// if err != nil {
//	// 	log.Fatalf("failed to listen: %v", err)
//	// }
//	// s := grpc.NewServer()
//
//	// pb.RegisterShippingServiceServer(s, &service{repo})
//	// reflection.Register(s)
//	// if err := s.Serve(lis); err != nil {
//	// 	log.Fatalf("failed to serve: %v", err)
//	// }
//	srv := micro.NewService(
//		micro.Name("go.micro.srv.consignment"),
//		micro.RegisterTTL(time.Second*30),
//		micro.RegisterInterval(time.Second*10),
//		micro.Version("latest"),
//	)
//
//	vesselClient := vesselProto.NewVesselService("go.micro.srv.vessel", srv.Client())
//
//	srv.Init()
//
//	pb.RegisterShippingServiceHandler(srv.Server(), &service{session, vesselClient})
//
//	if err := srv.Run(); err != nil {
//		fmt.Println(err)
//	}
//}
