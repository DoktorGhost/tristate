package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/DoktorGhost/tristate/reflection/gen/go/example/gen/data"
	"github.com/DoktorGhost/tristate/reflection/gen/go/example/gen/hello"
)

// server реализует сервис HelloService из hello.proto
type server struct {
	hello.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, req *data.Data) (*data.Data, error) {
	log.Printf("Got SayHello request: %v", req)
	// Просто возвращаем тот же объект
	return req, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hello.RegisterHelloServiceServer(s, &server{})

	reflection.Register(s)

	log.Println("gRPC server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
