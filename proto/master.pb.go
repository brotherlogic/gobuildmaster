// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        (unknown)
// source: master.proto

package gobuildmaster

import (
	context "context"
	proto1 "github.com/brotherlogic/gobuildslave/proto"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type Redundancy int32

const (
	Redundancy_UNKNOWN   Redundancy = 0
	Redundancy_SOLO      Redundancy = 1
	Redundancy_REDUNDANT Redundancy = 2
	Redundancy_GLOBAL    Redundancy = 3
)

// Enum value maps for Redundancy.
var (
	Redundancy_name = map[int32]string{
		0: "UNKNOWN",
		1: "SOLO",
		2: "REDUNDANT",
		3: "GLOBAL",
	}
	Redundancy_value = map[string]int32{
		"UNKNOWN":   0,
		"SOLO":      1,
		"REDUNDANT": 2,
		"GLOBAL":    3,
	}
)

func (x Redundancy) Enum() *Redundancy {
	p := new(Redundancy)
	*p = x
	return p
}

func (x Redundancy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Redundancy) Descriptor() protoreflect.EnumDescriptor {
	return file_master_proto_enumTypes[0].Descriptor()
}

func (Redundancy) Type() protoreflect.EnumType {
	return &file_master_proto_enumTypes[0]
}

func (x Redundancy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Redundancy.Descriptor instead.
func (Redundancy) EnumDescriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{0}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Intents  []*Intent  `protobuf:"bytes,1,rep,name=intents,proto3" json:"intents,omitempty"`
	Nintents []*NIntent `protobuf:"bytes,2,rep,name=nintents,proto3" json:"nintents,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetIntents() []*Intent {
	if x != nil {
		return x.Intents
	}
	return nil
}

func (x *Config) GetNintents() []*NIntent {
	if x != nil {
		return x.Nintents
	}
	return nil
}

type Intent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec  *proto1.JobSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	Count int32           `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Intent) Reset() {
	*x = Intent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Intent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Intent) ProtoMessage() {}

func (x *Intent) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Intent.ProtoReflect.Descriptor instead.
func (*Intent) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{1}
}

