package server

import (
	"database/sql"
	"fmt"

	"github.com/odedro987/tiyuli-server/expense-api/internal/db"
	pb "github.com/odedro987/tiyuli-server/expense-api/proto"
)

type Server struct {
	pb.UnimplementedExpenseServiceServer
	db *sql.DB
}

func NewServer() (*Server, error) {
	s := Server{}

	db, err := db.NewDB()
	if err != nil {
		return nil, fmt.Errorf("failed to init server db: %w", err)
	}
	s.db = db

	return &s, nil
}
