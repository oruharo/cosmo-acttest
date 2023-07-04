//
//Cosmo Dashboard API
//Manipulate cosmo dashboard resource API

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: dashboard/v1alpha1/template_service.proto

package dashboardv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserAddonTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Items   []*Template `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetUserAddonTemplatesResponse) Reset() {
	*x = GetUserAddonTemplatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1alpha1_template_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserAddonTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserAddonTemplatesResponse) ProtoMessage() {}

func (x *GetUserAddonTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1alpha1_template_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserAddonTemplatesResponse.ProtoReflect.Descriptor instead.
func (*GetUserAddonTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_v1alpha1_template_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserAddonTemplatesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetUserAddonTemplatesResponse) GetItems() []*Template {
	if x != nil {
		return x.Items
	}
	return nil
}

type GetWorkspaceTemplatesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string      `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Items   []*Template `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetWorkspaceTemplatesResponse) Reset() {
	*x = GetWorkspaceTemplatesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dashboard_v1alpha1_template_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWorkspaceTemplatesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWorkspaceTemplatesResponse) ProtoMessage() {}

func (x *GetWorkspaceTemplatesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dashboard_v1alpha1_template_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWorkspaceTemplatesResponse.ProtoReflect.Descriptor instead.
func (*GetWorkspaceTemplatesResponse) Descriptor() ([]byte, []int) {
	return file_dashboard_v1alpha1_template_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetWorkspaceTemplatesResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetWorkspaceTemplatesResponse) GetItems() []*Template {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_dashboard_v1alpha1_template_service_proto protoreflect.FileDescriptor

var file_dashboard_v1alpha1_template_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x64, 0x61, 0x73,
	0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x64, 0x61,
	0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6d, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x6f, 0x6e, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x6d,
	0x0a, 0x1d, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xd9, 0x01,
	0x0a, 0x0f, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x62, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x64, 0x64, 0x6f,
	0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x31, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x6f, 0x6e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x31, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xe8, 0x01, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x14, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2d, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x64, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x44, 0x58, 0x58, 0xaa, 0x02, 0x12, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x12, 0x44, 0x61, 0x73, 0x68, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1e,
	0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x13, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dashboard_v1alpha1_template_service_proto_rawDescOnce sync.Once
	file_dashboard_v1alpha1_template_service_proto_rawDescData = file_dashboard_v1alpha1_template_service_proto_rawDesc
)

func file_dashboard_v1alpha1_template_service_proto_rawDescGZIP() []byte {
	file_dashboard_v1alpha1_template_service_proto_rawDescOnce.Do(func() {
		file_dashboard_v1alpha1_template_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_dashboard_v1alpha1_template_service_proto_rawDescData)
	})
	return file_dashboard_v1alpha1_template_service_proto_rawDescData
}

var file_dashboard_v1alpha1_template_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dashboard_v1alpha1_template_service_proto_goTypes = []interface{}{
	(*GetUserAddonTemplatesResponse)(nil), // 0: dashboard.v1alpha1.GetUserAddonTemplatesResponse
	(*GetWorkspaceTemplatesResponse)(nil), // 1: dashboard.v1alpha1.GetWorkspaceTemplatesResponse
	(*Template)(nil),                      // 2: dashboard.v1alpha1.Template
	(*emptypb.Empty)(nil),                 // 3: google.protobuf.Empty
}
var file_dashboard_v1alpha1_template_service_proto_depIdxs = []int32{
	2, // 0: dashboard.v1alpha1.GetUserAddonTemplatesResponse.items:type_name -> dashboard.v1alpha1.Template
	2, // 1: dashboard.v1alpha1.GetWorkspaceTemplatesResponse.items:type_name -> dashboard.v1alpha1.Template
	3, // 2: dashboard.v1alpha1.TemplateService.GetUserAddonTemplates:input_type -> google.protobuf.Empty
	3, // 3: dashboard.v1alpha1.TemplateService.GetWorkspaceTemplates:input_type -> google.protobuf.Empty
	0, // 4: dashboard.v1alpha1.TemplateService.GetUserAddonTemplates:output_type -> dashboard.v1alpha1.GetUserAddonTemplatesResponse
	1, // 5: dashboard.v1alpha1.TemplateService.GetWorkspaceTemplates:output_type -> dashboard.v1alpha1.GetWorkspaceTemplatesResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_dashboard_v1alpha1_template_service_proto_init() }
func file_dashboard_v1alpha1_template_service_proto_init() {
	if File_dashboard_v1alpha1_template_service_proto != nil {
		return
	}
	file_dashboard_v1alpha1_template_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dashboard_v1alpha1_template_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserAddonTemplatesResponse); i {
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
		file_dashboard_v1alpha1_template_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWorkspaceTemplatesResponse); i {
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
			RawDescriptor: file_dashboard_v1alpha1_template_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dashboard_v1alpha1_template_service_proto_goTypes,
		DependencyIndexes: file_dashboard_v1alpha1_template_service_proto_depIdxs,
		MessageInfos:      file_dashboard_v1alpha1_template_service_proto_msgTypes,
	}.Build()
	File_dashboard_v1alpha1_template_service_proto = out.File
	file_dashboard_v1alpha1_template_service_proto_rawDesc = nil
	file_dashboard_v1alpha1_template_service_proto_goTypes = nil
	file_dashboard_v1alpha1_template_service_proto_depIdxs = nil
}
