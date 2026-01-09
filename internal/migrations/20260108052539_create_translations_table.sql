-- +goose Up
-- +goose StatementBegin
CREATE TABLE translations (
    id SERIAL PRIMARY KEY,
    faq_id INT REFERENCES faqs(id) ON DELETE CASCADE,
    language VARCHAR(10) NOT NULL,
    question TEXT NOT NULL,
    answer TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE translations;
-- +goose StatementEnd
