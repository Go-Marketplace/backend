-- +goose Up
CREATE TABLE IF NOT EXISTS carts (
    user_id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO carts (user_id, created_at, updated_at)
VALUES
    (
        'efb5b1a0-2222-4106-a2bc-577bd4b287d1',
        'now()',
        'now()'
    ),
    (
        'efb5b1a0-2222-4106-a2bc-577bd4b287d2',
        'now()',
        'now()'
    );

CREATE TABLE IF NOT EXISTS cartlines (
    user_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    PRIMARY KEY (user_id, product_id),
    FOREIGN KEY (user_id) REFERENCES carts(user_id) ON DELETE CASCADE
);

INSERT INTO cartlines (user_id, product_id, quantity, created_at, updated_at)
VALUES
    (
        'efb5b1a0-2222-4106-a2bc-577bd4b287d1',
        'efb5b1a0-2222-4106-a2bc-577bd4b287d5',
        1,
        'now()',
        'now()'
    ),
    (
        'efb5b1a0-2222-4106-a2bc-577bd4b287d1',
        'efb5b1a0-2222-4106-a2bc-577bd4b287d7',
        10,
        'now()',
        'now()'
    );

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down

DROP TABLE IF EXISTS cartlines;

DROP TABLE IF EXISTS carts;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
