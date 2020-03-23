// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/http/gzip/v3/gzip.proto

package envoy_extensions_filters_http_gzip_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/compressor/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

var Gzip_CompressionStrategy_name = map[int32]string{
	0: "DEFAULT",
	1: "FILTERED",
	2: "HUFFMAN",
	3: "RLE",
}

var Gzip_CompressionStrategy_value = map[string]int32{
	"DEFAULT":  0,
	"FILTERED": 1,
	"HUFFMAN":  2,
	"RLE":      3,
}

func (x Gzip_CompressionStrategy) String() string {
	return proto.EnumName(Gzip_CompressionStrategy_name, int32(x))
}

func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0, 0}
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

var Gzip_CompressionLevel_Enum_name = map[int32]string{
	0: "DEFAULT",
	1: "BEST",
	2: "SPEED",
}

var Gzip_CompressionLevel_Enum_value = map[string]int32{
	"DEFAULT": 0,
	"BEST":    1,
	"SPEED":   2,
}

func (x Gzip_CompressionLevel_Enum) String() string {
	return proto.EnumName(Gzip_CompressionLevel_Enum_name, int32(x))
}

func (Gzip_CompressionLevel_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0, 0, 0}
}

type Gzip struct {
	MemoryLevel                                     *wrappers.UInt32Value      `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel,proto3" json:"memory_level,omitempty"`
	HiddenEnvoyDeprecatedContentLength              *wrappers.UInt32Value      `protobuf:"bytes,2,opt,name=hidden_envoy_deprecated_content_length,json=hiddenEnvoyDeprecatedContentLength,proto3" json:"hidden_envoy_deprecated_content_length,omitempty"` // Deprecated: Do not use.
	CompressionLevel                                Gzip_CompressionLevel_Enum `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.extensions.filters.http.gzip.v3.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	CompressionStrategy                             Gzip_CompressionStrategy   `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.extensions.filters.http.gzip.v3.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	HiddenEnvoyDeprecatedContentType                []string                   `protobuf:"bytes,6,rep,name=hidden_envoy_deprecated_content_type,json=hiddenEnvoyDeprecatedContentType,proto3" json:"hidden_envoy_deprecated_content_type,omitempty"`                                                   // Deprecated: Do not use.
	HiddenEnvoyDeprecatedDisableOnEtagHeader        bool                       `protobuf:"varint,7,opt,name=hidden_envoy_deprecated_disable_on_etag_header,json=hiddenEnvoyDeprecatedDisableOnEtagHeader,proto3" json:"hidden_envoy_deprecated_disable_on_etag_header,omitempty"`                      // Deprecated: Do not use.
	HiddenEnvoyDeprecatedRemoveAcceptEncodingHeader bool                       `protobuf:"varint,8,opt,name=hidden_envoy_deprecated_remove_accept_encoding_header,json=hiddenEnvoyDeprecatedRemoveAcceptEncodingHeader,proto3" json:"hidden_envoy_deprecated_remove_accept_encoding_header,omitempty"` // Deprecated: Do not use.
	WindowBits                                      *wrappers.UInt32Value      `protobuf:"bytes,9,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	Compressor                                      *v3.Compressor             `protobuf:"bytes,10,opt,name=compressor,proto3" json:"compressor,omitempty"`
	XXX_NoUnkeyedLiteral                            struct{}                   `json:"-"`
	XXX_unrecognized                                []byte                     `json:"-"`
	XXX_sizecache                                   int32                      `json:"-"`
}

func (m *Gzip) Reset()         { *m = Gzip{} }
func (m *Gzip) String() string { return proto.CompactTextString(m) }
func (*Gzip) ProtoMessage()    {}
func (*Gzip) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0}
}

func (m *Gzip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip.Unmarshal(m, b)
}
func (m *Gzip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip.Marshal(b, m, deterministic)
}
func (m *Gzip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip.Merge(m, src)
}
func (m *Gzip) XXX_Size() int {
	return xxx_messageInfo_Gzip.Size(m)
}
func (m *Gzip) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip proto.InternalMessageInfo

func (m *Gzip) GetMemoryLevel() *wrappers.UInt32Value {
	if m != nil {
		return m.MemoryLevel
	}
	return nil
}

// Deprecated: Do not use.
func (m *Gzip) GetHiddenEnvoyDeprecatedContentLength() *wrappers.UInt32Value {
	if m != nil {
		return m.HiddenEnvoyDeprecatedContentLength
	}
	return nil
}

func (m *Gzip) GetCompressionLevel() Gzip_CompressionLevel_Enum {
	if m != nil {
		return m.CompressionLevel
	}
	return Gzip_CompressionLevel_DEFAULT
}

func (m *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if m != nil {
		return m.CompressionStrategy
	}
	return Gzip_DEFAULT
}

// Deprecated: Do not use.
func (m *Gzip) GetHiddenEnvoyDeprecatedContentType() []string {
	if m != nil {
		return m.HiddenEnvoyDeprecatedContentType
	}
	return nil
}

// Deprecated: Do not use.
func (m *Gzip) GetHiddenEnvoyDeprecatedDisableOnEtagHeader() bool {
	if m != nil {
		return m.HiddenEnvoyDeprecatedDisableOnEtagHeader
	}
	return false
}

// Deprecated: Do not use.
func (m *Gzip) GetHiddenEnvoyDeprecatedRemoveAcceptEncodingHeader() bool {
	if m != nil {
		return m.HiddenEnvoyDeprecatedRemoveAcceptEncodingHeader
	}
	return false
}

func (m *Gzip) GetWindowBits() *wrappers.UInt32Value {
	if m != nil {
		return m.WindowBits
	}
	return nil
}

func (m *Gzip) GetCompressor() *v3.Compressor {
	if m != nil {
		return m.Compressor
	}
	return nil
}

type Gzip_CompressionLevel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gzip_CompressionLevel) Reset()         { *m = Gzip_CompressionLevel{} }
func (m *Gzip_CompressionLevel) String() string { return proto.CompactTextString(m) }
func (*Gzip_CompressionLevel) ProtoMessage()    {}
func (*Gzip_CompressionLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_75f828f0702c619c, []int{0, 0}
}

func (m *Gzip_CompressionLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip_CompressionLevel.Unmarshal(m, b)
}
func (m *Gzip_CompressionLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip_CompressionLevel.Marshal(b, m, deterministic)
}
func (m *Gzip_CompressionLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip_CompressionLevel.Merge(m, src)
}
func (m *Gzip_CompressionLevel) XXX_Size() int {
	return xxx_messageInfo_Gzip_CompressionLevel.Size(m)
}
func (m *Gzip_CompressionLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip_CompressionLevel.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip_CompressionLevel proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.extensions.filters.http.gzip.v3.Gzip_CompressionStrategy", Gzip_CompressionStrategy_name, Gzip_CompressionStrategy_value)
	proto.RegisterEnum("envoy.extensions.filters.http.gzip.v3.Gzip_CompressionLevel_Enum", Gzip_CompressionLevel_Enum_name, Gzip_CompressionLevel_Enum_value)
	proto.RegisterType((*Gzip)(nil), "envoy.extensions.filters.http.gzip.v3.Gzip")
	proto.RegisterType((*Gzip_CompressionLevel)(nil), "envoy.extensions.filters.http.gzip.v3.Gzip.CompressionLevel")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/http/gzip/v3/gzip.proto", fileDescriptor_75f828f0702c619c)
}

var fileDescriptor_75f828f0702c619c = []byte{
	// 666 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x5f, 0x6f, 0xd3, 0x3c,
	0x14, 0x87, 0xdf, 0xb4, 0xdd, 0xd6, 0xb8, 0xd3, 0x4b, 0xc8, 0x90, 0x88, 0x26, 0x34, 0x95, 0x8a,
	0x41, 0x84, 0x50, 0x82, 0x5a, 0xf1, 0x47, 0x13, 0x08, 0x35, 0x6b, 0xca, 0x86, 0x0a, 0x4c, 0x59,
	0x07, 0x97, 0x51, 0x9a, 0x9c, 0xa5, 0x96, 0x52, 0x3b, 0x72, 0xdc, 0x6c, 0x1d, 0x5c, 0x21, 0x2e,
	0xf8, 0x0c, 0x7c, 0x1f, 0xbe, 0xd4, 0xae, 0x50, 0x9c, 0xac, 0x2d, 0xa3, 0xdb, 0x2a, 0xae, 0xda,
	0xc8, 0xfe, 0x3d, 0x8f, 0x7d, 0xce, 0x31, 0x7a, 0x0a, 0x24, 0xa5, 0x13, 0x13, 0x4e, 0x39, 0x90,
	0x04, 0x53, 0x92, 0x98, 0xc7, 0x38, 0xe2, 0xc0, 0x12, 0x73, 0xc8, 0x79, 0x6c, 0x86, 0x67, 0x38,
	0x36, 0xd3, 0x96, 0xf8, 0x35, 0x62, 0x46, 0x39, 0x55, 0xb7, 0x45, 0xc2, 0x98, 0x25, 0x8c, 0x22,
	0x61, 0x64, 0x09, 0x43, 0xec, 0x4c, 0x5b, 0x9b, 0xaf, 0xae, 0x07, 0xfb, 0x74, 0x14, 0x33, 0x48,
	0x12, 0xca, 0x32, 0xfc, 0xec, 0x2b, 0x97, 0x6c, 0x6e, 0x85, 0x94, 0x86, 0x11, 0x98, 0xe2, 0x6b,
	0x30, 0x3e, 0x36, 0x4f, 0x98, 0x17, 0xc7, 0x99, 0x24, 0x5f, 0xbf, 0x3f, 0x0e, 0x62, 0xcf, 0xf4,
	0x08, 0xa1, 0xdc, 0xe3, 0x82, 0x9e, 0x02, 0xcb, 0x34, 0x98, 0x84, 0xc5, 0x96, 0xbb, 0xa9, 0x17,
	0xe1, 0xc0, 0xe3, 0x60, 0x5e, 0xfc, 0xc9, 0x17, 0x1a, 0xdf, 0x65, 0x54, 0x79, 0x7b, 0x86, 0x63,
	0xf5, 0x1d, 0x5a, 0x1f, 0xc1, 0x88, 0xb2, 0x89, 0x1b, 0x41, 0x0a, 0x91, 0x26, 0xd5, 0x25, 0xbd,
	0xd6, 0xbc, 0x67, 0xe4, 0x6e, 0xe3, 0xc2, 0x6d, 0x1c, 0xed, 0x13, 0xde, 0x6a, 0x7e, 0xf2, 0xa2,
	0x31, 0x58, 0xf2, 0xb9, 0xb5, 0xfa, 0xb8, 0xa2, 0xc9, 0xba, 0xe4, 0xd4, 0xf2, 0x70, 0x2f, 0xcb,
	0xaa, 0x1c, 0x3d, 0x1c, 0xe2, 0x20, 0x00, 0xe2, 0x8a, 0x7b, 0xbb, 0x01, 0xc4, 0x0c, 0x7c, 0x8f,
	0x43, 0xe0, 0xfa, 0x94, 0x70, 0x20, 0xdc, 0x8d, 0x80, 0x84, 0x7c, 0xa8, 0x95, 0x96, 0xb0, 0x94,
	0x34, 0xc9, 0x69, 0xe4, 0x3c, 0x3b, 0xc3, 0x75, 0xa6, 0xb4, 0xdd, 0x1c, 0xd6, 0x13, 0x2c, 0xf5,
	0x14, 0xdd, 0xbe, 0x28, 0x1d, 0xa6, 0xa4, 0xb8, 0x46, 0xb9, 0x2e, 0xe9, 0xff, 0x37, 0xdb, 0xc6,
	0x52, 0x7d, 0x32, 0xb2, 0x4a, 0x18, 0xbb, 0x33, 0x88, 0xb8, 0x8e, 0x61, 0x93, 0xf1, 0xc8, 0xaa,
	0x9e, 0x5b, 0x2b, 0xdf, 0xa4, 0x92, 0x22, 0x39, 0x8a, 0x7f, 0x69, 0x83, 0xfa, 0x15, 0xdd, 0x99,
	0x37, 0x27, 0x9c, 0x79, 0x1c, 0xc2, 0x89, 0x56, 0x11, 0xf2, 0x37, 0xff, 0x28, 0x3f, 0x2c, 0x30,
	0x73, 0xea, 0x0d, 0xff, 0xef, 0x65, 0xd5, 0x41, 0x0f, 0x6e, 0xaa, 0x36, 0x9f, 0xc4, 0xa0, 0xad,
	0xd6, 0xcb, 0xba, 0x2c, 0xaa, 0x59, 0xbf, 0xae, 0x9a, 0xfd, 0x49, 0x0c, 0x6a, 0x80, 0x8c, 0xab,
	0x98, 0x01, 0x4e, 0xbc, 0x41, 0x04, 0x2e, 0x25, 0x2e, 0x70, 0x2f, 0x74, 0x87, 0xe0, 0x05, 0xc0,
	0xb4, 0xb5, 0xba, 0xa4, 0x57, 0x05, 0x5d, 0x5f, 0x48, 0xef, 0xe4, 0xb1, 0x8f, 0xc4, 0xe6, 0x5e,
	0xb8, 0x27, 0x32, 0x2a, 0x43, 0xcf, 0xae, 0xb2, 0x30, 0x18, 0xd1, 0x14, 0x5c, 0xcf, 0xf7, 0x21,
	0xe6, 0x2e, 0x10, 0x9f, 0x06, 0x98, 0x4c, 0x65, 0xd5, 0xa9, 0xcc, 0x5c, 0x28, 0x73, 0x44, 0xba,
	0x2d, 0xc2, 0x76, 0x91, 0x2d, 0x9c, 0x7b, 0xa8, 0x76, 0x82, 0x49, 0x40, 0x4f, 0xdc, 0x01, 0xe6,
	0x89, 0x26, 0x2f, 0x3f, 0xe6, 0xb7, 0x74, 0xd9, 0x41, 0x79, 0xd6, 0xc2, 0x3c, 0x51, 0x3f, 0x23,
	0x34, 0x7b, 0xaa, 0x1a, 0x12, 0xa0, 0x17, 0x37, 0xf4, 0x7a, 0xee, 0x6d, 0xa7, 0xad, 0x69, 0xb3,
	0x29, 0x73, 0xe6, 0x50, 0x9b, 0x5f, 0x90, 0x72, 0x79, 0x06, 0x1b, 0x3a, 0xaa, 0x64, 0x63, 0xa8,
	0xd6, 0xd0, 0x5a, 0xc7, 0xee, 0xb6, 0x8f, 0x7a, 0x7d, 0xe5, 0x3f, 0xb5, 0x8a, 0x2a, 0x96, 0x7d,
	0xd8, 0x57, 0x24, 0x55, 0x46, 0x2b, 0x87, 0x07, 0xb6, 0xdd, 0x51, 0x4a, 0x3b, 0xaf, 0x7f, 0xfe,
	0xfa, 0xb1, 0xf5, 0x12, 0x3d, 0xcf, 0x0f, 0xe2, 0x53, 0x72, 0x8c, 0xc3, 0xe2, 0x10, 0xf3, 0xf3,
	0xd6, 0x5c, 0x3c, 0xec, 0x8d, 0x2e, 0xda, 0x58, 0x30, 0x83, 0x7f, 0x7a, 0xd7, 0x51, 0xb5, 0xbb,
	0xdf, 0xeb, 0xdb, 0x8e, 0xdd, 0x51, 0xa4, 0x6c, 0x69, 0xef, 0xa8, 0xdb, 0x7d, 0xdf, 0xfe, 0xa0,
	0x94, 0xd4, 0x35, 0x54, 0x76, 0x7a, 0xb6, 0x52, 0xde, 0x79, 0x92, 0x1d, 0xe3, 0x11, 0xda, 0x5e,
	0xea, 0x18, 0x96, 0x85, 0x5a, 0x98, 0xe6, 0xb5, 0x8b, 0x19, 0x3d, 0x9d, 0x2c, 0xf7, 0x64, 0x2c,
	0x39, 0x0b, 0x1f, 0x64, 0x3d, 0x3b, 0x90, 0x06, 0xab, 0xa2, 0x79, 0xad, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xb4, 0x20, 0x87, 0x5e, 0xc6, 0x05, 0x00, 0x00,
}