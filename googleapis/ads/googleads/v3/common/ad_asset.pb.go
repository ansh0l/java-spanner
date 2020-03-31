// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/common/ad_asset.proto

package common

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
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

// A text asset used inside an ad.
type AdTextAsset struct {
	// Asset text.
	Text *wrappers.StringValue `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The pinned field of the asset. This restricts the asset to only serve
	// within this field. Multiple assets can be pinned to the same field. An
	// asset that is unpinned or pinned to a different field will not serve in a
	// field where some other asset has been pinned.
	PinnedField          enums.ServedAssetFieldTypeEnum_ServedAssetFieldType `protobuf:"varint,2,opt,name=pinned_field,json=pinnedField,proto3,enum=google.ads.googleads.v3.enums.ServedAssetFieldTypeEnum_ServedAssetFieldType" json:"pinned_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                            `json:"-"`
	XXX_unrecognized     []byte                                              `json:"-"`
	XXX_sizecache        int32                                               `json:"-"`
}

func (m *AdTextAsset) Reset()         { *m = AdTextAsset{} }
func (m *AdTextAsset) String() string { return proto.CompactTextString(m) }
func (*AdTextAsset) ProtoMessage()    {}
func (*AdTextAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa6efb5f889d94f, []int{0}
}

func (m *AdTextAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdTextAsset.Unmarshal(m, b)
}
func (m *AdTextAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdTextAsset.Marshal(b, m, deterministic)
}
func (m *AdTextAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdTextAsset.Merge(m, src)
}
func (m *AdTextAsset) XXX_Size() int {
	return xxx_messageInfo_AdTextAsset.Size(m)
}
func (m *AdTextAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_AdTextAsset.DiscardUnknown(m)
}

var xxx_messageInfo_AdTextAsset proto.InternalMessageInfo

func (m *AdTextAsset) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *AdTextAsset) GetPinnedField() enums.ServedAssetFieldTypeEnum_ServedAssetFieldType {
	if m != nil {
		return m.PinnedField
	}
	return enums.ServedAssetFieldTypeEnum_UNSPECIFIED
}

// An image asset used inside an ad.
type AdImageAsset struct {
	// The Asset resource name of this image.
	Asset                *wrappers.StringValue `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdImageAsset) Reset()         { *m = AdImageAsset{} }
func (m *AdImageAsset) String() string { return proto.CompactTextString(m) }
func (*AdImageAsset) ProtoMessage()    {}
func (*AdImageAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa6efb5f889d94f, []int{1}
}

func (m *AdImageAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdImageAsset.Unmarshal(m, b)
}
func (m *AdImageAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdImageAsset.Marshal(b, m, deterministic)
}
func (m *AdImageAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdImageAsset.Merge(m, src)
}
func (m *AdImageAsset) XXX_Size() int {
	return xxx_messageInfo_AdImageAsset.Size(m)
}
func (m *AdImageAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_AdImageAsset.DiscardUnknown(m)
}

var xxx_messageInfo_AdImageAsset proto.InternalMessageInfo

func (m *AdImageAsset) GetAsset() *wrappers.StringValue {
	if m != nil {
		return m.Asset
	}
	return nil
}

// A video asset used inside an ad.
type AdVideoAsset struct {
	// The Asset resource name of this video.
	Asset                *wrappers.StringValue `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdVideoAsset) Reset()         { *m = AdVideoAsset{} }
func (m *AdVideoAsset) String() string { return proto.CompactTextString(m) }
func (*AdVideoAsset) ProtoMessage()    {}
func (*AdVideoAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa6efb5f889d94f, []int{2}
}

func (m *AdVideoAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdVideoAsset.Unmarshal(m, b)
}
func (m *AdVideoAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdVideoAsset.Marshal(b, m, deterministic)
}
func (m *AdVideoAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdVideoAsset.Merge(m, src)
}
func (m *AdVideoAsset) XXX_Size() int {
	return xxx_messageInfo_AdVideoAsset.Size(m)
}
func (m *AdVideoAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_AdVideoAsset.DiscardUnknown(m)
}

