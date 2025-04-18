.PHONY:
build:
	gofmt -l -s -w .
	go build -o bin/dbadmin ./cmd/dbadmin

.PHONY: run
run:
	./bin/dbadmin

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...