package main

import (
	"context"
	"log"
	"net"

	pb "github.com/odedro987/tiyuli-server/expense-api/proto"
	pbExpense "github.com/odedro987/tiyuli-server/expense-api/proto/expense"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedTiyuliServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatalf("failed to listen on port 443: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTiyuliServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) NewExpense(ctx context.Context, in *pbExpense.NewExpenseRequest) (*pbExpense.NewExpenseResponse, error) {
	if in.Amount < 0 {
		return nil, nil
	}
	return &pbExpense.NewExpenseResponse{}, nil
}
