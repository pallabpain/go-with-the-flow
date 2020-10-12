package main

import (
	"fmt"
	"log"
	"net"

	"github.com/pallabpain/go-with-the-flow/01-go-grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	address := ":9001"

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer listener.Close()
	fmt.Printf("gRPC server is listening on %s...", address)

	chatServer := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &chatServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error: %v", err)
	}
}
