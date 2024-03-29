// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: authorizer.proto

package grpc

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

// AuthorizerClient is the client API for Authorizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizerClient interface {
	AuthorizeByToken(ctx context.Context, in *AuthorizeToken, opts ...grpc.CallOption) (*AuthorizeResponse, error)
}

type authorizerClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizerClient(cc grpc.ClientConnInterface) AuthorizerClient {
	return &authorizerClient{cc}
}

func (c *authorizerClient) AuthorizeByToken(ctx context.Context, in *AuthorizeToken, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/authorizer.Authorizer/AuthorizeByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizerServer is the server API for Authorizer service.
// All implementations must embed UnimplementedAuthorizerServer
// for forward compatibility
type AuthorizerServer interface {
	AuthorizeByToken(context.Context, *AuthorizeToken) (*AuthorizeResponse, error)
	mustEmbedUnimplementedAuthorizerServer()
}

// UnimplementedAuthorizerServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizerServer struct {
}

func (UnimplementedAuthorizerServer) AuthorizeByToken(context.Context, *AuthorizeToken) (*AuthorizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeByToken not implemented")
}
func (UnimplementedAuthorizerServer) mustEmbedUnimplementedAuthorizerServer() {}

// UnsafeAuthorizerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizerServer will
// result in compilation errors.
type UnsafeAuthorizerServer interface {
	mustEmbedUnimplementedAuthorizerServer()
}

func RegisterAuthorizerServer(s grpc.ServiceRegistrar, srv AuthorizerServer) {
	s.RegisterService(&Authorizer_ServiceDesc, srv)
}

func _Authorizer_AuthorizeByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizerServer).AuthorizeByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorizer.Authorizer/AuthorizeByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizerServer).AuthorizeByToken(ctx, req.(*AuthorizeToken))
	}
	return interceptor(ctx, in, info, handler)
}

// Authorizer_ServiceDesc is the grpc.ServiceDesc for Authorizer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authorizer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorizer.Authorizer",
	HandlerType: (*AuthorizerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizeByToken",
			Handler:    _Authorizer_AuthorizeByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorizer.proto",
}
