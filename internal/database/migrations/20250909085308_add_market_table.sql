-- +goose Up
-- +goose StatementBegin
CREATE TABLE market (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    symbol VARCHAR(16) NOT NULL
);
INSERT INTO market
    (name, symbol)
VALUES
    ('Bitcoin', 'BTC'),
    ('Etherium', 'ETH'),
    ('Solana', 'SOL');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE market;
-- +goose StatementEnd
