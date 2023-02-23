// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: trace/v1/trace.proto

package tracev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Trace_Type int32

const (
	Trace_TYPE_UNSPECIFIED Trace_Type = 0
	Trace_TYPE_AIR         Trace_Type = 1
	Trace_TYPE_SEA         Trace_Type = 2
	Trace_TYPE_LAND        Trace_Type = 3
)

// Enum value maps for Trace_Type.
var (
	Trace_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "TYPE_AIR",
		2: "TYPE_SEA",
		3: "TYPE_LAND",
	}
	Trace_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"TYPE_AIR":         1,
		"TYPE_SEA":         2,
		"TYPE_LAND":        3,
	}
)

func (x Trace_Type) Enum() *Trace_Type {
	p := new(Trace_Type)
	*p = x
	return p
}

func (x Trace_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Trace_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_trace_v1_trace_proto_enumTypes[0].Descriptor()
}

func (Trace_Type) Type() protoreflect.EnumType {
	return &file_trace_v1_trace_proto_enumTypes[0]
}

func (x Trace_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Trace_Type.Descriptor instead.
func (Trace_Type) EnumDescriptor() ([]byte, []int) {
	return file_trace_v1_trace_proto_rawDescGZIP(), []int{0, 0}
}

// METADATA
type Trace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        Trace_Type      `protobuf:"varint,1,opt,name=type,proto3,enum=orbyc.trace.v1.Trace_Type" json:"type,omitempty"`
	Departure   *Trace_Location `protobuf:"bytes,2,opt,name=departure,proto3,oneof" json:"departure,omitempty"`
	Destination *Trace_Location `protobuf:"bytes,3,opt,name=destination,proto3" json:"destination,omitempty"`
	Description *string         `protobuf:"bytes,4,opt,name=description,proto3,oneof" json:"description,omitempty"`
	// Types that are assignable to Footprint:
	//
	//	*Trace_Co2E
	//	*Trace_Co2O
	Footprint isTrace_Footprint      `protobuf_oneof:"footprint"`
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	Duration  *durationpb.Duration   `protobuf:"bytes,8,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Trace) Reset() {
	*x = Trace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trace_v1_trace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_trace_v1_trace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_trace_v1_trace_proto_rawDescGZIP(), []int{0}
}

func (x *Trace) GetType() Trace_Type {
	if x != nil {
		return x.Type
	}
	return Trace_TYPE_UNSPECIFIED
}

func (x *Trace) GetDeparture() *Trace_Location {
	if x != nil {
		return x.Departure
	}
	return nil
}

func (x *Trace) GetDestination() *Trace_Location {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *Trace) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (m *Trace) GetFootprint() isTrace_Footprint {
	if m != nil {
		return m.Footprint
	}
	return nil
}

func (x *Trace) GetCo2E() uint64 {
	if x, ok := x.GetFootprint().(*Trace_Co2E); ok {
		return x.Co2E
	}
	return 0
}

func (x *Trace) GetCo2O() uint64 {
	if x, ok := x.GetFootprint().(*Trace_Co2O); ok {
		return x.Co2O
	}
	return 0
}

func (x *Trace) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Trace) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type isTrace_Footprint interface {
	isTrace_Footprint()
}

type Trace_Co2E struct {
	Co2E uint64 `protobuf:"varint,5,opt,name=co2e,proto3,oneof"` // carbon emissions
}

type Trace_Co2O struct {
	Co2O uint64 `protobuf:"varint,6,opt,name=co2o,proto3,oneof"` // carbon offset
}

func (*Trace_Co2E) isTrace_Footprint() {}

func (*Trace_Co2O) isTrace_Footprint() {}

type Trace_Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Place string  `protobuf:"bytes,1,opt,name=place,proto3" json:"place,omitempty"`
	Lat   *string `protobuf:"bytes,2,opt,name=lat,proto3,oneof" json:"lat,omitempty"`
	Lng   *string `protobuf:"bytes,3,opt,name=lng,proto3,oneof" json:"lng,omitempty"`
}

func (x *Trace_Location) Reset() {
	*x = Trace_Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trace_v1_trace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trace_Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace_Location) ProtoMessage() {}

func (x *Trace_Location) ProtoReflect() protoreflect.Message {
	mi := &file_trace_v1_trace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace_Location.ProtoReflect.Descriptor instead.
func (*Trace_Location) Descriptor() ([]byte, []int) {
	return file_trace_v1_trace_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Trace_Location) GetPlace() string {
	if x != nil {
		return x.Place
	}
	return ""
}

func (x *Trace_Location) GetLat() string {
	if x != nil && x.Lat != nil {
		return *x.Lat
	}
	return ""
}

func (x *Trace_Location) GetLng() string {
	if x != nil && x.Lng != nil {
		return *x.Lng
	}
	return ""
}

var File_trace_v1_trace_proto protoreflect.FileDescriptor

var file_trace_v1_trace_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x72, 0x62, 0x79, 0x63, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x04, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x63,
	0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x6f, 0x72, 0x62, 0x79, 0x63, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x41, 0x0a, 0x09, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x62, 0x79, 0x63, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x01, 0x52, 0x09, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x40, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x72, 0x62, 0x79,
	0x63, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x14, 0x0a,
	0x04, 0x63, 0x6f, 0x32, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x04, 0x63,
	0x6f, 0x32, 0x65, 0x12, 0x14, 0x0a, 0x04, 0x63, 0x6f, 0x32, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x32, 0x6f, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x5e, 0x0a, 0x08, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x15, 0x0a,
	0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6c, 0x61,
	0x74, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03, 0x6c, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x01, 0x52, 0x03, 0x6c, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f,
	0x6c, 0x61, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6c, 0x6e, 0x67, 0x22, 0x47, 0x0a, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x41, 0x49, 0x52, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x45, 0x41, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x41,
	0x4e, 0x44, 0x10, 0x03, 0x42, 0x0b, 0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x74, 0x70, 0x72, 0x69, 0x6e,
	0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x75, 0x72, 0x65, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0xa9, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6f, 0x72, 0x62, 0x79, 0x63, 0x2e, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x54, 0x72, 0x61, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6f, 0x72, 0x62, 0x79, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x72, 0x61, 0x63,
	0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x4f, 0x54, 0x58, 0xaa, 0x02, 0x0e, 0x4f, 0x72, 0x62, 0x79,
	0x63, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x4f, 0x72, 0x62,
	0x79, 0x63, 0x5c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x4f, 0x72,
	0x62, 0x79, 0x63, 0x5c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x10, 0x4f, 0x72, 0x62, 0x79, 0x63,
	0x3a, 0x3a, 0x54, 0x72, 0x61, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_trace_v1_trace_proto_rawDescOnce sync.Once
	file_trace_v1_trace_proto_rawDescData = file_trace_v1_trace_proto_rawDesc
)

func file_trace_v1_trace_proto_rawDescGZIP() []byte {
	file_trace_v1_trace_proto_rawDescOnce.Do(func() {
		file_trace_v1_trace_proto_rawDescData = protoimpl.X.CompressGZIP(file_trace_v1_trace_proto_rawDescData)
	})
	return file_trace_v1_trace_proto_rawDescData
}

var file_trace_v1_trace_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_trace_v1_trace_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_trace_v1_trace_proto_goTypes = []interface{}{
	(Trace_Type)(0),               // 0: orbyc.trace.v1.Trace.Type
	(*Trace)(nil),                 // 1: orbyc.trace.v1.Trace
	(*Trace_Location)(nil),        // 2: orbyc.trace.v1.Trace.Location
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 4: google.protobuf.Duration
}
var file_trace_v1_trace_proto_depIdxs = []int32{
	0, // 0: orbyc.trace.v1.Trace.type:type_name -> orbyc.trace.v1.Trace.Type
	2, // 1: orbyc.trace.v1.Trace.departure:type_name -> orbyc.trace.v1.Trace.Location
	2, // 2: orbyc.trace.v1.Trace.destination:type_name -> orbyc.trace.v1.Trace.Location
	3, // 3: orbyc.trace.v1.Trace.started_at:type_name -> google.protobuf.Timestamp
	4, // 4: orbyc.trace.v1.Trace.duration:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_trace_v1_trace_proto_init() }
func file_trace_v1_trace_proto_init() {
	if File_trace_v1_trace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trace_v1_trace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trace); i {
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
		file_trace_v1_trace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trace_Location); i {
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
	file_trace_v1_trace_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Trace_Co2E)(nil),
		(*Trace_Co2O)(nil),
	}
	file_trace_v1_trace_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_trace_v1_trace_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_trace_v1_trace_proto_goTypes,
		DependencyIndexes: file_trace_v1_trace_proto_depIdxs,
		EnumInfos:         file_trace_v1_trace_proto_enumTypes,
		MessageInfos:      file_trace_v1_trace_proto_msgTypes,
	}.Build()
	File_trace_v1_trace_proto = out.File
	file_trace_v1_trace_proto_rawDesc = nil
	file_trace_v1_trace_proto_goTypes = nil
	file_trace_v1_trace_proto_depIdxs = nil
}
