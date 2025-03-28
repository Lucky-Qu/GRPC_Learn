//声明协议版本

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: person.proto

//声明分包名称

package person

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TestService_TestFunc1_FullMethodName = "/person.TestService/TestFunc1"
	TestService_TestFunc2_FullMethodName = "/person.TestService/TestFunc2"
	TestService_TestFunc3_FullMethodName = "/person.TestService/TestFunc3"
	TestService_TestFunc4_FullMethodName = "/person.TestService/TestFunc4"
)

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	// 使用rpc声明时遵循格式为：rpc 方法名称(入参)returns(出参)
	// rpc声明的方法根据出入参是否为流分为以下四类
	// 出入参不为流
	TestFunc1(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (*PersonRes, error)
	// 入参为流
	TestFunc2(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PersonReq, PersonRes], error)
	// 出参为流
	TestFunc3(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PersonRes], error)
	// 出入参为流
	TestFunc4(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[PersonReq, PersonRes], error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestFunc1(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (*PersonRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PersonRes)
	err := c.cc.Invoke(ctx, TestService_TestFunc1_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testServiceClient) TestFunc2(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[PersonReq, PersonRes], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[0], TestService_TestFunc2_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PersonReq, PersonRes]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_TestFunc2Client = grpc.ClientStreamingClient[PersonReq, PersonRes]

func (c *testServiceClient) TestFunc3(ctx context.Context, in *PersonReq, opts ...grpc.CallOption) (grpc.ServerStreamingClient[PersonRes], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[1], TestService_TestFunc3_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PersonReq, PersonRes]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_TestFunc3Client = grpc.ServerStreamingClient[PersonRes]

func (c *testServiceClient) TestFunc4(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[PersonReq, PersonRes], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TestService_ServiceDesc.Streams[2], TestService_TestFunc4_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[PersonReq, PersonRes]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_TestFunc4Client = grpc.BidiStreamingClient[PersonReq, PersonRes]

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility.
type TestServiceServer interface {
	// 使用rpc声明时遵循格式为：rpc 方法名称(入参)returns(出参)
	// rpc声明的方法根据出入参是否为流分为以下四类
	// 出入参不为流
	TestFunc1(context.Context, *PersonReq) (*PersonRes, error)
	// 入参为流
	TestFunc2(grpc.ClientStreamingServer[PersonReq, PersonRes]) error
	// 出参为流
	TestFunc3(*PersonReq, grpc.ServerStreamingServer[PersonRes]) error
	// 出入参为流
	TestFunc4(grpc.BidiStreamingServer[PersonReq, PersonRes]) error
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestServiceServer struct{}

func (UnimplementedTestServiceServer) TestFunc1(context.Context, *PersonReq) (*PersonRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestFunc1 not implemented")
}
func (UnimplementedTestServiceServer) TestFunc2(grpc.ClientStreamingServer[PersonReq, PersonRes]) error {
	return status.Errorf(codes.Unimplemented, "method TestFunc2 not implemented")
}
func (UnimplementedTestServiceServer) TestFunc3(*PersonReq, grpc.ServerStreamingServer[PersonRes]) error {
	return status.Errorf(codes.Unimplemented, "method TestFunc3 not implemented")
}
func (UnimplementedTestServiceServer) TestFunc4(grpc.BidiStreamingServer[PersonReq, PersonRes]) error {
	return status.Errorf(codes.Unimplemented, "method TestFunc4 not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}
func (UnimplementedTestServiceServer) testEmbeddedByValue()                     {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	// If the following call pancis, it indicates UnimplementedTestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_TestFunc1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestFunc1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_TestFunc1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestFunc1(ctx, req.(*PersonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _TestService_TestFunc2_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).TestFunc2(&grpc.GenericServerStream[PersonReq, PersonRes]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_TestFunc2Server = grpc.ClientStreamingServer[PersonReq, PersonRes]

func _TestService_TestFunc3_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PersonReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServiceServer).TestFunc3(m, &grpc.GenericServerStream[PersonReq, PersonRes]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_TestFunc3Server = grpc.ServerStreamingServer[PersonRes]

func _TestService_TestFunc4_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServiceServer).TestFunc4(&grpc.GenericServerStream[PersonReq, PersonRes]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TestService_TestFunc4Server = grpc.BidiStreamingServer[PersonReq, PersonRes]

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "person.TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestFunc1",
			Handler:    _TestService_TestFunc1_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestFunc2",
			Handler:       _TestService_TestFunc2_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TestFunc3",
			Handler:       _TestService_TestFunc3_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "TestFunc4",
			Handler:       _TestService_TestFunc4_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "person.proto",
}
