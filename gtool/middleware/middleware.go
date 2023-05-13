package middleware

import "google.golang.org/grpc"

// InterceptorChain 创建一个拦截器链
var InterceptorChain = grpc.ChainUnaryInterceptor(PanicInterceptor, ValidateInterceptor)
