-- relayer version: 1.0.0

-- Create Users table if not exists
CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    username TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);