package util

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"go.uber.org/zap/buffer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

// GrpcUrl 反射服务
type GrpcUrl struct {
	url        string                   // grpc服务地址 ip:端口
	clientConn *grpc.ClientConn         // 连接
	refClient  *grpcreflect.Client      //
	refSource  grpcurl.DescriptorSource //
}

func NewGrpcUrl(ctx context.Context, url string) (*GrpcUrl, error) {
	var (
		clientConn *grpc.ClientConn
		err        error
	)
	if clientConn, err = grpc.DialContext(
		ctx,
		fmt.Sprintf("%s", url),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock(),
	); err != nil {
		return nil, err
	}
	refClient := grpcreflect.NewClient(context.Background(), reflectpb.NewServerReflectionClient(clientConn))
	return &GrpcUrl{
		url:        url,
		clientConn: clientConn,
		refClient:  refClient,
		refSource:  grpcurl.DescriptorSourceFromServer(context.Background(), refClient),
	}, nil
}

// Close 关闭连接
func (r *GrpcUrl) Close() {
	r.refClient.Reset()
	_ = r.clientConn.Close()
}

var (
	options = grpcurl.FormatOptions{
		AllowUnknownFields: true,
	}
)

// CallUnaryMethod 反射调用一元rpc方法
func (r *GrpcUrl) CallUnaryMethod(ctx context.Context, serviceName, methodName, param string) (map[string]interface{}, error) {
	rf, formatter, err := grpcurl.RequestParserAndFormatter(grpcurl.FormatJSON, r.refSource, strings.NewReader(param), options)
	if err != nil {
		return nil, err
	}
	// 元数据
	var md []string
	// 响应
	response := NewResponse()
	handler := &grpcurl.DefaultEventHandler{
		VerbosityLevel: 2,
		Out:            response,
		Formatter:      formatter,
	}
	// 反射调用
	if err = grpcurl.InvokeRPC(ctx, r.refSource, r.clientConn, fmt.Sprintf("%s/%s", serviceName, methodName), md, handler, rf.Next); err != nil {
		return nil, err
	}
	// 处理返回值
	res := make(map[string]interface{})
	_ = json.Unmarshal(response.Body(), &res)
	return res, nil
}

var (
	responseHeaderPre  = "\nResponse headers received:"
	responseContentPre = "\nResponse contents:"
	responseTrailerPre = "\nResponse trailers received:"
)

func NewResponse() *Response {
	return &Response{
		header:    make(map[string]string),
		bodyWrite: false,
		body:      &buffer.Buffer{},
	}
}

type Response struct {
	header    map[string]string
	bodyWrite bool
	body      *buffer.Buffer
}

func (r *Response) Write(p []byte) (n int, err error) {
	str := string(p)
	if strings.HasPrefix(str, responseHeaderPre) || strings.HasPrefix(str, responseTrailerPre) {
		headers := strings.Split(str, "\n")
		if len(headers) == 2 && strings.HasPrefix(headers[1], "(empty)") {
			return len(p), nil
		}
		for index, header := range headers {
			if index == 0 {
				continue
			}
			values := strings.Split(header, ":")
			var v string
			if len(values) > 1 {
				v = values[1]
			}
			r.header[values[0]] = v
		}
	}
	if strings.HasPrefix(str, responseContentPre) {
		r.bodyWrite = true
		return len(p), nil
	}
	if r.bodyWrite {
		r.body.Write(p)
		r.bodyWrite = false
	}
	return len(p), nil
}

func (r *Response) Body() []byte {
	return r.body.Bytes()
}

func (r *Response) Header() map[string]string {
	return r.header
}
