// Copyright 2021 Google LLC
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/maps/fleetengine/delivery/v1/tasks.proto

package delivery

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	v1 "google.golang.org/genproto/googleapis/maps/fleetengine/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The type of a Task.
type Task_Type int32

const (
	// Default, the task type is not known.
	Task_TYPE_UNSPECIFIED Task_Type = 0
	// A pickup task is the action taken for picking up a shipment from an end
	// customer. Depot or feeder vehicle pickups should use the `SCHEDULED_STOP`
	// type.
	Task_PICKUP Task_Type = 1
	// A delivery task is the action taken for delivering a shipment to an end
	// customer. Depot or feeder vehicle dropoffs should use the
	// `SCHEDULED_STOP` type.
	Task_DELIVERY Task_Type = 2
	// A scheduled stop task is used for planning purposes. For example, it may
	// represent picking up or dropping off shipments at feeder vehicles or
	// depots. It should not be used for any shipments that are picked up or
	// dropped off from an end customer.
	Task_SCHEDULED_STOP Task_Type = 3
	// A task that indicates unavailability (e.g. driver breaks or vehicle
	// refueling).
	Task_UNAVAILABLE Task_Type = 4
)

// Enum value maps for Task_Type.
var (
	Task_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "PICKUP",
		2: "DELIVERY",
		3: "SCHEDULED_STOP",
		4: "UNAVAILABLE",
	}
	Task_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"PICKUP":           1,
		"DELIVERY":         2,
		"SCHEDULED_STOP":   3,
		"UNAVAILABLE":      4,
	}
)

func (x Task_Type) Enum() *Task_Type {
	p := new(Task_Type)
	*p = x
	return p
}

func (x Task_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_enumTypes[0].Descriptor()
}

func (Task_Type) Type() protoreflect.EnumType {
	return &file_google_maps_fleetengine_delivery_v1_tasks_proto_enumTypes[0]
}

func (x Task_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_Type.Descriptor instead.
func (Task_Type) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescGZIP(), []int{0, 0}
}

// The state of a Task indicating its progression.
type Task_State int32

const (
	// Default, used for an unspecified or unrecognized Task state.
	Task_STATE_UNSPECIFIED Task_State = 0
	// The task has not yet been assigned a delivery vehicle, or the delivery
	// vehicle has not yet passed the task's assigned vehicle stop.
	Task_OPEN Task_State = 1
	// When the vehicle this task was assigned to passes the vehicle stop of
	// this task.
	Task_CLOSED Task_State = 2
)

// Enum value maps for Task_State.
var (
	Task_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "OPEN",
		2: "CLOSED",
	}
	Task_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"OPEN":              1,
		"CLOSED":            2,
	}
)

func (x Task_State) Enum() *Task_State {
	p := new(Task_State)
	*p = x
	return p
}

func (x Task_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_State) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_enumTypes[1].Descriptor()
}

func (Task_State) Type() protoreflect.EnumType {
	return &file_google_maps_fleetengine_delivery_v1_tasks_proto_enumTypes[1]
}

func (x Task_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_State.Descriptor instead.
func (Task_State) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescGZIP(), []int{0, 1}
}

// The outcome of attempting to execute a task. When TaskState is closed,
// this indicates whether it was completed successfully or not.
type Task_TaskOutcome int32

const (
	// Task outcome before being set.
	Task_TASK_OUTCOME_UNSPECIFIED Task_TaskOutcome = 0
	// Task was completed successfully.
	Task_SUCCEEDED Task_TaskOutcome = 1
	// Task could not be completed or was cancelled.
	Task_FAILED Task_TaskOutcome = 2
)

// Enum value maps for Task_TaskOutcome.
var (
	Task_TaskOutcome_name = map[int32]string{
		0: "TASK_OUTCOME_UNSPECIFIED",
		1: "SUCCEEDED",
		2: "FAILED",
	}
	Task_TaskOutcome_value = map[string]int32{
		"TASK_OUTCOME_UNSPECIFIED": 0,
		"SUCCEEDED":                1,
		"FAILED":                   2,
	}
)

