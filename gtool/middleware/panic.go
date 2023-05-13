package middleware

import (
	"context"
	"log"
	"runtime/debug"

	// "runtime/debug"
	"google.golang.org/grpc"
)

func PanicInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			// 从 metadata 中获取请求 ID
			// md, ok := metadata.FromIncomingContext(ctx)
			// var reqID string
			// if ok {
			// 	reqIDList := md["request-id"]
			// 	if len(reqIDList) > 0 {
			// 		reqID = reqIDList[0]
			// 	}
			// }

			// 打印错误日志
			// log.Printf("[PANIC] request ID=%s, error=%v", reqID, r)
			log.Printf("[PANIC], value: %+v, stack:\n%s", r, debug.Stack())

			// 返回一个错误给客户端
			// err := status.Errorf(codes.Internal, "internal server error")
			// panic(err)

			// err := status.Errorf(codes.Internal, "internal server error: %v", r)
			// md, ok := metadata.FromIncomingContext(ctx)
			// if !ok {
			// 	md = metadata.New(nil)
			// }
			// md.Set("panic-error", fmt.Sprintf("%v", r))
			// grpc.SetHeader(ctx, md)
			// grpc.SendHeader(ctx, md)
			// grpc.SetTrailer(ctx, md)
			// panic(err)

		}
	}()
	return handler(ctx, req)
}
