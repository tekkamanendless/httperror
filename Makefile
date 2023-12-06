all: test

.PHONY: test
test:
	go vet ./...
	go test ./...

