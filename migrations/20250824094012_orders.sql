-- +goose Up
-- +goose StatementBegin
CREATE TABLE orders (
    id VARCHAR(10) PRIMARY KEY,
    amount BIGINT NOT NULL,
    status VARCHAR(20) NOT NULL,
    currency VARCHAR(3) NOT NULL,
    order_details JSON,
    merchant_id VARCHAR(10) NOT NULL
);
CREATE INDEX idx_orders_id ON orders(id);
CREATE INDEX idx_orders_merchant_id ON orders(merchant_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
