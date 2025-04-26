package server

import (
	"context"

	proto "github.com/vnFuhung2903/vcs-openapi-grpc/proto"
)

type Server struct {
	proto.UnimplementedBookServer
}

func (s *Server) GetBook(ctx context.Context, req *proto.BookRequest) (*proto.BookResponse, error) {
	if req.Title == "Harry Potter and the Sorcerer's Stone" || req.Title == "" {
		return &proto.BookResponse{
			Title:       "Harry Potter and the Sorcerer's Stone",
			Description: "An orphaned boy enrolls in a school of wizardry, where he learns the truth about himself, his family and the terrible evil that haunts the magical world.",
			Author:      "J.K. Rowling",
			Publisher:   "Bloomsbury",
			Year:        "1997",
		}, nil
	} else {
		return &proto.BookResponse{
			Title: req.Title,
		}, nil
	}
}
