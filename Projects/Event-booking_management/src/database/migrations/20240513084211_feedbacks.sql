-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS feedbacks (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users (id),
    event_id UUID REFERENCES events (id),
    rating NUMERIC,
    comments TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
