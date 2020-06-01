// Copyright 2019 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.2
// source: google/devtools/resultstore/v2/configured_target.proto

package resultstore

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Each ConfiguredTarget represents data for a given configuration of a given
// target in a given Invocation.
// Every ConfiguredTarget should have at least one Action as a child resource
// before the invocation is finalized. Refer to the Action's documentation for
// more info on this.
type ConfiguredTarget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name.  Its format must be:
	// invocations/${INVOCATION_ID}/targets/${url_encode(TARGET_ID)}/configuredTargets/${url_encode(CONFIG_ID)}
	// where ${CONFIG_ID} must match the ID of an existing Configuration under
	// this Invocation.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The resource ID components that identify the ConfiguredTarget. They must
	// match the resource name after proper encoding.
	Id *ConfiguredTarget_Id `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The aggregate status for this configuration of this target. If testing
	// was not requested, set this to the build status (e.g. BUILT or
	// FAILED_TO_BUILD).
	StatusAttributes *StatusAttributes `protobuf:"bytes,3,opt,name=status_attributes,json=statusAttributes,proto3" json:"status_attributes,omitempty"`
	// Captures the start time and duration of this configured target.
	Timing *Timing `protobuf:"bytes,4,opt,name=timing,proto3" json:"timing,omitempty"`
	// Test specific attributes for this ConfiguredTarget.
	TestAttributes *ConfiguredTestAttributes `protobuf:"bytes,6,opt,name=test_attributes,json=testAttributes,proto3" json:"test_attributes,omitempty"`
	// Arbitrary name-value pairs.
	// This is implemented as a multi-map. Multiple properties are allowed with
	// the same key. Properties will be returned in lexicographical order by key.
	Properties []*Property `protobuf:"bytes,7,rep,name=properties,proto3" json:"properties,omitempty"`
	// A list of file references for configured target level files.
	// The file IDs must be unique within this list. Duplicate file IDs will
	// result in an error. Files will be returned in lexicographical order by ID.
	Files []*File `protobuf:"bytes,8,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ConfiguredTarget) Reset() {
	*x = ConfiguredTarget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_configured_target_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfiguredTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfiguredTarget) ProtoMessage() {}

func (x *ConfiguredTarget) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_configured_target_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfiguredTarget.ProtoReflect.Descriptor instead.
func (*ConfiguredTarget) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_configured_target_proto_rawDescGZIP(), []int{0}
}

func (x *ConfiguredTarget) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConfiguredTarget) GetId() *ConfiguredTarget_Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ConfiguredTarget) GetStatusAttributes() *StatusAttributes {
	if x != nil {
		return x.StatusAttributes
	}
	return nil
}

func (x *ConfiguredTarget) GetTiming() *Timing {
	if x != nil {
		return x.Timing
	}
	return nil
}

func (x *ConfiguredTarget) GetTestAttributes() *ConfiguredTestAttributes {
	if x != nil {
		return x.TestAttributes
	}
	return nil
}

func (x *ConfiguredTarget) GetProperties() []*Property {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *ConfiguredTarget) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

// Attributes that apply only to test actions under this configured target.
type ConfiguredTestAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total number of test runs. For example, in bazel this is specified with
	// --runs_per_test. Zero if runs_per_test is not used.
	TotalRunCount int32 `protobuf:"varint,2,opt,name=total_run_count,json=totalRunCount,proto3" json:"total_run_count,omitempty"`
	// Total number of test shards. Zero if shard count was not specified.
	TotalShardCount int32 `protobuf:"varint,3,opt,name=total_shard_count,json=totalShardCount,proto3" json:"total_shard_count,omitempty"`
	// How long test is allowed to run.
	TimeoutDuration *duration.Duration `protobuf:"bytes,5,opt,name=timeout_duration,json=timeoutDuration,proto3" json:"timeout_duration,omitempty"`
}

func (x *ConfiguredTestAttributes) Reset() {
	*x = ConfiguredTestAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_configured_target_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfiguredTestAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfiguredTestAttributes) ProtoMessage() {}

func (x *ConfiguredTestAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_configured_target_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfiguredTestAttributes.ProtoReflect.Descriptor instead.
func (*ConfiguredTestAttributes) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_configured_target_proto_rawDescGZIP(), []int{1}
}

func (x *ConfiguredTestAttributes) GetTotalRunCount() int32 {
	if x != nil {
		return x.TotalRunCount
	}
	return 0
}

func (x *ConfiguredTestAttributes) GetTotalShardCount() int32 {
	if x != nil {
		return x.TotalShardCount
	}
	return 0
}

func (x *ConfiguredTestAttributes) GetTimeoutDuration() *duration.Duration {
	if x != nil {
		return x.TimeoutDuration
	}
	return nil
}

