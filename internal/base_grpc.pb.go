// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: ipc/base.proto

package internal

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BaseIpcClient is the client API for BaseIpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseIpcClient interface {
	Capabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesResponse, error)
	ConsumeTask(ctx context.Context, in *ConsumeTaskRequest, opts ...grpc.CallOption) (BaseIpc_ConsumeTaskClient, error)
	FinishWithResult(ctx context.Context, in *ExecuteResult, opts ...grpc.CallOption) (*FeedResultResponse, error)
	FinishWithFailure(ctx context.Context, in *ExecuteFailure, opts ...grpc.CallOption) (*FeedResultResponse, error)
	GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error)
	SaveLog(ctx context.Context, in *SaveLogRequest, opts ...grpc.CallOption) (*SaveLogResponse, error)
}

type baseIpcClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseIpcClient(cc grpc.ClientConnInterface) BaseIpcClient {
	return &baseIpcClient{cc}
}

func (c *baseIpcClient) Capabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesResponse, error) {
	out := new(CapabilitiesResponse)
	err := c.cc.Invoke(ctx, "/base_ipc.BaseIpc/Capabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseIpcClient) ConsumeTask(ctx context.Context, in *ConsumeTaskRequest, opts ...grpc.CallOption) (BaseIpc_ConsumeTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &BaseIpc_ServiceDesc.Streams[0], "/base_ipc.BaseIpc/ConsumeTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &baseIpcConsumeTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BaseIpc_ConsumeTaskClient interface {
	Recv() (*ConsumeTaskResponse, error)
	grpc.ClientStream
}

type baseIpcConsumeTaskClient struct {
	grpc.ClientStream
}

func (x *baseIpcConsumeTaskClient) Recv() (*ConsumeTaskResponse, error) {
	m := new(ConsumeTaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *baseIpcClient) FinishWithResult(ctx context.Context, in *ExecuteResult, opts ...grpc.CallOption) (*FeedResultResponse, error) {
	out := new(FeedResultResponse)
	err := c.cc.Invoke(ctx, "/base_ipc.BaseIpc/FinishWithResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseIpcClient) FinishWithFailure(ctx context.Context, in *ExecuteFailure, opts ...grpc.CallOption) (*FeedResultResponse, error) {
	out := new(FeedResultResponse)
	err := c.cc.Invoke(ctx, "/base_ipc.BaseIpc/FinishWithFailure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseIpcClient) GetSchema(ctx context.Context, in *GetSchemaRequest, opts ...grpc.CallOption) (*GetSchemaResponse, error) {
	out := new(GetSchemaResponse)
	err := c.cc.Invoke(ctx, "/base_ipc.BaseIpc/GetSchema", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *baseIpcClient) SaveLog(ctx context.Context, in *SaveLogRequest, opts ...grpc.CallOption) (*SaveLogResponse, error) {
	out := new(SaveLogResponse)
	err := c.cc.Invoke(ctx, "/base_ipc.BaseIpc/SaveLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BaseIpcServer is the server API for BaseIpc service.
// All implementations must embed UnimplementedBaseIpcServer
// for forward compatibility
type BaseIpcServer interface {
	Capabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesResponse, error)
	ConsumeTask(*ConsumeTaskRequest, BaseIpc_ConsumeTaskServer) error
	FinishWithResult(context.Context, *ExecuteResult) (*FeedResultResponse, error)
	FinishWithFailure(context.Context, *ExecuteFailure) (*FeedResultResponse, error)
	GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error)
	SaveLog(context.Context, *SaveLogRequest) (*SaveLogResponse, error)
	mustEmbedUnimplementedBaseIpcServer()
}

// UnimplementedBaseIpcServer must be embedded to have forward compatible implementations.
type UnimplementedBaseIpcServer struct {
}

func (UnimplementedBaseIpcServer) Capabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Capabilities not implemented")
}
func (UnimplementedBaseIpcServer) ConsumeTask(*ConsumeTaskRequest, BaseIpc_ConsumeTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method ConsumeTask not implemented")
}
func (UnimplementedBaseIpcServer) FinishWithResult(context.Context, *ExecuteResult) (*FeedResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishWithResult not implemented")
}
func (UnimplementedBaseIpcServer) FinishWithFailure(context.Context, *ExecuteFailure) (*FeedResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishWithFailure not implemented")
}
func (UnimplementedBaseIpcServer) GetSchema(context.Context, *GetSchemaRequest) (*GetSchemaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSchema not implemented")
}
func (UnimplementedBaseIpcServer) SaveLog(context.Context, *SaveLogRequest) (*SaveLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLog not implemented")
}
func (UnimplementedBaseIpcServer) mustEmbedUnimplementedBaseIpcServer() {}

// UnsafeBaseIpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseIpcServer will
// result in compilation errors.
type UnsafeBaseIpcServer interface {
	mustEmbedUnimplementedBaseIpcServer()
}

func RegisterBaseIpcServer(s grpc.ServiceRegistrar, srv BaseIpcServer) {
	s.RegisterService(&BaseIpc_ServiceDesc, srv)
}

func _BaseIpc_Capabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseIpcServer).Capabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base_ipc.BaseIpc/Capabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseIpcServer).Capabilities(ctx, req.(*CapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseIpc_ConsumeTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConsumeTaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BaseIpcServer).ConsumeTask(m, &baseIpcConsumeTaskServer{stream})
}

type BaseIpc_ConsumeTaskServer interface {
	Send(*ConsumeTaskResponse) error
	grpc.ServerStream
}

type baseIpcConsumeTaskServer struct {
	grpc.ServerStream
}

func (x *baseIpcConsumeTaskServer) Send(m *ConsumeTaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BaseIpc_FinishWithResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseIpcServer).FinishWithResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base_ipc.BaseIpc/FinishWithResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseIpcServer).FinishWithResult(ctx, req.(*ExecuteResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseIpc_FinishWithFailure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteFailure)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseIpcServer).FinishWithFailure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base_ipc.BaseIpc/FinishWithFailure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseIpcServer).FinishWithFailure(ctx, req.(*ExecuteFailure))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseIpc_GetSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSchemaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseIpcServer).GetSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base_ipc.BaseIpc/GetSchema",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseIpcServer).GetSchema(ctx, req.(*GetSchemaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BaseIpc_SaveLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BaseIpcServer).SaveLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base_ipc.BaseIpc/SaveLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BaseIpcServer).SaveLog(ctx, req.(*SaveLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BaseIpc_ServiceDesc is the grpc.ServiceDesc for BaseIpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BaseIpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "base_ipc.BaseIpc",
	HandlerType: (*BaseIpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Capabilities",
			Handler:    _BaseIpc_Capabilities_Handler,
		},
		{
			MethodName: "FinishWithResult",
			Handler:    _BaseIpc_FinishWithResult_Handler,
		},
		{
			MethodName: "FinishWithFailure",
			Handler:    _BaseIpc_FinishWithFailure_Handler,
		},
		{
			MethodName: "GetSchema",
			Handler:    _BaseIpc_GetSchema_Handler,
		},
		{
			MethodName: "SaveLog",
			Handler:    _BaseIpc_SaveLog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConsumeTask",
			Handler:       _BaseIpc_ConsumeTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "ipc/base.proto",
}
