-- +goose Up
-- +goose StatementBegin
CREATE TABLE items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    reception_id UUID NOT NULL REFERENCES reception(id) ON DELETE CASCADE,
    type TEXT NOT NULL CHECK (type IN ('электроника', 'одежда', 'обувь')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE items;
-- +goose StatementEnd
