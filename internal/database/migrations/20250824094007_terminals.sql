-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS terminals (
    id VARCHAR(10) PRIMARY KEY,
    merchant_id VARCHAR(10) NOT NULL,
    provider_id VARCHAR(10) NOT NULL,
    INDEX idx_terminals_id (id),
    INDEX idx_terminals_merchant_id (merchant_id),
    INDEX idx_terminals_provider_id (provider_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS terminals;
-- +goose StatementEnd
