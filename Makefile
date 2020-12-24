mod-init:
	go mod init github.com/ecshreve/gelpers

mod-tidy:
	go mod tidy
	
build:
	go build -o bin/gelpers github.com/ecshreve/gelpers/cmd/gelpers

install:
	go install -i github.com/ecshreve/gelpers/cmd/gelpers

run-only:
	bin/gelpers

run: build run-only

test:
	go test github.com/ecshreve/gelpers/...

testv:
	go test -v github.com/ecshreve/gelpers/...

testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/gelpers/...