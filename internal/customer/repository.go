package customer

import (
	"database/sql"
	"log"
	"time"
)

func Create(db *sql.DB, logger *log.Logger, name string, email *string, address string, billableCurrency string) (Customer, error) {
	var customer Customer

	// Prepare the query
	query := `INSERT INTO customers (customer_name, customer_email, customer_address, customer_billable_currency, 
						customer_added_date, customer_last_modified_date)
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING customer_id;`

	// Execute the query
	err := db.QueryRow(query, name, email, address, billableCurrency, time.Now(), time.Now()).Scan(&customer.ID)
	if err != nil {
		logger.Println("Error inserting into database:", err)
		return customer, err
	}

	// Populate the rest of the customer fields
	customer.Name = name
	customer.Email = email
	customer.Address = address
	customer.BillableCurrency = billableCurrency
	customer.AddedDate = time.Now()
	customer.LastModifiedDate = time.Now()

	logger.Println("Successfully created customer:", customer)

	return customer, nil
}
