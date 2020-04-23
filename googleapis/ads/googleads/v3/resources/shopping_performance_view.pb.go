// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/resources/shopping_performance_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Shopping performance view.
// Provides Shopping campaign statistics aggregated at several product dimension
// levels. Product dimension values from Merchant Center such as brand,
// category, custom attributes, product condition and product type will reflect
// the state of each dimension as of the date and time when the corresponding
// event was recorded.
type ShoppingPerformanceView struct {
	// Output only. The resource name of the Shopping performance view.
	// Shopping performance view resource names have the form:
	// `customers/{customer_id}/shoppingPerformanceView`
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShoppingPerformanceView) Reset()         { *m = ShoppingPerformanceView{} }
func (m *ShoppingPerformanceView) String() string { return proto.CompactTextString(m) }
func (*ShoppingPerformanceView) ProtoMessage()    {}
func (*ShoppingPerformanceView) Descriptor() ([]byte, []int) {
	return fileDescriptor_f485d5704344bc86, []int{0}
}

func (m *ShoppingPerformanceView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShoppingPerformanceView.Unmarshal(m, b)
}
func (m *ShoppingPerformanceView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShoppingPerformanceView.Marshal(b, m, deterministic)
}
func (m *ShoppingPerformanceView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShoppingPerformanceView.Merge(m, src)
}
func (m *ShoppingPerformanceView) XXX_Size() int {
	return xxx_messageInfo_ShoppingPerformanceView.Size(m)
}
func (m *ShoppingPerformanceView) XXX_DiscardUnknown() {
	xxx_messageInfo_ShoppingPerformanceView.DiscardUnknown(m)
}

var xxx_messageInfo_ShoppingPerformanceView proto.InternalMessageInfo

func (m *ShoppingPerformanceView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*ShoppingPerformanceView)(nil), "google.ads.googleads.v3.resources.ShoppingPerformanceView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/resources/shopping_performance_view.proto", fileDescriptor_f485d5704344bc86)
}

var fileDescriptor_f485d5704344bc86 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x69, 0x07, 0x82, 0x45, 0x2f, 0xbb, 0xa8, 0x63, 0xa0, 0x13, 0x06, 0x1e, 0x24, 0x11,
	0x7b, 0x91, 0x78, 0xca, 0x2e, 0x03, 0x0f, 0x32, 0x26, 0xf4, 0x20, 0x95, 0x9a, 0xb5, 0x59, 0x17,
	0x58, 0xf3, 0x95, 0xa4, 0xeb, 0x0e, 0xe2, 0x03, 0xf8, 0x1a, 0x1e, 0x7d, 0x14, 0xdf, 0x41, 0xf0,
	0xbc, 0x47, 0xf0, 0x24, 0x5b, 0x9b, 0x6c, 0x08, 0x53, 0xbc, 0xfd, 0xe1, 0xfb, 0xe5, 0xf7, 0xfd,
	0xf9, 0x88, 0x47, 0x53, 0x80, 0x74, 0xca, 0x31, 0x4b, 0x34, 0xae, 0xe2, 0x32, 0x95, 0x3e, 0x56,
	0x5c, 0xc3, 0x4c, 0xc5, 0x5c, 0x63, 0x3d, 0x81, 0x3c, 0x17, 0x32, 0x8d, 0x72, 0xae, 0xc6, 0xa0,
	0x32, 0x26, 0x63, 0x1e, 0x95, 0x82, 0xcf, 0x51, 0xae, 0xa0, 0x80, 0x66, 0xa7, 0x7a, 0x87, 0x58,
	0xa2, 0x91, 0x55, 0xa0, 0xd2, 0x47, 0x56, 0xd1, 0x3a, 0x36, 0x5b, 0x72, 0x81, 0xc7, 0x82, 0x4f,
	0x93, 0x68, 0xc4, 0x27, 0xac, 0x14, 0xa0, 0x2a, 0x47, 0xeb, 0x68, 0x03, 0x30, 0xcf, 0xea, 0x51,
	0x7b, 0x63, 0xc4, 0xa4, 0x84, 0x82, 0x15, 0x02, 0xa4, 0xae, 0xa6, 0xa7, 0x1f, 0x8e, 0x77, 0x70,
	0x57, 0x17, 0x1c, 0xac, 0xfb, 0x05, 0x82, 0xcf, 0x9b, 0x0f, 0xde, 0xbe, 0x71, 0x45, 0x92, 0x65,
	0xfc, 0xd0, 0x39, 0x71, 0xce, 0x76, 0x7b, 0x57, 0x9f, 0xb4, 0xf1, 0x45, 0x2f, 0xbd, 0x8b, 0x75,
	0xd9, 0x3a, 0xe5, 0x42, 0xa3, 0x18, 0x32, 0xbc, 0x45, 0x38, 0xdc, 0x33, 0xba, 0x5b, 0x96, 0x71,
	0x12, 0x2f, 0xe8, 0xe3, 0xff, 0x25, 0xcd, 0xf3, 0x78, 0xa6, 0x0b, 0xc8, 0xb8, 0xd2, 0xf8, 0xc9,
	0xc4, 0x67, 0x7b, 0xe4, 0x1f, 0x74, 0xef, 0xc5, 0xf5, 0xba, 0x31, 0x64, 0xe8, 0xcf, 0x1b, 0xf7,
	0xda, 0x5b, 0x16, 0x0e, 0x96, 0x77, 0x1a, 0x38, 0xf7, 0x37, 0xb5, 0x22, 0x85, 0x29, 0x93, 0x29,
	0x02, 0x95, 0xe2, 0x94, 0xcb, 0xd5, 0x15, 0xf1, 0xba, 0xfa, 0x2f, 0x1f, 0xe1, 0xda, 0xa6, 0x57,
	0xb7, 0xd1, 0xa7, 0xf4, 0xcd, 0xed, 0xf4, 0x2b, 0x25, 0x4d, 0x34, 0xaa, 0xe2, 0x32, 0x05, 0x3e,
	0x1a, 0x1a, 0xf2, 0xdd, 0x30, 0x21, 0x4d, 0x74, 0x68, 0x99, 0x30, 0xf0, 0x43, 0xcb, 0x2c, 0xdc,
	0x6e, 0x35, 0x20, 0x84, 0x26, 0x9a, 0x10, 0x4b, 0x11, 0x12, 0xf8, 0x84, 0x58, 0x6e, 0xb4, 0xb3,
	0x2a, 0xeb, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x30, 0x09, 0x1d, 0x72, 0xb4, 0x02, 0x00, 0x00,
}