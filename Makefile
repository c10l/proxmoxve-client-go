vet:
	go vet ./...
.PHONY: vet

test: vet
	go test ./...
.PHONY: test

testnocache: vet
	go test -count=1 ./...
.PHONY: testnocache

cleanup:
	go run ./tools/cleanup.go
.PHONY: cleanup
