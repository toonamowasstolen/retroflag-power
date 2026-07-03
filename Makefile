BINARY := retroflag-powerd

# Workshop commands for local development.
.PHONY: help build test version run clean check

help:
	@printf '%s\n' \
		'Workshop commands:' \
		'  make help     Show available commands' \
		'  make build    Build ./retroflag-powerd' \
		'  make test     Run all tests' \
		'  make version  Show daemon name and version' \
		'  make run      Run the daemon locally' \
		'  make clean    Remove the built binary' \
		'  make check    Run tests, build, and version'

build:
	go build -o ./$(BINARY) ./cmd/retroflag-powerd

test:
	go test ./...

version:
	go run ./cmd/retroflag-powerd --version

run:
	go run ./cmd/retroflag-powerd

clean:
	rm -f ./$(BINARY)

check: test build version
