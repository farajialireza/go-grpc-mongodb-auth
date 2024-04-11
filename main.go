package main

import (
	"log"
	"net"

	"go-grpc-mongodb-auth/grpcstuff/pb"

	"google.golang.org/grpc"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
}

func main() {
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, authServer{})
	lis, err := net.Listen("tcp", ":1111")
	if err != nil {
		log.Fatal("Error creating listener: ", err.Error())
	}

	log.Fatal("Serving gRPC: ", s.Serve(lis))

}
