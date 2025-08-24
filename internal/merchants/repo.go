package merchants

import (
	"database/sql"
	"fmt"
	"time"
)

// MerchantModel represents th merchant data structure
type MerchantModel struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// MerchantRepo handles database operations for merchants
type MerchantRepo struct {
	db *sql.DB
}

// NewMerchantRepo creates a new merchant repository
func NewMerchantRepo(db *sql.DB) *MerchantRepo {
	return &MerchantRepo{
		db: db,
	}
}

// CreateMerchant creates a new merchant in the database
func (pr *MerchantRepo) CreateMerchant(merchant *MerchantModel) error {
	query := `
		INSERT INTO merchants (id, name,category,status, created_at, updated_at) 
		VALUES (?, ?, ?, ?,?,?)
	`

	now := time.Now()
	_, err := pr.db.Exec(query, merchant.ID, merchant.Name, merchant.Category, merchant.Status, now, now)
	if err != nil {
		return fmt.Errorf("failed to create merchant: %w", err)
	}

	merchant.CreatedAt = now
	merchant.UpdatedAt = now
	return nil
}
