GREP ?= ""

GOBINDATA := go run github.com/kevinburke/go-bindata/go-bindata

SQLITE_TAGS="sqlite_foreign_keys sqlite_json sqlite_fts5"

tidy:
	@ go mod tidy

test: tidy generate
	@ rm -rf internal/postgres/tmp internal/sqlite/tmp
	@ go test -v ./... -tags $(SQLITE_TAGS) -failfast -run $(GREP)

precommit: test

generate:
	@ $(GOBINDATA) -nometadata -o internal/templates/templates.go -pkg templates -ignore="templates\.go" internal/templates/...
.PHONY: generate

install:
	@ go install -tags $(SQLITE_TAGS) ./cmd/pogo