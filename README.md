# 一、gRPC服务开启反射

gprc在注册时需要开启反射服务，不然将无法使用，golang开启反射如下:

```golang
func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    // 创建 RPC 服务容器
    grpcServer := grpc.NewServer()
    
    // 为 User 服务注册业务实现 将 User 服务绑定到 RPC 服务容器上
    user.RegisterUserServer(grpcServer, &UserService{})
    
    // 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
    reflection.Register(grpcServer)

    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

  

go-zero框架只需要在etc文件里加入 `Mode: dev` 就能开启反射



# 二、使用

![使用.jpg](https://baker-blog.oss-cn-chengdu.aliyuncs.com/article/b149e17ede10d961.jpg)

# 三、反射调用

## 一元RPC

```bash
curl --location --request POST 'http://localhost:10580/gRPCTool/call' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "localhost:50051",
    "serviceName": "grpc.examples.echo.Echo",
    "methodName": "UnaryEcho",
    "data": "{\"message\":\"baker\"}"
}'
```



## 服务器端流式RPC

```bash
curl --location --request POST 'http://localhost:10580/gRPCTool/callServerStream' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "url": "localhost:50051",
    "serviceName": "grpc.examples.echo.Echo",
    "methodName": "ServerStreamingEcho",
    "data": "{\"message\":\"baker\"}"
}'
```

## 客户端流式

```bash
curl --location --request POST 'http://localhost:10580/gRPCTool/callClientStream' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "localhost:50051",
    "serviceName": "grpc.examples.echo.Echo",
    "methodName": "ClientStreamingEcho",
    "data": [
        "{\"message\":\"2\"}",
        "{\"message\":\"2\"}"
    ]
}'
```





# 四、参考
- https://github.com/hXoreyer/grpc-tool

