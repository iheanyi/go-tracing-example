all: generate

generate: generate_pinger generate_ponger

generate_pinger:
	protoc --proto_path=. --twirp_out=. --go_out=. ./rpc/pinger/service.proto

generate_ponger:
	protoc --proto_path=. --twirp_out=. --go_out=. ./rpc/ponger/service.proto
