package customer

import (
	"bernardolsp/invoice-generator/helpers"
	"database/sql"
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
		c.handlePost(c.DB, c.Logger, w, r)
	}
}

func (c CustomerStruct) handlePost(db *sql.DB, logger *log.Logger, w http.ResponseWriter, r *http.Request) {
	var customer Customer

	if err := helpers.DecodeJSONBody(r, &customer); err != nil {
		logger.Println("Error decoding JSON:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdCustomer, err := c.create(customer.Name, customer.Email, customer.Address, customer.BillableCurrency)
	if err != nil {
		http.Error(w, "Failed to create customer", http.StatusInternalServerError)
		return
	}

	if err := helpers.EncodeJSONResponse(w, createdCustomer); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	// Get customers from the database

}
