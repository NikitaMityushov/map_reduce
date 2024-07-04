PHONY: generate-structs

generate-structs:
	mkdir -p gen/go/mr_rpc
	protoc -I mr_rpc ./mr_rpc/mr_rpc.proto --go_out=./gen/go/mr_rpc/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/mr_rpc/ --go-grpc_opt=paths=source_relative