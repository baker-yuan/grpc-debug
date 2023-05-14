module gRpcTool

go 1.19

require (
	github.com/elazarl/go-bindata-assetfs v1.0.1
	github.com/envoyproxy/protoc-gen-validate v0.9.1
	github.com/fullstorydev/grpcurl v1.8.7
	github.com/gogo/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.2
	github.com/jhump/protoreflect v1.12.0
	go.uber.org/zap v1.24.0
	golang.org/x/net v0.7.0
	google.golang.org/appengine v1.6.7
	google.golang.org/genproto v0.0.0-20230223222841-637eb2293923
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
)

replace (
	github.com/fullstorydev/grpcurl => /Users/yuanyu/code/go-study/grpcurl
	google.golang.org/grpc => /Users/yuanyu/code/go-study/grpc-go
)
