GREP ?= ""

test: templates
	@ rm -rf internal/postgres/tmp internal/sqlite/tmp
	@ go test -v ./... -failfast -run $(GREP)

precommit: test

templates:
	@ go-bindata -nometadata -o internal/templates/templates.go -pkg templates -ignore="templates\.go" internal/templates/...
.PHONY: templates

install: test
	@ go install ./cmd/pogo