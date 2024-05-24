-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bookings (
    id UUID PRIMARY KEY,
    event_id UUID REFERENCES events (id),
    user_id UUID REFERENCES users (id),
    event_name TEXT,
    number_of_tickets INTEGER,
    total_amount NUMERIC,
    created_at TIMESTAMP WITHOUT TIME ZONE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
