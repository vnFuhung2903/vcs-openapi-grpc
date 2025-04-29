package client

import (
	"log"

	"github.com/vnFuhung2903/vcs-openapi-grpc/proto"
	"google.golang.org/grpc"
)

func NewClient() proto.BookClient {
	cc, err := grpc.NewClient("localhost:50069")
	if err != nil {
		log.Fatalf("Grpc client error: %v", err)
	}
	defer cc.Close()

	client := proto.NewBookClient(cc)
	return client
}
