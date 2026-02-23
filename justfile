# List available recipes
default:
    @just --list

# Run the test suite
test:
    go test -v -count=1 ./...

# Build the binary
build:
    go build -o rep .

# Install to ~/.local/bin (symlink to built binary)
install: build
    mkdir -p ~/.local/bin
    ln -sf {{justfile_directory()}}/rep ~/.local/bin/rep
