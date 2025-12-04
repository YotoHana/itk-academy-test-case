-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS wallets(
    wallet_uuid UUID PRIMARY KEY,
    balance INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS wallets;
-- +goose StatementEnd
