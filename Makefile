
.PHONY: build dep clean lint
.SILENT: tidy


build: go build -o event main.go

clean: go clean

test: go test ./...

dep: go mod download

lint: golangci-lint run --enable-all

build_and_run: build run