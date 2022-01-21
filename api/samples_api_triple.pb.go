/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.5
// - protoc             v3.14.0
// source: samples_api.proto

package api

import (
	context "context"
	fmt "fmt"
)

import (
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"

	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"

	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc_go.CallOption) (*User, common.ErrorWithAttachment)
	// Sends a greeting via stream
	SayHelloStream(ctx context.Context, opts ...grpc_go.CallOption) (Greeter_SayHelloStreamClient, error)
}

type greeterClient struct {
	cc *triple.TripleConn
}

type GreeterClientImpl struct {
	SayHello       func(ctx context.Context, in *HelloRequest) (*User, error)
	SayHelloStream func(ctx context.Context) (Greeter_SayHelloStreamClient, error)
}

func (c *GreeterClientImpl) GetDubboStub(cc *triple.TripleConn) GreeterClient {
	return NewGreeterClient(cc)
}

func (c *GreeterClientImpl) XXX_InterfaceName() string {
	return "org.apache.dubbogo.samples.api.Greeter"
}

func NewGreeterClient(cc *triple.TripleConn) GreeterClient {
	return &greeterClient{cc}
}

//客户端的接口调用
func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc_go.CallOption) (*User, common.ErrorWithAttachment) {
	out := new(User)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SayHello", in, out)
}

func (c *greeterClient) SayHelloStream(ctx context.Context, opts ...grpc_go.CallOption) (Greeter_SayHelloStreamClient, error) {
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	stream, err := c.cc.NewStream(ctx, "/"+interfaceKey+"/SayHelloStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*User, error)
	grpc_go.ClientStream
}

type greeterSayHelloStreamClient struct {
	grpc_go.ClientStream
}

func (x *greeterSayHelloStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloStreamClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*User, error)
	// Sends a greeting via stream
	SayHelloStream(Greeter_SayHelloStreamServer) error
	mustEmbedUnimplementedGreeterServer()
}

// 服务端的接口调用
// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) SayHelloStream(Greeter_SayHelloStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStream not implemented")
}
func (s *UnimplementedGreeterServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedGreeterServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedGreeterServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &Greeter_ServiceDesc
}
func (s *UnimplementedGreeterServer) XXX_InterfaceName() string {
	return "org.apache.dubbogo.samples.api.Greeter"
}

func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc_go.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("SayHello", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloStream_Handler(srv interface{}, stream grpc_go.ServerStream) error {
	_, ok := srv.(dubbo3.Dubbo3GrpcService)
	invo := invocation.NewRPCInvocation("SayHelloStream", nil, nil)
	if !ok {
		fmt.Println(invo)
		return nil
	}
	return srv.(GreeterServer).SayHelloStream(&greeterSayHelloStreamServer{stream})
}

type Greeter_SayHelloStreamServer interface {
	Send(*User) error
	Recv() (*HelloRequest, error)
	grpc_go.ServerStream
}

type greeterSayHelloStreamServer struct {
	grpc_go.ServerStream
}

func (x *greeterSayHelloStreamServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc_go.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "org.apache.dubbogo.samples.api.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams: []grpc_go.StreamDesc{
		{
			StreamName:    "SayHelloStream",
			Handler:       _Greeter_SayHelloStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "samples_api.proto",
}
