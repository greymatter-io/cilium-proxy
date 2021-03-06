// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/cluster/v2alpha/outlier_detection_event.proto

package envoy_data_cluster_v2alpha

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Type of ejection that took place
type OutlierEjectionType int32

const (
	// In case upstream host returns certain number of consecutive 5xx.
	// If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is *false*, all type of errors are treated as HTTP 5xx errors.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	// details.
	OutlierEjectionType_CONSECUTIVE_5XX OutlierEjectionType = 0
	// In case upstream host returns certain number of consecutive gateway errors
	OutlierEjectionType_CONSECUTIVE_GATEWAY_FAILURE OutlierEjectionType = 1
	// Runs over aggregated success rate statistics from every host in cluster
	// and selects hosts for which ratio of successful replies deviates from other hosts
	// in the cluster.
	// If
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is *false*, all errors (externally and locally generated) are used to calculate success rate
	// statistics. See :ref:`Cluster outlier detection <arch_overview_outlier_detection>`
	// documentation for details.
	OutlierEjectionType_SUCCESS_RATE OutlierEjectionType = 2
	// Consecutive local origin failures: Connection failures, resets, timeouts, etc
	// This type of ejection happens only when
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to *true*.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	OutlierEjectionType_CONSECUTIVE_LOCAL_ORIGIN_FAILURE OutlierEjectionType = 3
	// Runs over aggregated success rate statistics for local origin failures
	// for all hosts in the cluster and selects hosts for which success rate deviates from other
	// hosts in the cluster. This type of ejection happens only when
	// :ref:`outlier_detection.split_external_local_origin_errors<envoy_api_field_cluster.OutlierDetection.split_external_local_origin_errors>`
	// is set to *true*.
	// See :ref:`Cluster outlier detection <arch_overview_outlier_detection>` documentation for
	OutlierEjectionType_SUCCESS_RATE_LOCAL_ORIGIN OutlierEjectionType = 4
	// Runs over aggregated success rate statistics from every host in cluster and selects hosts for
	// which ratio of failed replies is above configured value.
	OutlierEjectionType_FAILURE_PERCENTAGE OutlierEjectionType = 5
	// Runs over aggregated success rate statistics for local origin failures from every host in
	// cluster and selects hosts for which ratio of failed replies is above configured value.
	OutlierEjectionType_FAILURE_PERCENTAGE_LOCAL_ORIGIN OutlierEjectionType = 6
)

var OutlierEjectionType_name = map[int32]string{
	0: "CONSECUTIVE_5XX",
	1: "CONSECUTIVE_GATEWAY_FAILURE",
	2: "SUCCESS_RATE",
	3: "CONSECUTIVE_LOCAL_ORIGIN_FAILURE",
	4: "SUCCESS_RATE_LOCAL_ORIGIN",
	5: "FAILURE_PERCENTAGE",
	6: "FAILURE_PERCENTAGE_LOCAL_ORIGIN",
}

var OutlierEjectionType_value = map[string]int32{
	"CONSECUTIVE_5XX":                  0,
	"CONSECUTIVE_GATEWAY_FAILURE":      1,
	"SUCCESS_RATE":                     2,
	"CONSECUTIVE_LOCAL_ORIGIN_FAILURE": 3,
	"SUCCESS_RATE_LOCAL_ORIGIN":        4,
	"FAILURE_PERCENTAGE":               5,
	"FAILURE_PERCENTAGE_LOCAL_ORIGIN":  6,
}

func (x OutlierEjectionType) String() string {
	return proto.EnumName(OutlierEjectionType_name, int32(x))
}

func (OutlierEjectionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{0}
}

// Represents possible action applied to upstream host
type Action int32

const (
	// In case host was excluded from service
	Action_EJECT Action = 0
	// In case host was brought back into service
	Action_UNEJECT Action = 1
)

var Action_name = map[int32]string{
	0: "EJECT",
	1: "UNEJECT",
}

var Action_value = map[string]int32{
	"EJECT":   0,
	"UNEJECT": 1,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{1}
}

