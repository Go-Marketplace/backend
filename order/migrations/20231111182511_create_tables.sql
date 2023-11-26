-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
    order_id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    total_price BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO orders (order_id, user_id, total_price, created_at, updated_at)
VALUES
  ('efb5b1a0-2222-4106-a2bc-577bd4b287d1', 'efb5b1a0-2222-4106-a2bc-577bd4b287d3', 100, 'now()', 'now()'),
  ('efb5b1a0-2222-4106-a2bc-577bd4b287d2', 'efb5b1a0-2222-4106-a2bc-577bd4b287d3', 150, 'now()', 'now()');

CREATE TABLE IF NOT EXISTS orderlines (
    orderline_id UUID NOT NULL PRIMARY KEY,
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    name TEXT NOT NULL,
    price BIGINT NOT NULL,
    quantity BIGINT NOT NULL,
    status INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL

    FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE
);

INSERT INTO orderlines (orderline_id, order_id, product_id, name, price, quantity, status, created_at, updated_at)
VALUES
  ('efb5b1a0-2222-4106-a2bc-577bd4b287d4', 'efb5b1a0-2222-4106-a2bc-577bd4b287d1', 'efb5b1a0-2222-4106-a2bc-577bd4b287d1', 'Fish', 100, 2, 1, 'now()', 'now()'),
  ('efb5b1a0-2222-4106-a2bc-577bd4b287d5', 'efb5b1a0-2222-4106-a2bc-577bd4b287d2', 'efb5b1a0-2222-4106-a2bc-577bd4b287d1', 'Meat', 120, 3, 2, 'now()', 'now()');

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS products;

DROP TABLE IF EXISTS cartlines;

DROP TABLE IF EXISTS orders;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
