// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: pb/task.proto

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

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskClient interface {
	// 普通一元方法
	UnaryTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error)
	// 服务端推送流
	ServerStreamingTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (Task_ServerStreamingTaskClient, error)
	// 客户端推送流
	ClientStreamingTask(ctx context.Context, opts ...grpc.CallOption) (Task_ClientStreamingTaskClient, error)
	// 双向推送流
	BidirectionalStreamingTask(ctx context.Context, opts ...grpc.CallOption) (Task_BidirectionalStreamingTaskClient, error)
}

type taskClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskClient(cc grpc.ClientConnInterface) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) UnaryTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (*TaskResponse, error) {
	out := new(TaskResponse)
	err := c.cc.Invoke(ctx, "/pb.Task/UnaryTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskClient) ServerStreamingTask(ctx context.Context, in *TaskRequest, opts ...grpc.CallOption) (Task_ServerStreamingTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &Task_ServiceDesc.Streams[0], "/pb.Task/ServerStreamingTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskServerStreamingTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Task_ServerStreamingTaskClient interface {
	Recv() (*TaskResponse, error)
	grpc.ClientStream
}

type taskServerStreamingTaskClient struct {
	grpc.ClientStream
}

func (x *taskServerStreamingTaskClient) Recv() (*TaskResponse, error) {
	m := new(TaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskClient) ClientStreamingTask(ctx context.Context, opts ...grpc.CallOption) (Task_ClientStreamingTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &Task_ServiceDesc.Streams[1], "/pb.Task/ClientStreamingTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskClientStreamingTaskClient{stream}
	return x, nil
}

type Task_ClientStreamingTaskClient interface {
	Send(*TaskRequest) error
	CloseAndRecv() (*TaskResponse, error)
	grpc.ClientStream
}

type taskClientStreamingTaskClient struct {
	grpc.ClientStream
}

func (x *taskClientStreamingTaskClient) Send(m *TaskRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskClientStreamingTaskClient) CloseAndRecv() (*TaskResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskClient) BidirectionalStreamingTask(ctx context.Context, opts ...grpc.CallOption) (Task_BidirectionalStreamingTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &Task_ServiceDesc.Streams[2], "/pb.Task/BidirectionalStreamingTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskBidirectionalStreamingTaskClient{stream}
	return x, nil
}

type Task_BidirectionalStreamingTaskClient interface {
	Send(*TaskRequest) error
	Recv() (*TaskResponse, error)
	grpc.ClientStream
}

type taskBidirectionalStreamingTaskClient struct {
	grpc.ClientStream
}

func (x *taskBidirectionalStreamingTaskClient) Send(m *TaskRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *taskBidirectionalStreamingTaskClient) Recv() (*TaskResponse, error) {
	m := new(TaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TaskServer is the server API for Task service.
// All implementations must embed UnimplementedTaskServer
// for forward compatibility
type TaskServer interface {
	// 普通一元方法
	UnaryTask(context.Context, *TaskRequest) (*TaskResponse, error)
	// 服务端推送流
	ServerStreamingTask(*TaskRequest, Task_ServerStreamingTaskServer) error
	// 客户端推送流
	ClientStreamingTask(Task_ClientStreamingTaskServer) error
	// 双向推送流
	BidirectionalStreamingTask(Task_BidirectionalStreamingTaskServer) error
	mustEmbedUnimplementedTaskServer()
}

// UnimplementedTaskServer must be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (UnimplementedTaskServer) UnaryTask(context.Context, *TaskRequest) (*TaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryTask not implemented")
}
func (UnimplementedTaskServer) ServerStreamingTask(*TaskRequest, Task_ServerStreamingTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingTask not implemented")
}
func (UnimplementedTaskServer) ClientStreamingTask(Task_ClientStreamingTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamingTask not implemented")
}
func (UnimplementedTaskServer) BidirectionalStreamingTask(Task_BidirectionalStreamingTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamingTask not implemented")
}
func (UnimplementedTaskServer) mustEmbedUnimplementedTaskServer() {}

// UnsafeTaskServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskServer will
// result in compilation errors.
type UnsafeTaskServer interface {
	mustEmbedUnimplementedTaskServer()
}

func RegisterTaskServer(s grpc.ServiceRegistrar, srv TaskServer) {
	s.RegisterService(&Task_ServiceDesc, srv)
}

func _Task_UnaryTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskServer).UnaryTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Task/UnaryTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskServer).UnaryTask(ctx, req.(*TaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Task_ServerStreamingTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskServer).ServerStreamingTask(m, &taskServerStreamingTaskServer{stream})
}

type Task_ServerStreamingTaskServer interface {
	Send(*TaskResponse) error
	grpc.ServerStream
}

type taskServerStreamingTaskServer struct {
	grpc.ServerStream
}

func (x *taskServerStreamingTaskServer) Send(m *TaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Task_ClientStreamingTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServer).ClientStreamingTask(&taskClientStreamingTaskServer{stream})
}

type Task_ClientStreamingTaskServer interface {
	SendAndClose(*TaskResponse) error
	Recv() (*TaskRequest, error)
	grpc.ServerStream
}

type taskClientStreamingTaskServer struct {
	grpc.ServerStream
}

func (x *taskClientStreamingTaskServer) SendAndClose(m *TaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskClientStreamingTaskServer) Recv() (*TaskRequest, error) {
	m := new(TaskRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Task_BidirectionalStreamingTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TaskServer).BidirectionalStreamingTask(&taskBidirectionalStreamingTaskServer{stream})
}

type Task_BidirectionalStreamingTaskServer interface {
	Send(*TaskResponse) error
	Recv() (*TaskRequest, error)
	grpc.ServerStream
}

type taskBidirectionalStreamingTaskServer struct {
	grpc.ServerStream
}

func (x *taskBidirectionalStreamingTaskServer) Send(m *TaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *taskBidirectionalStreamingTaskServer) Recv() (*TaskRequest, error) {
	m := new(TaskRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Task_ServiceDesc is the grpc.ServiceDesc for Task service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Task_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Task",
	HandlerType: (*TaskServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryTask",
			Handler:    _Task_UnaryTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamingTask",
			Handler:       _Task_ServerStreamingTask_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamingTask",
			Handler:       _Task_ClientStreamingTask_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamingTask",
			Handler:       _Task_BidirectionalStreamingTask_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pb/task.proto",
}