// [#next-free-field: 12]
type OutlierDetectionEvent struct {
	// In case of eject represents type of ejection that took place.
	Type OutlierEjectionType `protobuf:"varint,1,opt,name=type,proto3,enum=envoy.data.cluster.v2alpha.OutlierEjectionType" json:"type,omitempty"`
	// Timestamp for event.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The time in seconds since the last action (either an ejection or unejection) took place.
	SecsSinceLastAction *wrappers.UInt64Value `protobuf:"bytes,3,opt,name=secs_since_last_action,json=secsSinceLastAction,proto3" json:"secs_since_last_action,omitempty"`
	// The :ref:`cluster <envoy_api_msg_Cluster>` that owns the ejected host.
	ClusterName string `protobuf:"bytes,4,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	// The URL of the ejected host. E.g., ``tcp://1.2.3.4:80``.
	UpstreamUrl string `protobuf:"bytes,5,opt,name=upstream_url,json=upstreamUrl,proto3" json:"upstream_url,omitempty"`
	// The action that took place.
	Action Action `protobuf:"varint,6,opt,name=action,proto3,enum=envoy.data.cluster.v2alpha.Action" json:"action,omitempty"`
	// If ``action`` is ``eject``, specifies the number of times the host has been ejected (local to
	// that Envoy and gets reset if the host gets removed from the upstream cluster for any reason and
	// then re-added).
	NumEjections uint32 `protobuf:"varint,7,opt,name=num_ejections,json=numEjections,proto3" json:"num_ejections,omitempty"`
	// If ``action`` is ``eject``, specifies if the ejection was enforced. ``true`` means the host was
	// ejected. ``false`` means the event was logged but the host was not actually ejected.
	Enforced bool `protobuf:"varint,8,opt,name=enforced,proto3" json:"enforced,omitempty"`
	// Types that are valid to be assigned to Event:
	//	*OutlierDetectionEvent_EjectSuccessRateEvent
	//	*OutlierDetectionEvent_EjectConsecutiveEvent
	//	*OutlierDetectionEvent_EjectFailurePercentageEvent
	Event                isOutlierDetectionEvent_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *OutlierDetectionEvent) Reset()         { *m = OutlierDetectionEvent{} }
func (m *OutlierDetectionEvent) String() string { return proto.CompactTextString(m) }
func (*OutlierDetectionEvent) ProtoMessage()    {}
func (*OutlierDetectionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{0}
}

func (m *OutlierDetectionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierDetectionEvent.Unmarshal(m, b)
}
func (m *OutlierDetectionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierDetectionEvent.Marshal(b, m, deterministic)
}
func (m *OutlierDetectionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierDetectionEvent.Merge(m, src)
}
func (m *OutlierDetectionEvent) XXX_Size() int {
	return xxx_messageInfo_OutlierDetectionEvent.Size(m)
}
func (m *OutlierDetectionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierDetectionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierDetectionEvent proto.InternalMessageInfo

func (m *OutlierDetectionEvent) GetType() OutlierEjectionType {
	if m != nil {
		return m.Type
	}
	return OutlierEjectionType_CONSECUTIVE_5XX
}

func (m *OutlierDetectionEvent) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *OutlierDetectionEvent) GetSecsSinceLastAction() *wrappers.UInt64Value {
	if m != nil {
		return m.SecsSinceLastAction
	}
	return nil
}

func (m *OutlierDetectionEvent) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

func (m *OutlierDetectionEvent) GetUpstreamUrl() string {
	if m != nil {
		return m.UpstreamUrl
	}
	return ""
}

func (m *OutlierDetectionEvent) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_EJECT
}

func (m *OutlierDetectionEvent) GetNumEjections() uint32 {
	if m != nil {
		return m.NumEjections
	}
	return 0
}

func (m *OutlierDetectionEvent) GetEnforced() bool {
	if m != nil {
		return m.Enforced
	}
	return false
}

type isOutlierDetectionEvent_Event interface {
	isOutlierDetectionEvent_Event()
}

type OutlierDetectionEvent_EjectSuccessRateEvent struct {
	EjectSuccessRateEvent *OutlierEjectSuccessRate `protobuf:"bytes,9,opt,name=eject_success_rate_event,json=ejectSuccessRateEvent,proto3,oneof"`
}

type OutlierDetectionEvent_EjectConsecutiveEvent struct {
	EjectConsecutiveEvent *OutlierEjectConsecutive `protobuf:"bytes,10,opt,name=eject_consecutive_event,json=ejectConsecutiveEvent,proto3,oneof"`
}

type OutlierDetectionEvent_EjectFailurePercentageEvent struct {
	EjectFailurePercentageEvent *OutlierEjectFailurePercentage `protobuf:"bytes,11,opt,name=eject_failure_percentage_event,json=ejectFailurePercentageEvent,proto3,oneof"`
}

func (*OutlierDetectionEvent_EjectSuccessRateEvent) isOutlierDetectionEvent_Event() {}

func (*OutlierDetectionEvent_EjectConsecutiveEvent) isOutlierDetectionEvent_Event() {}

func (*OutlierDetectionEvent_EjectFailurePercentageEvent) isOutlierDetectionEvent_Event() {}

func (m *OutlierDetectionEvent) GetEvent() isOutlierDetectionEvent_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *OutlierDetectionEvent) GetEjectSuccessRateEvent() *OutlierEjectSuccessRate {
	if x, ok := m.GetEvent().(*OutlierDetectionEvent_EjectSuccessRateEvent); ok {
		return x.EjectSuccessRateEvent
	}
	return nil
}

func (m *OutlierDetectionEvent) GetEjectConsecutiveEvent() *OutlierEjectConsecutive {
	if x, ok := m.GetEvent().(*OutlierDetectionEvent_EjectConsecutiveEvent); ok {
		return x.EjectConsecutiveEvent
	}
	return nil
}

func (m *OutlierDetectionEvent) GetEjectFailurePercentageEvent() *OutlierEjectFailurePercentage {
	if x, ok := m.GetEvent().(*OutlierDetectionEvent_EjectFailurePercentageEvent); ok {
		return x.EjectFailurePercentageEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*OutlierDetectionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*OutlierDetectionEvent_EjectSuccessRateEvent)(nil),
		(*OutlierDetectionEvent_EjectConsecutiveEvent)(nil),
		(*OutlierDetectionEvent_EjectFailurePercentageEvent)(nil),
	}
}

type OutlierEjectSuccessRate struct {
	// Host’s success rate at the time of the ejection event on a 0-100 range.
	HostSuccessRate uint32 `protobuf:"varint,1,opt,name=host_success_rate,json=hostSuccessRate,proto3" json:"host_success_rate,omitempty"`
	// Average success rate of the hosts in the cluster at the time of the ejection event on a 0-100
	// range.
	ClusterAverageSuccessRate uint32 `protobuf:"varint,2,opt,name=cluster_average_success_rate,json=clusterAverageSuccessRate,proto3" json:"cluster_average_success_rate,omitempty"`
	// Success rate ejection threshold at the time of the ejection event.
	ClusterSuccessRateEjectionThreshold uint32   `protobuf:"varint,3,opt,name=cluster_success_rate_ejection_threshold,json=clusterSuccessRateEjectionThreshold,proto3" json:"cluster_success_rate_ejection_threshold,omitempty"`
	XXX_NoUnkeyedLiteral                struct{} `json:"-"`
	XXX_unrecognized                    []byte   `json:"-"`
	XXX_sizecache                       int32    `json:"-"`
}

func (m *OutlierEjectSuccessRate) Reset()         { *m = OutlierEjectSuccessRate{} }
func (m *OutlierEjectSuccessRate) String() string { return proto.CompactTextString(m) }
func (*OutlierEjectSuccessRate) ProtoMessage()    {}
func (*OutlierEjectSuccessRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{1}
}

func (m *OutlierEjectSuccessRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierEjectSuccessRate.Unmarshal(m, b)
}
func (m *OutlierEjectSuccessRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierEjectSuccessRate.Marshal(b, m, deterministic)
}
func (m *OutlierEjectSuccessRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierEjectSuccessRate.Merge(m, src)
}
func (m *OutlierEjectSuccessRate) XXX_Size() int {
	return xxx_messageInfo_OutlierEjectSuccessRate.Size(m)
}
func (m *OutlierEjectSuccessRate) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierEjectSuccessRate.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierEjectSuccessRate proto.InternalMessageInfo

func (m *OutlierEjectSuccessRate) GetHostSuccessRate() uint32 {
	if m != nil {
		return m.HostSuccessRate
	}
	return 0
}

func (m *OutlierEjectSuccessRate) GetClusterAverageSuccessRate() uint32 {
	if m != nil {
		return m.ClusterAverageSuccessRate
	}
	return 0
}

func (m *OutlierEjectSuccessRate) GetClusterSuccessRateEjectionThreshold() uint32 {
	if m != nil {
		return m.ClusterSuccessRateEjectionThreshold
	}
	return 0
}

type OutlierEjectConsecutive struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutlierEjectConsecutive) Reset()         { *m = OutlierEjectConsecutive{} }
func (m *OutlierEjectConsecutive) String() string { return proto.CompactTextString(m) }
func (*OutlierEjectConsecutive) ProtoMessage()    {}
func (*OutlierEjectConsecutive) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{2}
}

func (m *OutlierEjectConsecutive) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierEjectConsecutive.Unmarshal(m, b)
}
func (m *OutlierEjectConsecutive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierEjectConsecutive.Marshal(b, m, deterministic)
}
func (m *OutlierEjectConsecutive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierEjectConsecutive.Merge(m, src)
}
func (m *OutlierEjectConsecutive) XXX_Size() int {
	return xxx_messageInfo_OutlierEjectConsecutive.Size(m)
}
func (m *OutlierEjectConsecutive) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierEjectConsecutive.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierEjectConsecutive proto.InternalMessageInfo

type OutlierEjectFailurePercentage struct {
	// Host's success rate at the time of the ejection event on a 0-100 range.
	HostSuccessRate      uint32   `protobuf:"varint,1,opt,name=host_success_rate,json=hostSuccessRate,proto3" json:"host_success_rate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutlierEjectFailurePercentage) Reset()         { *m = OutlierEjectFailurePercentage{} }
func (m *OutlierEjectFailurePercentage) String() string { return proto.CompactTextString(m) }
func (*OutlierEjectFailurePercentage) ProtoMessage()    {}
func (*OutlierEjectFailurePercentage) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e03c92c55863094, []int{3}
}

