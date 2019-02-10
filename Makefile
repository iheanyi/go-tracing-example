all: generate


generate:
	protoc --twirp_out=. --go_out=. ./rpc/pinger/service.proto