func (x *Intent) GetSpec() *proto1.JobSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Intent) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type NIntent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Job        *proto1.Job `protobuf:"bytes,1,opt,name=job,proto3" json:"job,omitempty"`
	Count      int32       `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Redundancy Redundancy  `protobuf:"varint,3,opt,name=redundancy,proto3,enum=gobuildmaster.Redundancy" json:"redundancy,omitempty"`
	NoMaster   bool        `protobuf:"varint,4,opt,name=no_master,json=noMaster,proto3" json:"no_master,omitempty"`
}

func (x *NIntent) Reset() {
	*x = NIntent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NIntent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NIntent) ProtoMessage() {}

func (x *NIntent) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NIntent.ProtoReflect.Descriptor instead.
func (*NIntent) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{2}
}

func (x *NIntent) GetJob() *proto1.Job {
	if x != nil {
		return x.Job
	}
	return nil
}

func (x *NIntent) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *NIntent) GetRedundancy() Redundancy {
	if x != nil {
		return x.Redundancy
	}
	return Redundancy_UNKNOWN
}

func (x *NIntent) GetNoMaster() bool {
	if x != nil {
		return x.NoMaster
	}
	return false
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{3}
}

type CompareResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current *Config `protobuf:"bytes,1,opt,name=current,proto3" json:"current,omitempty"`
	Desired *Config `protobuf:"bytes,2,opt,name=desired,proto3" json:"desired,omitempty"`
}

func (x *CompareResponse) Reset() {
	*x = CompareResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_master_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareResponse) ProtoMessage() {}

func (x *CompareResponse) ProtoReflect() protoreflect.Message {
	mi := &file_master_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareResponse.ProtoReflect.Descriptor instead.
func (*CompareResponse) Descriptor() ([]byte, []int) {
	return file_master_proto_rawDescGZIP(), []int{4}
}

func (x *CompareResponse) GetCurrent() *Config {
	if x != nil {
		return x.Current
	}
	return nil
}

func (x *CompareResponse) GetDesired() *Config {
	if x != nil {
		return x.Desired
	}
	return nil
}

var File_master_proto protoreflect.FileDescriptor

var file_master_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d,
	0x67, 0x6f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x36, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x72, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x2f, 0x67, 0x6f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x6c,
	0x61, 0x76, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x2f, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x67, 0x6f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x32, 0x0a, 0x08, 0x6e, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x4e, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x6e, 0x69, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x49, 0x0a, 0x06, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x29,
	0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x6f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x9c, 0x01, 0x0a, 0x07, 0x4e, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x6a,
	0x6f, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x67, 0x6f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x73, 0x6c, 0x61, 0x76, 0x65, 0x2e, 0x4a, 0x6f, 0x62, 0x52, 0x03, 0x6a, 0x6f, 0x62,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x72, 0x65, 0x64, 0x75, 0x6e, 0x64,
	0x61, 0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x64, 0x75, 0x6e,
	0x64, 0x61, 0x6e, 0x63, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x64, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x63,
	0x79, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6e, 0x6f, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x22, 0x07,
	0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x73, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x07, 0x64,
	0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x67,
	0x6f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x07, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x2a, 0x3e, 0x0a, 0x0a,
	0x52, 0x65, 0x64, 0x75, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x79, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x4f, 0x4c, 0x4f, 0x10,
	0x01, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x45, 0x44, 0x55, 0x4e, 0x44, 0x41, 0x4e, 0x54, 0x10, 0x02,
	0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4c, 0x4f, 0x42, 0x41, 0x4c, 0x10, 0x03, 0x32, 0x52, 0x0a, 0x0d,
	0x47, 0x6f, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x41, 0x0a,
	0x07, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e,
	0x2e, 0x67, 0x6f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x3b, 0x67, 0x6f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_master_proto_rawDescOnce sync.Once
	file_master_proto_rawDescData = file_master_proto_rawDesc
)

func file_master_proto_rawDescGZIP() []byte {
	file_master_proto_rawDescOnce.Do(func() {
		file_master_proto_rawDescData = protoimpl.X.CompressGZIP(file_master_proto_rawDescData)
	})
	return file_master_proto_rawDescData
}

var file_master_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_master_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_master_proto_goTypes = []interface{}{
	(Redundancy)(0),         // 0: gobuildmaster.Redundancy
	(*Config)(nil),          // 1: gobuildmaster.Config
	(*Intent)(nil),          // 2: gobuildmaster.Intent
	(*NIntent)(nil),         // 3: gobuildmaster.NIntent
	(*Empty)(nil),           // 4: gobuildmaster.Empty
	(*CompareResponse)(nil), // 5: gobuildmaster.CompareResponse
	(*proto1.JobSpec)(nil),  // 6: gobuildslave.JobSpec
	(*proto1.Job)(nil),      // 7: gobuildslave.Job
}
var file_master_proto_depIdxs = []int32{
	2, // 0: gobuildmaster.Config.intents:type_name -> gobuildmaster.Intent
	3, // 1: gobuildmaster.Config.nintents:type_name -> gobuildmaster.NIntent
	6, // 2: gobuildmaster.Intent.spec:type_name -> gobuildslave.JobSpec
	7, // 3: gobuildmaster.NIntent.job:type_name -> gobuildslave.Job
	0, // 4: gobuildmaster.NIntent.redundancy:type_name -> gobuildmaster.Redundancy
	1, // 5: gobuildmaster.CompareResponse.current:type_name -> gobuildmaster.Config
	1, // 6: gobuildmaster.CompareResponse.desired:type_name -> gobuildmaster.Config
	4, // 7: gobuildmaster.GoBuildMaster.Compare:input_type -> gobuildmaster.Empty
	5, // 8: gobuildmaster.GoBuildMaster.Compare:output_type -> gobuildmaster.CompareResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_master_proto_init() }
func file_master_proto_init() {
	if File_master_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_master_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_master_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Intent); i {
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
		file_master_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NIntent); i {
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
		file_master_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_master_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareResponse); i {
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
			RawDescriptor: file_master_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_master_proto_goTypes,
		DependencyIndexes: file_master_proto_depIdxs,
		EnumInfos:         file_master_proto_enumTypes,
		MessageInfos:      file_master_proto_msgTypes,
	}.Build()
	File_master_proto = out.File
	file_master_proto_rawDesc = nil
	file_master_proto_goTypes = nil
	file_master_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GoBuildMasterClient is the client API for GoBuildMaster service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GoBuildMasterClient interface {
	Compare(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CompareResponse, error)
}

type goBuildMasterClient struct {
	cc grpc.ClientConnInterface
}

func NewGoBuildMasterClient(cc grpc.ClientConnInterface) GoBuildMasterClient {
	return &goBuildMasterClient{cc}
}

func (c *goBuildMasterClient) Compare(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CompareResponse, error) {
	out := new(CompareResponse)
	err := c.cc.Invoke(ctx, "/gobuildmaster.GoBuildMaster/Compare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoBuildMasterServer is the server API for GoBuildMaster service.
type GoBuildMasterServer interface {
	Compare(context.Context, *Empty) (*CompareResponse, error)
}

// UnimplementedGoBuildMasterServer can be embedded to have forward compatible implementations.
type UnimplementedGoBuildMasterServer struct {
}

func (*UnimplementedGoBuildMasterServer) Compare(context.Context, *Empty) (*CompareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Compare not implemented")
}

func RegisterGoBuildMasterServer(s *grpc.Server, srv GoBuildMasterServer) {
	s.RegisterService(&_GoBuildMaster_serviceDesc, srv)
}

func _GoBuildMaster_Compare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoBuildMasterServer).Compare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gobuildmaster.GoBuildMaster/Compare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoBuildMasterServer).Compare(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoBuildMaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gobuildmaster.GoBuildMaster",
	HandlerType: (*GoBuildMasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Compare",
			Handler:    _GoBuildMaster_Compare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}