func (m *OutlierEjectFailurePercentage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutlierEjectFailurePercentage.Unmarshal(m, b)
}
func (m *OutlierEjectFailurePercentage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutlierEjectFailurePercentage.Marshal(b, m, deterministic)
}
func (m *OutlierEjectFailurePercentage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutlierEjectFailurePercentage.Merge(m, src)
}
func (m *OutlierEjectFailurePercentage) XXX_Size() int {
	return xxx_messageInfo_OutlierEjectFailurePercentage.Size(m)
}
func (m *OutlierEjectFailurePercentage) XXX_DiscardUnknown() {
	xxx_messageInfo_OutlierEjectFailurePercentage.DiscardUnknown(m)
}

var xxx_messageInfo_OutlierEjectFailurePercentage proto.InternalMessageInfo

func (m *OutlierEjectFailurePercentage) GetHostSuccessRate() uint32 {
	if m != nil {
		return m.HostSuccessRate
	}
	return 0
}

func init() {
	proto.RegisterEnum("envoy.data.cluster.v2alpha.OutlierEjectionType", OutlierEjectionType_name, OutlierEjectionType_value)
	proto.RegisterEnum("envoy.data.cluster.v2alpha.Action", Action_name, Action_value)
	proto.RegisterType((*OutlierDetectionEvent)(nil), "envoy.data.cluster.v2alpha.OutlierDetectionEvent")
	proto.RegisterType((*OutlierEjectSuccessRate)(nil), "envoy.data.cluster.v2alpha.OutlierEjectSuccessRate")
	proto.RegisterType((*OutlierEjectConsecutive)(nil), "envoy.data.cluster.v2alpha.OutlierEjectConsecutive")
	proto.RegisterType((*OutlierEjectFailurePercentage)(nil), "envoy.data.cluster.v2alpha.OutlierEjectFailurePercentage")
}

