-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX unique_in_progress_reception ON reception(pvz_id) 
WHERE status = 'in_progress';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX unique_in_progress_reception;
-- +goose StatementEnd
