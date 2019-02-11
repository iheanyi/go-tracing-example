all: generate

generate: generate_pinger generate_ponger

generate_pinger:
	protoc -I . --go_out=plugins=grpc:. ./rpc/pinger/service.proto

generate_ponger:
	protoc -I . --go_out=plugins=grpc:. ./rpc/ponger/service.proto
