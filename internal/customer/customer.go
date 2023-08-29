package customer

import (
	"fmt"
	"time"
)

type Customer struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	Email            *string   `json:"email"`
	Address          string    `json:"address"`
	BillableCurrency string    `json:"billable_currency"`
	AddedDate        time.Time `json:"added_date"`
	LastModifiedDate time.Time `json:"last_modified_date"`
}

func String(c Customer) string {
	return fmt.Sprintf("Customer %d: %s", c.ID, c.Name)
}
