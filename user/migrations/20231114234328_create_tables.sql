-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    user_id UUID NOT NULL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    address TEXT,
    phone TEXT,
    role INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO users (
    user_id,
    first_name,
    last_name,
    password,
    email,
    address,
    phone,
    role,
    created_at,
    updated_at
) VALUES (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d3',
    'test_name',
    'test_surname',
    'password',
    'test@mail.ru',
    'test address',
    '+1234',
    1,
    'now()',
    'now()'
), (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d9',
    'test_name_2',
    'test_surname_2',
    'password_2',
    'test2@mail.ru',
    'test address 2',
    '+5678',
    2,
    'now()',
    'now()'
), (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d4',
    'admin',
    'admin',
    'adminadmin',
    'admin@mail.ru',
    'admin',
    '+123456789',
    3,
    'now()',
    'now()'
);

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE IF EXISTS users;

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