func (x Task_TaskOutcome) Enum() *Task_TaskOutcome {
	p := new(Task_TaskOutcome)
	*p = x
	return p
}

func (x Task_TaskOutcome) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_TaskOutcome) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_enumTypes[2].Descriptor()
}

func (Task_TaskOutcome) Type() protoreflect.EnumType {
	return &file_google_maps_fleetengine_delivery_v1_tasks_proto_enumTypes[2]
}

func (x Task_TaskOutcome) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_TaskOutcome.Descriptor instead.
func (Task_TaskOutcome) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescGZIP(), []int{0, 2}
}

// The identity of the source which populated the task_outcome_location.
type Task_TaskOutcomeLocationSource int32

const (
	// Task outcome before being set.
	Task_TASK_OUTCOME_LOCATION_SOURCE_UNSPECIFIED Task_TaskOutcomeLocationSource = 0
	// The provider specified the task_outcome_location.
	Task_PROVIDER Task_TaskOutcomeLocationSource = 2
	// The provider did not specify the task_outcome_location so Fleet Engine
	// used the last known vehicle location.
	Task_LAST_VEHICLE_LOCATION Task_TaskOutcomeLocationSource = 3
)

// Enum value maps for Task_TaskOutcomeLocationSource.
var (
	Task_TaskOutcomeLocationSource_name = map[int32]string{
		0: "TASK_OUTCOME_LOCATION_SOURCE_UNSPECIFIED",
		2: "PROVIDER",
		3: "LAST_VEHICLE_LOCATION",
	}
	Task_TaskOutcomeLocationSource_value = map[string]int32{
		"TASK_OUTCOME_LOCATION_SOURCE_UNSPECIFIED": 0,
		"PROVIDER":              2,
		"LAST_VEHICLE_LOCATION": 3,
	}
)

func (x Task_TaskOutcomeLocationSource) Enum() *Task_TaskOutcomeLocationSource {
	p := new(Task_TaskOutcomeLocationSource)
	*p = x
	return p
}

func (x Task_TaskOutcomeLocationSource) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Task_TaskOutcomeLocationSource) Descriptor() protoreflect.EnumDescriptor {
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_enumTypes[3].Descriptor()
}

func (Task_TaskOutcomeLocationSource) Type() protoreflect.EnumType {
	return &file_google_maps_fleetengine_delivery_v1_tasks_proto_enumTypes[3]
}

func (x Task_TaskOutcomeLocationSource) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Task_TaskOutcomeLocationSource.Descriptor instead.
func (Task_TaskOutcomeLocationSource) EnumDescriptor() ([]byte, []int) {
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescGZIP(), []int{0, 3}
}

