package entity

import (
	"fmt"

	"github.com/go-kit/kit/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	DB *gorm.DB
}

func NewConnection(user, password, dbname string, logger log.Logger) Connection {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", user, password, dbname)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		logger.Log("error", err)
		panic(err)
	}

	logger.Log("Database", "connected")
	return Connection{DB: db}
}
