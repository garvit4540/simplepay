-- +goose Up
-- +goose StatementBegin
CREATE TABLE providers (
   id VARCHAR(10) PRIMARY KEY,
   name VARCHAR(100) NOT NULL
);
CREATE INDEX idx_providers_id ON providers(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS providers;
-- +goose StatementEnd
