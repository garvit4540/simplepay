package terminals

import (
	"database/sql"
	"fmt"
	"time"
)

// TerminalModel represents th terminal data structure
type TerminalModel struct {
	ID         string    `json:"id"`
	MerchantId string    `json:"merchant_id"`
	ProviderId string    `json:"provider_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

// TerminalRepo handles database operations for terminals
type TerminalRepo struct {
	db *sql.DB
}

// NewTerminalRepo creates a new terminal repository
func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{
		db: db,
	}
}

// CreateTerminal creates a new terminal in the database
func (tr *TerminalRepo) CreateTerminal(terminal *TerminalModel) error {
	query := `
		INSERT INTO terminals (id, merchant_id,provider_id,created_at, updated_at) 
		VALUES (?, ?, ?, ?,?)
	`

	now := time.Now()
	_, err := tr.db.Exec(query, terminal.ID, terminal.MerchantId, terminal.ProviderId, now, now)
	if err != nil {
		return fmt.Errorf("failed to create terminal: %w", err)
	}

	terminal.CreatedAt = now
	terminal.UpdatedAt = now
	return nil
}
