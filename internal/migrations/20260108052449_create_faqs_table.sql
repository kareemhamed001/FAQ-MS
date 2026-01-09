-- +goose Up
-- +goose StatementBegin
CREATE TABLE faqs (
    id SERIAL PRIMARY KEY,
    category_id INT REFERENCES categories(id),
    store_id INT REFERENCES stores(id), -- Nullable for Global
    is_global BOOLEAN DEFAULT FALSE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE faqs;
-- +goose StatementEnd
