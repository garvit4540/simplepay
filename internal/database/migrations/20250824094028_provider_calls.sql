-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS provider_calls (
    id VARCHAR(10) PRIMARY KEY,
    payment_id VARCHAR(10) NOT NULL,
    provider_request TEXT,
    provider_response TEXT,
    INDEX idx_provider_calls_id (id),
    INDEX idx_provider_calls_payment_id (payment_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS provider_calls;
-- +goose StatementEnd
