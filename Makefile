brew:
	brew install aws-vault docker-credential-helper-ecr mysql-client@8.4 bufbuild/buf/buf

install_protoc: ## Installs tools for protobuf generation
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.19.1
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.19.1

make install_tools: brew install_protoc

make gen_protos:
	rm -rf pkg/proto
	buf generate protobuf -o pkg/proto

make run_tests:
	go test -v ./...

make run:
	go run cmd/main.go