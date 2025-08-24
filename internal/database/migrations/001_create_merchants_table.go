package migrations

import (
	"database/sql"
)

func init() {
	Migrations = append(Migrations, &Migration{
		Version: 1,
		Up: func(tx *sql.Tx) error {
			_, err := tx.Exec(`
				CREATE TABLE merchants (
					id VARCHAR(10) PRIMARY KEY,
					name VARCHAR(100) NOT NULL,
					category VARCHAR(20) NOT NULL,
					status VARCHAR(20) NOT NULL,
					details JSON
				)
			`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_merchants_id ON merchants(id)`)
			return err
		},
		Down: func(tx *sql.Tx) error {
			_, err := tx.Exec(`DROP TABLE IF EXISTS merchants`)
			return err
		},
	})
}
