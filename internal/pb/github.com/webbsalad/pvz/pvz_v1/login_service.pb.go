// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: api/pvz/login_service.proto

package pvz_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_pvz_login_service_proto protoreflect.FileDescriptor

const file_api_pvz_login_service_proto_rawDesc = "" +
	"\n" +
	"\x1bapi/pvz/login_service.proto\x12\x06pvz.v12\x0e\n" +
	"\fLoginServiceB(Z&github.com/webbsalad/pvz/pvz_v1;pvz_v1b\x06proto3"

var file_api_pvz_login_service_proto_goTypes = []any{}
var file_api_pvz_login_service_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_pvz_login_service_proto_init() }
func file_api_pvz_login_service_proto_init() {
	if File_api_pvz_login_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_pvz_login_service_proto_rawDesc), len(file_api_pvz_login_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_pvz_login_service_proto_goTypes,
		DependencyIndexes: file_api_pvz_login_service_proto_depIdxs,
	}.Build()
	File_api_pvz_login_service_proto = out.File
	file_api_pvz_login_service_proto_goTypes = nil
	file_api_pvz_login_service_proto_depIdxs = nil
}
