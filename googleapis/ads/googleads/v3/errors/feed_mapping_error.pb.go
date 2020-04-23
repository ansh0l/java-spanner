// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/errors/feed_mapping_error.proto

package errors

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

// Enum describing possible feed item errors.
type FeedMappingErrorEnum_FeedMappingError int32

const (
	// Enum unspecified.
	FeedMappingErrorEnum_UNSPECIFIED FeedMappingErrorEnum_FeedMappingError = 0
	// The received error code is not known in this version.
	FeedMappingErrorEnum_UNKNOWN FeedMappingErrorEnum_FeedMappingError = 1
	// The given placeholder field does not exist.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_FIELD FeedMappingErrorEnum_FeedMappingError = 2
	// The given criterion field does not exist.
	FeedMappingErrorEnum_INVALID_CRITERION_FIELD FeedMappingErrorEnum_FeedMappingError = 3
	// The given placeholder type does not exist.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_TYPE FeedMappingErrorEnum_FeedMappingError = 4
	// The given criterion type does not exist.
	FeedMappingErrorEnum_INVALID_CRITERION_TYPE FeedMappingErrorEnum_FeedMappingError = 5
	// A feed mapping must contain at least one attribute field mapping.
	FeedMappingErrorEnum_NO_ATTRIBUTE_FIELD_MAPPINGS FeedMappingErrorEnum_FeedMappingError = 7
	// The type of the feed attribute referenced in the attribute field mapping
	// must match the type of the placeholder field.
	FeedMappingErrorEnum_FEED_ATTRIBUTE_TYPE_MISMATCH FeedMappingErrorEnum_FeedMappingError = 8
	// A feed mapping for a system generated feed cannot be operated on.
	FeedMappingErrorEnum_CANNOT_OPERATE_ON_MAPPINGS_FOR_SYSTEM_GENERATED_FEED FeedMappingErrorEnum_FeedMappingError = 9
	// Only one feed mapping for a placeholder type is allowed per feed or
	// customer (depending on the placeholder type).
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_TYPE FeedMappingErrorEnum_FeedMappingError = 10
	// Only one feed mapping for a criterion type is allowed per customer.
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_CRITERION_TYPE FeedMappingErrorEnum_FeedMappingError = 11
	// Only one feed attribute mapping for a placeholder field is allowed
	// (depending on the placeholder type).
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_FIELD FeedMappingErrorEnum_FeedMappingError = 12
	// Only one feed attribute mapping for a criterion field is allowed
	// (depending on the criterion type).
	FeedMappingErrorEnum_MULTIPLE_MAPPINGS_FOR_CRITERION_FIELD FeedMappingErrorEnum_FeedMappingError = 13
	// This feed mapping may not contain any explicit attribute field mappings.
	FeedMappingErrorEnum_UNEXPECTED_ATTRIBUTE_FIELD_MAPPINGS FeedMappingErrorEnum_FeedMappingError = 14
	// Location placeholder feed mappings can only be created for Places feeds.
	FeedMappingErrorEnum_LOCATION_PLACEHOLDER_ONLY_FOR_PLACES_FEEDS FeedMappingErrorEnum_FeedMappingError = 15
	// Mappings for typed feeds cannot be modified.
	FeedMappingErrorEnum_CANNOT_MODIFY_MAPPINGS_FOR_TYPED_FEED FeedMappingErrorEnum_FeedMappingError = 16
	// The given placeholder type can only be mapped to system generated feeds.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_TYPE_FOR_NON_SYSTEM_GENERATED_FEED FeedMappingErrorEnum_FeedMappingError = 17
	// The given placeholder type cannot be mapped to a system generated feed
	// with the given type.
	FeedMappingErrorEnum_INVALID_PLACEHOLDER_TYPE_FOR_SYSTEM_GENERATED_FEED_TYPE FeedMappingErrorEnum_FeedMappingError = 18
	// The "field" oneof was not set in an AttributeFieldMapping.
	FeedMappingErrorEnum_ATTRIBUTE_FIELD_MAPPING_MISSING_FIELD FeedMappingErrorEnum_FeedMappingError = 19
)

