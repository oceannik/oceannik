CERTS_DIR := ~/.oceannik/certs

OCEAN_BINARY_NAME := ocean

BINARIES_OUTPUT_DIR := bin
OCEAN_CMD_SOURCE_DIR := cmd
PROTO_SOURCE_DIR := common/proto
GENERATED_CERTS_DIR := generated-certs

default: build

proto:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	$(PROTO_SOURCE_DIR)/*.proto

build-release:
	go build -o $(BINARIES_OUTPUT_DIR)/$(OCEAN_BINARY_NAME) -ldflags "-s -w" ./$(OCEAN_CMD_SOURCE_DIR)

build:
	go build -o $(BINARIES_OUTPUT_DIR)/$(OCEAN_BINARY_NAME) ./$(OCEAN_CMD_SOURCE_DIR)

gen-certs:
	scripts/generate-certs.sh

copy-certs:
	mkdir -p $(CERTS_DIR)
	cp -r $(GENERATED_CERTS_DIR)/* $(CERTS_DIR)/

server:
	go run ./$(OCEAN_CMD_SOURCE_DIR) server

clean:
	rm -r bin/
	rm $(PROTO_SOURCE_DIR)/*.pb.go

.PHONY:
	proto
	build-release
	build
	gen-certs
	copy-certs
	server
	clean
