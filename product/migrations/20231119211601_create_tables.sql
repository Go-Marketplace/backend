-- +goose Up
CREATE TABLE IF NOT EXISTS categories (
    category_id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL
);

INSERT INTO categories (name, description)
VALUES
    ('Electronics', 'Devices and gadgets powered by electricity'),
    ('Clothing', 'Apparel and accessories for men, women, and children'),
    ('Books', 'Literature and educational materials'),
    ('Toys and Games', 'Entertainment and play items for children'),
    ('Furniture', 'Furnishings and home furniture');

CREATE TABLE IF NOT EXISTS products (
    product_id UUID NOT NULL PRIMARY KEY,
    user_id UUID NOT NULL,
    category_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price BIGINT NOT NULL,
    weight BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,

    FOREIGN KEY (category_id) REFERENCES categories(category_id)
);

INSERT INTO products (
    product_id,
    user_id,
    category_id,
    name,
    description,
    price,
    weight,
    created_at,
    updated_at
) VALUES (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d1',
    'efb5b1a0-2222-4106-a2bc-577bd4b287d2',
    1,
    'Smartphone',
    'Cool smartphone',
    100,
    1000,
    'now()',
    'now()'
), (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d3',
    'efb5b1a0-2222-4106-a2bc-577bd4b287d2',
    2,
    'T-Shirt',
    'Red T-Shirt',
    120,
    200,
    'now()',
    'now()'
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS products;

DROP TABLE IF EXISTS categories;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
