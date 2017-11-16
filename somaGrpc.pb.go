// Code generated by protoc-gen-go. DO NOT EDIT.
// source: somaGrpc.proto

/*
Package gserver is a generated protocol buffer package.

It is generated from these files:
	somaGrpc.proto

It has these top-level messages:
	UserQuery
	UserInfo
*/
package gserver

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

type UserQuery struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *UserQuery) Reset()                    { *m = UserQuery{} }
func (m *UserQuery) String() string            { return proto.CompactTextString(m) }
func (*UserQuery) ProtoMessage()               {}
func (*UserQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserQuery) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UserInfo struct {
	LoginName   string `protobuf:"bytes,1,opt,name=loginName" json:"loginName,omitempty"`
	UserName    string `protobuf:"bytes,2,opt,name=userName" json:"userName,omitempty"`
	OrgName     string `protobuf:"bytes,3,opt,name=orgName" json:"orgName,omitempty"`
	BureauName  string `protobuf:"bytes,4,opt,name=bureauName" json:"bureauName,omitempty"`
	StationName string `protobuf:"bytes,5,opt,name=StationName" json:"StationName,omitempty"`
	ShopName    string `protobuf:"bytes,6,opt,name=ShopName" json:"ShopName,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserInfo) GetLoginName() string {
	if m != nil {
		return m.LoginName
	}
	return ""
}

func (m *UserInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfo) GetOrgName() string {
	if m != nil {
		return m.OrgName
	}
	return ""
}

func (m *UserInfo) GetBureauName() string {
	if m != nil {
		return m.BureauName
	}
	return ""
}

func (m *UserInfo) GetStationName() string {
	if m != nil {
		return m.StationName
	}
	return ""
}

func (m *UserInfo) GetShopName() string {
	if m != nil {
		return m.ShopName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserQuery)(nil), "gserver.UserQuery")
	proto.RegisterType((*UserInfo)(nil), "gserver.UserInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SomaRouteGuide service

type SomaRouteGuideClient interface {
	// rpc GetFeature(Point) returns (HelloReply) {}
	// rpc GetEngWorkRecords(Days) returns (stream Records) {}
	AnalyticsUserInfoHandler(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*UserInfo, error)
}

type somaRouteGuideClient struct {
	cc *grpc.ClientConn
}

func NewSomaRouteGuideClient(cc *grpc.ClientConn) SomaRouteGuideClient {
	return &somaRouteGuideClient{cc}
}

func (c *somaRouteGuideClient) AnalyticsUserInfoHandler(ctx context.Context, in *UserQuery, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := grpc.Invoke(ctx, "/gserver.SomaRouteGuide/AnalyticsUserInfoHandler", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SomaRouteGuide service

type SomaRouteGuideServer interface {
	// rpc GetFeature(Point) returns (HelloReply) {}
	// rpc GetEngWorkRecords(Days) returns (stream Records) {}
	AnalyticsUserInfoHandler(context.Context, *UserQuery) (*UserInfo, error)
}

func RegisterSomaRouteGuideServer(s *grpc.Server, srv SomaRouteGuideServer) {
	s.RegisterService(&_SomaRouteGuide_serviceDesc, srv)
}

func _SomaRouteGuide_AnalyticsUserInfoHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SomaRouteGuideServer).AnalyticsUserInfoHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gserver.SomaRouteGuide/AnalyticsUserInfoHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SomaRouteGuideServer).AnalyticsUserInfoHandler(ctx, req.(*UserQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _SomaRouteGuide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gserver.SomaRouteGuide",
	HandlerType: (*SomaRouteGuideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnalyticsUserInfoHandler",
			Handler:    _SomaRouteGuide_AnalyticsUserInfoHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "somaGrpc.proto",
}

func init() { proto.RegisterFile("somaGrpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x8d, 0xd6, 0xb6, 0x19, 0xa1, 0xe0, 0x9c, 0x42, 0x11, 0x2d, 0x39, 0x79, 0xca, 0x41,
	0x7f, 0x81, 0x78, 0xa8, 0x5e, 0x04, 0x1b, 0xfa, 0x03, 0xb6, 0xed, 0x18, 0x17, 0x92, 0x9d, 0x30,
	0xbb, 0x2b, 0xf4, 0xbf, 0xf9, 0xe3, 0x24, 0x53, 0xd3, 0xa6, 0xb7, 0x79, 0xef, 0x7b, 0xf0, 0x98,
	0x07, 0x33, 0xcf, 0x8d, 0x59, 0x4a, 0xbb, 0x2d, 0x5a, 0xe1, 0xc0, 0x38, 0xa9, 0x3c, 0xc9, 0x0f,
	0x49, 0xfe, 0x00, 0xe9, 0xda, 0x93, 0x7c, 0x46, 0x92, 0x3d, 0x22, 0x8c, 0x9c, 0x69, 0x28, 0x4b,
	0x16, 0xc9, 0x63, 0xba, 0xd2, 0x3b, 0xff, 0x4d, 0x60, 0xda, 0x25, 0xde, 0xdd, 0x17, 0xe3, 0x1d,
	0xa4, 0x35, 0x57, 0xd6, 0x7d, 0x9c, 0x52, 0x27, 0x03, 0xe7, 0x30, 0x8d, 0x9e, 0x44, 0xe1, 0xa5,
	0xc2, 0xa3, 0xc6, 0x0c, 0x26, 0x2c, 0x95, 0xa2, 0x2b, 0x45, 0xbd, 0xc4, 0x7b, 0x80, 0x4d, 0x14,
	0x32, 0x51, 0xe1, 0x48, 0xe1, 0xc0, 0xc1, 0x05, 0xdc, 0x94, 0xc1, 0x04, 0xcb, 0x87, 0xd6, 0x6b,
	0x0d, 0x0c, 0xad, 0xae, 0xb7, 0xfc, 0xe6, 0x56, 0xf1, 0xf8, 0xd0, 0xdb, 0xeb, 0xa7, 0x35, 0xcc,
	0x4a, 0x6e, 0xcc, 0x8a, 0x63, 0xa0, 0x65, 0xb4, 0x3b, 0xc2, 0x57, 0xc8, 0x5e, 0x9c, 0xa9, 0xf7,
	0xc1, 0x6e, 0x7d, 0xff, 0xd8, 0x9b, 0x71, 0xbb, 0x9a, 0x04, 0xb1, 0xf8, 0xdf, 0xa5, 0x38, 0x8e,
	0x32, 0xbf, 0x3d, 0xf3, 0xba, 0x74, 0x7e, 0xb1, 0x19, 0xeb, 0x8c, 0xcf, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x31, 0x50, 0xd4, 0x27, 0x58, 0x01, 0x00, 0x00,
}
