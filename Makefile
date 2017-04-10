run:
	go run cmd/pogo/pogo.go -db postgres://localhost:5432/pogo?sslmode=disable -schema jack -path model

tempo:
	go run cmd/pogo/pogo.go -db postgres://localhost:5432/tempo_dev?sslmode=disable -path tempo
.PHONY: tempo

migrate:
	@migrate -path migration -url $(POSTGRES_URL) down
	@migrate -path migration -url $(POSTGRES_URL) up

test:
	@go test -v ./postgres/postgres_test.go
