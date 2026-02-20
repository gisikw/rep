# List available recipes
default:
    @just --list

# Run the test suite
test:
    go test -v -count=1 ./...

# Build the binary
build:
    go build -o crane .
