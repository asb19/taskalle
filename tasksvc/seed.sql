-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Insert 30 tasks, randomly assigned to Alice or Bob
INSERT INTO tasks (title, description, status, assigned_to)
SELECT
  'Task ' || i,
  'Description for task ' || i,
  (ARRAY['pending','inprogress','done'])[floor(random()*3 + 1)],
  (ARRAY[
    '11111111-1111-1111-1111-111111111111'::uuid,
    '22222222-2222-2222-2222-222222222222'::uuid
  ])[floor(random()*2 + 1)]
FROM generate_series(1,30) AS s(i);
