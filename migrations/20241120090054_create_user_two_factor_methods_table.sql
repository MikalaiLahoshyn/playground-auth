-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_two_factor_methods (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    two_factor_type_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (two_factor_type_id) REFERENCES two_factor_types(id) ON DELETE CASCADE,
    UNIQUE (user_id, two_factor_type_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_two_factor_methods;
-- +goose StatementEnd
