// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: platform/core_api/auth.proto

package core_api

import (
	basic "github.com/xh-polaris/platform-core-api/biz/application/dto/basic"
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

type SignInReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthType   string    `protobuf:"bytes,1,opt,name=authType,proto3" form:"authType" json:"authType" query:"authType"`
	AuthId     string    `protobuf:"bytes,2,opt,name=authId,proto3" form:"authId" json:"authId" query:"authId"` // authType为邮件时填邮箱，为电话时填电话号，为微信时填APP_ID
	Password   *string   `protobuf:"bytes,3,opt,name=password,proto3,oneof" form:"password" json:"password" query:"password"`
	VerifyCode *string   `protobuf:"bytes,4,opt,name=verifyCode,proto3,oneof" form:"verifyCode" json:"verifyCode" query:"verifyCode"`
	AppId      basic.APP `protobuf:"varint,5,opt,name=appId,proto3,enum=basic.APP" form:"appId" json:"appId" query:"appId"`
	DeviceId   string    `protobuf:"bytes,6,opt,name=deviceId,proto3" form:"deviceId" json:"deviceId" query:"deviceId"`
}

func (x *SignInReq) Reset() {
	*x = SignInReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_core_api_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInReq) ProtoMessage() {}

func (x *SignInReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_core_api_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInReq.ProtoReflect.Descriptor instead.
func (*SignInReq) Descriptor() ([]byte, []int) {
	return file_platform_core_api_auth_proto_rawDescGZIP(), []int{0}
}

func (x *SignInReq) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *SignInReq) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

func (x *SignInReq) GetPassword() string {
	if x != nil && x.Password != nil {
		return *x.Password
	}
	return ""
}

func (x *SignInReq) GetVerifyCode() string {
	if x != nil && x.VerifyCode != nil {
		return *x.VerifyCode
	}
	return ""
}

func (x *SignInReq) GetAppId() basic.APP {
	if x != nil {
		return x.AppId
	}
	return basic.APP(0)
}

func (x *SignInReq) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type SignInResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string `protobuf:"bytes,1,opt,name=userId,proto3" form:"userId" json:"userId" query:"userId"`
	AccessToken  string `protobuf:"bytes,2,opt,name=accessToken,proto3" form:"accessToken" json:"accessToken" query:"accessToken"`
	AccessExpire int64  `protobuf:"varint,3,opt,name=accessExpire,proto3" form:"accessExpire" json:"accessExpire" query:"accessExpire"`
}

func (x *SignInResp) Reset() {
	*x = SignInResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_core_api_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInResp) ProtoMessage() {}

func (x *SignInResp) ProtoReflect() protoreflect.Message {
	mi := &file_platform_core_api_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInResp.ProtoReflect.Descriptor instead.
func (*SignInResp) Descriptor() ([]byte, []int) {
	return file_platform_core_api_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SignInResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SignInResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *SignInResp) GetAccessExpire() int64 {
	if x != nil {
		return x.AccessExpire
	}
	return 0
}

type SetPasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" form:"password" json:"password" query:"password"`
}

func (x *SetPasswordReq) Reset() {
	*x = SetPasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_core_api_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordReq) ProtoMessage() {}

