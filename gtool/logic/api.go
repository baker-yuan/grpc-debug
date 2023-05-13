package logic

import (
	"context"
	"gRpcTool/pb"
)

// API gRPC反射逻辑层API
type API interface {
	// ServerInfo 获取服务信息
	ServerInfo(ctx context.Context, req *pb.ServerInfoReq) (*pb.Server, error)
	// MethodParam 获取方法参数
	MethodParam(context.Context, *pb.MethodParamReq) (map[string]interface{}, error)
	// CallUnaryMethod 一元RPC
	CallUnaryMethod(context.Context, *pb.CallMethodReq) (map[string]interface{}, error)
	// CallServerStreamMethod 服务器端流式RPC
	CallServerStreamMethod(ctx context.Context, in *pb.CallMethodReq) ([]map[string]interface{}, error)
	// CallClientStreamMethod 客户端流式RPC
	CallClientStreamMethod(context.Context, *pb.CallClientStreamMethodReq) (map[string]interface{}, error)
	// CallBidirectionalStreamMethod 双向流式RPC
	CallBidirectionalStreamMethod(ctx context.Context, req *pb.CallBidirectionalStreamMethodReq) ([]map[string]interface{}, error)
}

// New 实例化
func New() API {
	return newGRPCToolImpl()
}
