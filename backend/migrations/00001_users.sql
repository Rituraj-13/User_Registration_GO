-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    firstname VARCHAR(50) NOT NULL,
    lastname VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(200) NOT NULL
)

-- +goose StatementEnd
-- +goose Down

-- +goose StatementBegin

DROP TABLE users;

-- +goose StatementEnd