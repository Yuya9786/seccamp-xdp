// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.22.3
// source: protobuf/scmlb.proto

package rpc

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

type HealthRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HealthRequest) Reset() {
	*x = HealthRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scmlb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthRequest) ProtoMessage() {}

func (x *HealthRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scmlb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthRequest.ProtoReflect.Descriptor instead.
func (*HealthRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_scmlb_proto_rawDescGZIP(), []int{0}
}

type StatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatRequest) Reset() {
	*x = StatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scmlb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatRequest) ProtoMessage() {}

func (x *StatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scmlb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatRequest.ProtoReflect.Descriptor instead.
func (*StatRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_scmlb_proto_rawDescGZIP(), []int{1}
}

type StatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ifaces []*Interface `protobuf:"bytes,1,rep,name=ifaces,proto3" json:"ifaces,omitempty"`
}

func (x *StatResponse) Reset() {
	*x = StatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scmlb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatResponse) ProtoMessage() {}

func (x *StatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scmlb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatResponse.ProtoReflect.Descriptor instead.
func (*StatResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_scmlb_proto_rawDescGZIP(), []int{2}
}

func (x *StatResponse) GetIfaces() []*Interface {
	if x != nil {
		return x.Ifaces
	}
	return nil
}

type Interface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Index    int32          `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Upstream bool           `protobuf:"varint,4,opt,name=upstream,proto3" json:"upstream,omitempty"`
	Counter  *PacketCounter `protobuf:"bytes,5,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *Interface) Reset() {
	*x = Interface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scmlb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interface) ProtoMessage() {}

func (x *Interface) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scmlb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interface.ProtoReflect.Descriptor instead.
func (*Interface) Descriptor() ([]byte, []int) {
	return file_protobuf_scmlb_proto_rawDescGZIP(), []int{3}
}

func (x *Interface) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Interface) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Interface) GetUpstream() bool {
	if x != nil {
		return x.Upstream
	}
	return false
}

func (x *Interface) GetCounter() *PacketCounter {
	if x != nil {
		return x.Counter
	}
	return nil
}

type PacketCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Icmp int32 `protobuf:"varint,1,opt,name=icmp,proto3" json:"icmp,omitempty"`
	Tcp  int32 `protobuf:"varint,2,opt,name=tcp,proto3" json:"tcp,omitempty"`
	Udp  int32 `protobuf:"varint,3,opt,name=udp,proto3" json:"udp,omitempty"`
}

func (x *PacketCounter) Reset() {
	*x = PacketCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_scmlb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PacketCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PacketCounter) ProtoMessage() {}

func (x *PacketCounter) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_scmlb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PacketCounter.ProtoReflect.Descriptor instead.
func (*PacketCounter) Descriptor() ([]byte, []int) {
	return file_protobuf_scmlb_proto_rawDescGZIP(), []int{4}
}

func (x *PacketCounter) GetIcmp() int32 {
	if x != nil {
		return x.Icmp
	}
	return 0
}

func (x *PacketCounter) GetTcp() int32 {
	if x != nil {
		return x.Tcp
	}
	return 0
}

func (x *PacketCounter) GetUdp() int32 {
	if x != nil {
		return x.Udp
	}
	return 0
}

var File_protobuf_scmlb_proto protoreflect.FileDescriptor

var file_protobuf_scmlb_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x63, 0x6d, 0x6c, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x63, 0x6d, 0x6c, 0x62, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0f, 0x0a,
	0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d,
	0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a,
	0x0c, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x06, 0x69, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x73, 0x63, 0x6d, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61,
	0x63, 0x65, 0x52, 0x06, 0x69, 0x66, 0x61, 0x63, 0x65, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x09, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x75, 0x70, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x31,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x73, 0x63, 0x6d, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x22, 0x47, 0x0a, 0x0d, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x63, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x69, 0x63, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x63, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x74, 0x63, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x64, 0x70, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75, 0x64, 0x70, 0x32, 0x7c, 0x0a, 0x08, 0x53, 0x63,
	0x6d, 0x4c, 0x62, 0x41, 0x70, 0x69, 0x12, 0x39, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x12, 0x17, 0x2e, 0x73, 0x63, 0x6d, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x35, 0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x15, 0x2e, 0x73, 0x63, 0x6d, 0x6c,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x73, 0x63, 0x6d, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x72, 0x61, 0x73, 0x73, 0x79, 0x69, 0x2f,
	0x73, 0x65, 0x63, 0x63, 0x61, 0x6d, 0x70, 0x2d, 0x78, 0x64, 0x70, 0x2f, 0x73, 0x63, 0x6d, 0x6c,
	0x62, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_scmlb_proto_rawDescOnce sync.Once
	file_protobuf_scmlb_proto_rawDescData = file_protobuf_scmlb_proto_rawDesc
)

func file_protobuf_scmlb_proto_rawDescGZIP() []byte {
	file_protobuf_scmlb_proto_rawDescOnce.Do(func() {
		file_protobuf_scmlb_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_scmlb_proto_rawDescData)
	})
	return file_protobuf_scmlb_proto_rawDescData
}

var file_protobuf_scmlb_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobuf_scmlb_proto_goTypes = []interface{}{
	(*HealthRequest)(nil), // 0: scmlb.v1.HealthRequest
	(*StatRequest)(nil),   // 1: scmlb.v1.StatRequest
	(*StatResponse)(nil),  // 2: scmlb.v1.StatResponse
	(*Interface)(nil),     // 3: scmlb.v1.Interface
	(*PacketCounter)(nil), // 4: scmlb.v1.PacketCounter
	(*emptypb.Empty)(nil), // 5: google.protobuf.Empty
}
var file_protobuf_scmlb_proto_depIdxs = []int32{
	3, // 0: scmlb.v1.StatResponse.ifaces:type_name -> scmlb.v1.Interface
	4, // 1: scmlb.v1.Interface.counter:type_name -> scmlb.v1.PacketCounter
	0, // 2: scmlb.v1.ScmLbApi.Health:input_type -> scmlb.v1.HealthRequest
	1, // 3: scmlb.v1.ScmLbApi.Stat:input_type -> scmlb.v1.StatRequest
	5, // 4: scmlb.v1.ScmLbApi.Health:output_type -> google.protobuf.Empty
	2, // 5: scmlb.v1.ScmLbApi.Stat:output_type -> scmlb.v1.StatResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuf_scmlb_proto_init() }
func file_protobuf_scmlb_proto_init() {
	if File_protobuf_scmlb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_scmlb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthRequest); i {
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
		file_protobuf_scmlb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatRequest); i {
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
		file_protobuf_scmlb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatResponse); i {
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
		file_protobuf_scmlb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interface); i {
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
		file_protobuf_scmlb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PacketCounter); i {
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
			RawDescriptor: file_protobuf_scmlb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_scmlb_proto_goTypes,
		DependencyIndexes: file_protobuf_scmlb_proto_depIdxs,
		MessageInfos:      file_protobuf_scmlb_proto_msgTypes,
	}.Build()
	File_protobuf_scmlb_proto = out.File
	file_protobuf_scmlb_proto_rawDesc = nil
	file_protobuf_scmlb_proto_goTypes = nil
	file_protobuf_scmlb_proto_depIdxs = nil
}
