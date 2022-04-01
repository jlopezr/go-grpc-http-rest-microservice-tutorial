protoc todo-service.proto --proto_path=api/proto/v1 --proto_path=third-party \
    --go_out=./pkg/api/v1 --go_opt=paths=source_relative \
    --go-grpc_out=./pkg/api/v1 --go-grpc_opt=paths=source_relative \
    --grpc-gateway_out ./pkg/api/v1 --grpc-gateway_opt paths=source_relative --grpc-gateway_opt logtostderr=true \
    --openapiv2_out ./pkg/api/v1 --openapiv2_opt logtostderr=true \