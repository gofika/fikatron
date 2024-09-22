// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: core/TronInventoryItems.proto

package core

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

type InventoryItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Items [][]byte `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *InventoryItems) Reset() {
	*x = InventoryItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_TronInventoryItems_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InventoryItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InventoryItems) ProtoMessage() {}

func (x *InventoryItems) ProtoReflect() protoreflect.Message {
	mi := &file_core_TronInventoryItems_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InventoryItems.ProtoReflect.Descriptor instead.
func (*InventoryItems) Descriptor() ([]byte, []int) {
	return file_core_TronInventoryItems_proto_rawDescGZIP(), []int{0}
}

func (x *InventoryItems) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *InventoryItems) GetItems() [][]byte {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_core_TronInventoryItems_proto protoreflect.FileDescriptor

var file_core_TronInventoryItems_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x54, 0x72, 0x6f, 0x6e, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x3a, 0x0a, 0x0e, 0x49, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x50, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x72, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x42, 0x12, 0x54, 0x72, 0x6f, 0x6e, 0x49, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x5a, 0x29, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x6f, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_TronInventoryItems_proto_rawDescOnce sync.Once
	file_core_TronInventoryItems_proto_rawDescData = file_core_TronInventoryItems_proto_rawDesc
)

func file_core_TronInventoryItems_proto_rawDescGZIP() []byte {
	file_core_TronInventoryItems_proto_rawDescOnce.Do(func() {
		file_core_TronInventoryItems_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_TronInventoryItems_proto_rawDescData)
	})
	return file_core_TronInventoryItems_proto_rawDescData
}

var file_core_TronInventoryItems_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_core_TronInventoryItems_proto_goTypes = []any{
	(*InventoryItems)(nil), // 0: protocol.InventoryItems
}
var file_core_TronInventoryItems_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_TronInventoryItems_proto_init() }
func file_core_TronInventoryItems_proto_init() {
	if File_core_TronInventoryItems_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_TronInventoryItems_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*InventoryItems); i {
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
			RawDescriptor: file_core_TronInventoryItems_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_TronInventoryItems_proto_goTypes,
		DependencyIndexes: file_core_TronInventoryItems_proto_depIdxs,
		MessageInfos:      file_core_TronInventoryItems_proto_msgTypes,
	}.Build()
	File_core_TronInventoryItems_proto = out.File
	file_core_TronInventoryItems_proto_rawDesc = nil
	file_core_TronInventoryItems_proto_goTypes = nil
	file_core_TronInventoryItems_proto_depIdxs = nil
}
