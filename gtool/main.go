package main

import (
	"context"
	"gRpcTool/middleware"
	"gRpcTool/pb"
	"gRpcTool/pb/swagger"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":10580")
	if err != nil {
		log.Fatalln("failed to listen:", err)
	}

	// create gRPC server
	s := grpc.NewServer(
		middleware.InterceptorChain,
	)
	pb.RegisterGRPCToolServer(s, newGRPCToolImpl())

	// gRPC-Gateway mux
	gwmux := runtime.NewServeMux()
	dops := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pb.RegisterGRPCToolHandlerFromEndpoint(context.Background(), gwmux, "127.0.0.1:10580", dops)
	if err != nil {
		log.Fatalln("failed to register gwmux:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)

	// 文档
	// mux.HandleFunc("/swagger/", swaggerFile)
	// swaggerUI(mux)

	// 定义HTTP server配置
	gwServer := &http.Server{
		Addr:    "127.0.0.1:10580",
		Handler: grpcHandlerFunc(s, mux), // 请求的统一入口
	}
	log.Println("Serving on http://127.0.0.1:10580")
	log.Fatalln(gwServer.Serve(lis)) // 启动HTTP服务
}

// grpcHandlerFunc 将gRPC请求和HTTP请求分别调用不同的handler处理
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			// 跨域支持
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if r.Method == "OPTIONS" {
				return
			}
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

// serveSwaggerUI: 提供UI支持
func swaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}

// swaggerFile: 提供对swagger.json文件的访问支持
func swaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	name := path.Join("pb", p)
	log.Printf("serving swagger-file: %s", name)
	http.ServeFile(w, r, name)
}
