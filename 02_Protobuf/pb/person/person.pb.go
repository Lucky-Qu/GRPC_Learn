//声明协议版本

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: person.proto

//声明分包名称

package person

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 按结构体来理解message,语法与Go相反，先声明类型，再声明变量
type Person struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// =号后面代表唯一标识
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Sex  bool   `protobuf:"varint,3,opt,name=sex,proto3" json:"sex,omitempty"`
	// 声明切片
	Test []string `protobuf:"bytes,4,rep,name=test,proto3" json:"test,omitempty"`
	// 声明map(第一个参数为key，只允许为string或int，第二个参数为value，无限制类型)
	TestMap       map[string]int32 `protobuf:"bytes,5,rep,name=testMap,proto3" json:"testMap,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Person) Reset() {
	*x = Person{}
	mi := &file_person_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetSex() bool {
	if x != nil {
		return x.Sex
	}
	return false
}

func (x *Person) GetTest() []string {
	if x != nil {
		return x.Test
	}
	return nil
}

func (x *Person) GetTestMap() map[string]int32 {
	if x != nil {
		return x.TestMap
	}
	return nil
}

// 支持嵌套
type Home struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	M             *HomeMaster            `protobuf:"bytes,2,opt,name=m,proto3" json:"m,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Home) Reset() {
	*x = Home{}
	mi := &file_person_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Home) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Home) ProtoMessage() {}

func (x *Home) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Home.ProtoReflect.Descriptor instead.
func (*Home) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{1}
}

func (x *Home) GetM() *HomeMaster {
	if x != nil {
		return x.M
	}
	return nil
}

type HomeMaster struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HomeMaster) Reset() {
	*x = HomeMaster{}
	mi := &file_person_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HomeMaster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HomeMaster) ProtoMessage() {}

func (x *HomeMaster) ProtoReflect() protoreflect.Message {
	mi := &file_person_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HomeMaster.ProtoReflect.Descriptor instead.
func (*HomeMaster) Descriptor() ([]byte, []int) {
	return file_person_proto_rawDescGZIP(), []int{1, 0}
}

func (x *HomeMaster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_person_proto protoreflect.FileDescriptor

var file_person_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0xc7, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x12, 0x35, 0x0a,
	0x07, 0x74, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x74, 0x65, 0x73,
	0x74, 0x4d, 0x61, 0x70, 0x1a, 0x3a, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x61, 0x70, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x47, 0x0a, 0x04, 0x48, 0x6f, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x01, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x48, 0x6f, 0x6d,
	0x65, 0x2e, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x52, 0x01, 0x6d, 0x1a, 0x1c, 0x0a, 0x06, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x30, 0x32, 0x5f,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x3b, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_person_proto_rawDescOnce sync.Once
	file_person_proto_rawDescData []byte
)

func file_person_proto_rawDescGZIP() []byte {
	file_person_proto_rawDescOnce.Do(func() {
		file_person_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_person_proto_rawDesc), len(file_person_proto_rawDesc)))
	})
	return file_person_proto_rawDescData
}

var file_person_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_person_proto_goTypes = []any{
	(*Person)(nil),     // 0: person.Person
	(*Home)(nil),       // 1: person.Home
	nil,                // 2: person.Person.TestMapEntry
	(*HomeMaster)(nil), // 3: person.Home.Master
}
var file_person_proto_depIdxs = []int32{
	2, // 0: person.Person.testMap:type_name -> person.Person.TestMapEntry
	3, // 1: person.Home.m:type_name -> person.Home.Master
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_person_proto_init() }
func file_person_proto_init() {
	if File_person_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_person_proto_rawDesc), len(file_person_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_person_proto_goTypes,
		DependencyIndexes: file_person_proto_depIdxs,
		MessageInfos:      file_person_proto_msgTypes,
	}.Build()
	File_person_proto = out.File
	file_person_proto_goTypes = nil
	file_person_proto_depIdxs = nil
}
