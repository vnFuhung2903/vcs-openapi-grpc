package main

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	"github.com/vnFuhung2903/vcs-openapi-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func checkGetBook(client proto.BookClient, req *proto.BookRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetBook(ctx, req)
	if err != nil {
		log.Fatalf("GetBook rpc error %v", err)
	}
	log.Println("GetBook rpc response:", res.GetTitle())
}

func checkListBooks(client proto.BookClient, req *proto.BookRequest) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := client.ListBooks(ctx, req)
	if err != nil {
		log.Fatalf("ListBooks rpc error %v", err)
	}

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				return
			}
			if err != nil {
				log.Fatalf("ListBooks rpc error %v", err)
			}
			log.Println("ListBooks rpc response:", res.GetTitle())
		}
	}()
}

func checkMultiGetBook(client proto.BookClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reqs := []*proto.BookRequest{
		{Chapter: "2"},
		{Chapter: "3"},
		{Chapter: "4"},
	}
	stream, err := client.MultiGetBook(ctx)
	if err != nil {
		log.Fatalf("MultiGetBook rpc error: %v", err)
	}

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			log.Fatalf("MultiGetBook rpc error: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("MultiGetBook rpc error: %v", err)
	}
	log.Println("MultiGetBook rpc response:", res.GetTitle())
}

func checkMultiListBooks(client proto.BookClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reqs := []*proto.BookRequest{
		{Chapter: "5"},
		{Chapter: "6"},
		{Chapter: "7"},
	}
	stream, err := client.MultiListBooks(ctx)
	if err != nil {
		log.Fatalf("MultiListBooks rpc error: %v", err)
	}

	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				wg.Done()
				return
			}
			if err != nil {
				log.Fatalf("MultiListBooks rpc error: %v", err)
			}
			log.Println("MultiListBooks rpc response:", res.GetTitle())
		}
	}()

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			log.Fatalf("Multi List Books error: %v", err)
		}
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Multi List Books error: %v", err)
	}
}

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	cc, err := grpc.NewClient("localhost:50051", opts...)
	if err != nil {
		log.Fatalf("Grpc client error: %v", err)
	}
	defer cc.Close()

	client := proto.NewBookClient(cc)
	req := proto.BookRequest{
		Chapter: "1",
	}
	checkGetBook(client, &req)
	checkListBooks(client, &req)
	checkMultiGetBook(client)
	checkMultiListBooks(client)
}
