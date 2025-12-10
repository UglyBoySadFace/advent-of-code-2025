.PHONY: help run test fmt clean

# Default target
help:
	@echo "Advent of Code 2025 - Go Project"
	@echo ""
	@echo "Available targets:"
	@echo "  make run DAY=XX     - Run solution for day XX (e.g., make run DAY=01)"
	@echo "  make test           - Run all tests"
	@echo "  make test DAY=XX    - Run tests for specific day"
	@echo "  make fmt            - Format all Go files"
	@echo "  make clean          - Clean build artifacts"
	@echo ""

# Run a specific day's solution
run:
	@if [ -z "$(DAY)" ]; then \
		echo "Error: Please specify DAY (e.g., make run DAY=01)"; \
		exit 1; \
	fi
	@cd day$(DAY) && go run main.go

# Run tests
test:
	@if [ -z "$(DAY)" ]; then \
		go test ./...; \
	else \
		cd day$(DAY) && go test -v; \
	fi

# Format all Go code
fmt:
	go fmt ./...

# Clean build artifacts
clean:
	go clean ./...
	find . -name "*.test" -type f -delete
	find . -name "*.out" -type f -delete
