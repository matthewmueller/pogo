run: bin
	go run cmd/pogo/pogo.go -db postgres://localhost:5432/pogo?sslmode=disable -schema jack -path model
.PHONY: run

tempo:  bin
	go run cmd/pogo/pogo.go -db postgres://localhost:5432/tempo_dev?sslmode=disable -path tempo
.PHONY: tempo

bin:
	@go-bindata -o bin/bin.go -pkg bin templates/
.PHONY: bin

install: bin
	@go install ./...

migrate:
	@migrate -path migration -url $(POSTGRES_URL) down
	@migrate -path migration -url $(POSTGRES_URL) up

test:
	@go test -v ./postgres/postgres_test.go
