.PHONY: test

vet:
	go vet ./...

test: vet
	go test ./...

cleanup:
	go run ./tools/cleanup.go
