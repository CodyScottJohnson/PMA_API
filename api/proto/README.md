## Compile Protocol Buffer Files
protoc --protocol_path=api/protocol/v1 --go_out=plugins=grpc:pkg/api/v1 *.proto