package invoice

import (
	"bernardolsp/invoice-generator/internal/customer"
	"fmt"
	"time"
)

type Invoice struct {
	ID            int
	Customer      customer.Customer // New field
	InvoiceDate   time.Time
	ExpireDate    time.Time
	Value         string
	BeneficiaryID int
}

func (i Invoice) String() string {
	return fmt.Sprintf("Invoice ID: %d, Customer Name: %s, Invoice Date: %v, Expire Date: %v, Value: %s, Beneficiary ID: %d",
		i.ID, i.Customer.Name, i.InvoiceDate, i.ExpireDate, i.Value, i.BeneficiaryID)
}
