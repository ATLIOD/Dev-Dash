# Testing DevDash

This document explains how to run the various tests in the DevDash project.

## Prerequisites

To run the integration tests, you must have the application and its database running in Docker.

```bash
docker compose up -d
```

## Integration Tests

Integration tests are located in `backend/Test/integration`. These tests perform real HTTP requests against the running API to verify the full lifecycle of system components.

### Running Integration Tests (Recommended)

We use a `Makefile` to simplify running integration tests and to ensure that Go's test caching doesn't give false positives.

From the `backend/` directory, run:

```bash
make test-integration
```

This command executes:
`go test -v -count=1 ./Test/integration/...`

The `-count=1` flag is crucial as it forces Go to bypass the test cache and hit the live API every time.

### Customizing the API URL

By default, the tests look for the API at `http://localhost:8080`. If the API is running elsewhere, set the `API_BASE_URL` environment variable:

```bash
API_BASE_URL="http://myserver:9000" make test-integration
```

## Unit Tests

Unit tests verify individual components in isolation without requiring a database or running API.

To run all unit tests:

```bash
go test ./internal/...
```
