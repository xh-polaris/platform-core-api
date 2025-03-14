// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.27.3
// source: platform/core_api/data.proto

package core_api

import (
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

type ReportEventItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventName string `protobuf:"bytes,1,opt,name=eventName,proto3" form:"eventName" json:"eventName" query:"eventName"`
	Tags      string `protobuf:"bytes,2,opt,name=tags,proto3" form:"tags" json:"tags" query:"tags"`
}

func (x *ReportEventItem) Reset() {
	*x = ReportEventItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_core_api_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEventItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEventItem) ProtoMessage() {}

func (x *ReportEventItem) ProtoReflect() protoreflect.Message {
	mi := &file_platform_core_api_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEventItem.ProtoReflect.Descriptor instead.
func (*ReportEventItem) Descriptor() ([]byte, []int) {
	return file_platform_core_api_data_proto_rawDescGZIP(), []int{0}
}

func (x *ReportEventItem) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *ReportEventItem) GetTags() string {
	if x != nil {
		return x.Tags
	}
	return ""
}

type ReportEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ReportEventItem `protobuf:"bytes,1,rep,name=data,proto3" form:"data" json:"data" query:"data"`
}

func (x *ReportEventRequest) Reset() {
	*x = ReportEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_core_api_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEventRequest) ProtoMessage() {}

func (x *ReportEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_platform_core_api_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEventRequest.ProtoReflect.Descriptor instead.
func (*ReportEventRequest) Descriptor() ([]byte, []int) {
	return file_platform_core_api_data_proto_rawDescGZIP(), []int{1}
}

func (x *ReportEventRequest) GetData() []*ReportEventItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type ReportEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Done bool `protobuf:"varint,1,opt,name=done,proto3" form:"done" json:"done" query:"done"`
}

func (x *ReportEventResponse) Reset() {
	*x = ReportEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_core_api_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportEventResponse) ProtoMessage() {}

func (x *ReportEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_platform_core_api_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportEventResponse.ProtoReflect.Descriptor instead.
func (*ReportEventResponse) Descriptor() ([]byte, []int) {
	return file_platform_core_api_data_proto_rawDescGZIP(), []int{2}
}

func (x *ReportEventResponse) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

var File_platform_core_api_data_proto protoreflect.FileDescriptor

var file_platform_core_api_data_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x22, 0x43, 0x0a, 0x0f, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x4c, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0x29, 0x0a, 0x13, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x42,
	0x87, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x42, 0x0c, 0x43, 0x6f, 0x72, 0x65,
	0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69,
	0x73, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_platform_core_api_data_proto_rawDescOnce sync.Once
	file_platform_core_api_data_proto_rawDescData = file_platform_core_api_data_proto_rawDesc
)

func file_platform_core_api_data_proto_rawDescGZIP() []byte {
	file_platform_core_api_data_proto_rawDescOnce.Do(func() {
		file_platform_core_api_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_core_api_data_proto_rawDescData)
	})
	return file_platform_core_api_data_proto_rawDescData
}

var file_platform_core_api_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_platform_core_api_data_proto_goTypes = []interface{}{
	(*ReportEventItem)(nil),     // 0: platform.core_api.ReportEventItem
	(*ReportEventRequest)(nil),  // 1: platform.core_api.ReportEventRequest
	(*ReportEventResponse)(nil), // 2: platform.core_api.ReportEventResponse
}
var file_platform_core_api_data_proto_depIdxs = []int32{
	0, // 0: platform.core_api.ReportEventRequest.data:type_name -> platform.core_api.ReportEventItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func file_platform_core_api_data_proto_init() {
	if File_platform_core_api_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_core_api_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEventItem); i {
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
		file_platform_core_api_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEventRequest); i {
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
		file_platform_core_api_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportEventResponse); i {
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
			RawDescriptor: file_platform_core_api_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_core_api_data_proto_goTypes,
		DependencyIndexes: file_platform_core_api_data_proto_depIdxs,
		MessageInfos:      file_platform_core_api_data_proto_msgTypes,
	}.Build()
	File_platform_core_api_data_proto = out.File
	file_platform_core_api_data_proto_rawDesc = nil
	file_platform_core_api_data_proto_goTypes = nil
	file_platform_core_api_data_proto_depIdxs = nil
}
