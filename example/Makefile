all: clean proto build

build:
	go build .

proto:
	protoc -I. -I../ --go_out=. --go-grpc_out=. proto/example.proto

clean:
	rm -rf example

.PHONY: build proto clean
