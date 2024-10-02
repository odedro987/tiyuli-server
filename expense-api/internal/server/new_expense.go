package server

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/odedro987/tiyuli-server/expense-api/internal/db"
	pb "github.com/odedro987/tiyuli-server/expense-api/proto"
	"github.com/odedro987/tiyuli-server/go-common/pkg/auth"
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

	// connect to db and insert expense
	now := time.Now().Unix()
	result, err := s.db.Exec(db.NewExpenseQuery, ctx.Value(auth.UserId{}), in.Name, in.Note, strings.Join(in.Types, ","), now, in.CurrencyCode, in.Amount)
	if err != nil {
		log.Println(err)
		return nil, grpcError.NewStatusWithDetails(
			codes.Internal,
			"failed to create expense",
			&grpcError.ErrorInfo{ErrorCode: "INSERT_ERROR"},
		)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}

	return &pb.NewExpenseResponse{Id: strconv.FormatInt(id, 10)}, nil
}
