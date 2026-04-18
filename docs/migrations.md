# Database migrations

Migrations are being managed by [golang-migrate](https://github.com/golang-migrate/migrate).  

Migration files live in:

- `backend/db/migrations/`

Each change is a pair of files:

- `NNNNNN_description.up.sql` — apply
- `NNNNNN_description.down.sql` — rollback (exists for manual use only. These are **NOT** run automatically by the API container)

The API container runs **`migrate ... up` on startup**.

---

## Prerequisites

- A running Postgres instance and a valid `DATABASE_URL` (see [Connection strings](#connection-strings)).

---

## Create a new migration (all options)

### Option A — `migrate` CLI installed locally (fastest if you already have it)

From the `backend/` directory:

```bash
migrate create -ext sql -dir db/migrations -seq create_projects
```

From the repository root:

```bash
migrate create -ext sql -dir backend/db/migrations -seq create_projects
```

Verify install:

```bash
migrate -version
```

### Option B — Install `migrate` with Go (no Docker image pull)

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Then run the same `migrate create ...` commands as Option A (often `~/go/bin/migrate` until `~/go/bin` is on `PATH`).

### Option C — Docker `migrate` image (good for cross-platform work)

**Important:** mount the correct folder.

- If your shell is in `backend/`, mount `"$PWD"` to `/work` (not `"$PWD/backend"`).

```bash
docker run --rm -v "$PWD:/work" -w /work migrate/migrate \
  create -ext sql -dir db/migrations -seq create_projects
```

From repository root:

```bash
docker run --rm -v "$PWD/backend:/work" -w /work migrate/migrate \
  create -ext sql -dir db/migrations -seq create_projects
```

### Option D — Manual file creation (fallback)

Create the next numbered pair yourself in `backend/db/migrations/`:

- `000003_whatever.up.sql`
- `000003_whatever.down.sql`

Keep numbering consistent with existing files.

---

## After creating files

1. Edit the generated `.up.sql` / `.down.sql` with real SQL.
2. Run migrations against your database (local or compose) — see [Apply migrations](#apply-migrations).

---

## Apply migrations

### Automatic (compose)

If you start the stack with Docker Compose, the API container applies pending **`up`** migrations on startup.

### Manual (host)

Example (Linux; Postgres exposed on localhost):

```bash
export DATABASE_URL='postgres://web:webpass@localhost:5432/dev_dashdb?sslmode=disable'
migrate -path backend/db/migrations -database "$DATABASE_URL" up
```

From `backend/`:

```bash
export DATABASE_URL='postgres://web:webpass@localhost:5432/dev_dashdb?sslmode=disable'
migrate -path db/migrations -database "$DATABASE_URL" up
```

### Manual rollback (explicit; avoid in production unless you mean it)

```bash
migrate -path backend/db/migrations -database "$DATABASE_URL" down 1
```

---

## Connection strings

### Local Postgres via Docker Compose (published to host port 5432)

Example:

`postgres://web:webpass@localhost:5432/dev_dashdb?sslmode=disable`

### From another container on the same Compose network

Use the DB service hostname (often `db`) instead of `localhost`:

`postgres://web:webpass@db:5432/dev_dashdb?sslmode=disable`

### macOS / Windows Docker Desktop notes

If you run `migrate` in a container and Postgres is on the host, you may need `host.docker.internal` instead of `localhost` (platform-dependent).

---

## Troubleshooting

### `zsh: parse error near '\n'` when copying examples

Angle brackets like `<migration_name>` are placeholders. Do **not** type them literally.

Use:

```bash
migrate create -ext sql -dir db/migrations -seq create_projects
```

### Docker pull/extract errors (`blob not found`)

This usually indicates a **local Docker image store corruption**, not a bad migration command.

Try:

1. Restart Docker Desktop.
2. Remove the image and re-pull:

   ```bash
   docker image rm migrate/migrate:latest
   docker pull migrate/migrate:latest
   ```

3. If still broken, prune and retry, or reset Docker Desktop data (last resort).

While Docker is unhealthy, use **Option B (`go install`)** or **Option D (manual files)** to keep working.