func init() {
	proto.RegisterFile("envoy/data/cluster/v2alpha/outlier_detection_event.proto", fileDescriptor_5e03c92c55863094)
}

var fileDescriptor_5e03c92c55863094 = []byte{
	// 835 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x95, 0xcf, 0x6f, 0x1b, 0x45,
	0x14, 0xc7, 0x33, 0x4e, 0xec, 0x24, 0x93, 0x84, 0x2e, 0x13, 0xa5, 0xd9, 0xb8, 0x4d, 0x63, 0x39,
	0x48, 0x58, 0x39, 0xec, 0x4a, 0x09, 0xa0, 0x72, 0xb4, 0xdd, 0x6d, 0x62, 0x14, 0x9c, 0xb0, 0x5e,
	0x97, 0x72, 0x40, 0xa3, 0xe9, 0xfa, 0x25, 0x59, 0xb4, 0xbf, 0x34, 0x33, 0x6b, 0xc8, 0xad, 0xe2,
	0x1f, 0xe0, 0xca, 0xdf, 0xc2, 0x91, 0x13, 0x57, 0xfe, 0x02, 0xfe, 0x07, 0x2e, 0xa0, 0x1e, 0x10,
	0xda, 0x99, 0x5d, 0x77, 0x9d, 0xa4, 0x51, 0xe1, 0xe6, 0x9d, 0xf7, 0x3e, 0xdf, 0xef, 0xbc, 0x37,
	0x6f, 0xc6, 0xf8, 0x29, 0xc4, 0xd3, 0xe4, 0xda, 0x9e, 0x30, 0xc9, 0x6c, 0x3f, 0xcc, 0x84, 0x04,
	0x6e, 0x4f, 0x0f, 0x59, 0x98, 0x5e, 0x31, 0x3b, 0xc9, 0x64, 0x18, 0x00, 0xa7, 0x13, 0x90, 0xe0,
	0xcb, 0x20, 0x89, 0x29, 0x4c, 0x21, 0x96, 0x56, 0xca, 0x13, 0x99, 0x90, 0xa6, 0x22, 0xad, 0x9c,
	0xb4, 0x0a, 0xd2, 0x2a, 0xc8, 0xe6, 0xde, 0x65, 0x92, 0x5c, 0x86, 0x60, 0xab, 0xcc, 0x57, 0xd9,
	0x85, 0x2d, 0x83, 0x08, 0x84, 0x64, 0x51, 0xaa, 0xe1, 0xe6, 0x93, 0x9b, 0x09, 0xdf, 0x73, 0x96,
	0xa6, 0xc0, 0x45, 0x19, 0xcf, 0x26, 0x29, 0xb3, 0x59, 0x1c, 0x27, 0x92, 0xe5, 0xce, 0xc2, 0x8e,
	0x82, 0x4b, 0xce, 0x24, 0x14, 0xf1, 0xdd, 0x5b, 0x71, 0x21, 0x99, 0xcc, 0x4a, 0x7c, 0x7b, 0xca,
	0xc2, 0x60, 0xc2, 0x24, 0xd8, 0xe5, 0x0f, 0x1d, 0x68, 0xff, 0xda, 0xc0, 0x5b, 0x67, 0xba, 0xac,
	0x67, 0x65, 0x55, 0x4e, 0x5e, 0x14, 0xf9, 0x12, 0x2f, 0xc9, 0xeb, 0x14, 0x4c, 0xd4, 0x42, 0x9d,
	0x0f, 0x0e, 0x6d, 0xeb, 0xdd, 0xd5, 0x59, 0x85, 0x80, 0xf3, 0x9d, 0xe6, 0xbd, 0xeb, 0x14, 0x7a,
	0x2b, 0x6f, 0x7a, 0xf5, 0x1f, 0x51, 0xcd, 0x40, 0xae, 0x92, 0x21, 0x4f, 0xf1, 0xea, 0xac, 0x66,
	0xb3, 0xd6, 0x42, 0x9d, 0xb5, 0xc3, 0xa6, 0xa5, 0x8b, 0xb6, 0xca, 0xa2, 0x2d, 0xaf, 0xcc, 0x70,
	0xdf, 0x26, 0x93, 0xaf, 0xf0, 0x43, 0x01, 0xbe, 0xa0, 0x22, 0x88, 0x7d, 0xa0, 0x21, 0x13, 0x92,
	0x32, 0xe5, 0x63, 0x2e, 0x2a, 0x99, 0xc7, 0xb7, 0x64, 0xc6, 0x83, 0x58, 0x7e, 0xf6, 0xc9, 0x0b,
	0x16, 0x66, 0xe0, 0x6e, 0xe6, 0xec, 0x28, 0x47, 0x4f, 0x99, 0x90, 0x5d, 0x05, 0x92, 0x03, 0xbc,
	0x5e, 0xd4, 0x40, 0x63, 0x16, 0x81, 0xb9, 0xd4, 0x42, 0x9d, 0xd5, 0xde, 0xf2, 0x9b, 0xde, 0x12,
	0xaf, 0xb5, 0x90, 0xbb, 0x56, 0x04, 0x87, 0x2c, 0x82, 0x3c, 0x37, 0x4b, 0x85, 0xe4, 0xc0, 0x22,
	0x9a, 0xf1, 0xd0, 0xac, 0xdf, 0xc8, 0x2d, 0x83, 0x63, 0x1e, 0x92, 0x67, 0xb8, 0x51, 0x6c, 0xad,
	0xa1, 0xba, 0xd6, 0xbe, 0xaf, 0x6b, 0x7a, 0x2f, 0x95, 0x46, 0x15, 0x2c, 0xd9, 0xc7, 0x1b, 0x71,
	0x16, 0x51, 0x28, 0xda, 0x29, 0xcc, 0xe5, 0x16, 0xea, 0x6c, 0xb8, 0xeb, 0x71, 0x16, 0x95, 0x2d,
	0x16, 0xa4, 0x89, 0x57, 0x20, 0xbe, 0x48, 0xb8, 0x0f, 0x13, 0x73, 0xa5, 0x85, 0x3a, 0x2b, 0xee,
	0xec, 0x9b, 0xc4, 0xd8, 0x54, 0x30, 0x15, 0x99, 0xef, 0x83, 0x10, 0x34, 0x1f, 0x14, 0x3d, 0xab,
	0xe6, 0xaa, 0xea, 0xd9, 0xd1, 0xfb, 0x1e, 0xe7, 0x48, 0x2b, 0xb8, 0x4c, 0xc2, 0xc9, 0x82, 0xbb,
	0x05, 0x37, 0xd6, 0xf4, 0xa8, 0x44, 0x78, 0x5b, 0xfb, 0xf9, 0x49, 0x2c, 0xc0, 0xcf, 0x64, 0x30,
	0x2d, 0xed, 0xf0, 0x7f, 0xb3, 0xeb, 0xbf, 0x15, 0x98, 0xd9, 0x55, 0xd6, 0xb4, 0xdd, 0x6b, 0x84,
	0x9f, 0x68, 0xbf, 0x0b, 0x16, 0x84, 0x19, 0x07, 0x9a, 0x02, 0xf7, 0x21, 0x96, 0xec, 0xb2, 0xb4,
	0x5d, 0x53, 0xb6, 0x9f, 0xbf, 0xaf, 0xed, 0x73, 0xad, 0x73, 0x3e, 0x93, 0x39, 0x59, 0x70, 0x1f,
	0xc1, 0x9d, 0x11, 0xb5, 0x85, 0xde, 0x3a, 0xae, 0x2b, 0x23, 0xb2, 0xf8, 0x77, 0x0f, 0xb5, 0xff,
	0x42, 0x78, 0xfb, 0x1d, 0x4d, 0x23, 0x47, 0xf8, 0xc3, 0xab, 0x44, 0xcc, 0x1f, 0x85, 0xba, 0x53,
	0x1b, 0x6a, 0x86, 0x0e, 0x6a, 0xe6, 0xc4, 0x7d, 0x90, 0x67, 0x54, 0xa1, 0x13, 0xfc, 0xb8, 0x9c,
	0x4f, 0x36, 0x05, 0x9e, 0xd7, 0x35, 0xc7, 0xd7, 0xe6, 0xf9, 0x9d, 0x22, 0xb9, 0xab, 0x73, 0xab,
	0x4a, 0xdf, 0xe2, 0x8f, 0x4b, 0xa5, 0xf9, 0x61, 0x28, 0x06, 0x89, 0xca, 0x2b, 0x0e, 0xe2, 0x2a,
	0x09, 0x27, 0xea, 0x36, 0x55, 0x44, 0xf7, 0x0b, 0xae, 0x7a, 0xda, 0xe5, 0x05, 0x2f, 0x99, 0xf6,
	0xce, 0x7c, 0xe1, 0x95, 0xa3, 0x6a, 0x7b, 0x78, 0xf7, 0xde, 0x16, 0xff, 0xaf, 0xce, 0x1c, 0xfc,
	0x81, 0xf0, 0xe6, 0x1d, 0xcf, 0x0d, 0xd9, 0xc4, 0x0f, 0xfa, 0x67, 0xc3, 0x91, 0xd3, 0x1f, 0x7b,
	0x83, 0x17, 0x0e, 0xfd, 0xf4, 0xe5, 0x4b, 0x63, 0x81, 0xec, 0xe1, 0x47, 0xd5, 0xc5, 0xe3, 0xae,
	0xe7, 0x7c, 0xdd, 0xfd, 0x86, 0x3e, 0xef, 0x0e, 0x4e, 0xc7, 0xae, 0x63, 0x20, 0x62, 0xe0, 0xf5,
	0xd1, 0xb8, 0xdf, 0x77, 0x46, 0x23, 0xea, 0x76, 0x3d, 0xc7, 0xa8, 0x91, 0x8f, 0x70, 0xab, 0x8a,
	0x9c, 0x9e, 0xf5, 0xbb, 0xa7, 0xf4, 0xcc, 0x1d, 0x1c, 0x0f, 0x86, 0x33, 0x6e, 0x91, 0xec, 0xe2,
	0x9d, 0x2a, 0x37, 0x97, 0x66, 0x2c, 0x91, 0x87, 0x98, 0x14, 0xb9, 0xf4, 0xdc, 0x71, 0xfb, 0xce,
	0xd0, 0xeb, 0x1e, 0x3b, 0x46, 0x9d, 0xec, 0xe3, 0xbd, 0xdb, 0xeb, 0xf3, 0x70, 0xe3, 0xa0, 0x85,
	0x1b, 0xc5, 0x2b, 0xb5, 0x8a, 0xeb, 0xce, 0x17, 0x4e, 0xdf, 0x33, 0x16, 0xc8, 0x1a, 0x5e, 0x1e,
	0x0f, 0xf5, 0x07, 0xea, 0x45, 0x7f, 0xfe, 0xfc, 0xcf, 0x4f, 0xf5, 0x6d, 0xb2, 0x75, 0xd7, 0x74,
	0x1f, 0xfd, 0xf2, 0xfa, 0xb7, 0xdf, 0x1b, 0x35, 0x03, 0xe1, 0x4e, 0x90, 0xe8, 0xf9, 0x4f, 0x79,
	0xf2, 0xc3, 0xf5, 0x3d, 0x57, 0xa1, 0xd7, 0xbc, 0xf3, 0x1f, 0xe0, 0x3c, 0x7f, 0x4e, 0xcf, 0xd1,
	0xab, 0x86, 0x7a, 0x57, 0x8f, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x88, 0xaf, 0x69, 0x81, 0x1a,
	0x07, 0x00, 0x00,
}
