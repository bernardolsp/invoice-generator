package invoice

func CreateInvoice(inv Invoice) error {
	// Code to save invoice into the database
	return nil
}

func (h *Application) GetInvoices() ([]Invoice, error) {
	// Encapsulate the logic to get invoices from the database
	h.Logger.Println("Getting invoices from the database...")
	rows, err := h.DB.Query("SELECT * FROM invoices")
	if err != nil {
		h.Logger.Println("Error querying the database", err)
		return nil, err
	}
	defer rows.Close()
	var invoices []Invoice
	for rows.Next() {
		var inv Invoice
		err := rows.Scan(&inv.ID, &inv.BeneficiaryID, &inv.Customer.Email,
			&inv.Customer.Name, &inv.InvoiceDate, &inv.ExpireDate)
		if err != nil {
			h.Logger.Println("Error scanning row", err)
			return nil, err
		}
		invoices = append(invoices, inv)
	}
	return invoices, nil
}
