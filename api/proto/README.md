## Compile Protocol Buffer Files
protoc --proto_path=api/proto/v1 --go_out=plugins=grpc:pkg/api/v1 projects.proto