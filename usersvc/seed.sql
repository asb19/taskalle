-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Insert 2 users with hardcoded UUIDs
INSERT INTO users (id, name, email)
VALUES
  ('11111111-1111-1111-1111-111111111111', 'Alice', 'alice@example.com'),
  ('22222222-2222-2222-2222-222222222222', 'Bob', 'bob@example.com')
ON CONFLICT (id) DO NOTHING;
