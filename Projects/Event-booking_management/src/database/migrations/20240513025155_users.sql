-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    "password" TEXT,
    mobile TEXT,
    role TEXT,
    created_at TIMESTAMP WITHOUT TIME ZONE
);

CREATE UNIQUE INDEX users_email_role_uidx on users using BTREE (email,"role");
CREATE UNIQUE INDEX users_mobile_role_uidx on users using BTREE (mobile,"role");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
