package server

import (
	"context"

	pb "github.com/odedro987/tiyuli-server/expense-api/proto"
	grpcError "github.com/odedro987/tiyuli-server/go-common/pkg/error"
	"google.golang.org/grpc/codes"
)

func (s *Server) NewExpense(ctx context.Context, in *pb.NewExpenseRequest) (*pb.NewExpenseResponse, error) {
	if in.Amount < 0 {
		return nil, grpcError.NewStatusWithDetails(
			codes.InvalidArgument,
			"amount should be positive",
			&grpcError.ErrorInfo{ErrorCode: "INVALID_AMOUNT"},
		)
	}

	if in.Name == "" {
		return nil, grpcError.NewStatusWithDetails(
			codes.InvalidArgument,
			"name must be defined",
			&grpcError.ErrorInfo{ErrorCode: "INVALID_NAME"},
		)
	}

	if len(in.Types) <= 0 {
		return nil, grpcError.NewStatusWithDetails(
			codes.InvalidArgument,
			"expense must have at least one type",
			&grpcError.ErrorInfo{ErrorCode: "INVALID_TYPES"},
		)
	}

	return &pb.NewExpenseResponse{Id: "1213"}, nil
}
