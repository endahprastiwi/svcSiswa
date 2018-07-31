// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package siswa

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type SiswaRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SiswaRequest) Reset()         { *m = SiswaRequest{} }
func (m *SiswaRequest) String() string { return proto.CompactTextString(m) }
func (*SiswaRequest) ProtoMessage()    {}

func (*SiswaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_siswa_71e208cbdc16936b, []int{0}
}
func (m *SiswaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SiswaRequest.Unmarshal(m, b)
}
func (m *SiswaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SiswaRequest.Marshal(b, m, deterministic)
}
func (dst *SiswaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SiswaRequest.Merge(dst, src)
}
func (m *SiswaRequest) XXX_Size() int {
	return xxx_messageInfo_SiswaRequest.Size(m)
}
func (m *SiswaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SiswaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SiswaRequest proto.InternalMessageInfo

func (m *SiswaRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type SiswaReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SiswaReply) Reset()         { *m = SiswaReply{} }
func (m *SiswaReply) String() string { return proto.CompactTextString(m) }
func (*SiswaReply) ProtoMessage()    {}
func (*SiswaReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_siswa_71e208cbdc16936b, []int{1}
}
func (m *SiswaReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SiswaReply.Unmarshal(m, b)
}
func (m *SiswaReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SiswaReply.Marshal(b, m, deterministic)
}
func (dst *SiswaReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SiswaReply.Merge(dst, src)
}
func (m *SiswaReply) XXX_Size() int {
	return xxx_messageInfo_SiswaReply.Size(m)
}
func (m *SiswaReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SiswaReply.DiscardUnknown(m)
}

var xxx_messageInfo_SiswaReply proto.InternalMessageInfo

func (m *SiswaReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SiswaRequest)(nil), "siswa.SiswaRequest")
	proto.RegisterType((*SiswaReply)(nil), "siswa.SiswaReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SiswaClient interface {
	// Sends a greeting
	ViewSiswa(ctx context.Context, in *SiswaRequest, opts ...grpc.CallOption) (*SiswaReply, error)
}

type siswaClient struct {
	cc *grpc.ClientConn
}

func NewSiswaClient(cc *grpc.ClientConn) SiswaClient {
	return &siswaClient{cc}
}

func (c *siswaClient) ViewSiswa(ctx context.Context, in *SiswaRequest, opts ...grpc.CallOption) (*SiswaReply, error) {
	out := new(SiswaReply)
	err := c.cc.Invoke(ctx, "/siswa.Siswa/ViewSiswa", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type SiswaServer interface {
	// Sends a greeting
	ViewSiswa(context.Context, *SiswaRequest) (*SiswaReply, error)
}

func RegisterSiswaServer(s *grpc.Server, srv SiswaServer) {
	s.RegisterService(&_Siswa_serviceDesc, srv)
}

func _Siswa_ViewSiswa_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SiswaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SiswaServer).ViewSiswa(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/siswa.Siswa/ViewSiswa",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SiswaServer).ViewSiswa(ctx, req.(*SiswaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Siswa_serviceDesc = grpc.ServiceDesc{
	ServiceName: "siswa.Siswa",
	HandlerType: (*SiswaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ViewSiswa",
			Handler:    _Siswa_ViewSiswa_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "siswa.proto",
}

func init() { proto.RegisterFile("siswa.proto", fileDescriptor_siswa_71e208cbdc16936b) }

var fileDescriptor_siswa_71e208cbdc16936b = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x28, 0x29, 0x71, 0xf1, 0x78, 0x80, 0x78, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42,
	0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92,
	0x1a, 0x17, 0x17, 0x54, 0x4d, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71,
	0x62, 0x3a, 0x4c, 0x11, 0x8c, 0x6b, 0xe4, 0xc9, 0xc5, 0xee, 0x5e, 0x94, 0x9a, 0x5a, 0x92, 0x5a,
	0x24, 0x64, 0xc7, 0xc5, 0x11, 0x9c, 0x58, 0x09, 0xd6, 0x25, 0x24, 0xa1, 0x87, 0xe4, 0x02, 0x64,
	0xcb, 0xa4, 0xc4, 0xb0, 0xc8, 0x14, 0xe4, 0x54, 0x2a, 0x31, 0x38, 0x19, 0x70, 0x49, 0x67, 0xe6,
	0xeb, 0xa5, 0x17, 0x15, 0x24, 0xeb, 0xa5, 0x56, 0x24, 0xe6, 0x16, 0xe4, 0xa4, 0x16, 0x23, 0xa9,
	0x75, 0xe2, 0x07, 0x2b, 0x0e, 0x07, 0xb1, 0x03, 0x40, 0x5e, 0x0a, 0x60, 0x4c, 0x62, 0x03, 0xfb,
	0xcd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x0f, 0xb7, 0xcd, 0xf2, 0xef, 0x00, 0x00, 0x00,
}