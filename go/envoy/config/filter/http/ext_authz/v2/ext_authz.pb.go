// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/ext_authz/v2/ext_authz.proto

package v2

import (
	fmt "fmt"
	core "github.com/cilium/proxy/go/envoy/api/v2/core"
	matcher "github.com/cilium/proxy/go/envoy/type/matcher"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ExtAuthz struct {
	// External authorization service configuration.
	//
	// Types that are valid to be assigned to Services:
	//	*ExtAuthz_GrpcService
	//	*ExtAuthz_HttpService
	Services isExtAuthz_Services `protobuf_oneof:"services"`
	//  Changes filter's behaviour on errors:
	//
	//  1. When set to true, the filter will *accept* client request even if the communication with
	//  the authorization service has failed, or if the authorization service has returned a HTTP 5xx
	//  error.
	//
	//  2. When set to false, ext-authz will *reject* client requests and return a *Forbidden*
	//  response if the communication with the authorization service has failed, or if the
	//  authorization service has returned a HTTP 5xx error.
	//
	// Note that errors can be *always* tracked in the :ref:`stats
	// <config_http_filters_ext_authz_stats>`.
	FailureModeAllow bool `protobuf:"varint,2,opt,name=failure_mode_allow,json=failureModeAllow,proto3" json:"failure_mode_allow,omitempty"`
	// Sets the package version the gRPC service should use. This is particularly
	// useful when transitioning from alpha to release versions assuming that both definitions are
	// semantically compatible. Deprecation note: This field is deprecated and should only be used for
	// version upgrade. See release notes for more details.
	UseAlpha             bool     `protobuf:"varint,4,opt,name=use_alpha,json=useAlpha,proto3" json:"use_alpha,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtAuthz) Reset()         { *m = ExtAuthz{} }
func (m *ExtAuthz) String() string { return proto.CompactTextString(m) }
func (*ExtAuthz) ProtoMessage()    {}
func (*ExtAuthz) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{0}
}

func (m *ExtAuthz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthz.Unmarshal(m, b)
}
func (m *ExtAuthz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthz.Marshal(b, m, deterministic)
}
func (m *ExtAuthz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthz.Merge(m, src)
}
func (m *ExtAuthz) XXX_Size() int {
	return xxx_messageInfo_ExtAuthz.Size(m)
}
func (m *ExtAuthz) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthz.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthz proto.InternalMessageInfo

type isExtAuthz_Services interface {
	isExtAuthz_Services()
}

type ExtAuthz_GrpcService struct {
	GrpcService *core.GrpcService `protobuf:"bytes,1,opt,name=grpc_service,json=grpcService,proto3,oneof"`
}

type ExtAuthz_HttpService struct {
	HttpService *HttpService `protobuf:"bytes,3,opt,name=http_service,json=httpService,proto3,oneof"`
}

func (*ExtAuthz_GrpcService) isExtAuthz_Services() {}

func (*ExtAuthz_HttpService) isExtAuthz_Services() {}

func (m *ExtAuthz) GetServices() isExtAuthz_Services {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ExtAuthz) GetGrpcService() *core.GrpcService {
	if x, ok := m.GetServices().(*ExtAuthz_GrpcService); ok {
		return x.GrpcService
	}
	return nil
}

func (m *ExtAuthz) GetHttpService() *HttpService {
	if x, ok := m.GetServices().(*ExtAuthz_HttpService); ok {
		return x.HttpService
	}
	return nil
}

func (m *ExtAuthz) GetFailureModeAllow() bool {
	if m != nil {
		return m.FailureModeAllow
	}
	return false
}

// Deprecated: Do not use.
func (m *ExtAuthz) GetUseAlpha() bool {
	if m != nil {
		return m.UseAlpha
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthz) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthz_GrpcService)(nil),
		(*ExtAuthz_HttpService)(nil),
	}
}

// HttpService is used for raw HTTP communication between the filter and the authorization service.
// When configured, the filter will parse the client request and use these attributes to call the
// authorization server. Depending on the response, the filter may reject or accept the client
// request. Note that in any of these events, metadata can be added, removed or overridden by the
// filter:
//
// *On authorization request*, a list of allowed request headers may be supplied. See
// :ref:`allowed_headers
// <envoy_api_field_config.filter.http.ext_authz.v2.AuthorizationRequest.allowed_headers>`
// for details. Additional headers metadata may be added to the authorization request. See
// :ref:`headers_to_add
// <envoy_api_field_config.filter.http.ext_authz.v2.AuthorizationRequest.headers_to_add>` for
// details.
//
// On authorization response status HTTP 200 OK, the filter will allow traffic to the upstream and
// additional headers metadata may be added to the original client request. See
// :ref:`allowed_upstream_headers
// <envoy_api_field_config.filter.http.ext_authz.v2.AuthorizationResponse.allowed_upstream_headers>`
// for details.
//
// On other authorization response statuses, the filter will not allow traffic. Additional headers
// metadata as well as body may be added to the client's response. See :ref:`allowed_client_headers
// <envoy_api_field_config.filter.http.ext_authz.v2.AuthorizationResponse.allowed_client_headers>`
// for details.
type HttpService struct {
	// Sets the HTTP server URI which the authorization requests must be sent to.
	ServerUri *core.HttpUri `protobuf:"bytes,1,opt,name=server_uri,json=serverUri,proto3" json:"server_uri,omitempty"`
	// Sets a prefix to the value of authorization request header *Path*.
	PathPrefix string `protobuf:"bytes,2,opt,name=path_prefix,json=pathPrefix,proto3" json:"path_prefix,omitempty"`
	// Settings used for controlling authorization request metadata.
	AuthorizationRequest *AuthorizationRequest `protobuf:"bytes,7,opt,name=authorization_request,json=authorizationRequest,proto3" json:"authorization_request,omitempty"`
	// Settings used for controlling authorization response metadata.
	AuthorizationResponse *AuthorizationResponse `protobuf:"bytes,8,opt,name=authorization_response,json=authorizationResponse,proto3" json:"authorization_response,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *HttpService) Reset()         { *m = HttpService{} }
func (m *HttpService) String() string { return proto.CompactTextString(m) }
func (*HttpService) ProtoMessage()    {}
func (*HttpService) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{1}
}

func (m *HttpService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpService.Unmarshal(m, b)
}
func (m *HttpService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpService.Marshal(b, m, deterministic)
}
func (m *HttpService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpService.Merge(m, src)
}
func (m *HttpService) XXX_Size() int {
	return xxx_messageInfo_HttpService.Size(m)
}
func (m *HttpService) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpService.DiscardUnknown(m)
}

var xxx_messageInfo_HttpService proto.InternalMessageInfo

func (m *HttpService) GetServerUri() *core.HttpUri {
	if m != nil {
		return m.ServerUri
	}
	return nil
}

func (m *HttpService) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *HttpService) GetAuthorizationRequest() *AuthorizationRequest {
	if m != nil {
		return m.AuthorizationRequest
	}
	return nil
}

func (m *HttpService) GetAuthorizationResponse() *AuthorizationResponse {
	if m != nil {
		return m.AuthorizationResponse
	}
	return nil
}

type AuthorizationRequest struct {
	// Authorization request will include the client request headers that have a correspondent match
	// in the :ref:`list <envoy_api_msg_type.matcher.ListStringMatcher>`. Note that in addition to the
	// user's supplied matchers:
	//
	// 1. *Host*, *Method*, *Path* and *Content-Length* are automatically included to the list.
	//
	// 2. *Content-Length* will be set to 0 and the request to the authorization service will not have
	// a message body.
	//
	AllowedHeaders *matcher.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_headers,json=allowedHeaders,proto3" json:"allowed_headers,omitempty"`
	// Sets a list of headers that will be included to the request to authorization service. Note that
	// client request of the same key will be overridden.
	HeadersToAdd         []*core.HeaderValue `protobuf:"bytes,2,rep,name=headers_to_add,json=headersToAdd,proto3" json:"headers_to_add,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *AuthorizationRequest) Reset()         { *m = AuthorizationRequest{} }
func (m *AuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizationRequest) ProtoMessage()    {}
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{2}
}

func (m *AuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationRequest.Unmarshal(m, b)
}
func (m *AuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationRequest.Merge(m, src)
}
func (m *AuthorizationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizationRequest.Size(m)
}
func (m *AuthorizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationRequest proto.InternalMessageInfo

func (m *AuthorizationRequest) GetAllowedHeaders() *matcher.ListStringMatcher {
	if m != nil {
		return m.AllowedHeaders
	}
	return nil
}

func (m *AuthorizationRequest) GetHeadersToAdd() []*core.HeaderValue {
	if m != nil {
		return m.HeadersToAdd
	}
	return nil
}

type AuthorizationResponse struct {
	// When this :ref:`list <envoy_api_msg_type.matcher.ListStringMatcher>` is set, authorization
	// response headers that have a correspondent match will be added to the original client request.
	// Note that coexistent headers will be overridden.
	AllowedUpstreamHeaders *matcher.ListStringMatcher `protobuf:"bytes,1,opt,name=allowed_upstream_headers,json=allowedUpstreamHeaders,proto3" json:"allowed_upstream_headers,omitempty"`
	// When this :ref:`list <envoy_api_msg_type.matcher.ListStringMatcher>`. is set, authorization
	// response headers that have a correspondent match will be added to the client's response. Note
	// that when this list is *not* set, all the authorization response headers, except *Authority
	// (Host)* will be in the response to the client. When a header is included in this list, *Path*,
	// *Status*, *Content-Length*, *WWWAuthenticate* and *Location* are automatically added.
	AllowedClientHeaders *matcher.ListStringMatcher `protobuf:"bytes,2,opt,name=allowed_client_headers,json=allowedClientHeaders,proto3" json:"allowed_client_headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AuthorizationResponse) Reset()         { *m = AuthorizationResponse{} }
func (m *AuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizationResponse) ProtoMessage()    {}
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{3}
}

func (m *AuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationResponse.Unmarshal(m, b)
}
func (m *AuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationResponse.Merge(m, src)
}
func (m *AuthorizationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizationResponse.Size(m)
}
func (m *AuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationResponse proto.InternalMessageInfo

func (m *AuthorizationResponse) GetAllowedUpstreamHeaders() *matcher.ListStringMatcher {
	if m != nil {
		return m.AllowedUpstreamHeaders
	}
	return nil
}

func (m *AuthorizationResponse) GetAllowedClientHeaders() *matcher.ListStringMatcher {
	if m != nil {
		return m.AllowedClientHeaders
	}
	return nil
}

// Extra settings on a per virtualhost/route/weighted-cluster level.
type ExtAuthzPerRoute struct {
	// Types that are valid to be assigned to Override:
	//	*ExtAuthzPerRoute_Disabled
	//	*ExtAuthzPerRoute_CheckSettings
	Override             isExtAuthzPerRoute_Override `protobuf_oneof:"override"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *ExtAuthzPerRoute) Reset()         { *m = ExtAuthzPerRoute{} }
func (m *ExtAuthzPerRoute) String() string { return proto.CompactTextString(m) }
func (*ExtAuthzPerRoute) ProtoMessage()    {}
func (*ExtAuthzPerRoute) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{4}
}

func (m *ExtAuthzPerRoute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtAuthzPerRoute.Unmarshal(m, b)
}
func (m *ExtAuthzPerRoute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtAuthzPerRoute.Marshal(b, m, deterministic)
}
func (m *ExtAuthzPerRoute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtAuthzPerRoute.Merge(m, src)
}
func (m *ExtAuthzPerRoute) XXX_Size() int {
	return xxx_messageInfo_ExtAuthzPerRoute.Size(m)
}
func (m *ExtAuthzPerRoute) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtAuthzPerRoute.DiscardUnknown(m)
}

var xxx_messageInfo_ExtAuthzPerRoute proto.InternalMessageInfo

type isExtAuthzPerRoute_Override interface {
	isExtAuthzPerRoute_Override()
}

type ExtAuthzPerRoute_Disabled struct {
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3,oneof"`
}

type ExtAuthzPerRoute_CheckSettings struct {
	CheckSettings *CheckSettings `protobuf:"bytes,2,opt,name=check_settings,json=checkSettings,proto3,oneof"`
}

func (*ExtAuthzPerRoute_Disabled) isExtAuthzPerRoute_Override() {}

func (*ExtAuthzPerRoute_CheckSettings) isExtAuthzPerRoute_Override() {}

func (m *ExtAuthzPerRoute) GetOverride() isExtAuthzPerRoute_Override {
	if m != nil {
		return m.Override
	}
	return nil
}

func (m *ExtAuthzPerRoute) GetDisabled() bool {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_Disabled); ok {
		return x.Disabled
	}
	return false
}

func (m *ExtAuthzPerRoute) GetCheckSettings() *CheckSettings {
	if x, ok := m.GetOverride().(*ExtAuthzPerRoute_CheckSettings); ok {
		return x.CheckSettings
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExtAuthzPerRoute) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExtAuthzPerRoute_Disabled)(nil),
		(*ExtAuthzPerRoute_CheckSettings)(nil),
	}
}

