package logic

import (
	"context"
	"encoding/json"
	"gRpcTool/pb"
	"io"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/dynamic/grpcdynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/protobuf/types/descriptorpb"
)

// RefServer 反射服务
type RefServer struct {
	RefClient *grpcreflect.Client // 反射客户端
	url       string              // grpc服务地址 ip:端口
	channel   *grpc.ClientConn    // 连接
}

func NewRefServer(ctx context.Context, url string) (*RefServer, error) {
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	refCli := &RefServer{
		RefClient: grpcreflect.NewClient(ctx, reflectpb.NewServerReflectionClient(cc)),
		url:       url,
		channel:   cc,
	}
	return refCli, nil
}

func (r *RefServer) ListService() (*pb.Server, error) {
	var (
		listServices []string
		err          error
	)
	if listServices, err = r.RefClient.ListServices(); err != nil {
		return nil, err
	}
	// 去掉系统service
	for i := 0; i < len(listServices); i++ {
		if listServices[i] == "grpc.health.v1.Health" || listServices[i] == "grpc.reflection.v1alpha.ServerReflection" {
			listServices = append(listServices[:i], listServices[i+1:]...)
			i--
		}
	}
	// 获取service信息
	var (
		serviceDescriptor *desc.ServiceDescriptor
		server            *pb.Server
		methods           []*desc.MethodDescriptor
	)
	server = &pb.Server{
		Key:      proto.String(r.url),
		Url:      proto.String(r.url),
		Services: make([]*pb.Service, 0),
	}
	for _, s := range listServices {
		if serviceDescriptor, err = r.RefClient.ResolveService(s); err != nil {
			return nil, err
		}
		service := &pb.Service{
			ServiceName: proto.String(serviceDescriptor.GetFullyQualifiedName()),
		}
		service.Key = proto.String(server.GetKey() + ":" + service.GetServiceName())
		methods = serviceDescriptor.GetMethods()
		for _, m := range methods {
			method := &pb.Method{
				MethodName: proto.String(m.GetName()),
				MethodType: getMethodType(m).Enum(),
				InputType:  m.GetInputType().GetName(),
				OutputType: m.GetOutputType().GetName(),
			}
			method.Key = proto.String(service.GetKey() + ":" + method.GetMethodName())
			service.Methods = append(service.Methods, method)
		}
		server.Services = append(server.Services, service)
	}
	return server, nil
}

// 获取方法类型
func getMethodType(m *desc.MethodDescriptor) pb.MethodType {
	if m.IsClientStreaming() && m.IsServerStreaming() {
		return pb.MethodType_BidirectionalStreaming
	}
	if m.IsClientStreaming() {
		return pb.MethodType_ClientStreaming
	}
	if m.IsServerStreaming() {
		return pb.MethodType_ServerStreaming
	}
	return pb.MethodType_Unary
}

// getMethodDescriptor 获取方法描述符
func getMethodDescriptor(refClient *grpcreflect.Client, serviceName, methodName string) (*desc.MethodDescriptor, error) {
	var st *desc.MethodDescriptor
	s, e := refClient.ResolveService(serviceName)
	if e != nil {
		return nil, e
	}
	for _, v := range s.GetMethods() {
		if v.GetName() == methodName {
			st = v
			break
		}
	}
	return st, nil
}

// Close 关闭连接
func (r *RefServer) Close() {
	r.channel.Close()
}

// GetParams 获取方法参数
func (r *RefServer) GetParams(serviceName, methodName string) (map[string]interface{}, error) {
	methodDesc, err := getMethodDescriptor(r.RefClient, serviceName, methodName)
	if err != nil {
		return nil, err
	}
	param := make(map[string]interface{})
	for _, fieldDescriptor := range methodDesc.GetInputType().GetFields() {
		fieldName := fieldDescriptor.GetName()
		// 数组返回nil
		if fieldDescriptor.IsRepeated() {
			param[fieldName] = nil
			continue
		}
		// 消息类型，递归调用convertMessageToMap函数将该消息类型转换为一个字典
		switch fieldDescriptor.GetType() {
		case descriptorpb.FieldDescriptorProto_Type(descriptor.FieldDescriptorProto_TYPE_MESSAGE):
			param[fieldName] = convertMessageToMap(fieldDescriptor.GetMessageType())
			continue
		}
		// 基本类型
		param[fieldName] = fieldDescriptor.GetDefaultValue()
	}
	return param, nil
}

