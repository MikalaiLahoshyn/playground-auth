-- +goose Up
-- +goose StatementBegin
CREATE TABLE two_factor_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    description TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS two_factor_types;
-- +goose StatementEnd
