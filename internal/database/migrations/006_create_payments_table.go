package migrations

import (
	"database/sql"
)

func init() {
	Migrations = append(Migrations, &Migration{
		Version: 6,
		Up: func(tx *sql.Tx) error {
			_, err := tx.Exec(`
				CREATE TABLE payments (
					id VARCHAR(10) PRIMARY KEY,
					order_id VARCHAR(10) NOT NULL,
					merchant_id VARCHAR(10) NOT NULL,
					amount BIGINT NOT NULL,
					currency VARCHAR(3) NOT NULL,
					status VARCHAR(20) NOT NULL,
					provider_id VARCHAR(10) NOT NULL,
					forced_provider VARCHAR(10),
					terminal_id VARCHAR(10),
					FOREIGN KEY (order_id) REFERENCES orders(id),
					FOREIGN KEY (merchant_id) REFERENCES merchants(id),
					FOREIGN KEY (provider_id) REFERENCES providers(id),
					FOREIGN KEY (terminal_id) REFERENCES terminals(id)
				)
			`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_payments_id ON payments(id)`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_payments_order_id ON payments(order_id)`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_payments_merchant_id ON payments(merchant_id)`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_payments_provider_id ON payments(provider_id)`)
			return err
		},
		Down: func(tx *sql.Tx) error {
			_, err := tx.Exec(`DROP TABLE IF EXISTS payments`)
			return err
		},
	})
}
