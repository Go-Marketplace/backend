-- +goose Up
CREATE TABLE IF NOT EXISTS carts (
    cart_id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO carts (cart_id, user_id, created_at, updated_at)
VALUES
    ('efb5b1a0-2222-4106-a2bc-577bd4b287d1', 'efb5b1a0-2222-4106-a2bc-577bd4b287d2', 'now()', 'now()'),
    ('efb5b1a0-2222-4106-a2bc-577bd4b287d3', 'efb5b1a0-2222-4106-a2bc-577bd4b287d9', 'now()', 'now()');

CREATE TABLE IF NOT EXISTS cartlines (
    cartline_id UUID NOT NULL PRIMARY KEY,
    cart_id UUID NOT NULL,
    product_id UUID NOT NULL,
    name TEXT NOT NULL,
    quantity BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    FOREIGN KEY (cart_id) REFERENCES carts(cart_id) ON DELETE CASCADE
);

INSERT INTO cartlines (cartline_id, cart_id, product_id, name, quantity, created_at, updated_at)
VALUES
    ('efb5b1a0-2222-4106-a2bc-577bd4b287d4', 'efb5b1a0-2222-4106-a2bc-577bd4b287d1', 'efb5b1a0-2222-4106-a2bc-577bd4b287d5', 'cartline', 1, 'now()', 'now()'),
    ('efb5b1a0-2222-4106-a2bc-577bd4b287d6', 'efb5b1a0-2222-4106-a2bc-577bd4b287d1', 'efb5b1a0-2222-4106-a2bc-577bd4b287d7', 'cartline2', 1, 'now()', 'now()');

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS carts;

DROP TABLE IF EXISTS cartlines;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
