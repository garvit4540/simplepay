-- +goose Up
-- +goose StatementBegin
CREATE TABLE payments (
    id VARCHAR(10) PRIMARY KEY,
    order_id VARCHAR(10) NOT NULL,
    merchant_id VARCHAR(10) NOT NULL,
    amount BIGINT NOT NULL,
    currency VARCHAR(3) NOT NULL,
    status VARCHAR(20) NOT NULL,
    provider_id VARCHAR(10) NOT NULL,
    forced_provider VARCHAR(10),
    terminal_id VARCHAR(10)
)
CREATE INDEX idx_payments_id ON payments(id);
CREATE INDEX idx_payments_order_id ON payments(order_id);
CREATE INDEX idx_payments_merchant_id ON payments(merchant_id);
CREATE INDEX idx_payments_provider_id ON payments(provider_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments
-- +goose StatementEnd
