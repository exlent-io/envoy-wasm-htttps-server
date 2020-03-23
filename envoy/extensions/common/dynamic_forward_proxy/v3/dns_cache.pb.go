// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/common/dynamic_forward_proxy/v3/dns_cache.proto

package envoy_extensions_common_dynamic_forward_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type DnsCacheConfig struct {
	Name                  string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DnsLookupFamily       v3.Cluster_DnsLookupFamily `protobuf:"varint,2,opt,name=dns_lookup_family,json=dnsLookupFamily,proto3,enum=envoy.config.cluster.v3.Cluster_DnsLookupFamily" json:"dns_lookup_family,omitempty"`
	DnsRefreshRate        *duration.Duration         `protobuf:"bytes,3,opt,name=dns_refresh_rate,json=dnsRefreshRate,proto3" json:"dns_refresh_rate,omitempty"`
	HostTtl               *duration.Duration         `protobuf:"bytes,4,opt,name=host_ttl,json=hostTtl,proto3" json:"host_ttl,omitempty"`
	MaxHosts              *wrappers.UInt32Value      `protobuf:"bytes,5,opt,name=max_hosts,json=maxHosts,proto3" json:"max_hosts,omitempty"`
	DnsFailureRefreshRate *v3.Cluster_RefreshRate    `protobuf:"bytes,6,opt,name=dns_failure_refresh_rate,json=dnsFailureRefreshRate,proto3" json:"dns_failure_refresh_rate,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                   `json:"-"`
	XXX_unrecognized      []byte                     `json:"-"`
	XXX_sizecache         int32                      `json:"-"`
}

func (m *DnsCacheConfig) Reset()         { *m = DnsCacheConfig{} }
func (m *DnsCacheConfig) String() string { return proto.CompactTextString(m) }
func (*DnsCacheConfig) ProtoMessage()    {}
func (*DnsCacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f57893b6dd868364, []int{0}
}

func (m *DnsCacheConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsCacheConfig.Unmarshal(m, b)
}
func (m *DnsCacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsCacheConfig.Marshal(b, m, deterministic)
}
func (m *DnsCacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsCacheConfig.Merge(m, src)
}
func (m *DnsCacheConfig) XXX_Size() int {
	return xxx_messageInfo_DnsCacheConfig.Size(m)
}
func (m *DnsCacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsCacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsCacheConfig proto.InternalMessageInfo

func (m *DnsCacheConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsCacheConfig) GetDnsLookupFamily() v3.Cluster_DnsLookupFamily {
	if m != nil {
		return m.DnsLookupFamily
	}
	return v3.Cluster_AUTO
}

func (m *DnsCacheConfig) GetDnsRefreshRate() *duration.Duration {
	if m != nil {
		return m.DnsRefreshRate
	}
	return nil
}

func (m *DnsCacheConfig) GetHostTtl() *duration.Duration {
	if m != nil {
		return m.HostTtl
	}
	return nil
}

func (m *DnsCacheConfig) GetMaxHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxHosts
	}
	return nil
}

func (m *DnsCacheConfig) GetDnsFailureRefreshRate() *v3.Cluster_RefreshRate {
	if m != nil {
		return m.DnsFailureRefreshRate
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsCacheConfig)(nil), "envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/common/dynamic_forward_proxy/v3/dns_cache.proto", fileDescriptor_f57893b6dd868364)
}

var fileDescriptor_f57893b6dd868364 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xdf, 0x6e, 0xd3, 0x3c,
	0x18, 0xc6, 0xe7, 0x7e, 0x5d, 0xd7, 0xe6, 0x83, 0x52, 0x22, 0x21, 0xc2, 0x40, 0x53, 0x41, 0x42,
	0xaa, 0x26, 0x64, 0x4f, 0xc9, 0x19, 0x12, 0x13, 0x72, 0xab, 0x01, 0x12, 0x07, 0x5b, 0x04, 0x9c,
	0x46, 0x5e, 0xe2, 0xb4, 0x11, 0x8e, 0x1d, 0xd9, 0x4e, 0xd6, 0x9e, 0x22, 0x0e, 0xb8, 0x06, 0xc4,
	0x15, 0xec, 0x2a, 0x38, 0xe1, 0xa6, 0x7a, 0x84, 0x6c, 0xa7, 0xb0, 0x8a, 0x3f, 0x13, 0x67, 0x6f,
	0xfc, 0xfa, 0xf9, 0xf9, 0x7d, 0x1e, 0xbd, 0xf1, 0x9e, 0x53, 0xde, 0x88, 0x15, 0xa2, 0x4b, 0x4d,
	0xb9, 0x2a, 0x04, 0x57, 0x28, 0x15, 0x65, 0x29, 0x38, 0xca, 0x56, 0x9c, 0x94, 0x45, 0x9a, 0xe4,
	0x42, 0x5e, 0x10, 0x99, 0x25, 0x95, 0x14, 0xcb, 0x15, 0x6a, 0x22, 0x94, 0x71, 0x95, 0xa4, 0x24,
	0x5d, 0x50, 0x58, 0x49, 0xa1, 0x85, 0x7f, 0x64, 0x09, 0xf0, 0x27, 0x01, 0x3a, 0x02, 0xfc, 0x2d,
	0x01, 0x36, 0xd1, 0xfe, 0x63, 0xf7, 0x66, 0x2a, 0x78, 0x5e, 0xcc, 0x51, 0xca, 0x6a, 0xa5, 0xa9,
	0x34, 0xe8, 0xb6, 0x74, 0xe0, 0xfd, 0x83, 0xb9, 0x10, 0x73, 0x46, 0x91, 0xfd, 0x3a, 0xaf, 0x73,
	0x94, 0xd5, 0x92, 0xe8, 0x42, 0xf0, 0x3f, 0xf5, 0x2f, 0x24, 0xa9, 0x2a, 0x2a, 0x55, 0xdb, 0x7f,
	0x58, 0x67, 0x15, 0x41, 0x84, 0x73, 0xa1, 0xad, 0x4c, 0xa1, 0x86, 0x4a, 0x33, 0x61, 0xc1, 0xe7,
	0xed, 0x95, 0xbb, 0x0d, 0x61, 0x45, 0x46, 0x34, 0x45, 0x9b, 0xc2, 0x35, 0x1e, 0x7d, 0xe9, 0x7a,
	0xc3, 0x19, 0x57, 0x53, 0xe3, 0x73, 0x6a, 0xe7, 0xf4, 0xef, 0x7b, 0x5d, 0x4e, 0x4a, 0x1a, 0x80,
	0x31, 0x98, 0x0c, 0xf0, 0xde, 0x1a, 0x77, 0x65, 0x67, 0x0c, 0x62, 0x7b, 0xe8, 0xe7, 0xde, 0x6d,
	0x93, 0x0b, 0x13, 0xe2, 0x7d, 0x5d, 0x25, 0x39, 0x29, 0x0b, 0xb6, 0x0a, 0x3a, 0x63, 0x30, 0x19,
	0x86, 0x47, 0xd0, 0x05, 0xe4, 0xec, 0xc2, 0x8d, 0xc7, 0x26, 0x82, 0xd3, 0xb6, 0x9c, 0x71, 0xf5,
	0xda, 0x0a, 0x4f, 0xac, 0x0e, 0xf7, 0xd7, 0x78, 0xf7, 0x03, 0xe8, 0x8c, 0x40, 0x7c, 0x2b, 0xdb,
	0x6e, 0xf9, 0x67, 0xde, 0xc8, 0xbc, 0x23, 0x69, 0x2e, 0xa9, 0x5a, 0x24, 0x92, 0x68, 0x1a, 0xfc,
	0x37, 0x06, 0x93, 0xff, 0xc3, 0x7b, 0xd0, 0xc5, 0x01, 0x37, 0x71, 0xc0, 0x59, 0x1b, 0x17, 0xbe,
	0xb1, 0xc6, 0x83, 0x4b, 0xd0, 0x0b, 0xbb, 0xa3, 0xaf, 0x1f, 0x9f, 0xc5, 0xc3, 0x8c, 0xab, 0xd8,
	0xe9, 0x63, 0xa2, 0xa9, 0x7f, 0xec, 0xf5, 0x17, 0x42, 0xe9, 0x44, 0x6b, 0x16, 0x74, 0xaf, 0x43,
	0x99, 0xd1, 0x2e, 0x41, 0xe7, 0x70, 0x27, 0xde, 0x33, 0xa2, 0x37, 0x9a, 0xf9, 0xd8, 0x1b, 0x94,
	0x64, 0x99, 0x98, 0x4f, 0x15, 0xec, 0x5a, 0xc0, 0x83, 0x5f, 0x00, 0x6f, 0x5f, 0x71, 0x1d, 0x85,
	0xef, 0x08, 0xab, 0xa9, 0x8d, 0xee, 0xb0, 0x33, 0xde, 0x89, 0xfb, 0x25, 0x59, 0xbe, 0x34, 0x32,
	0x9f, 0x7a, 0x81, 0xb1, 0x95, 0x93, 0x82, 0xd5, 0x92, 0x6e, 0xdb, 0xeb, 0x59, 0xe4, 0x93, 0x6b,
	0x53, 0xbc, 0xe2, 0x29, 0xbe, 0x93, 0x71, 0x75, 0xe2, 0x60, 0x57, 0x8e, 0x9f, 0xbe, 0xf8, 0xfc,
	0xed, 0xd3, 0x01, 0x6e, 0x77, 0xfe, 0x07, 0xea, 0xaf, 0xdb, 0x1a, 0x12, 0x56, 0x2d, 0x08, 0xdc,
	0xde, 0x05, 0x7c, 0xe6, 0x1d, 0x17, 0xc2, 0x4d, 0xe4, 0xee, 0xfd, 0xeb, 0x3f, 0x80, 0x6f, 0x6e,
	0x88, 0xa7, 0x26, 0xa2, 0x53, 0x70, 0xde, 0xb3, 0x59, 0x45, 0xdf, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xec, 0x1a, 0x84, 0xf3, 0x91, 0x03, 0x00, 0x00,
}