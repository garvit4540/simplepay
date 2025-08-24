package orders

import (
	"database/sql"
	"fmt"
	"time"
)

type OrderModel struct {
	ID           string    `json:"id"`
	Amount       int64     `json:"amount"`
	Status       string    `json:"status"`
	Currency     string    `json:"currency"`
	OrderDetails string    `json:"order_details"`
	MerchantID   string    `json:"merchant_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

// OrdersRepo handles database operations for merchants
type OrdersRepo struct {
	repo *sql.DB
}

// NewOrdersRepo creates a new merchant repository
func NewOrdersRepo(db *sql.DB) *OrdersRepo {
	return &OrdersRepo{
		repo: db,
	}
}

func (or *OrdersRepo) CreateOrder(order *OrderModel) error {
	query := `
		INSERT INTO orders (id, amount, status, currency, order_details, merchant_id, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	now := time.Now()
	_, err := or.repo.Exec(query, order.ID, order.Amount, order.Status, order.Currency, order.OrderDetails, order.MerchantID, now, now)
	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	order.CreatedAt = now
	order.UpdatedAt = now
	return nil
}
