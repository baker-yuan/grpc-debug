package main

import (
	"context"
	"encoding/json"
	"gRpcTool/logic"
	"gRpcTool/pb"
	"net/http"

	"google.golang.org/appengine/log"
	"google.golang.org/protobuf/proto"
)

// GRPCToolImpl gRPC反射实现
type GRPCToolImpl struct {
	pb.UnimplementedGRPCToolServer
	logicAPI logic.API
}

func newGRPCToolImpl() *GRPCToolImpl {
	return &GRPCToolImpl{
		logicAPI: logic.New(),
	}
}

func (o GRPCToolImpl) ServerInfo(ctx context.Context, req *pb.ServerInfoReq) (*pb.ServerInfoRsp, error) {
	server, err := o.logicAPI.ServerInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp := &pb.ServerInfoRsp{
		Data:    server,
		Code:    proto.Uint32(http.StatusOK),
		Message: proto.String(http.StatusText(http.StatusOK)),
	}
	return rsp, err
}

func (o GRPCToolImpl) MethodParam(ctx context.Context, req *pb.MethodParamReq) (*pb.MethodParamRsp, error) {
	data, err := o.logicAPI.MethodParam(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp := &pb.MethodParamRsp{
		Data:    proto.String(objToJson(ctx, data)),
		Code:    proto.Uint32(http.StatusOK),
		Message: proto.String(http.StatusText(http.StatusOK)),
	}
	return rsp, err
}

func (o GRPCToolImpl) CallUnaryMethod(ctx context.Context, req *pb.CallMethodReq) (*pb.CallMethodRsp, error) {
	data, err := o.logicAPI.CallUnaryMethod(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp := &pb.CallMethodRsp{
		Data:    proto.String(objToJson(ctx, data)),
		Code:    proto.Uint32(http.StatusOK),
		Message: proto.String(http.StatusText(http.StatusOK)),
	}
	return rsp, err
}

func (o GRPCToolImpl) CallServerStreamMethod(ctx context.Context, req *pb.CallMethodReq) (*pb.CallServerStreamMethodRsp, error) {
	datas, err := o.logicAPI.CallServerStreamMethod(ctx, req)
	if err != nil {
		return nil, err
	}
	var res = make([]string, 0)
	for _, v := range datas {
		res = append(res, objToJson(ctx, v))
	}

	rsp := &pb.CallServerStreamMethodRsp{
		Data:    res,
		Code:    proto.Uint32(http.StatusOK),
		Message: proto.String(http.StatusText(http.StatusOK)),
	}
	return rsp, err
}
func (o GRPCToolImpl) CallClientStreamMethod(ctx context.Context, req *pb.CallClientStreamMethodReq) (*pb.CallMethodRsp, error) {
	data, err := o.logicAPI.CallClientStreamMethod(ctx, req)
	if err != nil {
		return nil, err
	}
	rsp := &pb.CallMethodRsp{
		Data:    proto.String(objToJson(ctx, data)),
		Code:    proto.Uint32(http.StatusOK),
		Message: proto.String(http.StatusText(http.StatusOK)),
	}
	return rsp, err
}

func objToJson(ctx context.Context, obj any) string {
	b, err := json.Marshal(obj)
	if err != nil {
		log.Errorf(ctx, "marshal fail req: %+v, err: %+v", obj, err)
		return ""
	}
	return string(b)
}
