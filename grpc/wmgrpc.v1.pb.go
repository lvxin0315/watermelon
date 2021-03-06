// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/wmgrpc.v1.proto

package wmgrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f29f673d8a61db, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Result struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7f29f673d8a61db, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "wmgrpc.empty")
	proto.RegisterType((*Result)(nil), "wmgrpc.result")
}

func init() { proto.RegisterFile("grpc/wmgrpc.v1.proto", fileDescriptor_e7f29f673d8a61db) }

var fileDescriptor_e7f29f673d8a61db = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x2f, 0x2a, 0x48,
	0xd6, 0x2f, 0xcf, 0x05, 0x51, 0x7a, 0x65, 0x86, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c,
	0x10, 0x01, 0x25, 0x76, 0x2e, 0xd6, 0xd4, 0xdc, 0x82, 0x92, 0x4a, 0x25, 0x19, 0x2e, 0xb6, 0xa2,
	0xd4, 0xe2, 0xd2, 0x9c, 0x12, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09, 0x46, 0x05,
	0x46, 0x0d, 0x9e, 0x20, 0x30, 0xdb, 0xc8, 0x82, 0x8b, 0xcb, 0xa3, 0xa4, 0xa4, 0x20, 0x38, 0xb5,
	0xa8, 0x2c, 0xb5, 0x48, 0x48, 0x8b, 0x8b, 0xdd, 0x3d, 0xb5, 0xc4, 0x25, 0xb1, 0x24, 0x51, 0x88,
	0x57, 0x0f, 0x6a, 0x32, 0xd8, 0x14, 0x29, 0x3e, 0x18, 0x17, 0x62, 0x96, 0x12, 0x83, 0x91, 0x23,
	0x17, 0x7f, 0x78, 0x6a, 0x52, 0x70, 0x7e, 0x72, 0x76, 0x6a, 0x09, 0x54, 0xbb, 0x1e, 0xf1, 0xda,
	0x35, 0x18, 0x0d, 0x18, 0x93, 0xd8, 0xc0, 0x4e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xcb,
	0x48, 0x73, 0xb1, 0xca, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HttpServerClient is the client API for HttpServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HttpServerClient interface {
	GetData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error)
}

type httpServerClient struct {
	cc *grpc.ClientConn
}

func NewHttpServerClient(cc *grpc.ClientConn) HttpServerClient {
	return &httpServerClient{cc}
}

func (c *httpServerClient) GetData(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/wmgrpc.HttpServer/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HttpServerServer is the server API for HttpServer service.
type HttpServerServer interface {
	GetData(context.Context, *Empty) (*Result, error)
}

// UnimplementedHttpServerServer can be embedded to have forward compatible implementations.
type UnimplementedHttpServerServer struct {
}

func (*UnimplementedHttpServerServer) GetData(ctx context.Context, req *Empty) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}

func RegisterHttpServerServer(s *grpc.Server, srv HttpServerServer) {
	s.RegisterService(&_HttpServer_serviceDesc, srv)
}

func _HttpServer_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HttpServerServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wmgrpc.HttpServer/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HttpServerServer).GetData(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _HttpServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wmgrpc.HttpServer",
	HandlerType: (*HttpServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _HttpServer_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/wmgrpc.v1.proto",
}

// WebSocketServerClient is the client API for WebSocketServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebSocketServerClient interface {
	GetData(ctx context.Context, opts ...grpc.CallOption) (WebSocketServer_GetDataClient, error)
}

type webSocketServerClient struct {
	cc *grpc.ClientConn
}

func NewWebSocketServerClient(cc *grpc.ClientConn) WebSocketServerClient {
	return &webSocketServerClient{cc}
}

func (c *webSocketServerClient) GetData(ctx context.Context, opts ...grpc.CallOption) (WebSocketServer_GetDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &_WebSocketServer_serviceDesc.Streams[0], "/wmgrpc.WebSocketServer/GetData", opts...)
	if err != nil {
		return nil, err
	}
	x := &webSocketServerGetDataClient{stream}
	return x, nil
}

type WebSocketServer_GetDataClient interface {
	Send(*Empty) error
	Recv() (*Result, error)
	grpc.ClientStream
}

type webSocketServerGetDataClient struct {
	grpc.ClientStream
}

func (x *webSocketServerGetDataClient) Send(m *Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *webSocketServerGetDataClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WebSocketServerServer is the server API for WebSocketServer service.
type WebSocketServerServer interface {
	GetData(WebSocketServer_GetDataServer) error
}

// UnimplementedWebSocketServerServer can be embedded to have forward compatible implementations.
type UnimplementedWebSocketServerServer struct {
}

func (*UnimplementedWebSocketServerServer) GetData(srv WebSocketServer_GetDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetData not implemented")
}

func RegisterWebSocketServerServer(s *grpc.Server, srv WebSocketServerServer) {
	s.RegisterService(&_WebSocketServer_serviceDesc, srv)
}

func _WebSocketServer_GetData_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WebSocketServerServer).GetData(&webSocketServerGetDataServer{stream})
}

type WebSocketServer_GetDataServer interface {
	Send(*Result) error
	Recv() (*Empty, error)
	grpc.ServerStream
}

type webSocketServerGetDataServer struct {
	grpc.ServerStream
}

func (x *webSocketServerGetDataServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func (x *webSocketServerGetDataServer) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _WebSocketServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wmgrpc.WebSocketServer",
	HandlerType: (*WebSocketServerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetData",
			Handler:       _WebSocketServer_GetData_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/wmgrpc.v1.proto",
}
