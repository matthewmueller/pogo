run: bin
	go run cmd/pogo/pogo.go -db postgres://localhost:5432/pogo?sslmode=disable -schema jack -path jack
.PHONY: run

examples:  bin
	# go run cmd/pogo/pogo.go -db postgres://localhost:5432/tempo_dev?sslmode=disable -path tempo
	# go run cmd/pogo/pogo.go -db postgres://localhost:5432/bot-ii?sslmode=disable -path jack
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
