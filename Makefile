.PHONY: test

vet:
	go vet ./...

test: vet
	go test ./...
