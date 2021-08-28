// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: crud.proto

package pb

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

type CRUDObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CRUDObject) Reset() {
	*x = CRUDObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CRUDObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CRUDObject) ProtoMessage() {}

func (x *CRUDObject) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CRUDObject.ProtoReflect.Descriptor instead.
func (*CRUDObject) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{0}
}

func (x *CRUDObject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CRUDObject) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type CreateCRUD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CreateCRUD) Reset() {
	*x = CreateCRUD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCRUD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCRUD) ProtoMessage() {}

func (x *CreateCRUD) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCRUD.ProtoReflect.Descriptor instead.
func (*CreateCRUD) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCRUD) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCRUD) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetCRUD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCRUD) Reset() {
	*x = GetCRUD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCRUD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCRUD) ProtoMessage() {}

func (x *GetCRUD) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCRUD.ProtoReflect.Descriptor instead.
func (*GetCRUD) Descriptor() ([]byte, []int) {
	return file_crud_proto_rawDescGZIP(), []int{2}
}

func (x *GetCRUD) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_crud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_crud_proto_msgTypes[3]
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
	return file_crud_proto_rawDescGZIP(), []int{3}
}

var File_crud_proto protoreflect.FileDescriptor

var file_crud_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x43, 0x52, 0x55, 0x44, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x36, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x52, 0x55, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x52, 0x55, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xcb, 0x01,
	0x0a, 0x04, 0x43, 0x52, 0x55, 0x44, 0x12, 0x32, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x52, 0x55, 0x44, 0x1a, 0x13, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x43, 0x52, 0x55, 0x44, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x10, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x52, 0x55, 0x44, 0x1a, 0x13, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x52,
	0x55, 0x44, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x32, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x13, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x43, 0x52, 0x55,
	0x44, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x13, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x43, 0x52, 0x55, 0x44, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2d, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x43, 0x52, 0x55, 0x44, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x1a, 0x0e, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_crud_proto_rawDescOnce sync.Once
	file_crud_proto_rawDescData = file_crud_proto_rawDesc
)

func file_crud_proto_rawDescGZIP() []byte {
	file_crud_proto_rawDescOnce.Do(func() {
		file_crud_proto_rawDescData = protoimpl.X.CompressGZIP(file_crud_proto_rawDescData)
	})
	return file_crud_proto_rawDescData
}

var file_crud_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_crud_proto_goTypes = []interface{}{
	(*CRUDObject)(nil), // 0: example.CRUDObject
	(*CreateCRUD)(nil), // 1: example.CreateCRUD
	(*GetCRUD)(nil),    // 2: example.GetCRUD
	(*Empty)(nil),      // 3: example.Empty
}
var file_crud_proto_depIdxs = []int32{
	1, // 0: example.CRUD.Create:input_type -> example.CreateCRUD
	2, // 1: example.CRUD.Get:input_type -> example.GetCRUD
	0, // 2: example.CRUD.Update:input_type -> example.CRUDObject
	0, // 3: example.CRUD.Delete:input_type -> example.CRUDObject
	0, // 4: example.CRUD.Create:output_type -> example.CRUDObject
	0, // 5: example.CRUD.Get:output_type -> example.CRUDObject
	0, // 6: example.CRUD.Update:output_type -> example.CRUDObject
	3, // 7: example.CRUD.Delete:output_type -> example.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crud_proto_init() }
func file_crud_proto_init() {
	if File_crud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_crud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CRUDObject); i {
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
		file_crud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCRUD); i {
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
		file_crud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCRUD); i {
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
		file_crud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_crud_proto_goTypes,
		DependencyIndexes: file_crud_proto_depIdxs,
		MessageInfos:      file_crud_proto_msgTypes,
	}.Build()
	File_crud_proto = out.File
	file_crud_proto_rawDesc = nil
	file_crud_proto_goTypes = nil
	file_crud_proto_depIdxs = nil
}
