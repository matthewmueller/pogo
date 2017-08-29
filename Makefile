run: templates
	@go run cmd/pogo/pogo.go --db postgres://localhost:5432/pogo?sslmode=disable --schema jack --path jack
.PHONY: run

examples:  templates
	# go run cmd/pogo/pogo.go -db postgres://localhost:5432/tempo_dev?sslmode=disable -path tempo
	# go run cmd/pogo/pogo.go -db postgres://localhost:5432/bot-ii?sslmode=disable -path jack
.PHONY: tempo

templates:
	@go-bindata -o templates/templates.go -pkg templates -ignore=templates.go templates/...
.PHONY: templates

install: templates
	@go install ./...

migrate:
	@migrate -path migration -url $(POSTGRES_URL) down
	@migrate -path migration -url $(POSTGRES_URL) up

test:
	@go test -v ./postgres/postgres_test.go