// A task in the Delivery API represents a single action to track.
// In general there is a distinction between shipment-related tasks and break
// tasks. A shipment can have multiple tasks associated with it; for
// example, one task for the pickup and one for the dropoff or transfer.
// Different tasks for a given shipment can be handled by different vehicles;
// for example, one vehicle handles the pickup and drives the shipment to the
// hub, while another drives the same shipment from the hub to the dropoff.
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// In the format `providers/{provider_id}/tasks/{task_id}`. The task_id must
	// be a unique identifier and not a `tracking_id`. To store a `tracking_id` of
	// a shipment use the `tracking_id` field. Multiple tasks can have the same
	// `tracking_id`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. Immutable. Defines the type of the task; for example, a break or shipment.
	Type Task_Type `protobuf:"varint,2,opt,name=type,proto3,enum=maps.fleetengine.delivery.v1.Task_Type" json:"type,omitempty"`
	// Required. The current execution state of the task.
	State Task_State `protobuf:"varint,3,opt,name=state,proto3,enum=maps.fleetengine.delivery.v1.Task_State" json:"state,omitempty"`
	// The outcome of the task.
	TaskOutcome Task_TaskOutcome `protobuf:"varint,9,opt,name=task_outcome,json=taskOutcome,proto3,enum=maps.fleetengine.delivery.v1.Task_TaskOutcome" json:"task_outcome,omitempty"`
	// The timestamp of when the task's outcome was set (from provider).
	TaskOutcomeTime *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=task_outcome_time,json=taskOutcomeTime,proto3" json:"task_outcome_time,omitempty"`
	// Location where the task's outcome was set. Updated as part of UpdateTask.
	// redacted as part of SearchTasks requests. If not explicitly updated by
	// provider then Fleet Engine will populate it by default with the last known
	// vehicle location (raw location).
	TaskOutcomeLocation *LocationInfo `protobuf:"bytes,11,opt,name=task_outcome_location,json=taskOutcomeLocation,proto3" json:"task_outcome_location,omitempty"`
	// Indicates where the value of the `task_outcome_location` came from.
	TaskOutcomeLocationSource Task_TaskOutcomeLocationSource `protobuf:"varint,12,opt,name=task_outcome_location_source,json=taskOutcomeLocationSource,proto3,enum=maps.fleetengine.delivery.v1.Task_TaskOutcomeLocationSource" json:"task_outcome_location_source,omitempty"`
	// Immutable. This field facilitates the storing of an ID for the customer to avoid
	// unnecessary or complicated mapping. Cannot be set for Tasks of type
	// `UNAVAILABLE` or `SCHEDULED_STOP`. IDs are subject to the
	// following normalization and restrictions:
	//
	// 1. IDs must be valid Unicode strings.
	// 2. IDs are limited to a maximum length of 64 characters.
	// 3. IDs will be normalized according to Unicode Normalization Form C
	// (http://www.unicode.org/reports/tr15/).
	// 4. IDs may not contain any of the following ASCII characters: '/', ':',
	// '\\', '?', or '#'.
	TrackingId string `protobuf:"bytes,4,opt,name=tracking_id,json=trackingId,proto3" json:"tracking_id,omitempty"`
	// Output only. The ID of the vehicle making this Task.
	DeliveryVehicleId string `protobuf:"bytes,5,opt,name=delivery_vehicle_id,json=deliveryVehicleId,proto3" json:"delivery_vehicle_id,omitempty"`
	// Immutable. The location where the task is to be completed.
	// Optional for `UNAVAILABLE` Tasks, required for all others.
	PlannedLocation *LocationInfo `protobuf:"bytes,6,opt,name=planned_location,json=plannedLocation,proto3" json:"planned_location,omitempty"`
	// Required. Immutable. Additional time to perform an action at this location.
	TaskDuration *durationpb.Duration `protobuf:"bytes,7,opt,name=task_duration,json=taskDuration,proto3" json:"task_duration,omitempty"`
	// Output only. Journey sharing specific fields. Not populated when state is `CLOSED`.
	JourneySharingInfo *Task_JourneySharingInfo `protobuf:"bytes,8,opt,name=journey_sharing_info,json=journeySharingInfo,proto3" json:"journey_sharing_info,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_fleetengine_delivery_v1_tasks_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_fleetengine_delivery_v1_tasks_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetType() Task_Type {
	if x != nil {
		return x.Type
	}
	return Task_TYPE_UNSPECIFIED
}

func (x *Task) GetState() Task_State {
	if x != nil {
		return x.State
	}
	return Task_STATE_UNSPECIFIED
}

func (x *Task) GetTaskOutcome() Task_TaskOutcome {
	if x != nil {
		return x.TaskOutcome
	}
	return Task_TASK_OUTCOME_UNSPECIFIED
}

func (x *Task) GetTaskOutcomeTime() *timestamppb.Timestamp {
	if x != nil {
		return x.TaskOutcomeTime
	}
	return nil
}

func (x *Task) GetTaskOutcomeLocation() *LocationInfo {
	if x != nil {
		return x.TaskOutcomeLocation
	}
	return nil
}

func (x *Task) GetTaskOutcomeLocationSource() Task_TaskOutcomeLocationSource {
	if x != nil {
		return x.TaskOutcomeLocationSource
	}
	return Task_TASK_OUTCOME_LOCATION_SOURCE_UNSPECIFIED
}

func (x *Task) GetTrackingId() string {
	if x != nil {
		return x.TrackingId
	}
	return ""
}

func (x *Task) GetDeliveryVehicleId() string {
	if x != nil {
		return x.DeliveryVehicleId
	}
	return ""
}

func (x *Task) GetPlannedLocation() *LocationInfo {
	if x != nil {
		return x.PlannedLocation
	}
	return nil
}

func (x *Task) GetTaskDuration() *durationpb.Duration {
	if x != nil {
		return x.TaskDuration
	}
	return nil
}

func (x *Task) GetJourneySharingInfo() *Task_JourneySharingInfo {
	if x != nil {
		return x.JourneySharingInfo
	}
	return nil
}

// Journey sharing specific fields.
type Task_JourneySharingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Tracking information for each stop that the assigned vehicle will
	// travel to before completing this task. This list may contain stops
	// from other tasks.
	//
	// The first segment,
	// `Task.journey_sharing_info.remaining_vehicle_journey_segments[0]`,
	// contains route information from the driver's last known location to the
	// upcoming VehicleStop. This current route information usually comes from
	// the driver app except for some cases noted in the documentation for
	// `DeliveyVehicle.current_route_segment`. The other segments in
	// `Task.journey_sharing_info.remaining_vehicle_journey_segments` are
	// populated by Fleet Engine and provide route information between the
	// remaining VehicleStops.
	RemainingVehicleJourneySegments []*VehicleJourneySegment `protobuf:"bytes,1,rep,name=remaining_vehicle_journey_segments,json=remainingVehicleJourneySegments,proto3" json:"remaining_vehicle_journey_segments,omitempty"`
	// Indicates the last reported location of the assigned vehicle
	// along the route.
	LastLocation *v1.VehicleLocation `protobuf:"bytes,2,opt,name=last_location,json=lastLocation,proto3" json:"last_location,omitempty"`
	// Indicates whether the vehicle's lastLocation can be snapped to
	// the `current_route_segment`. False if `last_location` or
	// `current_route_segment` do not exist. This value is computed by
	// Fleet Engine. Any update from clients will be ignored.
	LastLocationSnappable bool `protobuf:"varint,3,opt,name=last_location_snappable,json=lastLocationSnappable,proto3" json:"last_location_snappable,omitempty"`
}

func (x *Task_JourneySharingInfo) Reset() {
	*x = Task_JourneySharingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_maps_fleetengine_delivery_v1_tasks_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task_JourneySharingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task_JourneySharingInfo) ProtoMessage() {}

func (x *Task_JourneySharingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_maps_fleetengine_delivery_v1_tasks_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task_JourneySharingInfo.ProtoReflect.Descriptor instead.
func (*Task_JourneySharingInfo) Descriptor() ([]byte, []int) {
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Task_JourneySharingInfo) GetRemainingVehicleJourneySegments() []*VehicleJourneySegment {
	if x != nil {
		return x.RemainingVehicleJourneySegments
	}
	return nil
}

func (x *Task_JourneySharingInfo) GetLastLocation() *v1.VehicleLocation {
	if x != nil {
		return x.LastLocation
	}
	return nil
}

func (x *Task_JourneySharingInfo) GetLastLocationSnappable() bool {
	if x != nil {
		return x.LastLocationSnappable
	}
	return false
}

var File_google_maps_fleetengine_delivery_v1_tasks_proto protoreflect.FileDescriptor

var file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x0c, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x27, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xe0, 0x41, 0x02,
	0xe0, 0x41, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x51,
	0x0a, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65,
	0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x12, 0x46, 0x0a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0f, 0x74, 0x61, 0x73, 0x6b, 0x4f, 0x75,
	0x74, 0x63, 0x6f, 0x6d, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x5e, 0x0a, 0x15, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e,
	0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x13, 0x74, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x7d, 0x0a, 0x1c, 0x74, 0x61, 0x73,
	0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x3c, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69,
	0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x19, 0x74,
	0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0,
	0x41, 0x05, 0x52, 0x0a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x33,
	0x0a, 0x13, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x11, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c,
	0x65, 0x49, 0x64, 0x12, 0x5a, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x0f,
	0x70, 0x6c, 0x61, 0x6e, 0x6e, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x46, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6c, 0x0a, 0x14, 0x6a, 0x6f, 0x75, 0x72, 0x6e,
	0x65, 0x79, 0x5f, 0x73, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65,
	0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x03, 0xe0, 0x41,
	0x03, 0x52, 0x12, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e,
	0x67, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x9a, 0x02, 0x0a, 0x12, 0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65,
	0x79, 0x53, 0x68, 0x61, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x80, 0x01, 0x0a,
	0x22, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x65, 0x68, 0x69, 0x63,
	0x6c, 0x65, 0x5f, 0x6a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x5f, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x6d, 0x61, 0x70, 0x73,
	0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c,
	0x69, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x1f,
	0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65,
	0x4a, 0x6f, 0x75, 0x72, 0x6e, 0x65, 0x79, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x49, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c,
	0x65, 0x65, 0x74, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x17, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6e, 0x61, 0x70,
	0x70, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x6c, 0x61, 0x73,
	0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6e, 0x61, 0x70, 0x70, 0x61, 0x62,
	0x6c, 0x65, 0x22, 0x5b, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x50, 0x49, 0x43, 0x4b, 0x55, 0x50, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x59, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x43,
	0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x03, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x22,
	0x34, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c, 0x4f,
	0x53, 0x45, 0x44, 0x10, 0x02, 0x22, 0x46, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x4f, 0x55, 0x54,
	0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45, 0x44, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x22, 0x72, 0x0a,
	0x19, 0x54, 0x61, 0x73, 0x6b, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x28, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x4f, 0x55, 0x54, 0x43, 0x4f, 0x4d, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x56,
	0x49, 0x44, 0x45, 0x52, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x41, 0x53, 0x54, 0x5f, 0x56,
	0x45, 0x48, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x03, 0x3a, 0x47, 0xea, 0x41, 0x44, 0x0a, 0x1f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x7d, 0x2f, 0x74, 0x61,
	0x73, 0x6b, 0x73, 0x2f, 0x7b, 0x74, 0x61, 0x73, 0x6b, 0x7d, 0x42, 0x81, 0x01, 0x0a, 0x23, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x6d, 0x61, 0x70, 0x73, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x42, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x6d, 0x61, 0x70, 0x73, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b,
	0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0xa2, 0x02, 0x03, 0x43, 0x46, 0x45, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescOnce sync.Once
	file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescData = file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDesc
)

func file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescGZIP() []byte {
	file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescOnce.Do(func() {
		file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescData)
	})
	return file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDescData
}

var file_google_maps_fleetengine_delivery_v1_tasks_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_google_maps_fleetengine_delivery_v1_tasks_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_maps_fleetengine_delivery_v1_tasks_proto_goTypes = []interface{}{
	(Task_Type)(0),                      // 0: maps.fleetengine.delivery.v1.Task.Type
	(Task_State)(0),                     // 1: maps.fleetengine.delivery.v1.Task.State
	(Task_TaskOutcome)(0),               // 2: maps.fleetengine.delivery.v1.Task.TaskOutcome
	(Task_TaskOutcomeLocationSource)(0), // 3: maps.fleetengine.delivery.v1.Task.TaskOutcomeLocationSource
	(*Task)(nil),                        // 4: maps.fleetengine.delivery.v1.Task
	(*Task_JourneySharingInfo)(nil),     // 5: maps.fleetengine.delivery.v1.Task.JourneySharingInfo
	(*timestamppb.Timestamp)(nil),       // 6: google.protobuf.Timestamp
	(*LocationInfo)(nil),                // 7: maps.fleetengine.delivery.v1.LocationInfo
	(*durationpb.Duration)(nil),         // 8: google.protobuf.Duration
	(*VehicleJourneySegment)(nil),       // 9: maps.fleetengine.delivery.v1.VehicleJourneySegment
	(*v1.VehicleLocation)(nil),          // 10: maps.fleetengine.v1.VehicleLocation
}
var file_google_maps_fleetengine_delivery_v1_tasks_proto_depIdxs = []int32{
	0,  // 0: maps.fleetengine.delivery.v1.Task.type:type_name -> maps.fleetengine.delivery.v1.Task.Type
	1,  // 1: maps.fleetengine.delivery.v1.Task.state:type_name -> maps.fleetengine.delivery.v1.Task.State
	2,  // 2: maps.fleetengine.delivery.v1.Task.task_outcome:type_name -> maps.fleetengine.delivery.v1.Task.TaskOutcome
	6,  // 3: maps.fleetengine.delivery.v1.Task.task_outcome_time:type_name -> google.protobuf.Timestamp
	7,  // 4: maps.fleetengine.delivery.v1.Task.task_outcome_location:type_name -> maps.fleetengine.delivery.v1.LocationInfo
	3,  // 5: maps.fleetengine.delivery.v1.Task.task_outcome_location_source:type_name -> maps.fleetengine.delivery.v1.Task.TaskOutcomeLocationSource
	7,  // 6: maps.fleetengine.delivery.v1.Task.planned_location:type_name -> maps.fleetengine.delivery.v1.LocationInfo
	8,  // 7: maps.fleetengine.delivery.v1.Task.task_duration:type_name -> google.protobuf.Duration
	5,  // 8: maps.fleetengine.delivery.v1.Task.journey_sharing_info:type_name -> maps.fleetengine.delivery.v1.Task.JourneySharingInfo
	9,  // 9: maps.fleetengine.delivery.v1.Task.JourneySharingInfo.remaining_vehicle_journey_segments:type_name -> maps.fleetengine.delivery.v1.VehicleJourneySegment
	10, // 10: maps.fleetengine.delivery.v1.Task.JourneySharingInfo.last_location:type_name -> maps.fleetengine.v1.VehicleLocation
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_google_maps_fleetengine_delivery_v1_tasks_proto_init() }
func file_google_maps_fleetengine_delivery_v1_tasks_proto_init() {
	if File_google_maps_fleetengine_delivery_v1_tasks_proto != nil {
		return
	}
	file_google_maps_fleetengine_delivery_v1_delivery_vehicles_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_maps_fleetengine_delivery_v1_tasks_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_google_maps_fleetengine_delivery_v1_tasks_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task_JourneySharingInfo); i {
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
			RawDescriptor: file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_maps_fleetengine_delivery_v1_tasks_proto_goTypes,
		DependencyIndexes: file_google_maps_fleetengine_delivery_v1_tasks_proto_depIdxs,
		EnumInfos:         file_google_maps_fleetengine_delivery_v1_tasks_proto_enumTypes,
		MessageInfos:      file_google_maps_fleetengine_delivery_v1_tasks_proto_msgTypes,
	}.Build()
	File_google_maps_fleetengine_delivery_v1_tasks_proto = out.File
	file_google_maps_fleetengine_delivery_v1_tasks_proto_rawDesc = nil
	file_google_maps_fleetengine_delivery_v1_tasks_proto_goTypes = nil
	file_google_maps_fleetengine_delivery_v1_tasks_proto_depIdxs = nil
}
