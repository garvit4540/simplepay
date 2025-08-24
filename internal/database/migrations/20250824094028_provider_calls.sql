-- +goose Up
-- +goose StatementBegin
CREATE TABLE provider_calls (
    id VARCHAR(10) PRIMARY KEY,
    payment_id VARCHAR(10) NOT NULL,
    provider_request JSON,
    provider_response JSON
);
CREATE INDEX idx_provider_calls_id ON provider_calls(id);
CREATE INDEX idx_provider_calls_payment_id ON provider_calls(payment_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS provider_calls;
-- +goose StatementEnd
