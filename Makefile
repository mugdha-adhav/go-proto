generate-protobufs:
	protoc --proto_path=protobufs --go_out=pkg/protobufs --go_opt=paths=source_relative protobufs/main.proto

build:
	go build -o prog cmd/main.go

run: build
	./prog
