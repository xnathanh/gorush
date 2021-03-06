// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gorush.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	gorush.proto

It has these top-level messages:
	Alert
	NotificationRequest
	NotificationReply
	HealthCheckRequest
	HealthCheckResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto1.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

type Alert struct {
	Title    string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Body     string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	Subtitle string `protobuf:"bytes,3,opt,name=subtitle" json:"subtitle,omitempty"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto1.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Alert) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Alert) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Alert) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

type NotificationRequest struct {
	Tokens   []string `protobuf:"bytes,1,rep,name=tokens" json:"tokens,omitempty"`
	Platform int32    `protobuf:"varint,2,opt,name=platform" json:"platform,omitempty"`
	Message  string   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	Title    string   `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Topic    string   `protobuf:"bytes,5,opt,name=topic" json:"topic,omitempty"`
	Key      string   `protobuf:"bytes,6,opt,name=key" json:"key,omitempty"`
	Badge    int32    `protobuf:"varint,7,opt,name=badge" json:"badge,omitempty"`
	Category string   `protobuf:"bytes,8,opt,name=category" json:"category,omitempty"`
	Alert    *Alert   `protobuf:"bytes,9,opt,name=alert" json:"alert,omitempty"`
	Sound    string   `protobuf:"bytes,10,opt,name=sound" json:"sound,omitempty"`
	Userinfo string   `protobuf:"bytes,8,opt,name=userinfo" json:"userinfo,omitempty"`
}

func (m *NotificationRequest) Reset()                    { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string            { return proto1.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()               {}
func (*NotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NotificationRequest) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *NotificationRequest) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *NotificationRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *NotificationRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NotificationRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *NotificationRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *NotificationRequest) GetBadge() int32 {
	if m != nil {
		return m.Badge
	}
	return 0
}

func (m *NotificationRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *NotificationRequest) GetAlert() *Alert {
	if m != nil {
		return m.Alert
	}
	return nil
}

func (m *NotificationRequest) GetSound() string {
	if m != nil {
		return m.Sound
	}
	return ""
}

func (m *NotificationRequest) GetUserinfo() string {
	if m != nil {
		return m.Userinfo
	}
	return ""
}

