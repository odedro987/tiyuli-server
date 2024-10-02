package server

import (
	pb "github.com/odedro987/tiyuli-server/expense-api/proto"
)

type Server struct {
	pb.UnimplementedExpenseServiceServer
}
