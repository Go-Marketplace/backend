-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
    order_id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    status INTEGER NOT NULL,
    total_price BIGINT NOT NULL,
    shipping_cost BIGINT NOT NULL,
    delivery_address TEXT NOT NULL,
    delivery_type INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO orders (order_id, user_id, status, total_price, shipping_cost, delivery_address, delivery_type, created_at, updated_at)
VALUES
  ('efb5b1a0-2222-4106-a2bc-577bd4b287d1', 'efb5b1a0-2222-4106-a2bc-577bd4b287d3', 1, 100, 10, '123 Main St', 1, 'now()', 'now()'),
  ('efb5b1a0-2222-4106-a2bc-577bd4b287d2', 'efb5b1a0-2222-4106-a2bc-577bd4b287d3', 2, 150, 15, '456 Oak St', 2, 'now()', 'now()');

CREATE TABLE IF NOT EXISTS cartlines (
    cartline_id UUID NOT NULL PRIMARY KEY,
    order_id UUID NOT NULL,
    quantity BIGINT NOT NULL,

    FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE
);

INSERT INTO cartlines (cartline_id, order_id, quantity)
VALUES
  ('efb5b1a0-2222-4106-a2bc-577bd4b287d4', 'efb5b1a0-2222-4106-a2bc-577bd4b287d1', 2),
  ('efb5b1a0-2222-4106-a2bc-577bd4b287d5', 'efb5b1a0-2222-4106-a2bc-577bd4b287d2', 3);

CREATE TABLE IF NOT EXISTS products (
    product_id UUID NOT NULL PRIMARY KEY,
    cartline_id UUID NOT NULL UNIQUE,
    name TEXT NOT NULL,
    description TEXT,
    price BIGINT NOT NULL,

    FOREIGN KEY (cartline_id) REFERENCES cartlines(cartline_id) ON DELETE CASCADE
);

INSERT INTO products (product_id, cartline_id, name, description, price)
VALUES
  ('efb5b1a0-2222-4106-a2bc-577bd4b287d6', 'efb5b1a0-2222-4106-a2bc-577bd4b287d4', 'Product 1', 'Description 1', 50),
  ('efb5b1a0-2222-4106-a2bc-577bd4b287d7', 'efb5b1a0-2222-4106-a2bc-577bd4b287d5', 'Product 2', 'Description 2', 75);

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
