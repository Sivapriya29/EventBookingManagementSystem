-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    id UUID PRIMARY KEY,
    event_name TEXT,
    event_description TEXT,
    event_date TIMESTAMP WITHOUT TIME ZONE,
    event_time TIMESTAMP WITHOUT TIME ZONE,
    event_type TEXT,
    "location" TEXT,
    speaker_name TEXT,
    organizer_name TEXT,
    capacity INTEGER,
    per_person_price NUMERIC,
    created_at TIMESTAMP WITHOUT TIME ZONE,
    updated_at TIMESTAMP WITHOUT TIME ZONE,
    deleted_at TIMESTAMP WITHOUT TIME ZONE,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
