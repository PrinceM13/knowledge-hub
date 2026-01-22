CREATE TABLE health_checks(
    id SERIAL PRIMARY KEY,
    checked_at TIMESTAMP NOT NULL DEFAULT NOW()
);
