test: templates jack

precommit: test

examples: jack digby

jack: templates 
	@go test -v ./test/jack/jack_test.go

# digby: templates 
# 	@go run cmd/pogo/main.go --db $(DIGBY_POSTGRES_URL) --dir _examples/digby

# gambit: templates 
# 	@go run cmd/pogo/main.go --db $(GAMBIT_POSTGRES_URL) --schema 1 --dir _examples/gambit

templates:
	@go-bindata -nometadata -o templates/templates.go -pkg templates -ignore=templates.go templates/...
.PHONY: templates

install: templates
	@go install ./cmd/...