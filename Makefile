run:
	go run cmd/pogo/pogo.go -db postgres://localhost:5432/pogo?sslmode=disable

migrate:
	@migrate -path migration -url $(POSTGRES_URL) down
	@migrate -path migration -url $(POSTGRES_URL) up

test:
	@go test -v ./postgres/postgres_test.go
