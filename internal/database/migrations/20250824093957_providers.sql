-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS providers (
   id VARCHAR(10) PRIMARY KEY,
   name VARCHAR(100) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
   INDEX idx_providers_id (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS providers;
-- +goose StatementEnd
