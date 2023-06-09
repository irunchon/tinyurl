// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/tinyurl.proto

package tinyurl

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
	ShortenURL_GetShortURL_FullMethodName = "/tinyurl.ShortenURL/GetShortURL"
	ShortenURL_GetLongURL_FullMethodName  = "/tinyurl.ShortenURL/GetLongURL"
)

// ShortenURLClient is the client API for ShortenURL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShortenURLClient interface {
	GetShortURL(ctx context.Context, in *GetShortURLRequest, opts ...grpc.CallOption) (*GetShortURLResponse, error)
	GetLongURL(ctx context.Context, in *GetLongURLRequest, opts ...grpc.CallOption) (*GetLongURLResponse, error)
}

type shortenURLClient struct {
	cc grpc.ClientConnInterface
}

func NewShortenURLClient(cc grpc.ClientConnInterface) ShortenURLClient {
	return &shortenURLClient{cc}
}

func (c *shortenURLClient) GetShortURL(ctx context.Context, in *GetShortURLRequest, opts ...grpc.CallOption) (*GetShortURLResponse, error) {
	out := new(GetShortURLResponse)
	err := c.cc.Invoke(ctx, ShortenURL_GetShortURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortenURLClient) GetLongURL(ctx context.Context, in *GetLongURLRequest, opts ...grpc.CallOption) (*GetLongURLResponse, error) {
	out := new(GetLongURLResponse)
	err := c.cc.Invoke(ctx, ShortenURL_GetLongURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortenURLServer is the server API for ShortenURL service.
// All implementations must embed UnimplementedShortenURLServer
// for forward compatibility
type ShortenURLServer interface {
	GetShortURL(context.Context, *GetShortURLRequest) (*GetShortURLResponse, error)
	GetLongURL(context.Context, *GetLongURLRequest) (*GetLongURLResponse, error)
	mustEmbedUnimplementedShortenURLServer()
}

// UnimplementedShortenURLServer must be embedded to have forward compatible implementations.
type UnimplementedShortenURLServer struct {
}

func (UnimplementedShortenURLServer) GetShortURL(context.Context, *GetShortURLRequest) (*GetShortURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetShortURL not implemented")
}
func (UnimplementedShortenURLServer) GetLongURL(context.Context, *GetLongURLRequest) (*GetLongURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLongURL not implemented")
}
func (UnimplementedShortenURLServer) mustEmbedUnimplementedShortenURLServer() {}

// UnsafeShortenURLServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShortenURLServer will
// result in compilation errors.
type UnsafeShortenURLServer interface {
	mustEmbedUnimplementedShortenURLServer()
}

func RegisterShortenURLServer(s grpc.ServiceRegistrar, srv ShortenURLServer) {
	s.RegisterService(&ShortenURL_ServiceDesc, srv)
}

func _ShortenURL_GetShortURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetShortURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenURLServer).GetShortURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenURL_GetShortURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenURLServer).GetShortURL(ctx, req.(*GetShortURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ShortenURL_GetLongURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLongURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortenURLServer).GetLongURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShortenURL_GetLongURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortenURLServer).GetLongURL(ctx, req.(*GetLongURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShortenURL_ServiceDesc is the grpc.ServiceDesc for ShortenURL service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShortenURL_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tinyurl.ShortenURL",
	HandlerType: (*ShortenURLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetShortURL",
			Handler:    _ShortenURL_GetShortURL_Handler,
		},
		{
			MethodName: "GetLongURL",
			Handler:    _ShortenURL_GetLongURL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/tinyurl.proto",
}
