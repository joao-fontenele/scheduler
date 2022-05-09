GOCMD?=$(shell which go)
GOCMD=go1.18

init: install-deps
	touch .env

.PHONY: install-deps
install-deps:
	${GOCMD} install github.com/cespare/reflex@v0.3

.PHONY: run
run: tidy
	${GOCMD} run ./cmd/server/main.go

.PHONY: run-watch
run-watch:
	reflex -d none -s -r '\.go$$' make run

.PHONY: test
test:
	${GOCMD} test ./...

.PHONY: test-watch
test-watch:
	reflex -d none -s -r '\.go$$' make test

.PHONY: tidy
tidy:
	${GOCMD} mod tidy
