package main

import (
	"log"
	"net"

	"github.com/odedro987/tiyuli-server/expense-api/internal/server"
	pb "github.com/odedro987/tiyuli-server/expense-api/proto"
	"github.com/odedro987/tiyuli-server/go-common/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(auth.UnaryInterceptor))
	reflection.Register(s)

	expenseServer, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterExpenseServiceServer(s, expenseServer)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
