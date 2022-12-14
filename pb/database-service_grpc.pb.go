// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: database-service.proto

package pb

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

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetPerson(ctx context.Context, in *GetPersonReq, opts ...grpc.CallOption) (*GetPersonResp, error)
	GetRdsInstance(ctx context.Context, in *GetRdsInstanceReq, opts ...grpc.CallOption) (*GetRdsInstanceResp, error)
	GetRedisInstance(ctx context.Context, in *GetRedisInstanceReq, opts ...grpc.CallOption) (*GetRedisInstanceResp, error)
	GetDmsListOrder(ctx context.Context, in *GetDmsListOrderReq, opts ...grpc.CallOption) (*GetDmsListOrderResp, error)
	GetDmsListUsers(ctx context.Context, in *GetDmsListUsersReq, opts ...grpc.CallOption) (*GetDmsListUsersResp, error)
	GetDmsListInstance(ctx context.Context, in *GetDmsListInstanceReq, opts ...grpc.CallOption) (*GetDmsListInstanceResp, error)
	GetDmsSearchDatabase(ctx context.Context, in *GetDmsSearchDatabaseReq, opts ...grpc.CallOption) (*GetDmsSearchDatabaseResp, error)
	GetDbsBackupPlan(ctx context.Context, in *GetDbsBackupPlanReq, opts ...grpc.CallOption) (*GetDbsBackupPlanResp, error)
	GetDtsSyncJobName(ctx context.Context, in *GetDtsSyncJobNameReq, opts ...grpc.CallOption) (*GetDtsSyncJobNameResp, error)
	GetRdsSlowLog(ctx context.Context, in *GetRdsSlowLogReq, opts ...grpc.CallOption) (*GetRdsSlowLogResp, error)
}

type databaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseServiceClient(cc grpc.ClientConnInterface) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetPerson(ctx context.Context, in *GetPersonReq, opts ...grpc.CallOption) (*GetPersonResp, error) {
	out := new(GetPersonResp)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/GetPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetRdsInstance(ctx context.Context, in *GetRdsInstanceReq, opts ...grpc.CallOption) (*GetRdsInstanceResp, error) {
	out := new(GetRdsInstanceResp)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/GetRdsInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetRedisInstance(ctx context.Context, in *GetRedisInstanceReq, opts ...grpc.CallOption) (*GetRedisInstanceResp, error) {
	out := new(GetRedisInstanceResp)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/GetRedisInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetDmsListOrder(ctx context.Context, in *GetDmsListOrderReq, opts ...grpc.CallOption) (*GetDmsListOrderResp, error) {
	out := new(GetDmsListOrderResp)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/GetDmsListOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetDmsListUsers(ctx context.Context, in *GetDmsListUsersReq, opts ...grpc.CallOption) (*GetDmsListUsersResp, error) {
	out := new(GetDmsListUsersResp)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/GetDmsListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetDmsListInstance(ctx context.Context, in *GetDmsListInstanceReq, opts ...grpc.CallOption) (*GetDmsListInstanceResp, error) {
	out := new(GetDmsListInstanceResp)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/GetDmsListInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetDmsSearchDatabase(ctx context.Context, in *GetDmsSearchDatabaseReq, opts ...grpc.CallOption) (*GetDmsSearchDatabaseResp, error) {
	out := new(GetDmsSearchDatabaseResp)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/GetDmsSearchDatabase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetDbsBackupPlan(ctx context.Context, in *GetDbsBackupPlanReq, opts ...grpc.CallOption) (*GetDbsBackupPlanResp, error) {
	out := new(GetDbsBackupPlanResp)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/GetDbsBackupPlan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetDtsSyncJobName(ctx context.Context, in *GetDtsSyncJobNameReq, opts ...grpc.CallOption) (*GetDtsSyncJobNameResp, error) {
	out := new(GetDtsSyncJobNameResp)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/GetDtsSyncJobName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) GetRdsSlowLog(ctx context.Context, in *GetRdsSlowLogReq, opts ...grpc.CallOption) (*GetRdsSlowLogResp, error) {
	out := new(GetRdsSlowLogResp)
	err := c.cc.Invoke(ctx, "/databaseservice.DatabaseService/GetRdsSlowLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
// All implementations must embed UnimplementedDatabaseServiceServer
// for forward compatibility
type DatabaseServiceServer interface {
	Ping(context.Context, *Request) (*Response, error)
	GetPerson(context.Context, *GetPersonReq) (*GetPersonResp, error)
	GetRdsInstance(context.Context, *GetRdsInstanceReq) (*GetRdsInstanceResp, error)
	GetRedisInstance(context.Context, *GetRedisInstanceReq) (*GetRedisInstanceResp, error)
	GetDmsListOrder(context.Context, *GetDmsListOrderReq) (*GetDmsListOrderResp, error)
	GetDmsListUsers(context.Context, *GetDmsListUsersReq) (*GetDmsListUsersResp, error)
	GetDmsListInstance(context.Context, *GetDmsListInstanceReq) (*GetDmsListInstanceResp, error)
	GetDmsSearchDatabase(context.Context, *GetDmsSearchDatabaseReq) (*GetDmsSearchDatabaseResp, error)
	GetDbsBackupPlan(context.Context, *GetDbsBackupPlanReq) (*GetDbsBackupPlanResp, error)
	GetDtsSyncJobName(context.Context, *GetDtsSyncJobNameReq) (*GetDtsSyncJobNameResp, error)
	GetRdsSlowLog(context.Context, *GetRdsSlowLogReq) (*GetRdsSlowLogResp, error)
	mustEmbedUnimplementedDatabaseServiceServer()
}

// UnimplementedDatabaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDatabaseServiceServer struct {
}

func (UnimplementedDatabaseServiceServer) Ping(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedDatabaseServiceServer) GetPerson(context.Context, *GetPersonReq) (*GetPersonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPerson not implemented")
}
func (UnimplementedDatabaseServiceServer) GetRdsInstance(context.Context, *GetRdsInstanceReq) (*GetRdsInstanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRdsInstance not implemented")
}
func (UnimplementedDatabaseServiceServer) GetRedisInstance(context.Context, *GetRedisInstanceReq) (*GetRedisInstanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRedisInstance not implemented")
}
func (UnimplementedDatabaseServiceServer) GetDmsListOrder(context.Context, *GetDmsListOrderReq) (*GetDmsListOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDmsListOrder not implemented")
}
func (UnimplementedDatabaseServiceServer) GetDmsListUsers(context.Context, *GetDmsListUsersReq) (*GetDmsListUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDmsListUsers not implemented")
}
func (UnimplementedDatabaseServiceServer) GetDmsListInstance(context.Context, *GetDmsListInstanceReq) (*GetDmsListInstanceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDmsListInstance not implemented")
}
func (UnimplementedDatabaseServiceServer) GetDmsSearchDatabase(context.Context, *GetDmsSearchDatabaseReq) (*GetDmsSearchDatabaseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDmsSearchDatabase not implemented")
}
func (UnimplementedDatabaseServiceServer) GetDbsBackupPlan(context.Context, *GetDbsBackupPlanReq) (*GetDbsBackupPlanResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDbsBackupPlan not implemented")
}
func (UnimplementedDatabaseServiceServer) GetDtsSyncJobName(context.Context, *GetDtsSyncJobNameReq) (*GetDtsSyncJobNameResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDtsSyncJobName not implemented")
}
func (UnimplementedDatabaseServiceServer) GetRdsSlowLog(context.Context, *GetRdsSlowLogReq) (*GetRdsSlowLogResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRdsSlowLog not implemented")
}
func (UnimplementedDatabaseServiceServer) mustEmbedUnimplementedDatabaseServiceServer() {}

// UnsafeDatabaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServiceServer will
// result in compilation errors.
type UnsafeDatabaseServiceServer interface {
	mustEmbedUnimplementedDatabaseServiceServer()
}

func RegisterDatabaseServiceServer(s grpc.ServiceRegistrar, srv DatabaseServiceServer) {
	s.RegisterService(&DatabaseService_ServiceDesc, srv)
}

func _DatabaseService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).Ping(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPersonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/GetPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetPerson(ctx, req.(*GetPersonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetRdsInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRdsInstanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetRdsInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/GetRdsInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetRdsInstance(ctx, req.(*GetRdsInstanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetRedisInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRedisInstanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetRedisInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/GetRedisInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetRedisInstance(ctx, req.(*GetRedisInstanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetDmsListOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDmsListOrderReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetDmsListOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/GetDmsListOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetDmsListOrder(ctx, req.(*GetDmsListOrderReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetDmsListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDmsListUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetDmsListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/GetDmsListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetDmsListUsers(ctx, req.(*GetDmsListUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetDmsListInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDmsListInstanceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetDmsListInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/GetDmsListInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetDmsListInstance(ctx, req.(*GetDmsListInstanceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetDmsSearchDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDmsSearchDatabaseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetDmsSearchDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/GetDmsSearchDatabase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetDmsSearchDatabase(ctx, req.(*GetDmsSearchDatabaseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetDbsBackupPlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDbsBackupPlanReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetDbsBackupPlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/GetDbsBackupPlan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetDbsBackupPlan(ctx, req.(*GetDbsBackupPlanReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetDtsSyncJobName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDtsSyncJobNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetDtsSyncJobName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/GetDtsSyncJobName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetDtsSyncJobName(ctx, req.(*GetDtsSyncJobNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_GetRdsSlowLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRdsSlowLogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).GetRdsSlowLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/databaseservice.DatabaseService/GetRdsSlowLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).GetRdsSlowLog(ctx, req.(*GetRdsSlowLogReq))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseService_ServiceDesc is the grpc.ServiceDesc for DatabaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "databaseservice.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _DatabaseService_Ping_Handler,
		},
		{
			MethodName: "GetPerson",
			Handler:    _DatabaseService_GetPerson_Handler,
		},
		{
			MethodName: "GetRdsInstance",
			Handler:    _DatabaseService_GetRdsInstance_Handler,
		},
		{
			MethodName: "GetRedisInstance",
			Handler:    _DatabaseService_GetRedisInstance_Handler,
		},
		{
			MethodName: "GetDmsListOrder",
			Handler:    _DatabaseService_GetDmsListOrder_Handler,
		},
		{
			MethodName: "GetDmsListUsers",
			Handler:    _DatabaseService_GetDmsListUsers_Handler,
		},
		{
			MethodName: "GetDmsListInstance",
			Handler:    _DatabaseService_GetDmsListInstance_Handler,
		},
		{
			MethodName: "GetDmsSearchDatabase",
			Handler:    _DatabaseService_GetDmsSearchDatabase_Handler,
		},
		{
			MethodName: "GetDbsBackupPlan",
			Handler:    _DatabaseService_GetDbsBackupPlan_Handler,
		},
		{
			MethodName: "GetDtsSyncJobName",
			Handler:    _DatabaseService_GetDtsSyncJobName_Handler,
		},
		{
			MethodName: "GetRdsSlowLog",
			Handler:    _DatabaseService_GetRdsSlowLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "database-service.proto",
}
