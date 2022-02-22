set dotenv-load := false

# Show all available recipes
default:
  @just --list --unsorted

# Compile `awf` into an executable binary.
build:
	go build -o dist/server main.go

# Run the compiled `awf` server.
run: build
	./dist/server | sed 's/\x1b\[[0-9;]*m//g'

# Run `awf` with automatic reloading when code changes.
watch:
	$(go env GOPATH)/bin/reflex -d none -s -r '\.go$$' just run

# Build and tag the `awf` Docker image.
docker-build:
    docker build -t awf .

# Run the compiled `awf` server as a Docker container.
docker-run cmd="":
    docker run -it --rm -p 8888:3000 awf {{ cmd }}
