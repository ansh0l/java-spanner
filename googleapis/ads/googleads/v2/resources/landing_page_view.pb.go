// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/resources/landing_page_view.proto

package resources

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A landing page view with metrics aggregated at the unexpanded final URL
// level.
type LandingPageView struct {
	// Output only. The resource name of the landing page view.
	// Landing page view resource names have the form:
	//
	// `customers/{customer_id}/landingPageViews/{unexpanded_final_url_fingerprint}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The advertiser-specified final URL.
	UnexpandedFinalUrl   *wrappers.StringValue `protobuf:"bytes,2,opt,name=unexpanded_final_url,json=unexpandedFinalUrl,proto3" json:"unexpanded_final_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LandingPageView) Reset()         { *m = LandingPageView{} }
func (m *LandingPageView) String() string { return proto.CompactTextString(m) }
func (*LandingPageView) ProtoMessage()    {}
func (*LandingPageView) Descriptor() ([]byte, []int) {
	return fileDescriptor_902f5a0985d352f9, []int{0}
}

func (m *LandingPageView) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LandingPageView.Unmarshal(m, b)
}
func (m *LandingPageView) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LandingPageView.Marshal(b, m, deterministic)
}
func (m *LandingPageView) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LandingPageView.Merge(m, src)
}
func (m *LandingPageView) XXX_Size() int {
	return xxx_messageInfo_LandingPageView.Size(m)
}
func (m *LandingPageView) XXX_DiscardUnknown() {
	xxx_messageInfo_LandingPageView.DiscardUnknown(m)
}

var xxx_messageInfo_LandingPageView proto.InternalMessageInfo

func (m *LandingPageView) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *LandingPageView) GetUnexpandedFinalUrl() *wrappers.StringValue {
	if m != nil {
		return m.UnexpandedFinalUrl
	}
	return nil
}

func init() {
	proto.RegisterType((*LandingPageView)(nil), "google.ads.googleads.v2.resources.LandingPageView")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/resources/landing_page_view.proto", fileDescriptor_902f5a0985d352f9)
}

var fileDescriptor_902f5a0985d352f9 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x18, 0x86, 0xd9, 0x04, 0x04, 0xa3, 0x22, 0x84, 0x1e, 0x6a, 0x29, 0xda, 0x0a, 0x85, 0xc5, 0xc3,
	0x8c, 0xc4, 0x53, 0xc7, 0xd3, 0xe4, 0x60, 0x41, 0x44, 0x96, 0x2d, 0xcd, 0x41, 0x02, 0x61, 0x36,
	0xf3, 0xed, 0x74, 0x60, 0x32, 0x13, 0x66, 0x92, 0xac, 0x50, 0x7a, 0xf0, 0xaf, 0x78, 0x11, 0xfc,
	0x29, 0xfe, 0x8a, 0x9e, 0xfb, 0x13, 0x3c, 0x49, 0x36, 0x99, 0x6c, 0xa9, 0xa0, 0xbd, 0xbd, 0xe1,
	0x7d, 0xde, 0x77, 0xbe, 0x8f, 0x2f, 0xd1, 0xa9, 0x30, 0x46, 0x28, 0xc0, 0x8c, 0x3b, 0x3c, 0xc8,
	0x5e, 0x75, 0x09, 0xb6, 0xe0, 0x4c, 0x6b, 0x4b, 0x70, 0x58, 0x31, 0xcd, 0xa5, 0x16, 0x45, 0xcd,
	0x04, 0x14, 0x9d, 0x84, 0x0d, 0xaa, 0xad, 0x69, 0x4c, 0x7c, 0x3c, 0xf0, 0x88, 0x71, 0x87, 0xa6,
	0x28, 0xea, 0x12, 0x34, 0x45, 0x0f, 0x5e, 0xf9, 0xf6, 0x5a, 0xe2, 0xb5, 0x04, 0xc5, 0x8b, 0x15,
	0x5c, 0xb2, 0x4e, 0x1a, 0x3b, 0x74, 0x1c, 0xbc, 0xb8, 0x03, 0xf8, 0xd8, 0x68, 0xbd, 0x1c, 0xad,
	0xed, 0xd7, 0xaa, 0x5d, 0xe3, 0x8d, 0x65, 0x75, 0x0d, 0xd6, 0x8d, 0xfe, 0xe1, 0x9d, 0x28, 0xd3,
	0xda, 0x34, 0xac, 0x91, 0x46, 0x8f, 0xee, 0xeb, 0x1f, 0x41, 0xf4, 0xfc, 0xd3, 0x30, 0xf8, 0x82,
	0x09, 0xc8, 0x24, 0x6c, 0xe2, 0x8b, 0xe8, 0x99, 0x7f, 0xa3, 0xd0, 0xac, 0x82, 0xfd, 0xd9, 0xd1,
	0x6c, 0xfe, 0x38, 0x7d, 0x7b, 0x43, 0xc3, 0xdf, 0xf4, 0x4d, 0x34, 0xdf, 0x2d, 0x31, 0xaa, 0x5a,
	0x3a, 0x54, 0x9a, 0x0a, 0xdf, 0x2b, 0x5a, 0x3e, 0xf5, 0x35, 0x9f, 0x59, 0x05, 0xf1, 0x79, 0xb4,
	0xd7, 0x6a, 0xf8, 0x5a, 0x33, 0xcd, 0x81, 0x17, 0x6b, 0xa9, 0x99, 0x2a, 0x5a, 0xab, 0xf6, 0x83,
	0xa3, 0xd9, 0xfc, 0x49, 0x72, 0x38, 0x96, 0x21, 0xbf, 0x07, 0x3a, 0x6f, 0xac, 0xd4, 0x22, 0x63,
	0xaa, 0x85, 0x34, 0xbc, 0xa1, 0xe1, 0x32, 0xde, 0xc5, 0x3f, 0xf4, 0xe9, 0x0b, 0xab, 0xc8, 0xe5,
	0x2d, 0x85, 0x87, 0x4f, 0x14, 0x9f, 0x96, 0xad, 0x6b, 0x4c, 0x05, 0xd6, 0xe1, 0x2b, 0x2f, 0xaf,
	0xfd, 0xe5, 0x3c, 0xe5, 0xf0, 0xd5, 0x5f, 0xb7, 0xbc, 0x4e, 0xbf, 0x05, 0xd1, 0x49, 0x69, 0x2a,
	0xf4, 0xdf, 0x6b, 0xa6, 0x7b, 0xf7, 0x5e, 0x5d, 0xf4, 0x1b, 0x2d, 0x66, 0x5f, 0x3e, 0x8e, 0x51,
	0x61, 0x14, 0xd3, 0x02, 0x19, 0x2b, 0xb0, 0x00, 0xbd, 0xdd, 0x17, 0xef, 0xe6, 0xfe, 0xc7, 0x2f,
	0xf6, 0x7e, 0x52, 0xdf, 0x83, 0xf0, 0x8c, 0xd2, 0x9f, 0xc1, 0xf1, 0xd9, 0x50, 0x49, 0xb9, 0x43,
	0x83, 0xec, 0x55, 0x96, 0xa0, 0xa5, 0x27, 0x7f, 0x79, 0x26, 0xa7, 0xdc, 0xe5, 0x13, 0x93, 0x67,
	0x49, 0x3e, 0x31, 0xb7, 0xc1, 0xc9, 0x60, 0x10, 0x42, 0xb9, 0x23, 0x64, 0xa2, 0x08, 0xc9, 0x12,
	0x42, 0x26, 0x6e, 0xf5, 0x68, 0x3b, 0xec, 0xbb, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xba, 0x75,
	0x07, 0x0b, 0x0e, 0x03, 0x00, 0x00,
}