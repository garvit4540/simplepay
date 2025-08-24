package migrations

import (
	"database/sql"
)

func init() {
	Migrations = append(Migrations, &Migration{
		Version: 2,
		Up: func(tx *sql.Tx) error {
			_, err := tx.Exec(`
				CREATE TABLE keys (
					id VARCHAR(10) PRIMARY KEY,
					merchant_id VARCHAR(10) NOT NULL,
					key VARCHAR(10) NOT NULL,
					FOREIGN KEY (merchant_id) REFERENCES merchants(id)
				)
			`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_keys_id ON keys(id)`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_keys_merchant_id ON keys(merchant_id)`)
			return err
		},
		Down: func(tx *sql.Tx) error {
			_, err := tx.Exec(`DROP TABLE IF EXISTS keys`)
			return err
		},
	})
}
