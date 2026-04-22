.PHONY: test-integration

# Run integration tests, bypassing the Go test cache to ensure we check the live environment
test-integration:
	go test -v -count=1 ./Test/integration/...