// Extra settings for the check request. You can use this to provide extra context for the
// external authorization server on specific virtual hosts \ routes. For example, adding a context
// extension on the virtual host level can give the ext-authz server information on what virtual
// host is used without needing to parse the host header. If CheckSettings is specified in multiple
// per-filter-configs, they will be merged in order, and the result will be used.
type CheckSettings struct {
	// Context extensions to set on the CheckRequest's
	// :ref:`AttributeContext.context_extensions<envoy_api_field_service.auth.v2.AttributeContext.context_extensions>`
	//
	// Merge semantics for this field are such that keys from more specific configs override.
	//
	// .. note::
	//
	//   These settings are only applied to a filter configured with a
	//   :ref:`grpc_service<envoy_api_field_config.filter.http.ext_authz.v2.ExtAuthz.grpc_service>`.
	ContextExtensions    map[string]string `protobuf:"bytes,1,rep,name=context_extensions,json=contextExtensions,proto3" json:"context_extensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckSettings) Reset()         { *m = CheckSettings{} }
func (m *CheckSettings) String() string { return proto.CompactTextString(m) }
func (*CheckSettings) ProtoMessage()    {}
func (*CheckSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_861cd76675296973, []int{5}
}

func (m *CheckSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckSettings.Unmarshal(m, b)
}
func (m *CheckSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckSettings.Marshal(b, m, deterministic)
}
func (m *CheckSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckSettings.Merge(m, src)
}
func (m *CheckSettings) XXX_Size() int {
	return xxx_messageInfo_CheckSettings.Size(m)
}
func (m *CheckSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckSettings.DiscardUnknown(m)
}

var xxx_messageInfo_CheckSettings proto.InternalMessageInfo

func (m *CheckSettings) GetContextExtensions() map[string]string {
	if m != nil {
		return m.ContextExtensions
	}
	return nil
}

func init() {
	proto.RegisterType((*ExtAuthz)(nil), "envoy.config.filter.http.ext_authz.v2.ExtAuthz")
	proto.RegisterType((*HttpService)(nil), "envoy.config.filter.http.ext_authz.v2.HttpService")
	proto.RegisterType((*AuthorizationRequest)(nil), "envoy.config.filter.http.ext_authz.v2.AuthorizationRequest")
	proto.RegisterType((*AuthorizationResponse)(nil), "envoy.config.filter.http.ext_authz.v2.AuthorizationResponse")
	proto.RegisterType((*ExtAuthzPerRoute)(nil), "envoy.config.filter.http.ext_authz.v2.ExtAuthzPerRoute")
	proto.RegisterType((*CheckSettings)(nil), "envoy.config.filter.http.ext_authz.v2.CheckSettings")
	proto.RegisterMapType((map[string]string)(nil), "envoy.config.filter.http.ext_authz.v2.CheckSettings.ContextExtensionsEntry")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/ext_authz/v2/ext_authz.proto", fileDescriptor_861cd76675296973)
}

var fileDescriptor_861cd76675296973 = []byte{
	// 798 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xcd, 0x6e, 0xe4, 0x44,
	0x10, 0x5e, 0x7b, 0x66, 0xb2, 0x9e, 0xca, 0x0f, 0xb3, 0xad, 0x49, 0xb0, 0x22, 0xb4, 0x89, 0x46,
	0xac, 0xd8, 0x03, 0xb2, 0xa5, 0x59, 0x90, 0xf8, 0xbb, 0xcc, 0x64, 0x23, 0xa2, 0xc0, 0xae, 0x22,
	0x87, 0x80, 0x04, 0x07, 0xab, 0x63, 0x57, 0xc6, 0xcd, 0x3a, 0x6e, 0xd3, 0xdd, 0x36, 0x33, 0xb9,
	0x72, 0xe3, 0x59, 0x78, 0x00, 0xe0, 0xb4, 0xaf, 0xc0, 0x01, 0x89, 0x33, 0x27, 0xf6, 0x2d, 0x50,
	0xbb, 0xdb, 0x49, 0x76, 0x33, 0x48, 0x21, 0xb7, 0xee, 0xaa, 0xef, 0xfb, 0xea, 0xab, 0xea, 0xb2,
	0xe1, 0x43, 0x2c, 0x6a, 0xbe, 0x08, 0x13, 0x5e, 0x9c, 0xb1, 0x59, 0x78, 0xc6, 0x72, 0x85, 0x22,
	0xcc, 0x94, 0x2a, 0x43, 0x9c, 0xab, 0x98, 0x56, 0x2a, 0xbb, 0x08, 0xeb, 0xf1, 0xd5, 0x25, 0x28,
	0x05, 0x57, 0x9c, 0x3c, 0x6a, 0x68, 0x81, 0xa1, 0x05, 0x86, 0x16, 0x68, 0x5a, 0x70, 0x85, 0xac,
	0xc7, 0xdb, 0xef, 0x18, 0x75, 0x5a, 0x32, 0x2d, 0x92, 0x70, 0x81, 0xe1, 0x29, 0x95, 0x68, 0x44,
	0xb6, 0xdf, 0xbd, 0x99, 0x9d, 0x89, 0x32, 0x89, 0x25, 0x8a, 0x9a, 0x25, 0x2d, 0x6a, 0xf7, 0x26,
	0x4a, 0x17, 0x8a, 0x2b, 0xc1, 0x2c, 0x62, 0xc7, 0x20, 0xd4, 0xa2, 0xc4, 0xf0, 0x9c, 0xaa, 0x24,
	0x43, 0x11, 0x4a, 0x25, 0x58, 0x31, 0xb3, 0x80, 0xb7, 0x6b, 0x9a, 0xb3, 0x94, 0x2a, 0x0c, 0xdb,
	0x83, 0x4d, 0x0c, 0x67, 0x7c, 0xc6, 0x9b, 0x63, 0xa8, 0x4f, 0x26, 0x3a, 0xfa, 0xc9, 0x05, 0x6f,
	0x7f, 0xae, 0x26, 0xba, 0x0b, 0xb2, 0x07, 0x6b, 0xd7, 0x4d, 0xf9, 0xce, 0xae, 0xf3, 0x78, 0x75,
	0xfc, 0x30, 0x30, 0x03, 0xa0, 0x25, 0x0b, 0xea, 0x71, 0xa0, 0x5d, 0x05, 0x9f, 0x8b, 0x32, 0x39,
	0x36, 0xa8, 0x83, 0x7b, 0xd1, 0xea, 0xec, 0xea, 0x4a, 0xbe, 0x81, 0xb5, 0xc6, 0x73, 0x2b, 0xd2,
	0x69, 0x44, 0xc6, 0xc1, 0xad, 0xa6, 0x18, 0x1c, 0x28, 0x55, 0x5e, 0x13, 0xce, 0xae, 0xae, 0xe4,
	0x7d, 0x20, 0x67, 0x94, 0xe5, 0x95, 0xc0, 0xf8, 0x9c, 0xa7, 0x18, 0xd3, 0x3c, 0xe7, 0x3f, 0xfa,
	0xee, 0xae, 0xf3, 0xd8, 0x8b, 0x06, 0x36, 0xf3, 0x8c, 0xa7, 0x38, 0xd1, 0x71, 0xb2, 0x03, 0xfd,
	0x4a, 0x6a, 0x50, 0x99, 0x51, 0xbf, 0xab, 0x41, 0x53, 0xd7, 0x77, 0x22, 0xaf, 0x92, 0x38, 0xd1,
	0xb1, 0x29, 0x80, 0x67, 0x2d, 0xca, 0xd1, 0x3f, 0x2e, 0xac, 0x5e, 0xab, 0x4c, 0x3e, 0x06, 0xd0,
	0x39, 0x14, 0x7a, 0xf2, 0x76, 0x0c, 0xdb, 0x4b, 0xc6, 0xa0, 0x39, 0x27, 0x82, 0x45, 0x7d, 0x83,
	0x3e, 0x11, 0x8c, 0xec, 0xc0, 0x6a, 0x49, 0x55, 0x16, 0x97, 0x02, 0xcf, 0xd8, 0xbc, 0xb1, 0xd7,
	0x8f, 0x40, 0x87, 0x8e, 0x9a, 0x08, 0x29, 0x61, 0x53, 0x77, 0xcb, 0x05, 0xbb, 0xa0, 0x8a, 0xf1,
	0x22, 0x16, 0xf8, 0x43, 0x85, 0x52, 0xf9, 0xf7, 0x9b, 0x32, 0x9f, 0xde, 0x72, 0x50, 0x93, 0xeb,
	0x1a, 0x91, 0x91, 0x88, 0x86, 0x74, 0x49, 0x94, 0x48, 0xd8, 0x7a, 0xb3, 0xa2, 0x2c, 0x79, 0x21,
	0xd1, 0xf7, 0x9a, 0x92, 0x9f, 0xdd, 0xad, 0xa4, 0xd1, 0x88, 0x36, 0xe9, 0xb2, 0xf0, 0x61, 0xd7,
	0xeb, 0x0c, 0xba, 0x87, 0x5d, 0xaf, 0x3b, 0xe8, 0x1d, 0x76, 0xbd, 0xde, 0x60, 0xe5, 0xb0, 0xeb,
	0xad, 0x0c, 0xee, 0x8f, 0x7e, 0x71, 0x60, 0xb8, 0xcc, 0x3b, 0x79, 0x0e, 0x6f, 0x35, 0x2f, 0x8a,
	0x69, 0x9c, 0x21, 0x4d, 0x51, 0x48, 0x3b, 0xf8, 0x47, 0xd6, 0x9e, 0xde, 0xf9, 0xc0, 0xee, 0x7c,
	0xf0, 0x25, 0x93, 0xea, 0xb8, 0xd9, 0xfb, 0x67, 0x26, 0x12, 0x6d, 0x58, 0xf6, 0x81, 0x21, 0x93,
	0xa7, 0xb0, 0x61, 0x75, 0x62, 0xc5, 0x63, 0x9a, 0xa6, 0xbe, 0xbb, 0xdb, 0xf9, 0x8f, 0x75, 0x36,
	0x9c, 0xaf, 0x69, 0x5e, 0x61, 0xb4, 0x66, 0x59, 0x5f, 0xf1, 0x49, 0x9a, 0x8e, 0xfe, 0x74, 0x60,
	0x73, 0x69, 0xdf, 0x24, 0x06, 0xbf, 0xf5, 0x5b, 0x95, 0x52, 0x09, 0xa4, 0xe7, 0x77, 0x33, 0xbe,
	0x65, 0x65, 0x4e, 0xac, 0x4a, 0xdb, 0xc0, 0x77, 0xd0, 0x66, 0xe2, 0x24, 0x67, 0x58, 0xa8, 0x4b,
	0x79, 0xf7, 0xff, 0xc8, 0x0f, 0xad, 0xc8, 0x5e, 0xa3, 0x61, 0xc5, 0x47, 0xbf, 0x39, 0x30, 0x68,
	0xbf, 0xfb, 0x23, 0x14, 0x11, 0xaf, 0x14, 0x92, 0xf7, 0xc0, 0x4b, 0x99, 0xa4, 0xa7, 0x39, 0xa6,
	0x4d, 0x0b, 0xde, 0xb4, 0xff, 0xfb, 0xab, 0x97, 0x9d, 0xee, 0xf7, 0xae, 0xe7, 0x1c, 0xdc, 0x8b,
	0x2e, 0x93, 0x84, 0xc1, 0x46, 0x92, 0x61, 0xf2, 0x22, 0x96, 0xa8, 0x14, 0x2b, 0x66, 0xad, 0xa5,
	0x0f, 0x6e, 0xb9, 0x49, 0x7b, 0x9a, 0x7c, 0x6c, 0xb9, 0x53, 0xd0, 0x45, 0x7a, 0x3f, 0x3b, 0xee,
	0x40, 0x57, 0x59, 0x4f, 0x5e, 0x4b, 0x3e, 0x00, 0x8f, 0xd7, 0x28, 0x04, 0x4b, 0x91, 0xf4, 0x7e,
	0x7d, 0xf5, 0xb2, 0xe3, 0x8c, 0xfe, 0x70, 0x60, 0xfd, 0x35, 0x05, 0x72, 0x01, 0x24, 0xe1, 0x85,
	0xd2, 0x35, 0x70, 0xae, 0xb0, 0x90, 0x8c, 0x17, 0xfa, 0x15, 0xf4, 0x7b, 0x7f, 0x71, 0x17, 0x4f,
	0xc1, 0x9e, 0x91, 0xdb, 0xbf, 0x54, 0xdb, 0x2f, 0x94, 0x58, 0x44, 0x0f, 0x92, 0x37, 0xe3, 0xdb,
	0x4f, 0x61, 0x6b, 0x39, 0x98, 0x0c, 0xa0, 0xf3, 0x02, 0x17, 0xcd, 0x24, 0xfb, 0x91, 0x3e, 0x92,
	0x21, 0xf4, 0x6a, 0xbd, 0x64, 0xf6, 0xb7, 0x60, 0x2e, 0x9f, 0xb8, 0x1f, 0x39, 0xd3, 0xe7, 0x7f,
	0xfd, 0xfd, 0xd0, 0x81, 0x27, 0x8c, 0x1b, 0xb7, 0xa5, 0xe0, 0xf3, 0xc5, 0xed, 0x8c, 0x4f, 0xd7,
	0x2f, 0xdf, 0x51, 0xff, 0xd1, 0x8f, 0x9c, 0x6f, 0xdd, 0x7a, 0x7c, 0xba, 0xd2, 0xfc, 0xde, 0x9f,
	0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x27, 0x4d, 0x0d, 0xf4, 0x06, 0x00, 0x00,
}