type NotificationReply struct {
	Success bool  `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Counts  int32 `protobuf:"varint,2,opt,name=counts" json:"counts,omitempty"`
}

func (m *NotificationReply) Reset()                    { *m = NotificationReply{} }
func (m *NotificationReply) String() string            { return proto1.CompactTextString(m) }
func (*NotificationReply) ProtoMessage()               {}
func (*NotificationReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NotificationReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *NotificationReply) GetCounts() int32 {
	if m != nil {
		return m.Counts
	}
	return 0
}

type HealthCheckRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto1.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,enum=proto.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto1.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto1.RegisterType((*Alert)(nil), "proto.Alert")
	proto1.RegisterType((*NotificationRequest)(nil), "proto.NotificationRequest")
	proto1.RegisterType((*NotificationReply)(nil), "proto.NotificationReply")
	proto1.RegisterType((*HealthCheckRequest)(nil), "proto.HealthCheckRequest")
	proto1.RegisterType((*HealthCheckResponse)(nil), "proto.HealthCheckResponse")
	proto1.RegisterEnum("proto.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Gorush service

type GorushClient interface {
	Send(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationReply, error)
}

type gorushClient struct {
	cc *grpc.ClientConn
}

func NewGorushClient(cc *grpc.ClientConn) GorushClient {
	return &gorushClient{cc}
}

func (c *gorushClient) Send(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationReply, error) {
	out := new(NotificationReply)
	err := grpc.Invoke(ctx, "/proto.Gorush/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gorush service

type GorushServer interface {
	Send(context.Context, *NotificationRequest) (*NotificationReply, error)
}

func RegisterGorushServer(s *grpc.Server, srv GorushServer) {
	s.RegisterService(&_Gorush_serviceDesc, srv)
}

func _Gorush_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GorushServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Gorush/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GorushServer).Send(ctx, req.(*NotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gorush_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Gorush",
	HandlerType: (*GorushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Gorush_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gorush.proto",
}

// Client API for Health service

type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/proto.Health/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gorush.proto",
}

func init() { proto1.RegisterFile("gorush.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0x37, 0x6d, 0x93, 0xb6, 0xd3, 0x05, 0xca, 0x2c, 0x42, 0xa6, 0xa7, 0xca, 0xa7, 0x9e,
	0x7a, 0x28, 0x37, 0x0e, 0x2b, 0x10, 0x5a, 0xed, 0x22, 0x44, 0x56, 0x72, 0xf9, 0x73, 0x44, 0xae,
	0xeb, 0x4d, 0xa3, 0x66, 0xe3, 0x10, 0x3b, 0x48, 0x79, 0x08, 0x4e, 0xbc, 0x30, 0xb2, 0x1d, 0x17,
	0x2a, 0x85, 0x53, 0xfc, 0x9b, 0x19, 0xcf, 0x37, 0xf3, 0xc5, 0x70, 0x99, 0xa9, 0xba, 0xd1, 0x87,
	0x75, 0x55, 0x2b, 0xa3, 0x30, 0x76, 0x1f, 0xfa, 0x09, 0xe2, 0x77, 0x85, 0xac, 0x0d, 0xbe, 0x80,
	0xd8, 0xe4, 0xa6, 0x90, 0x24, 0x5a, 0x46, 0xab, 0x29, 0xf3, 0x80, 0x08, 0xa3, 0x9d, 0xda, 0xb7,
	0x64, 0xe0, 0x82, 0xee, 0x8c, 0x0b, 0x98, 0xe8, 0x66, 0xe7, 0x8b, 0x87, 0x2e, 0x7e, 0x62, 0xfa,
	0x6b, 0x00, 0x57, 0xa9, 0x32, 0xf9, 0x43, 0x2e, 0xb8, 0xc9, 0x55, 0xc9, 0xe4, 0x8f, 0x46, 0x6a,
	0x83, 0x2f, 0x21, 0x31, 0xea, 0x28, 0x4b, 0x4d, 0xa2, 0xe5, 0x70, 0x35, 0x65, 0x1d, 0xd9, 0x5e,
	0x55, 0xc1, 0xcd, 0x83, 0xaa, 0x1f, 0x9d, 0x46, 0xcc, 0x4e, 0x8c, 0x04, 0xc6, 0x8f, 0x52, 0x6b,
	0x9e, 0x05, 0x99, 0x80, 0x7f, 0x67, 0x1d, 0xfd, 0x3b, 0xab, 0x8d, 0xaa, 0x2a, 0x17, 0x24, 0xee,
	0xa2, 0x16, 0x70, 0x0e, 0xc3, 0xa3, 0x6c, 0x49, 0xe2, 0x62, 0xf6, 0x68, 0xeb, 0x76, 0x7c, 0x9f,
	0x49, 0x32, 0x76, 0x82, 0x1e, 0xec, 0x24, 0x82, 0x1b, 0x99, 0xa9, 0xba, 0x25, 0x13, 0xbf, 0x55,
	0x60, 0xa4, 0x10, 0x73, 0x6b, 0x12, 0x99, 0x2e, 0xa3, 0xd5, 0x6c, 0x73, 0xe9, 0x2d, 0x5c, 0x3b,
	0xe3, 0x98, 0x4f, 0xd9, 0xae, 0x5a, 0x35, 0xe5, 0x9e, 0x80, 0x57, 0x77, 0x40, 0x6f, 0xe0, 0xf9,
	0xb9, 0x1d, 0x55, 0xd1, 0xda, 0xc5, 0x74, 0x23, 0x84, 0xd4, 0xda, 0x99, 0x3d, 0x61, 0x01, 0xad,
	0x4d, 0x42, 0x35, 0xa5, 0xd1, 0x9d, 0x19, 0x1d, 0xd1, 0x35, 0xe0, 0x9d, 0xe4, 0x85, 0x39, 0xbc,
	0x3f, 0x48, 0x71, 0x0c, 0xa6, 0xda, 0x3e, 0xb2, 0xfe, 0x99, 0x8b, 0xf0, 0xd3, 0x02, 0xd2, 0xdf,
	0x11, 0x5c, 0x9d, 0x5d, 0xd0, 0x95, 0x2a, 0xb5, 0xc4, 0xb7, 0x90, 0x68, 0xc3, 0x4d, 0xe3, 0x85,
	0x9f, 0x6e, 0x56, 0xdd, 0x26, 0x3d, 0xb5, 0xeb, 0xad, 0xed, 0x55, 0x66, 0x5b, 0x57, 0xcf, 0xba,
	0x7b, 0xf4, 0x0d, 0x3c, 0x39, 0x4b, 0xe0, 0x0c, 0xc6, 0x5f, 0xd2, 0x8f, 0xe9, 0xfd, 0xb7, 0x74,
	0x7e, 0x61, 0x61, 0x7b, 0xc3, 0xbe, 0x7e, 0x48, 0x6f, 0xe7, 0x11, 0x3e, 0x83, 0x59, 0x7a, 0xff,
	0xf9, 0x7b, 0x08, 0x0c, 0x36, 0x77, 0x90, 0xdc, 0xba, 0x27, 0x88, 0xd7, 0x30, 0xda, 0xca, 0x72,
	0x8f, 0x8b, 0x4e, 0xbf, 0xe7, 0xc9, 0x2c, 0x48, 0x6f, 0xae, 0x2a, 0x5a, 0x7a, 0x61, 0x3b, 0xf9,
	0x91, 0xf1, 0x1a, 0x62, 0x37, 0x36, 0xbe, 0xea, 0x5b, 0xc5, 0x77, 0x5a, 0xfc, 0x7f, 0xcb, 0x5d,
	0xe2, 0x52, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xf6, 0xa7, 0x4d, 0x1d, 0x03, 0x00,
	0x00,
}
