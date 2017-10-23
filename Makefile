test: migrate generate testonly

generate: templates 
	@go run cmd/pogo/pogo.go --db $(POGO_POSTGRES_URL) --schema jack --path testjack
	
gambit: templates 
	@go run cmd/pogo/pogo.go --db $(GAMBIT_POSTGRES_URL) --schema 1 --path testgambit

testonly:
	@go test ./...
.PHONY: run

templates:
	@go-bindata -o templates/templates.go -pkg templates -ignore=templates.go templates/...
.PHONY: templates

install: templates
	@go install ./...

migrate:
	@migrate -path migration -url $(POGO_POSTGRES_URL) down
	@migrate -path migration -url $(POGO_POSTGRES_URL) up