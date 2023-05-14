package logic

import (
	"context"
	"gRpcTool/pb"
)

type gRPCToolImpl struct {
}

func newGRPCToolImpl() *gRPCToolImpl {
	return &gRPCToolImpl{}
}

func (g gRPCToolImpl) ServerInfo(ctx context.Context, req *pb.ServerInfoReq) (*pb.Server, error) {
	refClient, err := NewRefServer(ctx, req.GetUrl())
	if err != nil {
		return nil, err
	}
	defer refClient.Close()
	service, err := refClient.ListService()
	if err != nil {
		return nil, err
	}
	return service, nil
}

func (g gRPCToolImpl) MethodParam(ctx context.Context, req *pb.MethodParamReq) (map[string]interface{}, error) {
	refClient, err := NewRefServer(ctx, req.GetUrl())
	if err != nil {
		return nil, err
	}
	defer refClient.Close()
	return refClient.GetParams(req.GetServiceName(), req.GetMethodName())
}

func (g gRPCToolImpl) CallUnaryMethod(ctx context.Context, req *pb.CallMethodReq) (map[string]interface{}, error) {
	refClient, err := NewRefServer(ctx, req.GetUrl())
	if err != nil {
		return nil, err
	}
	defer refClient.Close()
	return refClient.CallUnaryMethod(req.GetServiceName(), req.GetMethodName(), req.GetData())
}

func (g gRPCToolImpl) CallServerStreamMethod(ctx context.Context, req *pb.CallMethodReq) ([]map[string]interface{}, error) {
	refClient, err := NewRefServer(ctx, req.GetUrl())
	if err != nil {
		return nil, err
	}
	defer refClient.Close()
	return refClient.CallServerStreamMethod(req.GetServiceName(), req.GetMethodName(), req.GetData())
}

func (g gRPCToolImpl) CallClientStreamMethod(ctx context.Context, req *pb.CallClientStreamMethodReq) (map[string]interface{}, error) {
	refClient, err := NewRefServer(ctx, req.GetUrl())
	if err != nil {
		return nil, err
	}
	defer refClient.Close()
	return refClient.CallClientStreamMethod(req.GetServiceName(), req.GetMethodName(), req.GetData())
}

// 客户端缓存 key=url value=RefServer
var refClientCache = make(map[string]*RefServer, 0)

func (g gRPCToolImpl) CallBidirectionalStreamMethod(ctx context.Context, req *pb.CallBidirectionalStreamMethodReq) ([]map[string]interface{}, error) {
	// 从缓存中获取
	var refClient *RefServer
	if _, exist := refClientCache[req.GetUrl()]; exist {
		refClient = refClientCache[req.GetUrl()]
	} else {
		c, err := NewRefServer(ctx, req.GetUrl())
		if err != nil {
			return nil, err
		}
		refClient = c
	}
	// 方法结束，关闭连接，清空缓存
	if req.GetComplete() {
		defer refClient.Close()
		defer func() {
			delete(refClientCache, req.GetUrl())
		}()
	}
	return refClient.CallBidirectionalStreamMethod(req.GetServiceName(), req.GetMethodName(), req.GetData(), req.GetComplete())
}
