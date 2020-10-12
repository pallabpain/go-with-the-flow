package chat

import (
	"log"

	"golang.org/x/net/context"
)

// Server represents a server implementation
type Server struct{}

// SendMessage is the server-side implementation
func (s *Server) SendMessage(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Received a message from the client: %s", in.Body)
	return &Message{Body: "The server received your message."}, nil
}
