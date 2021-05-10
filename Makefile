proto:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	proto/*.proto

build-release:
	go build -o bin/ocean -ldflags "-s -w" ./cli

build:
	go build -o bin/ocean ./cli

gen-certs:
	scripts/generate-certs.sh

copy-certs:
	mkdir -p ~/.oceannik/certs
	cp -r generated-certs/* ~/.oceannik/certs/

server:
	go run ./cli server

clean:
	rm -r bin/
	rm proto/*.pb.go

.PHONY: proto build-release build certs server clean
