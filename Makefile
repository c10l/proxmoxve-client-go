.PHONY: test

vet:
	go vet ./...

test: vet
	go test ./...

testnocache: vet
	go test -count=1 ./...

cleanup:
	go run ./tools/cleanup.go
