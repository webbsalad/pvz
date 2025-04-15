-- +goose Up
-- +goose StatementBegin
CREATE TABLE reception (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pvz_id UUID NOT NULL REFERENCES pvz(id) ON DELETE CASCADE,
    status TEXT NOT NULL DEFAULT 'in_progress' CHECK (status IN ('in_progress', 'close')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE reception;
-- +goose StatementEnd
