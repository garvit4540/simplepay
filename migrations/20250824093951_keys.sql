-- +goose Up
-- +goose StatementBegin
CREATE TABLE keys (
    id VARCHAR(10) PRIMARY KEY,
    merchant_id VARCHAR(10) NOT NULL,
    key VARCHAR(10) NOT NULL
);
CREATE INDEX idx_keys_id ON keys(id);
CREATE INDEX idx_keys_merchant_id ON keys(merchant_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS keys;
-- +goose StatementEnd
