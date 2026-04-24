BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

 CREATE TABLE IF NOT EXISTS projects (
   id BIGSERIAL PRIMARY KEY,
   uuid UUID NOT NULL UNIQUE DEFAULT gen_random_uuid(),
   name TEXT NOT NULL,
   description TEXT NULL,
   status TEXT NULL,
   stack TEXT NULL,
   repository_url TEXT NULL,
   deployment_url TEXT NULL,
   user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
   created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
 );

COMMIT;
