// db/db.go
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect(c Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName)

	return sql.Open("postgres", connStr)
}

func (d DatabaseConsumer) Migrate() error {
	tables := []string{"invoice", "customers", "beneficiaries"}

	for _, table := range tables {
		if err := d.migrate(table); err != nil {
			return err
		}

		if err := d.seeder(table); err != nil {
			return err
		}
	}
	return nil
}
