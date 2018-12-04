GREP?=""

test: templates jack.test

precommit: test

examples: jack.example

jack.test: templates
	@go test -v ./test/... -run $(GREP)

jack.example: templates
	@go run cmd/pogo/main.go --db $(POSTGRES_URL) --schema jack --dir _examples/jack/pogo

templates:
	@go-bindata -nometadata -o templates/templates.go -pkg templates -ignore=templates.go templates/...
.PHONY: templates

install: templates
	@go install ./cmd/pogo