-- +goose Up
-- +goose StatementBegin
ALTER TABLE pvz
RENAME COLUMN created_at TO registration_date;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE pvz
RENAME COLUMN registration_date TO created_at;
-- +goose StatementEnd
