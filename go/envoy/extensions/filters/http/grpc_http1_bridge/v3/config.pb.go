// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/grpc_http1_bridge/v3/config.proto

package envoy_extensions_filters_http_grpc_http1_bridge_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
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

// gRPC HTTP/1.1 Bridge filter config.
type Config struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_a9d1a8190e8dfb91, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Config)(nil), "envoy.extensions.filters.http.grpc_http1_bridge.v3.Config")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/grpc_http1_bridge/v3/config.proto", fileDescriptor_a9d1a8190e8dfb91)
}

var fileDescriptor_a9d1a8190e8dfb91 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x4f, 0xcb, 0xcc,
	0x29, 0x49, 0x2d, 0x2a, 0xd6, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x4f, 0x2f, 0x2a, 0x48, 0x8e, 0x07,
	0xb1, 0x0c, 0xe3, 0x93, 0x8a, 0x32, 0x53, 0xd2, 0x53, 0xf5, 0xcb, 0x8c, 0xf5, 0x93, 0xf3, 0xf3,
	0xd2, 0x32, 0xd3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x8c, 0xc0, 0x06, 0xe8, 0x21, 0x0c,
	0xd0, 0x83, 0x1a, 0xa0, 0x07, 0xd2, 0xa6, 0x87, 0x61, 0x80, 0x5e, 0x99, 0xb1, 0x94, 0x62, 0x69,
	0x4a, 0x41, 0xa2, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x09, 0xd8, 0xd2, 0xb2, 0xd4, 0x22,
	0x90, 0xe6, 0xcc, 0x3c, 0xa8, 0xb1, 0x4a, 0xae, 0x5c, 0x6c, 0xce, 0x60, 0x6b, 0xac, 0xac, 0x67,
	0x1d, 0xed, 0x90, 0x33, 0xe3, 0x32, 0x81, 0xd8, 0x03, 0xb5, 0x1b, 0x62, 0x07, 0x4e, 0x2b, 0x8c,
	0xf4, 0x20, 0x9a, 0x9d, 0x02, 0xb9, 0x1c, 0x32, 0xf3, 0xf5, 0xc0, 0x5a, 0x0b, 0x8a, 0xf2, 0x2b,
	0x2a, 0xf5, 0x48, 0x77, 0xad, 0x13, 0x37, 0xc4, 0xac, 0x00, 0x90, 0xbb, 0x02, 0x18, 0x93, 0xd8,
	0xc0, 0x0e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xd0, 0xea, 0x62, 0x3a, 0x01, 0x00,
	0x00,
}