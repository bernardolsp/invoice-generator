package db

import (
	"database/sql"
	"log"
)

type Config struct {
	Logger   *log.Logger
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type DatabaseConsumer struct {
	DB     *sql.DB
	Logger *log.Logger
}
