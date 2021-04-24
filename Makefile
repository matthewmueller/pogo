GREP ?= ""

tidy:
	@ go mod tidy

test: tidy generate
	@ rm -rf internal/postgres/tmp internal/sqlite/tmp
	@ go test -v ./... -failfast -run $(GREP)

precommit: test

generate:
	@ go-bindata -nometadata -o internal/templates/templates.go -pkg templates -ignore="templates\.go" internal/templates/...
.PHONY: generate

install: test
	@ go install ./cmd/pogo