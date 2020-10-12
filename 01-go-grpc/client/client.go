package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/pallabpain/go-with-the-flow/01-go-grpc/chat"

	"google.golang.org/grpc"
)

func main() {
	// gRPC server address to conenct
	address := "localhost:9001"
	opts := grpc.WithInsecure()

	conn, err := grpc.Dial(address, opts)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	defer conn.Close()

	client := chat.NewChatServiceClient(conn)

	message := &chat.Message{Body: os.Args[1]}

	response, err := client.SendMessage(context.Background(), message)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("[Reply from server]: %s\n", response.Body)
}
