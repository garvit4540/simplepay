-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS merchant_keys (
    id VARCHAR(10) PRIMARY KEY,
    merchant_id VARCHAR(10) NOT NULL,
    key_value VARCHAR(10) NOT NULL,
    INDEX idx_keys_id (id),
    INDEX idx_keys_merchant_id (merchant_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS merchant_keys;
-- +goose StatementEnd
