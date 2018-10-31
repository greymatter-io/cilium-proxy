// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/network/rate_limit/v2/rate_limit.proto

package v2

import (
	fmt "fmt"
	ratelimit "github.com/cilium/proxy/go/envoy/api/v2/ratelimit"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/lyft/protoc-gen-validate/validate"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RateLimit struct {
	// The prefix to use when emitting :ref:`statistics <config_network_filters_rate_limit_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// The rate limit domain to use in the rate limit service request.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// The rate limit descriptor list to use in the rate limit service request.
	Descriptors []*ratelimit.RateLimitDescriptor `protobuf:"bytes,3,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny      bool     `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_34e9a222968daa71, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.network.rate_limit.v2.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/network/rate_limit/v2/rate_limit.proto", fileDescriptor_34e9a222968daa71)
}

var fileDescriptor_34e9a222968daa71 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x6e, 0xa3, 0x40,
	0x10, 0x86, 0xb5, 0x18, 0xfb, 0xec, 0xa5, 0x38, 0x1d, 0x3a, 0xe9, 0x38, 0x17, 0x77, 0xdc, 0x15,
	0x27, 0xce, 0x91, 0x76, 0x25, 0x52, 0x25, 0x25, 0x72, 0x99, 0x48, 0x11, 0x5d, 0xd2, 0xa0, 0xb5,
	0x19, 0xd0, 0x2a, 0xc0, 0xa0, 0xf5, 0x42, 0xe2, 0xd7, 0x48, 0x95, 0x22, 0x4f, 0x92, 0x2a, 0x6f,
	0x92, 0x3a, 0x6f, 0x11, 0xc1, 0x82, 0x63, 0xa5, 0xfb, 0x87, 0x7f, 0xfe, 0x6f, 0xd8, 0x9f, 0x9e,
	0x43, 0xd5, 0xe2, 0x9e, 0x6f, 0xb1, 0xca, 0x64, 0xce, 0x33, 0x59, 0x68, 0x50, 0xbc, 0x02, 0x7d,
	0x87, 0xea, 0x96, 0x2b, 0xa1, 0x21, 0x29, 0x64, 0x29, 0x35, 0x6f, 0xc3, 0xa3, 0x89, 0xd5, 0x0a,
	0x35, 0xba, 0xff, 0xfb, 0x2c, 0x33, 0x59, 0x66, 0xb2, 0x6c, 0xc8, 0xb2, 0xa3, 0xed, 0x36, 0x5c,
	0xfe, 0x33, 0x67, 0x44, 0x2d, 0x47, 0x92, 0xc1, 0x1e, 0x94, 0x41, 0x2e, 0x7f, 0xe5, 0x88, 0x79,
	0x01, 0xbc, 0x9f, 0x36, 0x4d, 0xc6, 0xd3, 0x46, 0x09, 0x2d, 0xb1, 0x1a, 0xfc, 0x1f, 0xad, 0x28,
	0x64, 0x2a, 0x34, 0xf0, 0x51, 0x0c, 0xc6, 0xf7, 0x1c, 0x73, 0xec, 0x25, 0xef, 0x94, 0xf9, 0xfa,
	0xf7, 0xc9, 0xa2, 0x8b, 0x58, 0x68, 0xb8, 0xe8, 0x4e, 0xb8, 0x2b, 0xea, 0xec, 0xb4, 0xd0, 0x49,
	0xad, 0x20, 0x93, 0xf7, 0x1e, 0xf1, 0x49, 0xb0, 0x88, 0x16, 0xcf, 0x6f, 0x2f, 0x13, 0x5b, 0x59,
	0x3e, 0x89, 0x69, 0xe7, 0x5e, 0xf5, 0xa6, 0xfb, 0x87, 0xce, 0x52, 0x2c, 0x85, 0xac, 0x3c, 0xeb,
	0xf3, 0xda, 0x60, 0xb8, 0xd7, 0xd4, 0x49, 0x61, 0xb7, 0x55, 0xb2, 0xd6, 0xa8, 0x76, 0xde, 0xc4,
	0x9f, 0x04, 0x4e, 0x78, 0xc2, 0x4c, 0x29, 0xa2, 0x96, 0xac, 0x0d, 0xd9, 0xc7, 0xfb, 0x0e, 0xbf,
	0xb1, 0x3e, 0x64, 0x22, 0xda, 0x41, 0xa7, 0x0f, 0xc4, 0x9a, 0x93, 0xf8, 0x98, 0xe5, 0x9e, 0xd1,
	0x2f, 0x5a, 0x96, 0x80, 0x8d, 0xf6, 0x6c, 0x9f, 0x04, 0x4e, 0xf8, 0x93, 0x99, 0x62, 0xd8, 0x58,
	0x0c, 0x5b, 0x0f, 0xc5, 0x44, 0xf6, 0xe3, 0xeb, 0x6f, 0x12, 0x8f, 0xfb, 0xee, 0x8a, 0x7e, 0xcb,
	0x84, 0x2c, 0x1a, 0x05, 0x49, 0x89, 0x29, 0x24, 0x29, 0x54, 0x7b, 0x6f, 0xea, 0x93, 0x60, 0x1e,
	0x7f, 0x1d, 0x8c, 0x4b, 0x4c, 0x61, 0x0d, 0xd5, 0x3e, 0xb2, 0x6f, 0xac, 0x36, 0xdc, 0xcc, 0x7a,
	0xe6, 0xe9, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x6d, 0xf4, 0xe3, 0x0b, 0x02, 0x00, 0x00,
}
