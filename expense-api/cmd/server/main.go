package main

import (
	"context"
	"log"
	"net"

	pb "github.com/odedro987/tiyuli-server/expense-api/proto"
	"github.com/odedro987/tiyuli-server/go-common/pkg/auth"
	grpcError "github.com/odedro987/tiyuli-server/go-common/pkg/error"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedTiyuliServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(auth.UnaryInterceptor))
	reflection.Register(s)
	pb.RegisterTiyuliServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) NewExpense(ctx context.Context, in *pb.NewExpenseRequest) (*pb.NewExpenseResponse, error) {
	if in.Amount < 0 {
		return nil, grpcError.NewStatusWithDetails(
			codes.InvalidArgument,
			"amount should be positive",
			&grpcError.ErrorInfo{ErrorCode: "INVALID_AMOUNT"},
		).Err()
	}
	return &pb.NewExpenseResponse{}, nil
}
