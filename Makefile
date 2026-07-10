BINARY := retroflag-powerd

# Workshop commands for local development.
.PHONY: help build test version run clean check check-scripts check-links

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
		'  make check-scripts  Check portable shell script syntax and help' \
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

check: test build version check-scripts

check-scripts:
	sh -n scripts/gpi-case2-bundle-collector-field-lantern.sh
	sh -n scripts/gpi-case2-boot-power-trace-field-lantern.sh
	sh -n scripts/gpi-case2-install-relic-welcome-scroll-hook.sh
	sh -n scripts/gpi-case2-relic-welcome-scroll.sh
	sh -n scripts/gpi-case2-session-watch-lantern.sh
	sh -n scripts/gpi-case2-true-boot-trace-lantern.sh
	sh scripts/gpi-case2-install-relic-welcome-scroll-hook.sh --help >/dev/null
	HOME=/tmp/gpi-case2-relic-welcome-scroll-hook-status sh scripts/gpi-case2-install-relic-welcome-scroll-hook.sh --status --plain >/dev/null
	NO_COLOR=1 HOME=/tmp/gpi-case2-relic-welcome-scroll-hook-status sh scripts/gpi-case2-install-relic-welcome-scroll-hook.sh --status >/dev/null
	tmp=$$(mktemp -d /tmp/gpi-case2-relic-hook-dry-run.XXXXXX); cp scripts/gpi-case2-relic-welcome-scroll.sh "$$tmp/"; HOME="$$tmp" sh scripts/gpi-case2-install-relic-welcome-scroll-hook.sh --install --dry-run --plain >/dev/null
	tmp=$$(mktemp -d /tmp/gpi-case2-relic-hook.XXXXXX); cp scripts/gpi-case2-relic-welcome-scroll.sh "$$tmp/"; printf '%s\n' '# field shell map' > "$$tmp/.bashrc"; HOME="$$tmp" sh scripts/gpi-case2-install-relic-welcome-scroll-hook.sh --install --plain >/dev/null; grep -F 'gpi-case2 relic welcome scroll hook' "$$tmp/.bashrc" >/dev/null; test -f "$$tmp/.gpi-relic-welcome-scroll-hook.sh"; HOME="$$tmp" sh scripts/gpi-case2-install-relic-welcome-scroll-hook.sh --uninstall --plain >/dev/null; ! grep -F 'gpi-case2 relic welcome scroll hook' "$$tmp/.bashrc" >/dev/null
	sh scripts/gpi-case2-relic-welcome-scroll.sh --help >/dev/null
	sh scripts/gpi-case2-session-watch-lantern.sh --help >/dev/null
	sh scripts/gpi-case2-true-boot-trace-lantern.sh --help >/dev/null
	NO_COLOR= sh scripts/gpi-case2-relic-welcome-scroll.sh >/dev/null
	sh scripts/gpi-case2-relic-welcome-scroll.sh --plain >/dev/null
	NO_COLOR=1 sh scripts/gpi-case2-relic-welcome-scroll.sh >/dev/null
	sh scripts/gpi-case2-session-watch-lantern.sh --plain --duration 1 --interval 1 --output /tmp/gpi-case2-session-watch-lantern-plain-smoke.txt >/dev/null
	NO_COLOR=1 sh scripts/gpi-case2-session-watch-lantern.sh --duration 1 --interval 1 --output /tmp/gpi-case2-session-watch-lantern-nocolor-smoke.txt >/dev/null
	sh scripts/gpi-case2-true-boot-trace-lantern.sh --plain --duration 1 --interval 1 --output /tmp/gpi-case2-true-boot-trace-lantern-plain-smoke.txt >/dev/null
	NO_COLOR=1 sh scripts/gpi-case2-true-boot-trace-lantern.sh --duration 1 --interval 1 --output /tmp/gpi-case2-true-boot-trace-lantern-nocolor-smoke.txt >/dev/null

check-links:
	python3 scripts/check-markdown-links.py
