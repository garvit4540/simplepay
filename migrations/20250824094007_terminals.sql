-- +goose Up
-- +goose StatementBegin
CREATE TABLE terminals (
    id VARCHAR(10) PRIMARY KEY,
    merchant_id VARCHAR(10) NOT NULL,
    provider_id VARCHAR(10) NOT NULL
);
CREATE INDEX idx_terminals_id ON terminals(id);
CREATE INDEX idx_terminals_merchant_id ON terminals(merchant_id);
CREATE INDEX idx_terminals_provider_id ON terminals(provider_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS terminals;
-- +goose StatementEnd
