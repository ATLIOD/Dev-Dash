BEGIN;

DO $$
BEGIN
  IF EXISTS (
    SELECT 1
    FROM information_schema.columns
    WHERE table_name = 'users' AND column_name = 'password_hash'
  ) THEN
    ALTER TABLE users ALTER COLUMN password_hash DROP NOT NULL;
    ALTER TABLE users ALTER COLUMN password_hash SET DEFAULT '';
    ALTER TABLE users RENAME COLUMN password_hash TO password;
  END IF;
END $$;

COMMIT;
