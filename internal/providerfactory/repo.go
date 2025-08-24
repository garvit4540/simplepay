package providerfactory

import (
	"database/sql"
	"fmt"
	"time"
)

// ProviderModel represents the provider data structure
type ProviderModel struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// ProviderRepo handles database operations for providers
type ProviderRepo struct {
	db *sql.DB
}

// NewProviderRepo creates a new provider repository
func NewProviderRepo(db *sql.DB) *ProviderRepo {
	return &ProviderRepo{
		db: db,
	}
}

// CreateProvider creates a new provider in the database
func (pr *ProviderRepo) CreateProvider(provider *ProviderModel) error {
	query := `
		INSERT INTO providers (id, name, created_at, updated_at) 
		VALUES (?, ?, ?, ?)
	`

	now := time.Now()
	_, err := pr.db.Exec(query, provider.ID, provider.Name, now, now)
	if err != nil {
		return fmt.Errorf("failed to create provider: %w", err)
	}

	provider.CreatedAt = now
	provider.UpdatedAt = now
	return nil
}

// GetProviderByID retrieves a provider by its ID
func (pr *ProviderRepo) GetProviderByID(providerID string) (*ProviderModel, error) {
	query := `
		SELECT id, name, created_at, updated_at 
		FROM providers 
		WHERE id = ?
	`

	provider := &ProviderModel{}
	err := pr.db.QueryRow(query, providerID).Scan(
		&provider.ID,
		&provider.Name,
		&provider.CreatedAt,
		&provider.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("provider with ID '%s' not found", providerID)
		}
		return nil, fmt.Errorf("failed to get provider: %w", err)
	}

	return provider, nil
}

// GetProviderByName retrieves a provider by its name
func (pr *ProviderRepo) GetProviderByName(name string) (*ProviderModel, error) {
	query := `
		SELECT id, name, created_at, updated_at 
		FROM providers 
		WHERE name = ?
	`

	provider := &ProviderModel{}
	err := pr.db.QueryRow(query, name).Scan(
		&provider.ID,
		&provider.Name,
		&provider.CreatedAt,
		&provider.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("provider with name '%s' not found", name)
		}
		return nil, fmt.Errorf("failed to get provider: %w", err)
	}

	return provider, nil
}

// GetAllProviders retrieves all providers from the database
func (pr *ProviderRepo) GetAllProviders() ([]*ProviderModel, error) {
	query := `
		SELECT id, name, created_at, updated_at 
		FROM providers 
		ORDER BY created_at ASC
	`

	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query providers: %w", err)
	}
	defer rows.Close()

	var providers []*ProviderModel
	for rows.Next() {
		provider := &ProviderModel{}
		err := rows.Scan(
			&provider.ID,
			&provider.Name,
			&provider.CreatedAt,
			&provider.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan provider: %w", err)
		}
		providers = append(providers, provider)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating providers: %w", err)
	}

	return providers, nil
}
