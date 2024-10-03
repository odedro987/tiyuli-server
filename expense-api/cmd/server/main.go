package main

import (
	"log"
	"log/slog"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/odedro987/tiyuli-server/expense-api/internal/server"
	pb "github.com/odedro987/tiyuli-server/expense-api/proto"
	"github.com/odedro987/tiyuli-server/go-common/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Warn("failed loading environment variables", "err", err)
	}

	port := os.Getenv("TIYULI_EXPENSE_API_PORT")
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", port, err)
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
