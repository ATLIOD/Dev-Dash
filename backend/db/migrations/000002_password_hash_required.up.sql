BEGIN;

DO $$
BEGIN
  IF EXISTS (
    SELECT 1
    FROM information_schema.columns
    WHERE table_name = 'users' AND column_name = 'password'
  ) THEN
    ALTER TABLE users RENAME COLUMN password TO password_hash;
  END IF;
END $$;

ALTER TABLE users ALTER COLUMN password_hash DROP DEFAULT;
DELETE FROM users WHERE password_hash = '';
ALTER TABLE users ALTER COLUMN password_hash SET NOT NULL;

COMMIT;
