-- +goose Up
-- +goose StatementBegin
INSERT INTO wallets(wallet_uuid, balance) VALUES ('73fd60a4-ef54-484c-9de8-4beb3808da26', 10000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE wallets;
-- +goose StatementEnd
