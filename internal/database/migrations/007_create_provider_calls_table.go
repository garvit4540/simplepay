package migrations

import (
	"database/sql"
)

func init() {
	Migrations = append(Migrations, &Migration{
		Version: 7,
		Up: func(tx *sql.Tx) error {
			_, err := tx.Exec(`
				CREATE TABLE provider_calls (
					id VARCHAR(10) PRIMARY KEY,
					payment_id VARCHAR(10) NOT NULL,
					provider_request JSON,
					provider_response JSON,
					FOREIGN KEY (payment_id) REFERENCES payments(id)
				)
			`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_provider_calls_id ON provider_calls(id)`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_provider_calls_payment_id ON provider_calls(payment_id)`)
			return err
		},
		Down: func(tx *sql.Tx) error {
			_, err := tx.Exec(`DROP TABLE IF EXISTS provider_calls`)
			return err
		},
	})
}
