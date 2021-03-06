// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// FullNameBuilderClient is the client API for FullNameBuilder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FullNameBuilderClient interface {
	// Recebe um nome e sobrenome e retorna um nome completo
	GetFullName(ctx context.Context, in *FullNameRequest, opts ...grpc.CallOption) (*FullName, error)
	// Recebe um nome completo e retorna um stream com varios nomes
	GetNames(ctx context.Context, in *FullName, opts ...grpc.CallOption) (FullNameBuilder_GetNamesClient, error)
	// Recebe varios nomes e retorna um nome completo
	GetFullNameWithNames(ctx context.Context, opts ...grpc.CallOption) (FullNameBuilder_GetFullNameWithNamesClient, error)
	// Recebe varios nomes e retorna varios nomes
	GetNamesStream(ctx context.Context, opts ...grpc.CallOption) (FullNameBuilder_GetNamesStreamClient, error)
}

type fullNameBuilderClient struct {
	cc grpc.ClientConnInterface
}

func NewFullNameBuilderClient(cc grpc.ClientConnInterface) FullNameBuilderClient {
	return &fullNameBuilderClient{cc}
}

func (c *fullNameBuilderClient) GetFullName(ctx context.Context, in *FullNameRequest, opts ...grpc.CallOption) (*FullName, error) {
	out := new(FullName)
	err := c.cc.Invoke(ctx, "/helloworld.FullNameBuilder/GetFullName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fullNameBuilderClient) GetNames(ctx context.Context, in *FullName, opts ...grpc.CallOption) (FullNameBuilder_GetNamesClient, error) {
	stream, err := c.cc.NewStream(ctx, &FullNameBuilder_ServiceDesc.Streams[0], "/helloworld.FullNameBuilder/GetNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &fullNameBuilderGetNamesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FullNameBuilder_GetNamesClient interface {
	Recv() (*Name, error)
	grpc.ClientStream
}

type fullNameBuilderGetNamesClient struct {
	grpc.ClientStream
}

func (x *fullNameBuilderGetNamesClient) Recv() (*Name, error) {
	m := new(Name)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fullNameBuilderClient) GetFullNameWithNames(ctx context.Context, opts ...grpc.CallOption) (FullNameBuilder_GetFullNameWithNamesClient, error) {
	stream, err := c.cc.NewStream(ctx, &FullNameBuilder_ServiceDesc.Streams[1], "/helloworld.FullNameBuilder/GetFullNameWithNames", opts...)
	if err != nil {
		return nil, err
	}
	x := &fullNameBuilderGetFullNameWithNamesClient{stream}
	return x, nil
}

type FullNameBuilder_GetFullNameWithNamesClient interface {
	Send(*Name) error
	CloseAndRecv() (*FullName, error)
	grpc.ClientStream
}

type fullNameBuilderGetFullNameWithNamesClient struct {
	grpc.ClientStream
}

func (x *fullNameBuilderGetFullNameWithNamesClient) Send(m *Name) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fullNameBuilderGetFullNameWithNamesClient) CloseAndRecv() (*FullName, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(FullName)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fullNameBuilderClient) GetNamesStream(ctx context.Context, opts ...grpc.CallOption) (FullNameBuilder_GetNamesStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &FullNameBuilder_ServiceDesc.Streams[2], "/helloworld.FullNameBuilder/GetNamesStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &fullNameBuilderGetNamesStreamClient{stream}
	return x, nil
}

type FullNameBuilder_GetNamesStreamClient interface {
	Send(*Name) error
	Recv() (*Name, error)
	grpc.ClientStream
}

type fullNameBuilderGetNamesStreamClient struct {
	grpc.ClientStream
}

func (x *fullNameBuilderGetNamesStreamClient) Send(m *Name) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fullNameBuilderGetNamesStreamClient) Recv() (*Name, error) {
	m := new(Name)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FullNameBuilderServer is the server API for FullNameBuilder service.
// All implementations must embed UnimplementedFullNameBuilderServer
// for forward compatibility
type FullNameBuilderServer interface {
	// Recebe um nome e sobrenome e retorna um nome completo
	GetFullName(context.Context, *FullNameRequest) (*FullName, error)
	// Recebe um nome completo e retorna um stream com varios nomes
	GetNames(*FullName, FullNameBuilder_GetNamesServer) error
	// Recebe varios nomes e retorna um nome completo
	GetFullNameWithNames(FullNameBuilder_GetFullNameWithNamesServer) error
	// Recebe varios nomes e retorna varios nomes
	GetNamesStream(FullNameBuilder_GetNamesStreamServer) error
	mustEmbedUnimplementedFullNameBuilderServer()
}

// UnimplementedFullNameBuilderServer must be embedded to have forward compatible implementations.
type UnimplementedFullNameBuilderServer struct {
}

func (UnimplementedFullNameBuilderServer) GetFullName(context.Context, *FullNameRequest) (*FullName, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFullName not implemented")
}
func (UnimplementedFullNameBuilderServer) GetNames(*FullName, FullNameBuilder_GetNamesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNames not implemented")
}
func (UnimplementedFullNameBuilderServer) GetFullNameWithNames(FullNameBuilder_GetFullNameWithNamesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFullNameWithNames not implemented")
}
func (UnimplementedFullNameBuilderServer) GetNamesStream(FullNameBuilder_GetNamesStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNamesStream not implemented")
}
func (UnimplementedFullNameBuilderServer) mustEmbedUnimplementedFullNameBuilderServer() {}

// UnsafeFullNameBuilderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FullNameBuilderServer will
// result in compilation errors.
type UnsafeFullNameBuilderServer interface {
	mustEmbedUnimplementedFullNameBuilderServer()
}

func RegisterFullNameBuilderServer(s grpc.ServiceRegistrar, srv FullNameBuilderServer) {
	s.RegisterService(&FullNameBuilder_ServiceDesc, srv)
}

func _FullNameBuilder_GetFullName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FullNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FullNameBuilderServer).GetFullName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.FullNameBuilder/GetFullName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FullNameBuilderServer).GetFullName(ctx, req.(*FullNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FullNameBuilder_GetNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FullName)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FullNameBuilderServer).GetNames(m, &fullNameBuilderGetNamesServer{stream})
}

