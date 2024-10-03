package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewDB() (*sql.DB, error) {
	dbUser := os.Getenv("TIYULI_EXPENSE_DB_USER")
	dbPass := os.Getenv("TIYULI_EXPENSE_DB_PASS")
	dbAddress := os.Getenv("TIYULI_EXPENSE_DB_ADDRESS")
	dbPort := os.Getenv("TIYULI_EXPENSE_DB_PORT")
	dbName := os.Getenv("TIYULI_EXPENSE_DB_NAME")

	cfg := mysql.Config{
		User:   dbUser,
		Passwd: dbPass,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", dbAddress, dbPort),
		DBName: dbName,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, fmt.Errorf("failed to ping db: %w", pingErr)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
