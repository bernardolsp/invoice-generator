package customer

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type CustomerStruct struct {
	DB     *sql.DB
	Logger *log.Logger
}

func (c CustomerStruct) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handleGet(w, r)
	} else if r.Method == http.MethodPost {
		handlePost(c.DB, c.Logger, w, r)
	}
}

func handlePost(db *sql.DB, logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	var customer Customer

	// Decode the incoming Customer json
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		logger.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdCustomer, err := Create(db, logger, customer.Name, customer.Email, customer.Address, customer.BillableCurrency)
	if err != nil {
		http.Error(w, "Failed to create customer", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(createdCustomer)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	// ...
}