var FeedMappingErrorEnum_FeedMappingError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "INVALID_PLACEHOLDER_FIELD",
	3:  "INVALID_CRITERION_FIELD",
	4:  "INVALID_PLACEHOLDER_TYPE",
	5:  "INVALID_CRITERION_TYPE",
	7:  "NO_ATTRIBUTE_FIELD_MAPPINGS",
	8:  "FEED_ATTRIBUTE_TYPE_MISMATCH",
	9:  "CANNOT_OPERATE_ON_MAPPINGS_FOR_SYSTEM_GENERATED_FEED",
	10: "MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_TYPE",
	11: "MULTIPLE_MAPPINGS_FOR_CRITERION_TYPE",
	12: "MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_FIELD",
	13: "MULTIPLE_MAPPINGS_FOR_CRITERION_FIELD",
	14: "UNEXPECTED_ATTRIBUTE_FIELD_MAPPINGS",
	15: "LOCATION_PLACEHOLDER_ONLY_FOR_PLACES_FEEDS",
	16: "CANNOT_MODIFY_MAPPINGS_FOR_TYPED_FEED",
	17: "INVALID_PLACEHOLDER_TYPE_FOR_NON_SYSTEM_GENERATED_FEED",
	18: "INVALID_PLACEHOLDER_TYPE_FOR_SYSTEM_GENERATED_FEED_TYPE",
	19: "ATTRIBUTE_FIELD_MAPPING_MISSING_FIELD",
}

var FeedMappingErrorEnum_FeedMappingError_value = map[string]int32{
	"UNSPECIFIED":                  0,
	"UNKNOWN":                      1,
	"INVALID_PLACEHOLDER_FIELD":    2,
	"INVALID_CRITERION_FIELD":      3,
	"INVALID_PLACEHOLDER_TYPE":     4,
	"INVALID_CRITERION_TYPE":       5,
	"NO_ATTRIBUTE_FIELD_MAPPINGS":  7,
	"FEED_ATTRIBUTE_TYPE_MISMATCH": 8,
	"CANNOT_OPERATE_ON_MAPPINGS_FOR_SYSTEM_GENERATED_FEED":    9,
	"MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_TYPE":                  10,
	"MULTIPLE_MAPPINGS_FOR_CRITERION_TYPE":                    11,
	"MULTIPLE_MAPPINGS_FOR_PLACEHOLDER_FIELD":                 12,
	"MULTIPLE_MAPPINGS_FOR_CRITERION_FIELD":                   13,
	"UNEXPECTED_ATTRIBUTE_FIELD_MAPPINGS":                     14,
	"LOCATION_PLACEHOLDER_ONLY_FOR_PLACES_FEEDS":              15,
	"CANNOT_MODIFY_MAPPINGS_FOR_TYPED_FEED":                   16,
	"INVALID_PLACEHOLDER_TYPE_FOR_NON_SYSTEM_GENERATED_FEED":  17,
	"INVALID_PLACEHOLDER_TYPE_FOR_SYSTEM_GENERATED_FEED_TYPE": 18,
	"ATTRIBUTE_FIELD_MAPPING_MISSING_FIELD":                   19,
}

func (x FeedMappingErrorEnum_FeedMappingError) String() string {
	return proto.EnumName(FeedMappingErrorEnum_FeedMappingError_name, int32(x))
}

func (FeedMappingErrorEnum_FeedMappingError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a10df35593e929d8, []int{0, 0}
}

// Container for enum describing possible feed item errors.
type FeedMappingErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedMappingErrorEnum) Reset()         { *m = FeedMappingErrorEnum{} }
func (m *FeedMappingErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FeedMappingErrorEnum) ProtoMessage()    {}
func (*FeedMappingErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_a10df35593e929d8, []int{0}
}

