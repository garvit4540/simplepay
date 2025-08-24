package migrations

import (
	"database/sql"
)

func init() {
	Migrations = append(Migrations, &Migration{
		Version: 5,
		Up: func(tx *sql.Tx) error {
			_, err := tx.Exec(`
				CREATE TABLE orders (
					id VARCHAR(10) PRIMARY KEY,
					amount BIGINT NOT NULL,
					status VARCHAR(20) NOT NULL,
					currency VARCHAR(3) NOT NULL,
					order_details JSON,
					merchant_id VARCHAR(10) NOT NULL,
					FOREIGN KEY (merchant_id) REFERENCES merchants(id)
				)
			`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_orders_id ON orders(id)`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_orders_merchant_id ON orders(merchant_id)`)
			return err
		},
		Down: func(tx *sql.Tx) error {
			_, err := tx.Exec(`DROP TABLE IF EXISTS orders`)
			return err
		},
	})
}
