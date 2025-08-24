package keys

import (
	"database/sql"
	"fmt"
	"time"
)

// KeysModel represents th keys data structure
type KeysModel struct {
	ID         string    `json:"id"`
	MerchantId string    `json:"merchant_id"`
	KeyValue   string    `json:"key_value"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}

// KeysRepo handles database operations for keys
type KeysRepo struct {
	db *sql.DB
}

// NewKeysRepo creates a new key repository
func NewKeysRepo(db *sql.DB) *KeysRepo {
	return &KeysRepo{
		db: db,
	}
}

// CreateKeys creates a new key in the database
func (kr *KeysRepo) CreateKeys(key *KeysModel) error {
	query := `
		INSERT INTO merchant_keys (id, merchant_id,key_value,created_at, updated_at) 
		VALUES (?, ?, ?, ?,?)
	`

	now := time.Now()
	_, err := kr.db.Exec(query, key.ID, key.MerchantId, key.KeyValue, now, now)
	if err != nil {
		return fmt.Errorf("failed to create keys: %w", err)
	}

	key.CreatedAt = now
	key.UpdatedAt = now
	return nil
}

// GetKeyByMerchantId retrieves keys for a merchant by merchant id
func (kr *KeysRepo) GetKeyByMerchantId(merchantId string) (*KeysModel, error) {
	query := `
		SELECT id, merchant_id, key_value 
		FROM merchant_keys 
		WHERE merchant_id = ?
	`

	keys := &KeysModel{}
	err := kr.db.QueryRow(query, merchantId).Scan(
		&keys.ID,
		&keys.MerchantId,
		&keys.KeyValue,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("keys for merchantId '%s' not found", merchantId)
		}
		return nil, fmt.Errorf("failed to get keys: %w", err)
	}

	return keys, nil
}
