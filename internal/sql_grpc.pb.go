// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: ipc/sql.proto

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

// SqlIpcClient is the client API for SqlIpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SqlIpcClient interface {
	QuerySingleSql(ctx context.Context, in *ExecuteSqlRequest, opts ...grpc.CallOption) (*QuerySingleSqlResponse, error)
	ExecuteSingleSql(ctx context.Context, in *ExecuteSqlRequest, opts ...grpc.CallOption) (*ExecuteSingleSqlResponse, error)
	ExecuteMultipleSqls(ctx context.Context, in *ExecuteSqlRequest, opts ...grpc.CallOption) (*ExecuteMultipleSqlsResponse, error)
}

type sqlIpcClient struct {
	cc grpc.ClientConnInterface
}

func NewSqlIpcClient(cc grpc.ClientConnInterface) SqlIpcClient {
	return &sqlIpcClient{cc}
}

func (c *sqlIpcClient) QuerySingleSql(ctx context.Context, in *ExecuteSqlRequest, opts ...grpc.CallOption) (*QuerySingleSqlResponse, error) {
	out := new(QuerySingleSqlResponse)
	err := c.cc.Invoke(ctx, "/sql_ipc.SqlIpc/QuerySingleSql", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sqlIpcClient) ExecuteSingleSql(ctx context.Context, in *ExecuteSqlRequest, opts ...grpc.CallOption) (*ExecuteSingleSqlResponse, error) {
	out := new(ExecuteSingleSqlResponse)
	err := c.cc.Invoke(ctx, "/sql_ipc.SqlIpc/ExecuteSingleSql", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sqlIpcClient) ExecuteMultipleSqls(ctx context.Context, in *ExecuteSqlRequest, opts ...grpc.CallOption) (*ExecuteMultipleSqlsResponse, error) {
	out := new(ExecuteMultipleSqlsResponse)
	err := c.cc.Invoke(ctx, "/sql_ipc.SqlIpc/ExecuteMultipleSqls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SqlIpcServer is the server API for SqlIpc service.
// All implementations must embed UnimplementedSqlIpcServer
// for forward compatibility
type SqlIpcServer interface {
	QuerySingleSql(context.Context, *ExecuteSqlRequest) (*QuerySingleSqlResponse, error)
	ExecuteSingleSql(context.Context, *ExecuteSqlRequest) (*ExecuteSingleSqlResponse, error)
	ExecuteMultipleSqls(context.Context, *ExecuteSqlRequest) (*ExecuteMultipleSqlsResponse, error)
	mustEmbedUnimplementedSqlIpcServer()
}

// UnimplementedSqlIpcServer must be embedded to have forward compatible implementations.
type UnimplementedSqlIpcServer struct {
}

func (UnimplementedSqlIpcServer) QuerySingleSql(context.Context, *ExecuteSqlRequest) (*QuerySingleSqlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuerySingleSql not implemented")
}
func (UnimplementedSqlIpcServer) ExecuteSingleSql(context.Context, *ExecuteSqlRequest) (*ExecuteSingleSqlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteSingleSql not implemented")
}
func (UnimplementedSqlIpcServer) ExecuteMultipleSqls(context.Context, *ExecuteSqlRequest) (*ExecuteMultipleSqlsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteMultipleSqls not implemented")
}
func (UnimplementedSqlIpcServer) mustEmbedUnimplementedSqlIpcServer() {}

// UnsafeSqlIpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SqlIpcServer will
// result in compilation errors.
type UnsafeSqlIpcServer interface {
	mustEmbedUnimplementedSqlIpcServer()
}

func RegisterSqlIpcServer(s grpc.ServiceRegistrar, srv SqlIpcServer) {
	s.RegisterService(&SqlIpc_ServiceDesc, srv)
}

func _SqlIpc_QuerySingleSql_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteSqlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlIpcServer).QuerySingleSql(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sql_ipc.SqlIpc/QuerySingleSql",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlIpcServer).QuerySingleSql(ctx, req.(*ExecuteSqlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SqlIpc_ExecuteSingleSql_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteSqlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlIpcServer).ExecuteSingleSql(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sql_ipc.SqlIpc/ExecuteSingleSql",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlIpcServer).ExecuteSingleSql(ctx, req.(*ExecuteSqlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SqlIpc_ExecuteMultipleSqls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteSqlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SqlIpcServer).ExecuteMultipleSqls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sql_ipc.SqlIpc/ExecuteMultipleSqls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SqlIpcServer).ExecuteMultipleSqls(ctx, req.(*ExecuteSqlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SqlIpc_ServiceDesc is the grpc.ServiceDesc for SqlIpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SqlIpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sql_ipc.SqlIpc",
	HandlerType: (*SqlIpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QuerySingleSql",
			Handler:    _SqlIpc_QuerySingleSql_Handler,
		},
		{
			MethodName: "ExecuteSingleSql",
			Handler:    _SqlIpc_ExecuteSingleSql_Handler,
		},
		{
			MethodName: "ExecuteMultipleSqls",
			Handler:    _SqlIpc_ExecuteMultipleSqls_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ipc/sql.proto",
}
