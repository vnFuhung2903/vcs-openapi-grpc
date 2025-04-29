package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"

	proto "github.com/vnFuhung2903/vcs-openapi-grpc/proto"
	"google.golang.org/grpc"
)

type Book struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Author      string `json:"author,omitempty"`
	Publisher   string `son:"publisher,omitempty"`
	Year        string `json:"year,omitempty"`
}

type Server struct {
	proto.UnimplementedBookServer
}

func NewServer() {
	s := grpc.NewServer()
	proto.RegisterBookServer(s, &Server{})
}

func RetrieveBooks() []Book {
	file, err := os.Open("book.json")
	if err != nil {
		log.Fatalf("Open file error %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Read file error %v", err)
	}

	var books []Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		log.Fatalf("Unmarshal JSON file error %v", err)
	}

	return books
}

func (s *Server) GetBook(ctx context.Context, req *proto.BookRequest) (*proto.BookResponse, error) {
	books := RetrieveBooks()
	index, err := strconv.Atoi(req.Chapter)
	if err != nil {
		return nil, err
	}

	return &proto.BookResponse{
		Title:       books[index-1].Title,
		Description: books[index-1].Description,
		Author:      books[index-1].Author,
		Publisher:   books[index-1].Publisher,
		Year:        books[index-1].Year,
	}, nil
}

func (s *Server) ListBooks(_ *proto.BookRequest, stream proto.Book_ListBooksServer) error {
	books := RetrieveBooks()
	for _, book := range books {
		res := &proto.BookResponse{
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
			Publisher:   book.Publisher,
			Year:        book.Year,
		}
		err := stream.Send(res)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Server) MultiGetBook(stream proto.Book_MultiGetBookServer) error {
	books := RetrieveBooks()
	cntReq := 0
	for {
		req, err := stream.Recv()
		if err != nil && err != io.EOF {
			return err
		}

		index, err := strconv.Atoi(req.Chapter)
		if err != nil {
			return err
		}

		if err == io.EOF {
			return stream.SendAndClose(&proto.BookResponse{
				Title:       books[index-1].Title,
				Description: books[index-1].Description,
				Author:      books[index-1].Author,
				Publisher:   books[index-1].Publisher,
				Year:        books[index-1].Year,
			})
		}

		cntReq += 1
		if cntReq >= 7 {
			return stream.SendAndClose(&proto.BookResponse{
				Title:       books[cntReq-1].Title,
				Description: books[cntReq-1].Description,
				Author:      books[cntReq-1].Author,
				Publisher:   books[cntReq-1].Publisher,
				Year:        books[cntReq-1].Year,
			})
		}
	}
}

func (s *Server) MultiListBooks(stream proto.Book_MultiListBooksServer) error {
	books := RetrieveBooks()
	for {
		req, err := stream.Recv()
		if err != nil && err != io.EOF {
			return err
		}

		index, err := strconv.Atoi(req.Chapter)
		if err != nil {
			return err
		}

		if err == io.EOF {
			break
		}

		err = stream.Send(&proto.BookResponse{
			Title:       books[index-1].Title,
			Description: books[index-1].Description,
			Author:      books[index-1].Author,
			Publisher:   books[index-1].Publisher,
			Year:        books[index-1].Year,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
