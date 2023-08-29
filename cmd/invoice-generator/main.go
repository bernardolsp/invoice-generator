package main

import (
	"bernardolsp/invoice-generator/db"
	"bernardolsp/invoice-generator/internal/customer"
	"bernardolsp/invoice-generator/internal/invoice"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func bootstrap_db(logger log.Logger) (*sql.DB, error) {
	dbConfig := db.Config{
		Logger:   &logger,
		Host:     "localhost",
		Port:     5432,
		User:     "username",
		Password: "password",
		DBName:   "invoice_db",
	}
	logger.Println("Connecting to the database...")
	return db.Connect(dbConfig)
}

func bootstrap_router(logger log.Logger, application *invoice.Application, customer *customer.CustomerStruct) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/invoice", application.Handle).Methods("GET", "POST")
	api.HandleFunc("/customer", customer.Handle).Methods("GET", "POST")
	logger.Println("Starting on port 8080...")
	http.ListenAndServe(":8080", r)
	return r
}

func main() {
	http_logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	db_logger := log.New(os.Stdout, "db: ", log.LstdFlags)
	logger := log.New(os.Stdout, "main: ", log.LstdFlags)

	database, err := bootstrap_db(*db_logger)

	if err != nil {
		logger.Fatal("Could not connect to the database, ", err)
	}

	migrator := db.DatabaseConsumer{
		DB:     database,
		Logger: db_logger,
	}
	err = migrator.Migrate()
	if err != nil {
		logger.Fatal("Could not run migrations, ", err)
	}

	application := &invoice.Application{
		Logger: http_logger,
		DB:     database,
	}
	customer := &customer.CustomerStruct{
		Logger: http_logger,
		DB:     database,
	}

	bootstrap_router(*http_logger, application, customer)
}
