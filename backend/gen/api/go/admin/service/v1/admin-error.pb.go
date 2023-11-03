// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: admin/service/v1/admin-error.proto

package v1

import (
	_ "github.com/devexps/go-micro/v2/errors"
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

type AdminErrorReason int32

const (
	AdminErrorReason_NOT_LOGGED_IN AdminErrorReason = 0 // 401 code
)

// Enum value maps for AdminErrorReason.
var (
	AdminErrorReason_name = map[int32]string{
		0: "NOT_LOGGED_IN",
	}
	AdminErrorReason_value = map[string]int32{
		"NOT_LOGGED_IN": 0,
	}
)

func (x AdminErrorReason) Enum() *AdminErrorReason {
	p := new(AdminErrorReason)
	*p = x
	return p
}

func (x AdminErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdminErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_admin_service_v1_admin_error_proto_enumTypes[0].Descriptor()
}

func (AdminErrorReason) Type() protoreflect.EnumType {
	return &file_admin_service_v1_admin_error_proto_enumTypes[0]
}

func (x AdminErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdminErrorReason.Descriptor instead.
func (AdminErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_admin_service_v1_admin_error_proto_rawDescGZIP(), []int{0}
}

var File_admin_service_v1_admin_error_proto protoreflect.FileDescriptor

var file_admin_service_v1_admin_error_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x31, 0x0a, 0x10, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x17, 0x0a, 0x0d, 0x4e, 0x4f, 0x54, 0x5f, 0x4c, 0x4f, 0x47, 0x47, 0x45, 0x44, 0x5f, 0x49, 0x4e,
	0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x46,
	0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76,
	0x65, 0x78, 0x70, 0x73, 0x2f, 0x67, 0x6f, 0x2d, 0x6d, 0x6f, 0x6e, 0x6f, 0x6c, 0x69, 0x74, 0x68,
	0x69, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6f, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_service_v1_admin_error_proto_rawDescOnce sync.Once
	file_admin_service_v1_admin_error_proto_rawDescData = file_admin_service_v1_admin_error_proto_rawDesc
)

func file_admin_service_v1_admin_error_proto_rawDescGZIP() []byte {
	file_admin_service_v1_admin_error_proto_rawDescOnce.Do(func() {
		file_admin_service_v1_admin_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_service_v1_admin_error_proto_rawDescData)
	})
	return file_admin_service_v1_admin_error_proto_rawDescData
}

var file_admin_service_v1_admin_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_admin_service_v1_admin_error_proto_goTypes = []interface{}{
	(AdminErrorReason)(0), // 0: admin.service.v1.AdminErrorReason
}
var file_admin_service_v1_admin_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_admin_service_v1_admin_error_proto_init() }
func file_admin_service_v1_admin_error_proto_init() {
	if File_admin_service_v1_admin_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_service_v1_admin_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_admin_service_v1_admin_error_proto_goTypes,
		DependencyIndexes: file_admin_service_v1_admin_error_proto_depIdxs,
		EnumInfos:         file_admin_service_v1_admin_error_proto_enumTypes,
	}.Build()
	File_admin_service_v1_admin_error_proto = out.File
	file_admin_service_v1_admin_error_proto_rawDesc = nil
	file_admin_service_v1_admin_error_proto_goTypes = nil
	file_admin_service_v1_admin_error_proto_depIdxs = nil
}