func convertMessageToMap(message *desc.MessageDescriptor) map[string]interface{} {
	m := make(map[string]interface{})
	for _, fieldDescriptor := range message.GetFields() {
		fieldName := fieldDescriptor.GetName()
		if fieldDescriptor.IsRepeated() {
			// 数组返回 nil
			m[fieldName] = nil
			continue
		}
		switch fieldDescriptor.GetType() {
		case descriptorpb.FieldDescriptorProto_Type(descriptor.FieldDescriptorProto_TYPE_MESSAGE):
			m[fieldName] = convertMessageToMap(fieldDescriptor.GetMessageType())
			continue
		}
		m[fieldName] = fieldDescriptor.GetDefaultValue()
	}
	return m
}

// CallUnaryMethod 反射调用一元rpc方法
func (r *RefServer) CallUnaryMethod(serviceName, methodName, jsonString string) (map[string]interface{}, error) {
	methodDesc, err := getMethodDescriptor(r.RefClient, serviceName, methodName)
	if err != nil {
		return nil, err
	}

	req := dynamic.NewMessage(methodDesc.GetInputType())
	req.UnmarshalJSON([]byte(jsonString))
	stub := grpcdynamic.NewStub(r.channel)

	resp, err := stub.InvokeRpc(context.Background(), methodDesc, req)
	if err != nil {
		return nil, err
	}
	return messageToMap(resp), nil
}

// CallServerStreamMethod 反射调用服务器端流式RPC方法
func (r *RefServer) CallServerStreamMethod(serviceName, methodName, param string) ([]map[string]interface{}, error) {
	methodDesc, err := getMethodDescriptor(r.RefClient, serviceName, methodName)
	if err != nil {
		return nil, err
	}

	req := dynamic.NewMessage(methodDesc.GetInputType())
	req.UnmarshalJSON([]byte(param))
	stub := grpcdynamic.NewStub(r.channel)

	str, err := stub.InvokeRpcServerStream(context.Background(), methodDesc, req)
	if err != nil {
		return nil, err
	}
	ret := make([]map[string]interface{}, 0)

	for {
		resp, err := str.RecvMsg()
		if err != nil {
			if err == io.EOF {
				return ret, nil
			}
			return nil, err
		}
		ret = append(ret, messageToMap(resp))
	}
}

func (r *RefServer) CallClientStreamMethod(serviceName string, methodName string, params []string) (map[string]interface{}, error) {
	methodDesc, err := getMethodDescriptor(r.RefClient, serviceName, methodName)
	if err != nil {
		return nil, err
	}

	stub := grpcdynamic.NewStub(r.channel)
	streamClient, err := stub.InvokeRpcClientStream(context.Background(), methodDesc)
	if err != nil {
		return nil, err
	}
	var resp proto.Message
	for _, param := range params {
		req := dynamic.NewMessage(methodDesc.GetInputType())
		_ = req.UnmarshalJSON([]byte(param))
		err = streamClient.SendMsg(req)
		resp, err = streamClient.CloseAndReceive()
		if err != nil {
			return nil, err
		}
		return messageToMap(resp), nil

	}
	return nil, nil
}

func messageToMap(msg proto.Message) map[string]interface{} {
	res := make(map[string]interface{})
	js := msg.(*dynamic.Message)
	ty, _ := js.MarshalJSON()
	_ = json.Unmarshal(ty, &res)
	return res
}
