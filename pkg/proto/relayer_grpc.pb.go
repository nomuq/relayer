// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: relayer.proto

package proto

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

// RelayerClient is the client API for Relayer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelayerClient interface {
	GetAuthToken(ctx context.Context, in *GetAuthTokenRequest, opts ...grpc.CallOption) (*GetAuthTokenResponse, error)
	Authenticate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*User, error)
}

type relayerClient struct {
	cc grpc.ClientConnInterface
}

func NewRelayerClient(cc grpc.ClientConnInterface) RelayerClient {
	return &relayerClient{cc}
}

func (c *relayerClient) GetAuthToken(ctx context.Context, in *GetAuthTokenRequest, opts ...grpc.CallOption) (*GetAuthTokenResponse, error) {
	out := new(GetAuthTokenResponse)
	err := c.cc.Invoke(ctx, "/relayer.Relayer/GetAuthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relayerClient) Authenticate(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/relayer.Relayer/Authenticate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelayerServer is the server API for Relayer service.
// All implementations must embed UnimplementedRelayerServer
// for forward compatibility
type RelayerServer interface {
	GetAuthToken(context.Context, *GetAuthTokenRequest) (*GetAuthTokenResponse, error)
	Authenticate(context.Context, *Empty) (*User, error)
	mustEmbedUnimplementedRelayerServer()
}

// UnimplementedRelayerServer must be embedded to have forward compatible implementations.
type UnimplementedRelayerServer struct {
}

func (UnimplementedRelayerServer) GetAuthToken(context.Context, *GetAuthTokenRequest) (*GetAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthToken not implemented")
}
func (UnimplementedRelayerServer) Authenticate(context.Context, *Empty) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedRelayerServer) mustEmbedUnimplementedRelayerServer() {}

// UnsafeRelayerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelayerServer will
// result in compilation errors.
type UnsafeRelayerServer interface {
	mustEmbedUnimplementedRelayerServer()
}

func RegisterRelayerServer(s grpc.ServiceRegistrar, srv RelayerServer) {
	s.RegisterService(&Relayer_ServiceDesc, srv)
}

func _Relayer_GetAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayerServer).GetAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relayer.Relayer/GetAuthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayerServer).GetAuthToken(ctx, req.(*GetAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Relayer_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelayerServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/relayer.Relayer/Authenticate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelayerServer).Authenticate(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Relayer_ServiceDesc is the grpc.ServiceDesc for Relayer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Relayer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "relayer.Relayer",
	HandlerType: (*RelayerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthToken",
			Handler:    _Relayer_GetAuthToken_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _Relayer_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "relayer.proto",
}
