# easymarket-go-server

# generate all

kratos tool protoc api.proto

# generate gRPC

kratos tool protoc --grpc api.proto

# generate BM HTTP

kratos tool protoc --bm api.proto

# generate ecode

kratos tool protoc --ecode api.proto

# generate swagger

kratos tool protoc --swagger api.proto