// The resource ID components that identify the ConfiguredTarget.
type ConfiguredTarget_Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Invocation ID.
	InvocationId string `protobuf:"bytes,1,opt,name=invocation_id,json=invocationId,proto3" json:"invocation_id,omitempty"`
	// The Target ID.
	TargetId string `protobuf:"bytes,2,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	// The Configuration ID.
	ConfigurationId string `protobuf:"bytes,3,opt,name=configuration_id,json=configurationId,proto3" json:"configuration_id,omitempty"`
}

func (x *ConfiguredTarget_Id) Reset() {
	*x = ConfiguredTarget_Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_resultstore_v2_configured_target_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfiguredTarget_Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfiguredTarget_Id) ProtoMessage() {}

func (x *ConfiguredTarget_Id) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_resultstore_v2_configured_target_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfiguredTarget_Id.ProtoReflect.Descriptor instead.
func (*ConfiguredTarget_Id) Descriptor() ([]byte, []int) {
	return file_google_devtools_resultstore_v2_configured_target_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ConfiguredTarget_Id) GetInvocationId() string {
	if x != nil {
		return x.InvocationId
	}
	return ""
}

func (x *ConfiguredTarget_Id) GetTargetId() string {
	if x != nil {
		return x.TargetId
	}
	return ""
}

func (x *ConfiguredTarget_Id) GetConfigurationId() string {
	if x != nil {
		return x.ConfigurationId
	}
	return ""
}

var File_google_devtools_resultstore_v2_configured_target_proto protoreflect.FileDescriptor

var file_google_devtools_resultstore_v2_configured_target_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xe6, 0x04, 0x0a, 0x10, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65,
	0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x5d,
	0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x10, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x3e, 0x0a,
	0x06, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x54,
	0x69, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x61, 0x0a,
	0x0f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x52, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x12, 0x48, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65,
	0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x05, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x71, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x49, 0x64, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xb4, 0x01, 0x0a, 0x18, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f,
	0x72, 0x75, 0x6e, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a,
	0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x68, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x10, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x71, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_resultstore_v2_configured_target_proto_rawDescOnce sync.Once
	file_google_devtools_resultstore_v2_configured_target_proto_rawDescData = file_google_devtools_resultstore_v2_configured_target_proto_rawDesc
)

func file_google_devtools_resultstore_v2_configured_target_proto_rawDescGZIP() []byte {
	file_google_devtools_resultstore_v2_configured_target_proto_rawDescOnce.Do(func() {
		file_google_devtools_resultstore_v2_configured_target_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_resultstore_v2_configured_target_proto_rawDescData)
	})
	return file_google_devtools_resultstore_v2_configured_target_proto_rawDescData
}

var file_google_devtools_resultstore_v2_configured_target_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_devtools_resultstore_v2_configured_target_proto_goTypes = []interface{}{
	(*ConfiguredTarget)(nil),         // 0: google.devtools.resultstore.v2.ConfiguredTarget
	(*ConfiguredTestAttributes)(nil), // 1: google.devtools.resultstore.v2.ConfiguredTestAttributes
	(*ConfiguredTarget_Id)(nil),      // 2: google.devtools.resultstore.v2.ConfiguredTarget.Id
	(*StatusAttributes)(nil),         // 3: google.devtools.resultstore.v2.StatusAttributes
	(*Timing)(nil),                   // 4: google.devtools.resultstore.v2.Timing
	(*Property)(nil),                 // 5: google.devtools.resultstore.v2.Property
	(*File)(nil),                     // 6: google.devtools.resultstore.v2.File
	(*duration.Duration)(nil),        // 7: google.protobuf.Duration
}
var file_google_devtools_resultstore_v2_configured_target_proto_depIdxs = []int32{
	2, // 0: google.devtools.resultstore.v2.ConfiguredTarget.id:type_name -> google.devtools.resultstore.v2.ConfiguredTarget.Id
	3, // 1: google.devtools.resultstore.v2.ConfiguredTarget.status_attributes:type_name -> google.devtools.resultstore.v2.StatusAttributes
	4, // 2: google.devtools.resultstore.v2.ConfiguredTarget.timing:type_name -> google.devtools.resultstore.v2.Timing
	1, // 3: google.devtools.resultstore.v2.ConfiguredTarget.test_attributes:type_name -> google.devtools.resultstore.v2.ConfiguredTestAttributes
	5, // 4: google.devtools.resultstore.v2.ConfiguredTarget.properties:type_name -> google.devtools.resultstore.v2.Property
	6, // 5: google.devtools.resultstore.v2.ConfiguredTarget.files:type_name -> google.devtools.resultstore.v2.File
	7, // 6: google.devtools.resultstore.v2.ConfiguredTestAttributes.timeout_duration:type_name -> google.protobuf.Duration
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_devtools_resultstore_v2_configured_target_proto_init() }
func file_google_devtools_resultstore_v2_configured_target_proto_init() {
	if File_google_devtools_resultstore_v2_configured_target_proto != nil {
		return
	}
	file_google_devtools_resultstore_v2_common_proto_init()
	file_google_devtools_resultstore_v2_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_resultstore_v2_configured_target_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfiguredTarget); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_resultstore_v2_configured_target_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfiguredTestAttributes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_resultstore_v2_configured_target_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfiguredTarget_Id); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_resultstore_v2_configured_target_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_devtools_resultstore_v2_configured_target_proto_goTypes,
		DependencyIndexes: file_google_devtools_resultstore_v2_configured_target_proto_depIdxs,
		MessageInfos:      file_google_devtools_resultstore_v2_configured_target_proto_msgTypes,
	}.Build()
	File_google_devtools_resultstore_v2_configured_target_proto = out.File
	file_google_devtools_resultstore_v2_configured_target_proto_rawDesc = nil
	file_google_devtools_resultstore_v2_configured_target_proto_goTypes = nil
	file_google_devtools_resultstore_v2_configured_target_proto_depIdxs = nil
}
