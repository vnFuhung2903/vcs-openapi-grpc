// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/book.proto

// protoc --go_out=. --go-grpc_out=. proto/book.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Book_GetBook_FullMethodName        = "/proto.Book/GetBook"
	Book_ListBooks_FullMethodName      = "/proto.Book/ListBooks"
	Book_MultiGetBook_FullMethodName   = "/proto.Book/MultiGetBook"
	Book_MultiListBooks_FullMethodName = "/proto.Book/MultiListBooks"
)

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookClient interface {
	GetBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error)
	ListBooks(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[BookResponse], error)
	MultiGetBook(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[BookRequest, BookResponse], error)
	MultiListBooks(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[BookRequest, BookResponse], error)
}

type bookClient struct {
	cc grpc.ClientConnInterface
}

func NewBookClient(cc grpc.ClientConnInterface) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) GetBook(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (*BookResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BookResponse)
	err := c.cc.Invoke(ctx, Book_GetBook_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) ListBooks(ctx context.Context, in *BookRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[BookResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Book_ServiceDesc.Streams[0], Book_ListBooks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[BookRequest, BookResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Book_ListBooksClient = grpc.ServerStreamingClient[BookResponse]

func (c *bookClient) MultiGetBook(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[BookRequest, BookResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Book_ServiceDesc.Streams[1], Book_MultiGetBook_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[BookRequest, BookResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Book_MultiGetBookClient = grpc.ClientStreamingClient[BookRequest, BookResponse]

func (c *bookClient) MultiListBooks(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[BookRequest, BookResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Book_ServiceDesc.Streams[2], Book_MultiListBooks_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[BookRequest, BookResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Book_MultiListBooksClient = grpc.BidiStreamingClient[BookRequest, BookResponse]

// BookServer is the server API for Book service.
// All implementations must embed UnimplementedBookServer
// for forward compatibility.
type BookServer interface {
	GetBook(context.Context, *BookRequest) (*BookResponse, error)
	ListBooks(*BookRequest, grpc.ServerStreamingServer[BookResponse]) error
	MultiGetBook(grpc.ClientStreamingServer[BookRequest, BookResponse]) error
	MultiListBooks(grpc.BidiStreamingServer[BookRequest, BookResponse]) error
	mustEmbedUnimplementedBookServer()
}

// UnimplementedBookServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBookServer struct{}

func (UnimplementedBookServer) GetBook(context.Context, *BookRequest) (*BookResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBook not implemented")
}
func (UnimplementedBookServer) ListBooks(*BookRequest, grpc.ServerStreamingServer[BookResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ListBooks not implemented")
}
func (UnimplementedBookServer) MultiGetBook(grpc.ClientStreamingServer[BookRequest, BookResponse]) error {
	return status.Errorf(codes.Unimplemented, "method MultiGetBook not implemented")
}
func (UnimplementedBookServer) MultiListBooks(grpc.BidiStreamingServer[BookRequest, BookResponse]) error {
	return status.Errorf(codes.Unimplemented, "method MultiListBooks not implemented")
}
func (UnimplementedBookServer) mustEmbedUnimplementedBookServer() {}
func (UnimplementedBookServer) testEmbeddedByValue()              {}

// UnsafeBookServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookServer will
// result in compilation errors.
type UnsafeBookServer interface {
	mustEmbedUnimplementedBookServer()
}

func RegisterBookServer(s grpc.ServiceRegistrar, srv BookServer) {
	// If the following call pancis, it indicates UnimplementedBookServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Book_ServiceDesc, srv)
}

func _Book_GetBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Book_GetBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBook(ctx, req.(*BookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_ListBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BookRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BookServer).ListBooks(m, &grpc.GenericServerStream[BookRequest, BookResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Book_ListBooksServer = grpc.ServerStreamingServer[BookResponse]

func _Book_MultiGetBook_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BookServer).MultiGetBook(&grpc.GenericServerStream[BookRequest, BookResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Book_MultiGetBookServer = grpc.ClientStreamingServer[BookRequest, BookResponse]

func _Book_MultiListBooks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BookServer).MultiListBooks(&grpc.GenericServerStream[BookRequest, BookResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Book_MultiListBooksServer = grpc.BidiStreamingServer[BookRequest, BookResponse]

// Book_ServiceDesc is the grpc.ServiceDesc for Book service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Book_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBook",
			Handler:    _Book_GetBook_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBooks",
			Handler:       _Book_ListBooks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "MultiGetBook",
			Handler:       _Book_MultiGetBook_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "MultiListBooks",
			Handler:       _Book_MultiListBooks_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/book.proto",
}
