BINARY=gocut
BIN_DIR=bin

.PHONY: build test integration lint clean

build:
	@mkdir -p $(BIN_DIR)
	go build -o ${BIN_DIR}/${BINARY} ./cmd/gocut

test:
	go test -v ./...
	go test -v -tags=integration ./integration

lint:
	go vet ./...
	golangci-lint run ./...

clean:
	rm -rf ${BIN_DIR}
	rm -f gocut