PHONY: generate-structs

generate-structs:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	export PATH="$PATH:$(go env GOPATH)/bin"
	mkdir -p gen/go/mr_rpc_v1
	protoc -I ./src/main/proto ./src/main/proto/mr_rpc_v1.proto --go_out=./gen/go/mr_rpc_v1/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/mr_rpc_v1/ --go-grpc_opt=paths=source_relative