all: test

.PHONY: clean
clean:

.PHONY: test
test:
	go vet ./...
	go test ./...

