-- +goose Up
CREATE TABLE IF NOT EXISTS orders (
    order_id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO orders (
    order_id, 
    user_id,
    created_at,
    updated_at
) VALUES (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d1',
    'efb5b1a0-2222-4106-a2bc-577bd4b287d3',
    'now()',
    'now()'
), (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d2',
    'efb5b1a0-2222-4106-a2bc-577bd4b287d3',
    'now()',
    'now()'
);

CREATE TABLE IF NOT EXISTS orderlines (
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    name TEXT NOT NULL,
    price BIGINT NOT NULL,
    quantity BIGINT NOT NULL,
    status INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders(order_id) ON DELETE CASCADE
);

INSERT INTO orderlines (
    order_id,
    product_id,
    name,
    price,
    quantity,
    status,
    created_at,
    updated_at
) VALUES (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d1',
    'efb5b1a0-2222-4106-a2bc-577bd4b287d1',
    'Fish',
    100,
    2,
    1,
    'now()',
    'now()'
), (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d2',
    'efb5b1a0-2222-4106-a2bc-577bd4b287d2',
    'Meat',
    120,
    3,
    2,
    'now()',
    'now()'
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down

DROP TABLE IF EXISTS orderlines;

DROP TABLE IF EXISTS orders;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
