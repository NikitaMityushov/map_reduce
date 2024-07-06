PHONY: generate-structs

generate-structs:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	mkdir -p gen/go/mr_rpc_v1
	protoc -I mr_rpc ./mr_rpc/mr_rpc_v1.proto --go_out=./gen/go/mr_rpc_v1/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/mr_rpc_v1/ --go-grpc_opt=paths=source_relative