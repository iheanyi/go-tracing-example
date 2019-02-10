all: generate


generate:
	protoc --proto_path=. --twirp_out=. --go_out=. ./rpc/pinger/service.proto
