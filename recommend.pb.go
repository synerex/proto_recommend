// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: recommend.proto

package proto_recommend

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Recommend service message
type Recommend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecommendId    uint32           `protobuf:"varint,1,opt,name=recommend_id,json=recommendId,proto3" json:"recommend_id,omitempty"`
	RecommendName  string           `protobuf:"bytes,2,opt,name=recommend_name,json=recommendName,proto3" json:"recommend_name,omitempty"`
	RecommendSteps []*RecommendStep `protobuf:"bytes,3,rep,name=recommend_steps,json=recommendSteps,proto3" json:"recommend_steps,omitempty"`
}

func (x *Recommend) Reset() {
	*x = Recommend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommend_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Recommend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Recommend) ProtoMessage() {}

func (x *Recommend) ProtoReflect() protoreflect.Message {
	mi := &file_recommend_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Recommend.ProtoReflect.Descriptor instead.
func (*Recommend) Descriptor() ([]byte, []int) {
	return file_recommend_proto_rawDescGZIP(), []int{0}
}

func (x *Recommend) GetRecommendId() uint32 {
	if x != nil {
		return x.RecommendId
	}
	return 0
}

func (x *Recommend) GetRecommendName() string {
	if x != nil {
		return x.RecommendName
	}
	return ""
}

func (x *Recommend) GetRecommendSteps() []*RecommendStep {
	if x != nil {
		return x.RecommendSteps
	}
	return nil
}

type RecommendStep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransitOrder  uint32                 `protobuf:"varint,1,opt,name=transit_order,json=transitOrder,proto3" json:"transit_order,omitempty"`
	MobilityType  uint32                 `protobuf:"varint,2,opt,name=mobility_type,json=mobilityType,proto3" json:"mobility_type,omitempty"`
	FromStationId uint32                 `protobuf:"varint,3,opt,name=from_station_id,json=fromStationId,proto3" json:"from_station_id,omitempty"`
	ToStationId   uint32                 `protobuf:"varint,4,opt,name=to_station_id,json=toStationId,proto3" json:"to_station_id,omitempty"`
	Ts            *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=ts,proto3" json:"ts,omitempty"`
	IsAsap        bool                   `protobuf:"varint,6,opt,name=is_asap,json=isAsap,proto3" json:"is_asap,omitempty"`
}

func (x *RecommendStep) Reset() {
	*x = RecommendStep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recommend_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecommendStep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecommendStep) ProtoMessage() {}

func (x *RecommendStep) ProtoReflect() protoreflect.Message {
	mi := &file_recommend_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecommendStep.ProtoReflect.Descriptor instead.
func (*RecommendStep) Descriptor() ([]byte, []int) {
	return file_recommend_proto_rawDescGZIP(), []int{1}
}

func (x *RecommendStep) GetTransitOrder() uint32 {
	if x != nil {
		return x.TransitOrder
	}
	return 0
}

func (x *RecommendStep) GetMobilityType() uint32 {
	if x != nil {
		return x.MobilityType
	}
	return 0
}

func (x *RecommendStep) GetFromStationId() uint32 {
	if x != nil {
		return x.FromStationId
	}
	return 0
}

func (x *RecommendStep) GetToStationId() uint32 {
	if x != nil {
		return x.ToStationId
	}
	return 0
}

func (x *RecommendStep) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *RecommendStep) GetIsAsap() bool {
	if x != nil {
		return x.IsAsap
	}
	return false
}

var File_recommend_proto protoreflect.FileDescriptor

var file_recommend_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01,
	0x0a, 0x09, 0x52, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x72,
	0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x41, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x64, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x65, 0x70, 0x52, 0x0e, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x64, 0x53, 0x74, 0x65, 0x70, 0x73, 0x22, 0xea, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x65, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x69, 0x74, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x66,
	0x72, 0x6f, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0d,
	0x74, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x61, 0x73, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x73, 0x41, 0x73, 0x61, 0x70, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x6e, 0x65, 0x72, 0x65, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_recommend_proto_rawDescOnce sync.Once
	file_recommend_proto_rawDescData = file_recommend_proto_rawDesc
)

func file_recommend_proto_rawDescGZIP() []byte {
	file_recommend_proto_rawDescOnce.Do(func() {
		file_recommend_proto_rawDescData = protoimpl.X.CompressGZIP(file_recommend_proto_rawDescData)
	})
	return file_recommend_proto_rawDescData
}

var file_recommend_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_recommend_proto_goTypes = []interface{}{
	(*Recommend)(nil),             // 0: recommend.Recommend
	(*RecommendStep)(nil),         // 1: recommend.RecommendStep
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_recommend_proto_depIdxs = []int32{
	1, // 0: recommend.Recommend.recommend_steps:type_name -> recommend.RecommendStep
	2, // 1: recommend.RecommendStep.ts:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_recommend_proto_init() }
func file_recommend_proto_init() {
	if File_recommend_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_recommend_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Recommend); i {
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
		file_recommend_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecommendStep); i {
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
			RawDescriptor: file_recommend_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_recommend_proto_goTypes,
		DependencyIndexes: file_recommend_proto_depIdxs,
		MessageInfos:      file_recommend_proto_msgTypes,
	}.Build()
	File_recommend_proto = out.File
	file_recommend_proto_rawDesc = nil
	file_recommend_proto_goTypes = nil
	file_recommend_proto_depIdxs = nil
}
