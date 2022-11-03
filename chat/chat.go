package chat

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Recived message body from client: %s", message.Body)
	return &Message{Body: fmt.Sprintf("Hello from the server!, your message was: %s", message.Body)}, nil
}
