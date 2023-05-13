package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator interface {
	Validate() error
}

// ValidateInterceptor 一元拦截器
func ValidateInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 请求结构没有接入 Validator
	valid, ok := req.(Validator)
	if !ok {
		return handler(ctx, req)
	}

	// 验证通过
	err := valid.Validate()
	if err == nil {
		return handler(ctx, req)
	}

	return nil, status.Error(codes.InvalidArgument, err.Error())
}

// StreamInterceptor 流式拦截器
func StreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, ss)
}
