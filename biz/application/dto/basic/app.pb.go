// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.27.3
// source: basic/app.proto

package basic

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

type APP int32

const (
	APP_Unknown         APP = 0
	APP_Meowchat        APP = 1
	APP_MeowchatManager APP = 2
	APP_Meowcloud       APP = 3
)

// Enum value maps for APP.
var (
	APP_name = map[int32]string{
		0: "Unknown",
		1: "Meowchat",
		2: "MeowchatManager",
		3: "Meowcloud",
	}
	APP_value = map[string]int32{
		"Unknown":         0,
		"Meowchat":        1,
		"MeowchatManager": 2,
		"Meowcloud":       3,
	}
)

func (x APP) Enum() *APP {
	p := new(APP)
	*p = x
	return p
}

func (x APP) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (APP) Descriptor() protoreflect.EnumDescriptor {
	return file_basic_app_proto_enumTypes[0].Descriptor()
}

func (APP) Type() protoreflect.EnumType {
	return &file_basic_app_proto_enumTypes[0]
}

func (x APP) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use APP.Descriptor instead.
func (APP) EnumDescriptor() ([]byte, []int) {
	return file_basic_app_proto_rawDescGZIP(), []int{0}
}

var File_basic_app_proto protoreflect.FileDescriptor

var file_basic_app_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2a, 0x44, 0x0a, 0x03, 0x41, 0x50, 0x50, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x4d, 0x65, 0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x65,
	0x6f, 0x77, 0x63, 0x68, 0x61, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x4d, 0x65, 0x6f, 0x77, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x10, 0x03, 0x42, 0x55,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2e,
	0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x42, 0x08, 0x41, 0x70,
	0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_basic_app_proto_rawDescOnce sync.Once
	file_basic_app_proto_rawDescData = file_basic_app_proto_rawDesc
)

func file_basic_app_proto_rawDescGZIP() []byte {
	file_basic_app_proto_rawDescOnce.Do(func() {
		file_basic_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_basic_app_proto_rawDescData)
	})
	return file_basic_app_proto_rawDescData
}

var file_basic_app_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_basic_app_proto_goTypes = []interface{}{
	(APP)(0), // 0: basic.APP
}
var file_basic_app_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_basic_app_proto_init() }
func file_basic_app_proto_init() {
	if File_basic_app_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_basic_app_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_basic_app_proto_goTypes,
		DependencyIndexes: file_basic_app_proto_depIdxs,
		EnumInfos:         file_basic_app_proto_enumTypes,
	}.Build()
	File_basic_app_proto = out.File
	file_basic_app_proto_rawDesc = nil
	file_basic_app_proto_goTypes = nil
	file_basic_app_proto_depIdxs = nil
}