var xxx_messageInfo_AdVideoAsset proto.InternalMessageInfo

func (m *AdVideoAsset) GetAsset() *wrappers.StringValue {
	if m != nil {
		return m.Asset
	}
	return nil
}

// A media bundle asset used inside an ad.
type AdMediaBundleAsset struct {
	// The Asset resource name of this media bundle.
	Asset                *wrappers.StringValue `protobuf:"bytes,1,opt,name=asset,proto3" json:"asset,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *AdMediaBundleAsset) Reset()         { *m = AdMediaBundleAsset{} }
func (m *AdMediaBundleAsset) String() string { return proto.CompactTextString(m) }
func (*AdMediaBundleAsset) ProtoMessage()    {}
func (*AdMediaBundleAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fa6efb5f889d94f, []int{3}
}

func (m *AdMediaBundleAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdMediaBundleAsset.Unmarshal(m, b)
}
func (m *AdMediaBundleAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdMediaBundleAsset.Marshal(b, m, deterministic)
}
func (m *AdMediaBundleAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdMediaBundleAsset.Merge(m, src)
}
func (m *AdMediaBundleAsset) XXX_Size() int {
	return xxx_messageInfo_AdMediaBundleAsset.Size(m)
}
func (m *AdMediaBundleAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_AdMediaBundleAsset.DiscardUnknown(m)
}

var xxx_messageInfo_AdMediaBundleAsset proto.InternalMessageInfo

func (m *AdMediaBundleAsset) GetAsset() *wrappers.StringValue {
	if m != nil {
		return m.Asset
	}
	return nil
}

func init() {
	proto.RegisterType((*AdTextAsset)(nil), "google.ads.googleads.v3.common.AdTextAsset")
	proto.RegisterType((*AdImageAsset)(nil), "google.ads.googleads.v3.common.AdImageAsset")
	proto.RegisterType((*AdVideoAsset)(nil), "google.ads.googleads.v3.common.AdVideoAsset")
	proto.RegisterType((*AdMediaBundleAsset)(nil), "google.ads.googleads.v3.common.AdMediaBundleAsset")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/common/ad_asset.proto", fileDescriptor_0fa6efb5f889d94f)
}

var fileDescriptor_0fa6efb5f889d94f = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcd, 0x6a, 0xdc, 0x30,
	0x10, 0xc7, 0xf1, 0xf6, 0xe3, 0xa0, 0x5d, 0x7a, 0xf0, 0x29, 0x84, 0x10, 0x82, 0x4f, 0xb9, 0x54,
	0x2a, 0xeb, 0x9b, 0x72, 0x92, 0xfb, 0x91, 0x16, 0x5a, 0x08, 0x49, 0xf0, 0xa1, 0x18, 0x16, 0x25,
	0x9a, 0x08, 0x81, 0x2d, 0x09, 0x4b, 0xde, 0x26, 0xaf, 0xd3, 0x63, 0x4f, 0x7d, 0x8e, 0xbe, 0x48,
	0xa1, 0x4f, 0x51, 0xac, 0xb1, 0xf7, 0xd4, 0x2d, 0x21, 0xa7, 0x9d, 0x95, 0x7e, 0xfe, 0xcd, 0x7f,
	0xc4, 0x90, 0xd7, 0xda, 0x39, 0xdd, 0x02, 0x93, 0x2a, 0x30, 0x2c, 0xc7, 0x6a, 0x5b, 0xb2, 0x5b,
	0xd7, 0x75, 0xce, 0x32, 0xa9, 0x36, 0x32, 0x04, 0x88, 0xd4, 0xf7, 0x2e, 0xba, 0xfc, 0x18, 0x19,
	0x2a, 0x55, 0xa0, 0x3b, 0x9c, 0x6e, 0x4b, 0x8a, 0xf8, 0xe1, 0xd9, 0x3e, 0x1d, 0xd8, 0xa1, 0x0b,
	0x2c, 0x40, 0xbf, 0x85, 0xc9, 0xb8, 0xb9, 0x33, 0xd0, 0xaa, 0x4d, 0x7c, 0xf0, 0x80, 0xf2, 0xc3,
	0x49, 0xce, 0xd2, 0xbf, 0x9b, 0xe1, 0x8e, 0x7d, 0xeb, 0xa5, 0xf7, 0xd0, 0x87, 0xe9, 0xfe, 0x68,
	0x96, 0x7b, 0xc3, 0xa4, 0xb5, 0x2e, 0xca, 0x68, 0x9c, 0x9d, 0x6e, 0x8b, 0x9f, 0x19, 0x59, 0x0a,
	0x75, 0x0d, 0xf7, 0x51, 0x8c, 0xfa, 0xfc, 0x0d, 0x79, 0x1e, 0xe1, 0x3e, 0x1e, 0x64, 0x27, 0xd9,
	0xe9, 0x72, 0x7d, 0x34, 0xc5, 0xa5, 0xb3, 0x9c, 0x5e, 0xc5, 0xde, 0x58, 0x5d, 0xcb, 0x76, 0x80,
	0xcb, 0x44, 0xe6, 0x8e, 0xac, 0xbc, 0xb1, 0x16, 0x14, 0x46, 0x3b, 0x58, 0x9c, 0x64, 0xa7, 0xaf,
	0xd6, 0x9f, 0xe9, 0xbe, 0x99, 0xd3, 0x4c, 0xf4, 0x2a, 0xcd, 0x94, 0x7a, 0x7e, 0x18, 0x3f, 0xbb,
	0x7e, 0xf0, 0xf0, 0xde, 0x0e, 0xdd, 0x3f, 0x2f, 0x2e, 0x97, 0xd8, 0x21, 0x1d, 0x14, 0x15, 0x59,
	0x09, 0xf5, 0xa9, 0x93, 0x1a, 0x30, 0xf2, 0x9a, 0xbc, 0x48, 0x4f, 0xf3, 0xa8, 0xcc, 0x88, 0xa2,
	0xa3, 0x36, 0x0a, 0xdc, 0xd3, 0x1d, 0x1f, 0x49, 0x2e, 0xd4, 0x17, 0x50, 0x46, 0x56, 0x83, 0x55,
	0xed, 0xd3, 0xd3, 0x54, 0xbf, 0x33, 0x52, 0xdc, 0xba, 0x8e, 0xfe, 0x7f, 0x4d, 0xaa, 0x95, 0xc0,
	0x77, 0xb9, 0x18, 0x55, 0x17, 0xd9, 0xd7, 0x77, 0x13, 0xaf, 0x5d, 0x2b, 0xad, 0xa6, 0xae, 0xd7,
	0x4c, 0x83, 0x4d, 0x8d, 0xe6, 0x35, 0xf2, 0x26, 0xec, 0x5b, 0xd2, 0x33, 0xfc, 0xf9, 0xbe, 0x78,
	0x76, 0x2e, 0xc4, 0x8f, 0xc5, 0xf1, 0x39, 0xca, 0x84, 0x0a, 0x14, 0xcb, 0xb1, 0xaa, 0x4b, 0xfa,
	0x36, 0x61, 0xbf, 0x66, 0xa0, 0x11, 0x2a, 0x34, 0x3b, 0xa0, 0xa9, 0xcb, 0x06, 0x81, 0x3f, 0x8b,
	0x02, 0x4f, 0x39, 0x17, 0x2a, 0x70, 0xbe, 0x43, 0x38, 0xaf, 0x4b, 0xce, 0x11, 0xba, 0x79, 0x99,
	0xd2, 0x95, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd3, 0x61, 0x48, 0xbb, 0x41, 0x03, 0x00, 0x00,
}