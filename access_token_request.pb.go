// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: access_token_request.proto

package tokopedia

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

type AccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *AccessTokenRequest) Reset() {
	*x = AccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_access_token_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessTokenRequest) ProtoMessage() {}

func (x *AccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_access_token_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessTokenRequest.ProtoReflect.Descriptor instead.
func (*AccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_access_token_request_proto_rawDescGZIP(), []int{0}
}

func (x *AccessTokenRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *AccessTokenRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

var File_access_token_request_proto protoreflect.FileDescriptor

var file_access_token_request_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x6f,
	0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x22, 0x56, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x42,
	0x0d, 0x5a, 0x0b, 0x2e, 0x3b, 0x74, 0x6f, 0x6b, 0x6f, 0x70, 0x65, 0x64, 0x69, 0x61, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_access_token_request_proto_rawDescOnce sync.Once
	file_access_token_request_proto_rawDescData = file_access_token_request_proto_rawDesc
)

func file_access_token_request_proto_rawDescGZIP() []byte {
	file_access_token_request_proto_rawDescOnce.Do(func() {
		file_access_token_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_access_token_request_proto_rawDescData)
	})
	return file_access_token_request_proto_rawDescData
}

var file_access_token_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_access_token_request_proto_goTypes = []interface{}{
	(*AccessTokenRequest)(nil), // 0: tokopedia.AccessTokenRequest
}
var file_access_token_request_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_access_token_request_proto_init() }
func file_access_token_request_proto_init() {
	if File_access_token_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_access_token_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessTokenRequest); i {
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
			RawDescriptor: file_access_token_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_access_token_request_proto_goTypes,
		DependencyIndexes: file_access_token_request_proto_depIdxs,
		MessageInfos:      file_access_token_request_proto_msgTypes,
	}.Build()
	File_access_token_request_proto = out.File
	file_access_token_request_proto_rawDesc = nil
	file_access_token_request_proto_goTypes = nil
	file_access_token_request_proto_depIdxs = nil
}
