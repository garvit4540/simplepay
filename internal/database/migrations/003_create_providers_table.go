package migrations

import (
	"database/sql"
)

func init() {
	Migrations = append(Migrations, &Migration{
		Version: 3,
		Up: func(tx *sql.Tx) error {
			_, err := tx.Exec(`
				CREATE TABLE providers (
					id VARCHAR(10) PRIMARY KEY,
					name VARCHAR(100) NOT NULL
				)
			`)
			if err != nil {
				return err
			}

			_, err = tx.Exec(`CREATE INDEX idx_providers_id ON providers(id)`)
			return err
		},
		Down: func(tx *sql.Tx) error {
			_, err := tx.Exec(`DROP TABLE IF EXISTS providers`)
			return err
		},
	})
}
