generate: templates 
	@go run cmd/pogo/pogo.go --db $(POSTGRES_URL) --schema jack --path testjack
	
test:
	@go test ./...
.PHONY: run

templates:
	@go-bindata -o templates/templates.go -pkg templates -ignore=templates.go templates/...
.PHONY: templates

install: templates
	@go install ./...

migrate:
	@migrate -path migration -url $(POSTGRES_URL) down
	@migrate -path migration -url $(POSTGRES_URL) up