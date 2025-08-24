package migrations

import (
	"database/sql"
)

func init() {
	Migrations = append(Migrations, &Migration{
		Version: 4,
		Up: func(tx *sql.Tx) error {
			_, err := tx.Exec(`
				CREATE TABLE terminals (
					id VARCHAR(10) PRIMARY KEY,
					merchant_id VARCHAR(10) NOT NULL,
					provider_id VARCHAR(10) NOT NULL,
					FOREIGN KEY (merchant_id) REFERENCES merchants(id),
					FOREIGN KEY (provider_id) REFERENCES providers(id)
				)
			`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_terminals_id ON terminals(id)`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_terminals_merchant_id ON terminals(merchant_id)`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_terminals_provider_id ON terminals(provider_id)`)
			return err
		},
		Down: func(tx *sql.Tx) error {
			_, err := tx.Exec(`DROP TABLE IF EXISTS terminals`)
			return err
		},
	})
}
