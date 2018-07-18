test: migrate generate testonly

examples: jack digby

jack: templates 
	@go test -v ./test/jack/jack_test.go

digby: templates 
	@go run cmd/pogo/main.go --db $(DIGBY_POSTGRES_URL) --dir _examples/digby

gambit: templates 
	@go run cmd/pogo/main.go --db $(GAMBIT_POSTGRES_URL) --schema 1 --dir _examples/gambit

testonly:
	@go test ./...
.PHONY: run

templates:
	@go-bindata -nometadata -o templates/templates.go -pkg templates -ignore=templates.go templates/...
.PHONY: templates

install: templates
	@go install ./cmd/...

# migrate:
# 	@migrate --dir test/jack down --db $(JACK_POSTGRES_URL)
# 	@migrate --dir test/jack up --db $(JACK_POSTGRES_URL)
