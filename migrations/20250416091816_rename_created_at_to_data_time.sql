-- +goose Up
-- +goose StatementBegin
ALTER TABLE reception
RENAME COLUMN created_at TO date_time;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE reception
RENAME COLUMN date_time TO created_at;
-- +goose StatementEnd
