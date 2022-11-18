package main

import (
	"log"
	"net"

	"github.com/moms-spaghetti/followers/internal/server"
	"github.com/moms-spaghetti/followers/internal/storage"
	pb "github.com/moms-spaghetti/followers/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = ":50050" 
	protocol = "tcp"
)

func main() {
	lis, err := net.Listen(protocol, address)
	if err != nil {
		log.Fatalf("'unable to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := &server.Server{
		MockStorage: storage.CreateMockUserStorage(),
	}
	
	pb.RegisterUserServiceServer(s, srv)
	reflection.Register(s)
	
	log.Print("followers service running")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("unable to serve: %v", err)
	}	

}

