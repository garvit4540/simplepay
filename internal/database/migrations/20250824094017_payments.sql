-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
    id VARCHAR(10) PRIMARY KEY,
    order_id VARCHAR(10) NOT NULL,
    merchant_id VARCHAR(10) NOT NULL,
    amount BIGINT NOT NULL,
    currency VARCHAR(3) NOT NULL,
    status VARCHAR(20) NOT NULL,
    provider_id VARCHAR(10) NOT NULL,
    forced_provider VARCHAR(10),
    terminal_id VARCHAR(10),
    gateway_response TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    INDEX idx_payments_id (id),
    INDEX idx_payments_order_id (order_id),
    INDEX idx_payments_merchant_id (merchant_id),
    INDEX idx_payments_provider_id (provider_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments;
-- +goose StatementEnd
