package db

import "fmt"

// Receives a (*sql.DB, error)
// Executes a Migrate function
// The function runs a migration

func (d DatabaseConsumer) migrate(table string) error {
	d.Logger.Printf("Running migrations for table %s...", table)

	var sql string

	switch table {
	case "invoice":
		sql = `CREATE TABLE IF NOT EXISTS invoices (
			invoice_id SERIAL PRIMARY KEY,
			customer_id INT REFERENCES customers(customer_id),
			invoice_date DATE NOT NULL,
			invoice_expire_date DATE NOT NULL,
			invoice_value VARCHAR(255) NOT NULL,
			beneficiary_id INT REFERENCES beneficiaries(beneficiary_id)
	);`
	case "customers":
		sql = `CREATE TABLE IF NOT EXISTS customers (
			customer_id SERIAL PRIMARY KEY,
			customer_name VARCHAR(255) NOT NULL,
			customer_email VARCHAR(255),
			customer_address VARCHAR(255) NOT NULL,
			customer_billable_currency VARCHAR(10) NOT NULL,
			customer_added_date DATE NOT NULL,
			customer_last_modified_date DATE NOT NULL
	);`
	case "beneficiaries":
		sql = `CREATE TABLE IF NOT EXISTS beneficiaries (
			beneficiary_id SERIAL PRIMARY KEY,
			beneficiary_name VARCHAR(255) NOT NULL,
			beneficiary_address VARCHAR(255) NOT NULL,
			beneficiary_email VARCHAR(255),
			beneficiary_currency VARCHAR(10) NOT NULL,
			beneficiary_added_date DATE NOT NULL,
			beneficiary_last_modified_date DATE NOT NULL
	);`
	default:
		return fmt.Errorf("unknown table: %s", table)
	}

	result, err := d.DB.Exec(sql)
	if err != nil {
		return fmt.Errorf("error running migration: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error getting rows affected: %w", err)
	}

	d.Logger.Printf("Successfully ran migrations for table %s. %d rows impacted.", table, rows)
	return nil
}

func (d DatabaseConsumer) seeder(table string) error {
	return nil
}
