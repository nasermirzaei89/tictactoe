#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export CGO_ENABLED=0

default: build-terminal build-wasm build-server test

format:
	gofmt -s -w $(ROOT)

build-terminal:
	go build -o bin/terminal -ldflags="-s -w" $(ROOT)/cmd/terminal/*.go

build-wasm:
	GOOS=js GOARCH=wasm go build -o wasm/app.wasm $(ROOT)/cmd/wasm/*.go

build-server:
	go build -o bin/server -ldflags="-s -w" $(ROOT)/cmd/server/*.go

test:
	CGO_ENABLED=1 go test -race -coverprofile=coverage.txt -covermode=atomic `go list ./... | grep -v tictactoe/cmd/wasm`

run-server: build-wasm build-server
	API_ADDR=0.0.0.0:8080 $(ROOT)/bin/server $(ROOT)/wasm
