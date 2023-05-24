// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: proto/history.proto

package historypb

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

const (
	HistoryService_CreateHistory_FullMethodName = "/history.HistoryService/CreateHistory"
	HistoryService_ReadHistory_FullMethodName   = "/history.HistoryService/ReadHistory"
	HistoryService_UpdateHistory_FullMethodName = "/history.HistoryService/UpdateHistory"
	HistoryService_DeleteHistory_FullMethodName = "/history.HistoryService/DeleteHistory"
	HistoryService_ListHistorys_FullMethodName  = "/history.HistoryService/ListHistorys"
)

// HistoryServiceClient is the client API for HistoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HistoryServiceClient interface {
	CreateHistory(ctx context.Context, in *CreateHistoryReq, opts ...grpc.CallOption) (*CreateHistoryRes, error)
	ReadHistory(ctx context.Context, in *ReadHistoryReq, opts ...grpc.CallOption) (*ReadHistoryRes, error)
	UpdateHistory(ctx context.Context, in *UpdateHistoryReq, opts ...grpc.CallOption) (*UpdateHistoryRes, error)
	DeleteHistory(ctx context.Context, in *DeleteHistoryReq, opts ...grpc.CallOption) (*DeleteHistoryRes, error)
	ListHistorys(ctx context.Context, in *ListHistoriesReq, opts ...grpc.CallOption) (HistoryService_ListHistorysClient, error)
}

type historyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHistoryServiceClient(cc grpc.ClientConnInterface) HistoryServiceClient {
	return &historyServiceClient{cc}
}

func (c *historyServiceClient) CreateHistory(ctx context.Context, in *CreateHistoryReq, opts ...grpc.CallOption) (*CreateHistoryRes, error) {
	out := new(CreateHistoryRes)
	err := c.cc.Invoke(ctx, HistoryService_CreateHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyServiceClient) ReadHistory(ctx context.Context, in *ReadHistoryReq, opts ...grpc.CallOption) (*ReadHistoryRes, error) {
	out := new(ReadHistoryRes)
	err := c.cc.Invoke(ctx, HistoryService_ReadHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyServiceClient) UpdateHistory(ctx context.Context, in *UpdateHistoryReq, opts ...grpc.CallOption) (*UpdateHistoryRes, error) {
	out := new(UpdateHistoryRes)
	err := c.cc.Invoke(ctx, HistoryService_UpdateHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyServiceClient) DeleteHistory(ctx context.Context, in *DeleteHistoryReq, opts ...grpc.CallOption) (*DeleteHistoryRes, error) {
	out := new(DeleteHistoryRes)
	err := c.cc.Invoke(ctx, HistoryService_DeleteHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *historyServiceClient) ListHistorys(ctx context.Context, in *ListHistoriesReq, opts ...grpc.CallOption) (HistoryService_ListHistorysClient, error) {
	stream, err := c.cc.NewStream(ctx, &HistoryService_ServiceDesc.Streams[0], HistoryService_ListHistorys_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &historyServiceListHistorysClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HistoryService_ListHistorysClient interface {
	Recv() (*ListHistoriesRes, error)
	grpc.ClientStream
}

type historyServiceListHistorysClient struct {
	grpc.ClientStream
}

func (x *historyServiceListHistorysClient) Recv() (*ListHistoriesRes, error) {
	m := new(ListHistoriesRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HistoryServiceServer is the server API for HistoryService service.
// All implementations must embed UnimplementedHistoryServiceServer
// for forward compatibility
type HistoryServiceServer interface {
	CreateHistory(context.Context, *CreateHistoryReq) (*CreateHistoryRes, error)
	ReadHistory(context.Context, *ReadHistoryReq) (*ReadHistoryRes, error)
	UpdateHistory(context.Context, *UpdateHistoryReq) (*UpdateHistoryRes, error)
	DeleteHistory(context.Context, *DeleteHistoryReq) (*DeleteHistoryRes, error)
	ListHistorys(*ListHistoriesReq, HistoryService_ListHistorysServer) error
	mustEmbedUnimplementedHistoryServiceServer()
}

// UnimplementedHistoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHistoryServiceServer struct {
}

func (UnimplementedHistoryServiceServer) CreateHistory(context.Context, *CreateHistoryReq) (*CreateHistoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateHistory not implemented")
}
func (UnimplementedHistoryServiceServer) ReadHistory(context.Context, *ReadHistoryReq) (*ReadHistoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadHistory not implemented")
}
func (UnimplementedHistoryServiceServer) UpdateHistory(context.Context, *UpdateHistoryReq) (*UpdateHistoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateHistory not implemented")
}
func (UnimplementedHistoryServiceServer) DeleteHistory(context.Context, *DeleteHistoryReq) (*DeleteHistoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteHistory not implemented")
}
func (UnimplementedHistoryServiceServer) ListHistorys(*ListHistoriesReq, HistoryService_ListHistorysServer) error {
	return status.Errorf(codes.Unimplemented, "method ListHistorys not implemented")
}
func (UnimplementedHistoryServiceServer) mustEmbedUnimplementedHistoryServiceServer() {}

// UnsafeHistoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HistoryServiceServer will
// result in compilation errors.
type UnsafeHistoryServiceServer interface {
	mustEmbedUnimplementedHistoryServiceServer()
}

func RegisterHistoryServiceServer(s grpc.ServiceRegistrar, srv HistoryServiceServer) {
	s.RegisterService(&HistoryService_ServiceDesc, srv)
}

func _HistoryService_CreateHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServiceServer).CreateHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HistoryService_CreateHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServiceServer).CreateHistory(ctx, req.(*CreateHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoryService_ReadHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServiceServer).ReadHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HistoryService_ReadHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServiceServer).ReadHistory(ctx, req.(*ReadHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoryService_UpdateHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServiceServer).UpdateHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HistoryService_UpdateHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServiceServer).UpdateHistory(ctx, req.(*UpdateHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoryService_DeleteHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteHistoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HistoryServiceServer).DeleteHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HistoryService_DeleteHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HistoryServiceServer).DeleteHistory(ctx, req.(*DeleteHistoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _HistoryService_ListHistorys_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListHistoriesReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HistoryServiceServer).ListHistorys(m, &historyServiceListHistorysServer{stream})
}

type HistoryService_ListHistorysServer interface {
	Send(*ListHistoriesRes) error
	grpc.ServerStream
}

type historyServiceListHistorysServer struct {
	grpc.ServerStream
}

func (x *historyServiceListHistorysServer) Send(m *ListHistoriesRes) error {
	return x.ServerStream.SendMsg(m)
}

// HistoryService_ServiceDesc is the grpc.ServiceDesc for HistoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HistoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "history.HistoryService",
	HandlerType: (*HistoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateHistory",
			Handler:    _HistoryService_CreateHistory_Handler,
		},
		{
			MethodName: "ReadHistory",
			Handler:    _HistoryService_ReadHistory_Handler,
		},
		{
			MethodName: "UpdateHistory",
			Handler:    _HistoryService_UpdateHistory_Handler,
		},
		{
			MethodName: "DeleteHistory",
			Handler:    _HistoryService_DeleteHistory_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListHistorys",
			Handler:       _HistoryService_ListHistorys_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/history.proto",
}
