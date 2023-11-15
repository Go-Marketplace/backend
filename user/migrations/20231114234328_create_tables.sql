-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    user_id UUID NOT NULL PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

INSERT INTO users (
    user_id,
    first_name,
    last_name,
    password,
    email,
    created_at,
    updated_at
) VALUES (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d3',
    'test_name',
    'test_surname',
    'password',
    'test@mail.ru',
    'now()',
    'now()'
), (
    'efb5b1a0-2222-4106-a2bc-577bd4b287d9',
    'test_name_2',
    'test_surname_2',
    'password_2',
    'test2@mail.ru',
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
