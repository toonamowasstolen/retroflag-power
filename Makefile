BINARY := retroflag-powerd

# Workshop commands for local development.
.PHONY: help build test version run clean check check-links

help:
	@printf '%s\n' \
		'Workshop commands:' \
		'  make help     Show available commands' \
		'  make build    Build ./retroflag-powerd' \
		'  make test     Run all tests' \
		'  make version  Show daemon name and version' \
		'  make run      Run the daemon locally' \
		'  make clean    Remove the built binary' \
		'  make check    Run tests, build, and version' \
		'  make check-links  Check internal Markdown links and anchors'

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

check-links:
	python3 scripts/check-markdown-links.py
