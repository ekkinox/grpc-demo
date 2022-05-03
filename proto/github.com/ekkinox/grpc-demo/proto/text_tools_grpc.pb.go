// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: text_tools.proto

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

// TextToolsServiceClient is the client API for TextToolsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextToolsServiceClient interface {
	// Unary rpc
	TransformText(ctx context.Context, in *TransformTextRequest, opts ...grpc.CallOption) (*TransformTextResponse, error)
	// BiDi rpc
	TransformAndSplitText(ctx context.Context, opts ...grpc.CallOption) (TextToolsService_TransformAndSplitTextClient, error)
}

type textToolsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextToolsServiceClient(cc grpc.ClientConnInterface) TextToolsServiceClient {
	return &textToolsServiceClient{cc}
}

func (c *textToolsServiceClient) TransformText(ctx context.Context, in *TransformTextRequest, opts ...grpc.CallOption) (*TransformTextResponse, error) {
	out := new(TransformTextResponse)
	err := c.cc.Invoke(ctx, "/text_tools.TextToolsService/TransformText", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textToolsServiceClient) TransformAndSplitText(ctx context.Context, opts ...grpc.CallOption) (TextToolsService_TransformAndSplitTextClient, error) {
	stream, err := c.cc.NewStream(ctx, &TextToolsService_ServiceDesc.Streams[0], "/text_tools.TextToolsService/TransformAndSplitText", opts...)
	if err != nil {
		return nil, err
	}
	x := &textToolsServiceTransformAndSplitTextClient{stream}
	return x, nil
}

type TextToolsService_TransformAndSplitTextClient interface {
	Send(*TransformTextRequest) error
	Recv() (*TransformTextResponse, error)
	grpc.ClientStream
}

type textToolsServiceTransformAndSplitTextClient struct {
	grpc.ClientStream
}

func (x *textToolsServiceTransformAndSplitTextClient) Send(m *TransformTextRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *textToolsServiceTransformAndSplitTextClient) Recv() (*TransformTextResponse, error) {
	m := new(TransformTextResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TextToolsServiceServer is the server API for TextToolsService service.
// All implementations must embed UnimplementedTextToolsServiceServer
// for forward compatibility
type TextToolsServiceServer interface {
	// Unary rpc
	TransformText(context.Context, *TransformTextRequest) (*TransformTextResponse, error)
	// BiDi rpc
	TransformAndSplitText(TextToolsService_TransformAndSplitTextServer) error
	mustEmbedUnimplementedTextToolsServiceServer()
}

// UnimplementedTextToolsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTextToolsServiceServer struct {
}

func (UnimplementedTextToolsServiceServer) TransformText(context.Context, *TransformTextRequest) (*TransformTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransformText not implemented")
}
func (UnimplementedTextToolsServiceServer) TransformAndSplitText(TextToolsService_TransformAndSplitTextServer) error {
	return status.Errorf(codes.Unimplemented, "method TransformAndSplitText not implemented")
}
func (UnimplementedTextToolsServiceServer) mustEmbedUnimplementedTextToolsServiceServer() {}

// UnsafeTextToolsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextToolsServiceServer will
// result in compilation errors.
type UnsafeTextToolsServiceServer interface {
	mustEmbedUnimplementedTextToolsServiceServer()
}

func RegisterTextToolsServiceServer(s grpc.ServiceRegistrar, srv TextToolsServiceServer) {
	s.RegisterService(&TextToolsService_ServiceDesc, srv)
}

func _TextToolsService_TransformText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransformTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextToolsServiceServer).TransformText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/text_tools.TextToolsService/TransformText",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextToolsServiceServer).TransformText(ctx, req.(*TransformTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextToolsService_TransformAndSplitText_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TextToolsServiceServer).TransformAndSplitText(&textToolsServiceTransformAndSplitTextServer{stream})
}

type TextToolsService_TransformAndSplitTextServer interface {
	Send(*TransformTextResponse) error
	Recv() (*TransformTextRequest, error)
	grpc.ServerStream
}

type textToolsServiceTransformAndSplitTextServer struct {
	grpc.ServerStream
}

func (x *textToolsServiceTransformAndSplitTextServer) Send(m *TransformTextResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *textToolsServiceTransformAndSplitTextServer) Recv() (*TransformTextRequest, error) {
	m := new(TransformTextRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TextToolsService_ServiceDesc is the grpc.ServiceDesc for TextToolsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextToolsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "text_tools.TextToolsService",
	HandlerType: (*TextToolsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransformText",
			Handler:    _TextToolsService_TransformText_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TransformAndSplitText",
			Handler:       _TextToolsService_TransformAndSplitText_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "text_tools.proto",
}