type FullNameBuilder_GetNamesServer interface {
	Send(*Name) error
	grpc.ServerStream
}

type fullNameBuilderGetNamesServer struct {
	grpc.ServerStream
}

func (x *fullNameBuilderGetNamesServer) Send(m *Name) error {
	return x.ServerStream.SendMsg(m)
}

func _FullNameBuilder_GetFullNameWithNames_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FullNameBuilderServer).GetFullNameWithNames(&fullNameBuilderGetFullNameWithNamesServer{stream})
}

type FullNameBuilder_GetFullNameWithNamesServer interface {
	SendAndClose(*FullName) error
	Recv() (*Name, error)
	grpc.ServerStream
}

type fullNameBuilderGetFullNameWithNamesServer struct {
	grpc.ServerStream
}

func (x *fullNameBuilderGetFullNameWithNamesServer) SendAndClose(m *FullName) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fullNameBuilderGetFullNameWithNamesServer) Recv() (*Name, error) {
	m := new(Name)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FullNameBuilder_GetNamesStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FullNameBuilderServer).GetNamesStream(&fullNameBuilderGetNamesStreamServer{stream})
}

type FullNameBuilder_GetNamesStreamServer interface {
	Send(*Name) error
	Recv() (*Name, error)
	grpc.ServerStream
}

type fullNameBuilderGetNamesStreamServer struct {
	grpc.ServerStream
}

func (x *fullNameBuilderGetNamesStreamServer) Send(m *Name) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fullNameBuilderGetNamesStreamServer) Recv() (*Name, error) {
	m := new(Name)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FullNameBuilder_ServiceDesc is the grpc.ServiceDesc for FullNameBuilder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FullNameBuilder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.FullNameBuilder",
	HandlerType: (*FullNameBuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFullName",
			Handler:    _FullNameBuilder_GetFullName_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNames",
			Handler:       _FullNameBuilder_GetNames_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetFullNameWithNames",
			Handler:       _FullNameBuilder_GetFullNameWithNames_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetNamesStream",
			Handler:       _FullNameBuilder_GetNamesStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "app.proto",
}
