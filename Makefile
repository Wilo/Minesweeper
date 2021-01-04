# cross parameters
SHELL:=/bin/bash -O extglob
BINARY=app.bin
VERSION=0.1.0

LDFLAGS=-ldflags="-w -s -X main.Version=${VERSION}"
GCFLAGS=-gcflags "-N -l"

# Build step, generates the binary.
build:
	@echo "Building ${BINARY}"
	@env CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/${BINARY}
	@echo "binary file path bin/${BINARY}"

# Test all apps
# Run the test for all the directories.
test:
	@clear
	@go test -race -v ./...

# Run Test Coverage for all the directories.
coverage:
	@clear
	@go test -race -v ./... -coverprofile cover.out
	@go tool cover -html cover.out

# end test

# Clear binary folder.
clean clear:
	@echo "Remove binary app"
	@rm -frv bin/

# Build And DEBUG usign CGDB
debug:
	@go build ${GCFLAGS} -o bin/${BINARY} main.go
	@cgdb bin/${BINARY}