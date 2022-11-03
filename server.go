package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/andreaspb/grpc-go-svelte/chat"
)

func main() {
	log.Println("Doing st0ff...")
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil{
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}