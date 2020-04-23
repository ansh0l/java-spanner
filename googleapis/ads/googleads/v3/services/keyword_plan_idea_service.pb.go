// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v3/services/keyword_plan_idea_service.proto

package services

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google.golang.org/genproto/googleapis/ads/googleads/v3/common"
	enums "google.golang.org/genproto/googleapis/ads/googleads/v3/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Request message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeasRequest struct {
	// The ID of the customer with the recommendation.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// Required. The resource name of the language to target.
	// Required
	Language *wrappers.StringValue `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	// The resource names of the location to target.
	// Max 10
	GeoTargetConstants []*wrappers.StringValue `protobuf:"bytes,8,rep,name=geo_target_constants,json=geoTargetConstants,proto3" json:"geo_target_constants,omitempty"`
	// Targeting network.
	KeywordPlanNetwork enums.KeywordPlanNetworkEnum_KeywordPlanNetwork `protobuf:"varint,9,opt,name=keyword_plan_network,json=keywordPlanNetwork,proto3,enum=google.ads.googleads.v3.enums.KeywordPlanNetworkEnum_KeywordPlanNetwork" json:"keyword_plan_network,omitempty"`
	// The type of seed to generate keyword ideas.
	//
	// Types that are valid to be assigned to Seed:
	//	*GenerateKeywordIdeasRequest_KeywordAndUrlSeed
	//	*GenerateKeywordIdeasRequest_KeywordSeed
	//	*GenerateKeywordIdeasRequest_UrlSeed
	Seed                 isGenerateKeywordIdeasRequest_Seed `protobuf_oneof:"seed"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *GenerateKeywordIdeasRequest) Reset()         { *m = GenerateKeywordIdeasRequest{} }
func (m *GenerateKeywordIdeasRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeasRequest) ProtoMessage()    {}
func (*GenerateKeywordIdeasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{0}
}

func (m *GenerateKeywordIdeasRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeasRequest.Merge(m, src)
}
func (m *GenerateKeywordIdeasRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeasRequest.Size(m)
}
func (m *GenerateKeywordIdeasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeasRequest proto.InternalMessageInfo

func (m *GenerateKeywordIdeasRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *GenerateKeywordIdeasRequest) GetLanguage() *wrappers.StringValue {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetGeoTargetConstants() []*wrappers.StringValue {
	if m != nil {
		return m.GeoTargetConstants
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordPlanNetwork() enums.KeywordPlanNetworkEnum_KeywordPlanNetwork {
	if m != nil {
		return m.KeywordPlanNetwork
	}
	return enums.KeywordPlanNetworkEnum_UNSPECIFIED
}

type isGenerateKeywordIdeasRequest_Seed interface {
	isGenerateKeywordIdeasRequest_Seed()
}

type GenerateKeywordIdeasRequest_KeywordAndUrlSeed struct {
	KeywordAndUrlSeed *KeywordAndUrlSeed `protobuf:"bytes,2,opt,name=keyword_and_url_seed,json=keywordAndUrlSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_KeywordSeed struct {
	KeywordSeed *KeywordSeed `protobuf:"bytes,3,opt,name=keyword_seed,json=keywordSeed,proto3,oneof"`
}

type GenerateKeywordIdeasRequest_UrlSeed struct {
	UrlSeed *UrlSeed `protobuf:"bytes,5,opt,name=url_seed,json=urlSeed,proto3,oneof"`
}

func (*GenerateKeywordIdeasRequest_KeywordAndUrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_KeywordSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (*GenerateKeywordIdeasRequest_UrlSeed) isGenerateKeywordIdeasRequest_Seed() {}

func (m *GenerateKeywordIdeasRequest) GetSeed() isGenerateKeywordIdeasRequest_Seed {
	if m != nil {
		return m.Seed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordAndUrlSeed() *KeywordAndUrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed); ok {
		return x.KeywordAndUrlSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetKeywordSeed() *KeywordSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_KeywordSeed); ok {
		return x.KeywordSeed
	}
	return nil
}

func (m *GenerateKeywordIdeasRequest) GetUrlSeed() *UrlSeed {
	if x, ok := m.GetSeed().(*GenerateKeywordIdeasRequest_UrlSeed); ok {
		return x.UrlSeed
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*GenerateKeywordIdeasRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*GenerateKeywordIdeasRequest_KeywordAndUrlSeed)(nil),
		(*GenerateKeywordIdeasRequest_KeywordSeed)(nil),
		(*GenerateKeywordIdeasRequest_UrlSeed)(nil),
	}
}

// Keyword And Url Seed
type KeywordAndUrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,2,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordAndUrlSeed) Reset()         { *m = KeywordAndUrlSeed{} }
func (m *KeywordAndUrlSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordAndUrlSeed) ProtoMessage()    {}
func (*KeywordAndUrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{1}
}

func (m *KeywordAndUrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordAndUrlSeed.Unmarshal(m, b)
}
func (m *KeywordAndUrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordAndUrlSeed.Marshal(b, m, deterministic)
}
func (m *KeywordAndUrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordAndUrlSeed.Merge(m, src)
}
func (m *KeywordAndUrlSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordAndUrlSeed.Size(m)
}
func (m *KeywordAndUrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordAndUrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordAndUrlSeed proto.InternalMessageInfo

func (m *KeywordAndUrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *KeywordAndUrlSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Keyword Seed
type KeywordSeed struct {
	// Requires at least one keyword.
	Keywords             []*wrappers.StringValue `protobuf:"bytes,1,rep,name=keywords,proto3" json:"keywords,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeywordSeed) Reset()         { *m = KeywordSeed{} }
func (m *KeywordSeed) String() string { return proto.CompactTextString(m) }
func (*KeywordSeed) ProtoMessage()    {}
func (*KeywordSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{2}
}

func (m *KeywordSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordSeed.Unmarshal(m, b)
}
func (m *KeywordSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordSeed.Marshal(b, m, deterministic)
}
func (m *KeywordSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordSeed.Merge(m, src)
}
func (m *KeywordSeed) XXX_Size() int {
	return xxx_messageInfo_KeywordSeed.Size(m)
}
func (m *KeywordSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordSeed.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordSeed proto.InternalMessageInfo

func (m *KeywordSeed) GetKeywords() []*wrappers.StringValue {
	if m != nil {
		return m.Keywords
	}
	return nil
}

// Url Seed
type UrlSeed struct {
	// The URL to crawl in order to generate keyword ideas.
	Url                  *wrappers.StringValue `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UrlSeed) Reset()         { *m = UrlSeed{} }
func (m *UrlSeed) String() string { return proto.CompactTextString(m) }
func (*UrlSeed) ProtoMessage()    {}
func (*UrlSeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{3}
}

func (m *UrlSeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UrlSeed.Unmarshal(m, b)
}
func (m *UrlSeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UrlSeed.Marshal(b, m, deterministic)
}
func (m *UrlSeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UrlSeed.Merge(m, src)
}
func (m *UrlSeed) XXX_Size() int {
	return xxx_messageInfo_UrlSeed.Size(m)
}
func (m *UrlSeed) XXX_DiscardUnknown() {
	xxx_messageInfo_UrlSeed.DiscardUnknown(m)
}

var xxx_messageInfo_UrlSeed proto.InternalMessageInfo

func (m *UrlSeed) GetUrl() *wrappers.StringValue {
	if m != nil {
		return m.Url
	}
	return nil
}

// Response message for [KeywordIdeaService.GenerateKeywordIdeas][].
type GenerateKeywordIdeaResponse struct {
	// Results of generating keyword ideas.
	Results              []*GenerateKeywordIdeaResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GenerateKeywordIdeaResponse) Reset()         { *m = GenerateKeywordIdeaResponse{} }
func (m *GenerateKeywordIdeaResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResponse) ProtoMessage()    {}
func (*GenerateKeywordIdeaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{4}
}

func (m *GenerateKeywordIdeaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResponse.Merge(m, src)
}
func (m *GenerateKeywordIdeaResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResponse.Size(m)
}
func (m *GenerateKeywordIdeaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResponse proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResponse) GetResults() []*GenerateKeywordIdeaResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result of generating keyword ideas.
type GenerateKeywordIdeaResult struct {
	// Text of the keyword idea.
	// As in Keyword Plan historical metrics, this text may not be an actual
	// keyword, but the canonical form of multiple keywords.
	// See KeywordPlanKeywordHistoricalMetrics message in KeywordPlanService.
	Text *wrappers.StringValue `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// The historical metrics for the keyword
	KeywordIdeaMetrics   *common.KeywordPlanHistoricalMetrics `protobuf:"bytes,3,opt,name=keyword_idea_metrics,json=keywordIdeaMetrics,proto3" json:"keyword_idea_metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *GenerateKeywordIdeaResult) Reset()         { *m = GenerateKeywordIdeaResult{} }
func (m *GenerateKeywordIdeaResult) String() string { return proto.CompactTextString(m) }
func (*GenerateKeywordIdeaResult) ProtoMessage()    {}
func (*GenerateKeywordIdeaResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_26269aab04477620, []int{5}
}

func (m *GenerateKeywordIdeaResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Unmarshal(m, b)
}
func (m *GenerateKeywordIdeaResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Marshal(b, m, deterministic)
}
func (m *GenerateKeywordIdeaResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateKeywordIdeaResult.Merge(m, src)
}
func (m *GenerateKeywordIdeaResult) XXX_Size() int {
	return xxx_messageInfo_GenerateKeywordIdeaResult.Size(m)
}
func (m *GenerateKeywordIdeaResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateKeywordIdeaResult.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateKeywordIdeaResult proto.InternalMessageInfo

func (m *GenerateKeywordIdeaResult) GetText() *wrappers.StringValue {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *GenerateKeywordIdeaResult) GetKeywordIdeaMetrics() *common.KeywordPlanHistoricalMetrics {
	if m != nil {
		return m.KeywordIdeaMetrics
	}
	return nil
}

func init() {
	proto.RegisterType((*GenerateKeywordIdeasRequest)(nil), "google.ads.googleads.v3.services.GenerateKeywordIdeasRequest")
	proto.RegisterType((*KeywordAndUrlSeed)(nil), "google.ads.googleads.v3.services.KeywordAndUrlSeed")
	proto.RegisterType((*KeywordSeed)(nil), "google.ads.googleads.v3.services.KeywordSeed")
	proto.RegisterType((*UrlSeed)(nil), "google.ads.googleads.v3.services.UrlSeed")
	proto.RegisterType((*GenerateKeywordIdeaResponse)(nil), "google.ads.googleads.v3.services.GenerateKeywordIdeaResponse")
	proto.RegisterType((*GenerateKeywordIdeaResult)(nil), "google.ads.googleads.v3.services.GenerateKeywordIdeaResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v3/services/keyword_plan_idea_service.proto", fileDescriptor_26269aab04477620)
}

var fileDescriptor_26269aab04477620 = []byte{
	// 771 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x2e, 0x29, 0xd7, 0x92, 0x57, 0x45, 0x01, 0x2f, 0x8c, 0x56, 0xb5, 0x8d, 0x5a, 0x10, 0x7c,
	0x70, 0x0d, 0x94, 0x2c, 0xa4, 0x43, 0x5d, 0xba, 0x02, 0x4a, 0x15, 0xad, 0x64, 0x14, 0x35, 0x0c,
	0x3a, 0xd6, 0x21, 0x10, 0x40, 0xac, 0xc9, 0x31, 0x43, 0x88, 0xda, 0x55, 0x76, 0x97, 0x72, 0x7e,
	0xe0, 0x8b, 0xaf, 0x39, 0xe6, 0x0d, 0x72, 0xcc, 0x3b, 0xe4, 0x05, 0x7c, 0x35, 0x72, 0xcf, 0x21,
	0xa7, 0x3c, 0x45, 0x40, 0x72, 0x29, 0xc9, 0xb6, 0x14, 0x39, 0xbe, 0x0d, 0x67, 0xbf, 0xf9, 0xe6,
	0xdb, 0x99, 0xd9, 0x21, 0xfa, 0x2b, 0x60, 0x2c, 0x88, 0xc0, 0x24, 0xbe, 0x30, 0x33, 0x33, 0xb1,
	0x46, 0x0d, 0x53, 0x00, 0x1f, 0x85, 0x1e, 0x08, 0xb3, 0x0f, 0xcf, 0xcf, 0x19, 0xf7, 0xdd, 0x61,
	0x44, 0xa8, 0x1b, 0xfa, 0x40, 0x5c, 0x75, 0x64, 0x0c, 0x39, 0x93, 0x0c, 0x57, 0xb3, 0x30, 0x83,
	0xf8, 0xc2, 0x18, 0x33, 0x18, 0xa3, 0x86, 0x91, 0x33, 0xac, 0xef, 0xcd, 0xcb, 0xe1, 0xb1, 0xc1,
	0x80, 0xd1, 0x9b, 0x19, 0x32, 0x5f, 0xc6, 0x3d, 0x3f, 0x12, 0x68, 0x3c, 0xb8, 0x25, 0x8d, 0x82,
	0x3c, 0x67, 0xbc, 0xaf, 0x22, 0x37, 0xf3, 0xc8, 0x61, 0x68, 0x12, 0x4a, 0x99, 0x24, 0x32, 0x64,
	0x54, 0xa8, 0xd3, 0xad, 0xa9, 0xd3, 0xb3, 0x10, 0x22, 0xdf, 0x3d, 0x85, 0x27, 0x64, 0x14, 0x32,
	0xae, 0x00, 0x3f, 0x2b, 0x40, 0xfa, 0x75, 0x1a, 0x9f, 0x99, 0xe7, 0x9c, 0x0c, 0x87, 0xc0, 0x73,
	0x82, 0x1f, 0xa7, 0x08, 0xbc, 0x28, 0x04, 0x2a, 0xb3, 0x83, 0xda, 0xfb, 0x25, 0xb4, 0xd1, 0x06,
	0x0a, 0x9c, 0x48, 0xf8, 0x2f, 0x93, 0x77, 0xe0, 0x03, 0x11, 0x0e, 0x3c, 0x8d, 0x41, 0x48, 0xbc,
	0x85, 0xca, 0x5e, 0x2c, 0x24, 0x1b, 0x00, 0x77, 0x43, 0xbf, 0xa2, 0x55, 0xb5, 0x9d, 0x15, 0x07,
	0xe5, 0xae, 0x03, 0x1f, 0x37, 0x51, 0x29, 0x22, 0x34, 0x88, 0x49, 0x00, 0x95, 0x62, 0x55, 0xdb,
	0x29, 0xd7, 0x37, 0x55, 0x59, 0x8d, 0x5c, 0x8c, 0x71, 0x2c, 0x79, 0x48, 0x83, 0x2e, 0x89, 0x62,
	0x68, 0x15, 0x3e, 0xd8, 0xba, 0x33, 0x0e, 0xc1, 0x87, 0x68, 0x2d, 0x00, 0xe6, 0x4a, 0xc2, 0x03,
	0x90, 0xae, 0xc7, 0xa8, 0x90, 0x84, 0x4a, 0x51, 0x29, 0x55, 0x0b, 0x8b, 0xa8, 0x1c, 0x1c, 0x00,
	0x7b, 0x94, 0x06, 0xfe, 0x9d, 0xc7, 0xe1, 0x17, 0x68, 0x6d, 0x56, 0x95, 0x2b, 0x2b, 0x55, 0x6d,
	0xe7, 0xfb, 0x7a, 0xc7, 0x98, 0xd7, 0xfc, 0xb4, 0x41, 0x86, 0xaa, 0xc0, 0x51, 0x44, 0xe8, 0x61,
	0x16, 0xf8, 0x0f, 0x8d, 0x07, 0x33, 0xdc, 0x0e, 0xee, 0xdf, 0xf1, 0xe1, 0xb3, 0x49, 0x6e, 0x42,
	0x7d, 0x37, 0xe6, 0x91, 0x2b, 0x00, 0xfc, 0x8a, 0x9e, 0x96, 0xa5, 0x61, 0x2c, 0x1a, 0xbc, 0x3c,
	0x8f, 0x4d, 0xfd, 0x13, 0x1e, 0x1d, 0x03, 0xf8, 0x9d, 0x6f, 0x9c, 0xd5, 0xfe, 0x6d, 0x27, 0x76,
	0xd0, 0x77, 0x79, 0x9e, 0x94, 0xbf, 0x90, 0xf2, 0xff, 0x7a, 0x6f, 0x7e, 0xc5, 0x5c, 0xee, 0x4f,
	0x3e, 0xf1, 0xbf, 0xa8, 0x34, 0xd6, 0xfb, 0x6d, 0xca, 0xf7, 0xcb, 0x62, 0xbe, 0x89, 0xca, 0x62,
	0x9c, 0x99, 0xad, 0x65, 0xb4, 0x94, 0x70, 0xd4, 0x2e, 0xd0, 0xea, 0x9d, 0xdb, 0x60, 0x03, 0x15,
	0x62, 0x1e, 0xa5, 0x43, 0xb4, 0xa8, 0xb7, 0x09, 0x10, 0xef, 0xa1, 0x92, 0xd2, 0x28, 0x2a, 0xfa,
	0x3d, 0x06, 0x62, 0x8c, 0xae, 0xb5, 0x51, 0x79, 0xea, 0xb2, 0x37, 0x88, 0xb4, 0xaf, 0x22, 0xfa,
	0x03, 0x15, 0x1f, 0xa8, 0xbe, 0x26, 0x67, 0xbe, 0x2c, 0x07, 0xc4, 0x90, 0x51, 0x01, 0xf8, 0x04,
	0x15, 0x39, 0x88, 0x38, 0x92, 0xb9, 0xa4, 0xfd, 0xc5, 0x05, 0x9f, 0xcd, 0x17, 0x47, 0xd2, 0xc9,
	0xb9, 0x6a, 0xef, 0x34, 0xf4, 0xd3, 0x5c, 0x18, 0xfe, 0x0d, 0x2d, 0x49, 0x78, 0x26, 0xd5, 0x48,
	0x7e, 0xf9, 0x12, 0x29, 0x12, 0xd3, 0xc9, 0x50, 0xa7, 0xcb, 0x74, 0x00, 0x92, 0x87, 0x9e, 0x50,
	0x43, 0xf7, 0xe7, 0x5c, 0xcd, 0x6a, 0x2f, 0x4e, 0x3d, 0x9d, 0x4e, 0x28, 0x24, 0xe3, 0xa1, 0x47,
	0xa2, 0xff, 0x33, 0x8e, 0xf1, 0x23, 0x4a, 0x04, 0x2a, 0x5f, 0xfd, 0x95, 0x8e, 0x7e, 0x98, 0x0a,
	0x4a, 0x8e, 0x8e, 0xb3, 0xeb, 0xe3, 0x6b, 0x0d, 0xad, 0xcd, 0xda, 0x55, 0xb8, 0xf9, 0xa0, 0xca,
	0xe5, 0x3b, 0x6e, 0xbd, 0xf9, 0xd0, 0xc2, 0xa7, 0x8d, 0xac, 0x35, 0x2f, 0xaf, 0x3f, 0xbe, 0xd6,
	0x7f, 0xaf, 0xd5, 0xd3, 0x5f, 0x84, 0xda, 0x8c, 0xc2, 0x7c, 0x39, 0xb5, 0x37, 0x9b, 0xbb, 0x17,
	0x56, 0x30, 0x43, 0x81, 0xa5, 0xed, 0xae, 0x6f, 0x5c, 0xd9, 0x95, 0x49, 0x52, 0x65, 0x0d, 0x43,
	0x91, 0x54, 0xb0, 0x75, 0xa9, 0xa3, 0x6d, 0x8f, 0x0d, 0x16, 0x0a, 0x6c, 0x6d, 0xcc, 0xae, 0xd9,
	0x51, 0xd2, 0xd8, 0x23, 0xed, 0x71, 0x47, 0x11, 0x04, 0x2c, 0xd9, 0xbc, 0x06, 0xe3, 0x81, 0x19,
	0x00, 0x4d, 0xdb, 0x6e, 0x4e, 0x52, 0xce, 0xff, 0xab, 0xee, 0xe7, 0xc6, 0x1b, 0xbd, 0xd0, 0xb6,
	0xed, 0xb7, 0x7a, 0xb5, 0x9d, 0x11, 0xda, 0xbe, 0x30, 0x32, 0x33, 0xb1, 0xba, 0x0d, 0x43, 0x25,
	0x16, 0x57, 0x39, 0xa4, 0x67, 0xfb, 0xa2, 0x37, 0x86, 0xf4, 0xba, 0x8d, 0x5e, 0x0e, 0xf9, 0xa4,
	0x6f, 0x67, 0x7e, 0xcb, 0xb2, 0x7d, 0x61, 0x59, 0x63, 0x90, 0x65, 0x75, 0x1b, 0x96, 0x95, 0xc3,
	0x4e, 0x97, 0x53, 0x9d, 0x8d, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0xe8, 0x42, 0x8b, 0xfc,
	0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// KeywordPlanIdeaServiceClient is the client API for KeywordPlanIdeaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeywordPlanIdeaServiceClient interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error)
}

type keywordPlanIdeaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKeywordPlanIdeaServiceClient(cc grpc.ClientConnInterface) KeywordPlanIdeaServiceClient {
	return &keywordPlanIdeaServiceClient{cc}
}

func (c *keywordPlanIdeaServiceClient) GenerateKeywordIdeas(ctx context.Context, in *GenerateKeywordIdeasRequest, opts ...grpc.CallOption) (*GenerateKeywordIdeaResponse, error) {
	out := new(GenerateKeywordIdeaResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v3.services.KeywordPlanIdeaService/GenerateKeywordIdeas", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeywordPlanIdeaServiceServer is the server API for KeywordPlanIdeaService service.
type KeywordPlanIdeaServiceServer interface {
	// Returns a list of keyword ideas.
	GenerateKeywordIdeas(context.Context, *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error)
}

// UnimplementedKeywordPlanIdeaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeywordPlanIdeaServiceServer struct {
}

func (*UnimplementedKeywordPlanIdeaServiceServer) GenerateKeywordIdeas(ctx context.Context, req *GenerateKeywordIdeasRequest) (*GenerateKeywordIdeaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateKeywordIdeas not implemented")
}

func RegisterKeywordPlanIdeaServiceServer(s *grpc.Server, srv KeywordPlanIdeaServiceServer) {
	s.RegisterService(&_KeywordPlanIdeaService_serviceDesc, srv)
}

func _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateKeywordIdeasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v3.services.KeywordPlanIdeaService/GenerateKeywordIdeas",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeywordPlanIdeaServiceServer).GenerateKeywordIdeas(ctx, req.(*GenerateKeywordIdeasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeywordPlanIdeaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v3.services.KeywordPlanIdeaService",
	HandlerType: (*KeywordPlanIdeaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateKeywordIdeas",
			Handler:    _KeywordPlanIdeaService_GenerateKeywordIdeas_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v3/services/keyword_plan_idea_service.proto",
}