func (x *SetPasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_core_api_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordReq.ProtoReflect.Descriptor instead.
func (*SetPasswordReq) Descriptor() ([]byte, []int) {
	return file_platform_core_api_auth_proto_rawDescGZIP(), []int{2}
}

func (x *SetPasswordReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SetPasswordResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetPasswordResp) Reset() {
	*x = SetPasswordResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_core_api_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetPasswordResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetPasswordResp) ProtoMessage() {}

func (x *SetPasswordResp) ProtoReflect() protoreflect.Message {
	mi := &file_platform_core_api_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetPasswordResp.ProtoReflect.Descriptor instead.
func (*SetPasswordResp) Descriptor() ([]byte, []int) {
	return file_platform_core_api_auth_proto_rawDescGZIP(), []int{3}
}

type SendVerifyCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthType string `protobuf:"bytes,1,opt,name=authType,proto3" form:"authType" json:"authType" query:"authType"`
	AuthId   string `protobuf:"bytes,2,opt,name=authId,proto3" form:"authId" json:"authId" query:"authId"` // authType为邮件时填邮箱，为电话时填电话号
}

func (x *SendVerifyCodeReq) Reset() {
	*x = SendVerifyCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_core_api_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerifyCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerifyCodeReq) ProtoMessage() {}

func (x *SendVerifyCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_platform_core_api_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerifyCodeReq.ProtoReflect.Descriptor instead.
func (*SendVerifyCodeReq) Descriptor() ([]byte, []int) {
	return file_platform_core_api_auth_proto_rawDescGZIP(), []int{4}
}

func (x *SendVerifyCodeReq) GetAuthType() string {
	if x != nil {
		return x.AuthType
	}
	return ""
}

func (x *SendVerifyCodeReq) GetAuthId() string {
	if x != nil {
		return x.AuthId
	}
	return ""
}

type SendVerifyCodeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendVerifyCodeResp) Reset() {
	*x = SendVerifyCodeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_platform_core_api_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerifyCodeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerifyCodeResp) ProtoMessage() {}

func (x *SendVerifyCodeResp) ProtoReflect() protoreflect.Message {
	mi := &file_platform_core_api_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerifyCodeResp.ProtoReflect.Descriptor instead.
func (*SendVerifyCodeResp) Descriptor() ([]byte, []int) {
	return file_platform_core_api_auth_proto_rawDescGZIP(), []int{5}
}

var File_platform_core_api_auth_proto protoreflect.FileDescriptor

var file_platform_core_api_auth_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x1a, 0x0f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x05, 0x61, 0x70,
	0x70, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x41, 0x50, 0x50, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0x6a, 0x0a, 0x0a, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x22, 0x2c, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x11,
	0x0a, 0x0f, 0x53, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x47, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x49, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x65,
	0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x42, 0x84, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61, 0x72,
	0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x42, 0x09, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61, 0x72, 0x69, 0x73, 0x2f,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_platform_core_api_auth_proto_rawDescOnce sync.Once
	file_platform_core_api_auth_proto_rawDescData = file_platform_core_api_auth_proto_rawDesc
)

func file_platform_core_api_auth_proto_rawDescGZIP() []byte {
	file_platform_core_api_auth_proto_rawDescOnce.Do(func() {
		file_platform_core_api_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_platform_core_api_auth_proto_rawDescData)
	})
	return file_platform_core_api_auth_proto_rawDescData
}

var file_platform_core_api_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_platform_core_api_auth_proto_goTypes = []interface{}{
	(*SignInReq)(nil),          // 0: platform.core_api.SignInReq
	(*SignInResp)(nil),         // 1: platform.core_api.SignInResp
	(*SetPasswordReq)(nil),     // 2: platform.core_api.SetPasswordReq
	(*SetPasswordResp)(nil),    // 3: platform.core_api.SetPasswordResp
	(*SendVerifyCodeReq)(nil),  // 4: platform.core_api.SendVerifyCodeReq
	(*SendVerifyCodeResp)(nil), // 5: platform.core_api.SendVerifyCodeResp
	(basic.APP)(0),             // 6: basic.APP
}
var file_platform_core_api_auth_proto_depIdxs = []int32{
	6, // 0: platform.core_api.SignInReq.appId:type_name -> basic.APP
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func file_platform_core_api_auth_proto_init() {
	if File_platform_core_api_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_platform_core_api_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInReq); i {
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
		file_platform_core_api_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInResp); i {
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
		file_platform_core_api_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordReq); i {
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
		file_platform_core_api_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetPasswordResp); i {
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
		file_platform_core_api_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerifyCodeReq); i {
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
		file_platform_core_api_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerifyCodeResp); i {
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
	file_platform_core_api_auth_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_platform_core_api_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_platform_core_api_auth_proto_goTypes,
		DependencyIndexes: file_platform_core_api_auth_proto_depIdxs,
		MessageInfos:      file_platform_core_api_auth_proto_msgTypes,
	}.Build()
	File_platform_core_api_auth_proto = out.File
	file_platform_core_api_auth_proto_rawDesc = nil
	file_platform_core_api_auth_proto_goTypes = nil
	file_platform_core_api_auth_proto_depIdxs = nil
}
