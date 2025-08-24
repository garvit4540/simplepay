-- +goose Up
-- +goose StatementBegin
CREATE TABLE merchants (
    id VARCHAR(10) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(20) NOT NULL,
    status VARCHAR(20) NOT NULL,
    details JSON
);
CREATE INDEX idx_merchants_id ON merchants(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS merchants;
-- +goose StatementEnd
