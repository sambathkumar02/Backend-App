package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	//create a raw TCP listener
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal("Unable to listen on port 4000")
	}

	//Get instance of a Server strcut-> implementaion of DockerService
	server := Server{}

	//Create new GRPC server
	grpc_server := grpc.NewServer()

	//Register the GRPC service
	RegisterDockerServiceServer(grpc_server, &server)

	//Server GRPC
	if err := grpc_server.Serve(listener); err != nil {
		log.Fatal("Faied to create GRPC server!")
	}

}
