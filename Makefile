test:
	go test github.com/ecshreve/gelpers

testv:
	go test -v github.com/ecshreve/gelpers

testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/gelpers