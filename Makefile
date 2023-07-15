.PHONY: start test

start:
	go run main.go


test:
	go test -v -cover -short ./...