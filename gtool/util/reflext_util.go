package util

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
func (r *RefServer) CallUnaryMethod(serviceName, methodName, param string) (map[string]interface{}, error) {
	// 获取方法描述符
	methodDesc, err := getMethodDescriptor(r.RefClient, serviceName, methodName)
	if err != nil {
		return nil, err
	}

	// 使用grpcdynamic包中的NewStub方法创建了一个动态的gRPC客户端stub，可以通过这个stub来调用gRPC服务端中定义的方法。
	stub := grpcdynamic.NewStub(r.channel)

	// 调用一元rpc方法
	resp, err := stub.InvokeRpc(context.Background(), methodDesc, getProtoMessage(methodDesc, param))
	if err != nil {
		return nil, err
	}

	// 处理返回结果
	return messageToMap(resp), nil
}

// CallServerStreamMethod 反射调用服务器端流式RPC方法
func (r *RefServer) CallServerStreamMethod(serviceName, methodName, param string) ([]map[string]interface{}, error) {
	// 获取方法描述符
	methodDesc, err := getMethodDescriptor(r.RefClient, serviceName, methodName)
	if err != nil {
		return nil, err
	}

	stub := grpcdynamic.NewStub(r.channel)

	// 获取服务端流对象，serverStream对象可以用来接收服务端返回的多个响应结果。
	serverStream, err := stub.InvokeRpcServerStream(context.Background(), methodDesc, getProtoMessage(methodDesc, param))
	if err != nil {
		return nil, err
	}

	// 死循环接收服务端返回的消息
	ret := make([]map[string]interface{}, 0)
	for {
		var resp proto.Message
		resp, err = serverStream.RecvMsg()
		if err != nil {
			if err == io.EOF {
				return ret, nil
			}
			return nil, err
		}
		ret = append(ret, messageToMap(resp))
	}
}

// CallClientStreamMethod 反射调用客户端流式RPC方法
func (r *RefServer) CallClientStreamMethod(serviceName string, methodName string, params []string) (map[string]interface{}, error) {
	// 获取方法描述符
	methodDesc, err := getMethodDescriptor(r.RefClient, serviceName, methodName)
	if err != nil {
		return nil, err
	}

	// 获取客户端流对象
	stub := grpcdynamic.NewStub(r.channel)
	clientStream, err := stub.InvokeRpcClientStream(context.Background(), methodDesc)
	if err != nil {
		return nil, err
	}

	// 循环发送消息给服务端
	for _, param := range params {
		if err = clientStream.SendMsg(getProtoMessage(methodDesc, param)); err != nil {
			return nil, err
		}
	}
	var resp proto.Message
	resp, err = clientStream.CloseAndReceive()
	if err != nil {
		return nil, err
	}
	return messageToMap(resp), nil
}

// CallBidirectionalStreamMethod 双向流式RPC
func (r *RefServer) CallBidirectionalStreamMethod(serviceName string, methodName string, params []string, complete bool) ([]map[string]interface{}, error) {
	// 获取方法描述符
	methodDesc, err := getMethodDescriptor(r.RefClient, serviceName, methodName)
	if err != nil {
		return nil, err
	}

	// 获取双向流对象
	stub := grpcdynamic.NewStub(r.channel)
	bidiStream, err := stub.InvokeRpcBidiStream(context.Background(), methodDesc)
	if err != nil {
		return nil, err
	}

	// 点击停止发送请求
	if complete {
		err = bidiStream.CloseSend()
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	// 循环发送消息给服务端
	for _, param := range params {
		if err = bidiStream.SendMsg(getProtoMessage(methodDesc, param)); err != nil {
			return nil, err
		}
	}

	// 死循环接收服务端返回的消息
	ret := make([]map[string]interface{}, 0)
	for {
		var resp proto.Message
		// BidiStream.RecvMsg 方法是一个阻塞方法，它会等待直到收到下一个消息。
		// 如果你调用该方法后程序一直没有收到消息，那么该方法就会一直阻塞住程序，直到收到消息或者超时。
		// 如果你希望在等待消息的同时可以处理其他的任务，可以考虑使用异步方法，如BidiStream.ReadAsync()方法。该方法可以在等待消息的同时处理其他的任务，当收到消息时返回一个Task对象，你可以通过该对象来获取收到的消息。
		// 另外，你也可以在调用BidiStream.RecvMsg方法时设置超时时间，如果在指定的时间内没有收到消息，该方法会自动返回一个空消息。这可以避免程序被永久地阻塞住。
		resp, err = bidiStream.RecvMsg()
		if err != nil {
			if err == io.EOF {
				return ret, nil
			}
			return nil, err
		}
		ret = append(ret, messageToMap(resp))
	}

}

// 解析grpc返回的数据
func messageToMap(msg proto.Message) map[string]interface{} {
	// *dynamic.Message类型是一个protobuf库提供的动态消息类型，可以方便地对消息进行操作。
	// dMsg.MarshalJSON() 方法将消息转换成 JSON 格式的字节数组。
	dMsg := msg.(*dynamic.Message)
	jsonData, _ := dMsg.MarshalJSON()
	res := make(map[string]interface{})
	_ = json.Unmarshal(jsonData, &res)
	return res
}

// json入参转换成proto格式的参数
func getProtoMessage(methodDesc *desc.MethodDescriptor, param string) proto.Message {
	// 使用dynamic.NewMessage方法创建了一个空的消息对象req，该消息对象的类型是从methodDesc中获取的该方法的输入类型。
	// 接着，使用json.Unmarshal将json字符串解析到req中，填充请求消息的具体参数。
	req := dynamic.NewMessage(methodDesc.GetInputType())
	_ = req.UnmarshalJSON([]byte(param))
	return req
}
