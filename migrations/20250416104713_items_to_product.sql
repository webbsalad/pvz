-- +goose Up
-- +goose StatementBegin
ALTER TABLE items RENAME TO product;
ALTER TABLE product RENAME COLUMN created_at TO date_time;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE product RENAME COLUMN date_time TO created_at;
ALTER TABLE product RENAME TO items;
-- +goose StatementEnd