func (m *FeedMappingErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedMappingErrorEnum.Unmarshal(m, b)
}
func (m *FeedMappingErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedMappingErrorEnum.Marshal(b, m, deterministic)
}
func (m *FeedMappingErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedMappingErrorEnum.Merge(m, src)
}
func (m *FeedMappingErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FeedMappingErrorEnum.Size(m)
}
func (m *FeedMappingErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedMappingErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedMappingErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v3.errors.FeedMappingErrorEnum_FeedMappingError", FeedMappingErrorEnum_FeedMappingError_name, FeedMappingErrorEnum_FeedMappingError_value)
	proto.RegisterType((*FeedMappingErrorEnum)(nil), "google.ads.googleads.v3.errors.FeedMappingErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/errors/feed_mapping_error.proto", fileDescriptor_a10df35593e929d8)
}

var fileDescriptor_a10df35593e929d8 = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0x9a, 0x7e, 0x6d, 0x61, 0x0a, 0xd4, 0x0c, 0xff, 0x6d, 0x29, 0x28, 0xfc, 0x14, 0x8a,
	0x64, 0x2f, 0x8c, 0x28, 0x72, 0x57, 0x13, 0x7b, 0x9c, 0x8e, 0xb0, 0x67, 0x2c, 0x7b, 0x12, 0x08,
	0x8a, 0x34, 0x32, 0xd8, 0x58, 0x91, 0x1a, 0x3b, 0x8a, 0x43, 0x1f, 0x84, 0x47, 0x60, 0xc9, 0xa3,
	0xf0, 0x28, 0xac, 0xd8, 0xb2, 0x43, 0xe3, 0x49, 0x42, 0x93, 0x26, 0xed, 0x6a, 0xae, 0xe6, 0x9e,
	0x73, 0xe6, 0xde, 0x63, 0x1f, 0x70, 0x98, 0x15, 0x45, 0x76, 0x92, 0x1a, 0x71, 0x52, 0x1a, 0xaa,
	0x94, 0xd5, 0xa9, 0x69, 0xa4, 0xc3, 0x61, 0x31, 0x2c, 0x8d, 0x2f, 0x69, 0x9a, 0x88, 0x7e, 0x3c,
	0x18, 0xf4, 0xf2, 0x4c, 0x54, 0x77, 0xfa, 0x60, 0x58, 0x8c, 0x0a, 0xb8, 0xa7, 0xd0, 0x7a, 0x9c,
	0x94, 0xfa, 0x94, 0xa8, 0x9f, 0x9a, 0xba, 0x22, 0x6e, 0xef, 0x4e, 0x84, 0x07, 0x3d, 0x23, 0xce,
	0xf3, 0x62, 0x14, 0x8f, 0x7a, 0x45, 0x5e, 0x2a, 0x76, 0xfd, 0xdb, 0x3a, 0xb8, 0xed, 0xa6, 0x69,
	0xe2, 0x2b, 0x65, 0x2c, 0x39, 0x38, 0xff, 0xda, 0xaf, 0xff, 0x59, 0x03, 0xda, 0x7c, 0x03, 0x6e,
	0x81, 0xcd, 0x16, 0x8d, 0x02, 0x6c, 0x13, 0x97, 0x60, 0x47, 0xfb, 0x0f, 0x6e, 0x82, 0x8d, 0x16,
	0x7d, 0x47, 0xd9, 0x7b, 0xaa, 0xad, 0xc0, 0x87, 0xe0, 0x01, 0xa1, 0x6d, 0xe4, 0x11, 0x47, 0x04,
	0x1e, 0xb2, 0xf1, 0x31, 0xf3, 0x1c, 0x1c, 0x0a, 0x97, 0x60, 0xcf, 0xd1, 0x6a, 0x70, 0x07, 0xdc,
	0x9b, 0xb4, 0xed, 0x90, 0x70, 0x1c, 0x12, 0x46, 0xc7, 0xcd, 0x55, 0xb8, 0x0b, 0xee, 0x2f, 0xe2,
	0xf2, 0x4e, 0x80, 0xb5, 0xff, 0xe1, 0x36, 0xb8, 0x7b, 0x9e, 0x5a, 0xf5, 0xd6, 0xe0, 0x23, 0xb0,
	0x43, 0x99, 0x40, 0x9c, 0x87, 0xa4, 0xd1, 0xe2, 0x58, 0x29, 0x0a, 0x1f, 0x05, 0x01, 0xa1, 0xcd,
	0x48, 0xdb, 0x80, 0x8f, 0xc1, 0xae, 0x8b, 0xb1, 0x73, 0x06, 0x22, 0x99, 0xc2, 0x27, 0x91, 0x8f,
	0xb8, 0x7d, 0xac, 0x5d, 0x81, 0x6f, 0xc1, 0x6b, 0x1b, 0x51, 0xca, 0xb8, 0x60, 0x01, 0x0e, 0x11,
	0xc7, 0x82, 0xd1, 0xa9, 0x82, 0x70, 0x59, 0x28, 0xa2, 0x4e, 0xc4, 0xb1, 0x2f, 0x9a, 0x98, 0x56,
	0x7d, 0x47, 0x48, 0x45, 0xed, 0x2a, 0x3c, 0x00, 0xcf, 0xfd, 0x96, 0xc7, 0x49, 0xe0, 0xe1, 0x59,
	0xc2, 0xb9, 0x25, 0x00, 0x7c, 0x01, 0x9e, 0x2e, 0xc6, 0xce, 0xad, 0xb4, 0x09, 0x5f, 0x81, 0xfd,
	0xcb, 0x55, 0x95, 0x73, 0xd7, 0xe0, 0x4b, 0xf0, 0xec, 0x32, 0x59, 0x05, 0xbd, 0x0e, 0xf7, 0xc1,
	0x93, 0x16, 0xc5, 0x1f, 0x02, 0x6c, 0xf3, 0x19, 0x3f, 0xe6, 0x2c, 0xbb, 0x01, 0x75, 0x70, 0xe0,
	0x31, 0x1b, 0x71, 0x49, 0x3e, 0xfb, 0x26, 0xa3, 0x5e, 0xe7, 0xdf, 0x20, 0x51, 0xe5, 0x42, 0xa4,
	0x6d, 0xc9, 0x19, 0xc6, 0x06, 0xfa, 0xcc, 0x21, 0x6e, 0x67, 0x76, 0x10, 0xb9, 0xd5, 0xd8, 0x31,
	0x0d, 0x5a, 0xe0, 0xcd, 0xb2, 0x0f, 0x5d, 0x81, 0x29, 0xa3, 0x4b, 0xdc, 0xbe, 0x09, 0x8f, 0xc0,
	0xe1, 0x85, 0xdc, 0x85, 0x3c, 0x65, 0x2a, 0x94, 0x33, 0x2e, 0xd9, 0x58, 0xfe, 0x0a, 0x91, 0x3c,
	0x95, 0x4f, 0xb7, 0x1a, 0xbf, 0x57, 0x40, 0xfd, 0x73, 0xd1, 0xd7, 0x2f, 0x4e, 0x56, 0xe3, 0xce,
	0x7c, 0x3e, 0x02, 0x19, 0xa9, 0x60, 0xe5, 0xa3, 0x33, 0x26, 0x66, 0xc5, 0x49, 0x9c, 0x67, 0x7a,
	0x31, 0xcc, 0x8c, 0x2c, 0xcd, 0xab, 0xc0, 0x4d, 0xb2, 0x3d, 0xe8, 0x95, 0xcb, 0xa2, 0x7e, 0xa4,
	0x8e, 0xef, 0xb5, 0xd5, 0x26, 0x42, 0x3f, 0x6a, 0x7b, 0x4d, 0x25, 0x86, 0x92, 0x52, 0x57, 0xa5,
	0xac, 0xda, 0xa6, 0x5e, 0x3d, 0x59, 0xfe, 0x9c, 0x00, 0xba, 0x28, 0x29, 0xbb, 0x53, 0x40, 0xb7,
	0x6d, 0x76, 0x15, 0xe0, 0x57, 0xad, 0xae, 0x6e, 0x2d, 0x0b, 0x25, 0xa5, 0x65, 0x4d, 0x21, 0x96,
	0xd5, 0x36, 0x2d, 0x4b, 0x81, 0x3e, 0xad, 0x57, 0xd3, 0x99, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0xa5, 0xc4, 0xae, 0xc3, 0x87, 0x04, 0x00, 0x00,
}