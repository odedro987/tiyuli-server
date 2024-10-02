package error

import (
	pb "github.com/odedro987/tiyuli-server/go-common/proto/error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorInfo = pb.ErrorInfo

func NewStatusWithDetails(code codes.Code, msg string, details *ErrorInfo) error {
	sts := status.New(code, msg)
	newStatus, err := sts.WithDetails(details)
	if err != nil {
		return status.New(codes.Internal, "failed to attach error metadata").Err()
	}

	return newStatus.Err()
}
