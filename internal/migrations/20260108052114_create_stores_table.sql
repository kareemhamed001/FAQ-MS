-- +goose Up
-- +goose StatementBegin
create table stores(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    merchant_id INT REFERENCES users(id) on delete cascade,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table stores;
-- +goose StatementEnd
