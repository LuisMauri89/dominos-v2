package entity

import (
	"database/sql"
	"fmt"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/lib/pq"
)

// Connection - holds an open connection to database.
type Connection struct {
	DB *sql.DB
}

// NewConnection - returns new Connection based on arguments provided.
func NewConnection(user, password, dbname string, logger log.Logger) Connection {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		logger = level.NewFilter(logger, level.AllowInfo())
		level.Error(logger).Log("db connection error", err)
		panic(err)
	}

	logger.Log("Database", "connected")
	return Connection{DB: connection}
}
