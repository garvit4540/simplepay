package payments

import (
	"database/sql"
	"fmt"
	"time"
)

type PaymentModel struct {
	ID              string    `json:"id"`
	OrderID         string    `json:"order_id"`
	MerchantID      string    `json:"merchant_id"`
	Amount          int64     `json:"amount"`
	Currency        string    `json:"currency"`
	Status          string    `json:"status"`
	ProviderID      string    `json:"provider_id"`
	ForcedProvider  string    `json:"forced_provider,omitempty"`
	TerminalID      string    `json:"terminal_id,omitempty"`
	GatewayResponse string    `json:"gateway_response,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

type PaymentsRepo struct {
	db *sql.DB
}

func NewPaymentsRepo(db *sql.DB) *PaymentsRepo {
	return &PaymentsRepo{
		db: db,
	}
}

func (pr *PaymentsRepo) CreatePayment(payment *PaymentModel) error {
	query := `
		INSERT INTO payments 
		(id, order_id, merchant_id, amount, currency, status, provider_id, forced_provider, terminal_id,gateway_response, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)
	`

	now := time.Now()
	_, err := pr.db.Exec(query, payment.ID, payment.OrderID, payment.MerchantID, payment.Amount, payment.Currency, payment.Status, payment.ProviderID, payment.ForcedProvider, payment.TerminalID, payment.GatewayResponse, now, now)
	if err != nil {
		return fmt.Errorf("failed to create payment: %w", err)
	}

	payment.CreatedAt = now
	payment.UpdatedAt = now
	return nil
}

func (pr *PaymentsRepo) UpdatePayment(payment *PaymentModel) error {
	query := `
		UPDATE payments SET 
			order_id = ?, 
			merchant_id = ?, 
			amount = ?, 
			currency = ?, 
			status = ?, 
			provider_id = ?, 
			forced_provider = ?, 
			terminal_id = ?, 
			gateway_response = ?, 
			updated_at = ?
		WHERE id = ?
	`

	now := time.Now()
	_, err := pr.db.Exec(query, payment.OrderID, payment.MerchantID, payment.Amount, payment.Currency, payment.Status, payment.ProviderID, payment.ForcedProvider, payment.TerminalID, payment.GatewayResponse, now, payment.ID)
	if err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}

	payment.UpdatedAt = now
	return nil
}
