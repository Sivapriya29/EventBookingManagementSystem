-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
    id UUID PRIMARY KEY,
    booking_id UUID REFERENCES bookings(id),
    user_id UUID REFERENCES users (id),
    event_id UUID REFERENCES events (id),
    amount NUMERIC,
    card_number TEXT,
    expiry_month TEXT,
    expiry_year TEXT,
    cvv TEXT,
    card_holder TEXT,
    payment_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
