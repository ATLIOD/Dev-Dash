#!/bin/sh
set -eu

echo "Waiting for database to be ready..."
until pg_isready -d "${DATABASE_URL}" >/dev/null 2>&1; do
    sleep 1
done

echo "Running database up migrations..."
migrate -path /app/db/migrations -database "${DATABASE_URL}" up

echo "Starting API server..."
exec /app/api
