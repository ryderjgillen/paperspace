// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: port-service.proto

package grpc

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

type Protocol int32

const (
	Protocol_TCP Protocol = 0
	Protocol_UDP Protocol = 1
)

// Enum value maps for Protocol.
var (
	Protocol_name = map[int32]string{
		0: "TCP",
		1: "UDP",
	}
	Protocol_value = map[string]int32{
		"TCP": 0,
		"UDP": 1,
	}
)

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Protocol) Descriptor() protoreflect.EnumDescriptor {
	return file_port_service_proto_enumTypes[0].Descriptor()
}

func (Protocol) Type() protoreflect.EnumType {
	return &file_port_service_proto_enumTypes[0]
}

func (x Protocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Protocol.Descriptor instead.
func (Protocol) EnumDescriptor() ([]byte, []int) {
	return file_port_service_proto_rawDescGZIP(), []int{0}
}

type PortInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol Protocol `protobuf:"varint,1,opt,name=Protocol,proto3,enum=Protocol" json:"Protocol,omitempty"`
}

func (x *PortInfoRequest) Reset() {
	*x = PortInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortInfoRequest) ProtoMessage() {}

func (x *PortInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_port_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortInfoRequest.ProtoReflect.Descriptor instead.
func (*PortInfoRequest) Descriptor() ([]byte, []int) {
	return file_port_service_proto_rawDescGZIP(), []int{0}
}

func (x *PortInfoRequest) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_TCP
}

type IpPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpAddress string `protobuf:"bytes,1,opt,name=IpAddress,proto3" json:"IpAddress,omitempty"`
	Port      uint32 `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *IpPort) Reset() {
	*x = IpPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpPort) ProtoMessage() {}

func (x *IpPort) ProtoReflect() protoreflect.Message {
	mi := &file_port_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpPort.ProtoReflect.Descriptor instead.
func (*IpPort) Descriptor() ([]byte, []int) {
	return file_port_service_proto_rawDescGZIP(), []int{1}
}

func (x *IpPort) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *IpPort) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type PortInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source      *IpPort  `protobuf:"bytes,1,opt,name=Source,proto3" json:"Source,omitempty"`
	Destination *IpPort  `protobuf:"bytes,2,opt,name=Destination,proto3" json:"Destination,omitempty"`
	Protocol    Protocol `protobuf:"varint,3,opt,name=Protocol,proto3,enum=Protocol" json:"Protocol,omitempty"`
	Command     string   `protobuf:"bytes,4,opt,name=Command,proto3" json:"Command,omitempty"`
}

func (x *PortInfoResponse) Reset() {
	*x = PortInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_port_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortInfoResponse) ProtoMessage() {}

func (x *PortInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_port_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortInfoResponse.ProtoReflect.Descriptor instead.
func (*PortInfoResponse) Descriptor() ([]byte, []int) {
	return file_port_service_proto_rawDescGZIP(), []int{2}
}

func (x *PortInfoResponse) GetSource() *IpPort {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *PortInfoResponse) GetDestination() *IpPort {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *PortInfoResponse) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_TCP
}

func (x *PortInfoResponse) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

var File_port_service_proto protoreflect.FileDescriptor

var file_port_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0f, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x3a,
	0x0a, 0x06, 0x49, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x70, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x70, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x9f, 0x01, 0x0a, 0x10, 0x50,
	0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x49, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x29, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x49, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x0b,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2a, 0x1c, 0x0a, 0x08,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x43, 0x50, 0x10,
	0x00, 0x12, 0x07, 0x0a, 0x03, 0x55, 0x44, 0x50, 0x10, 0x01, 0x32, 0x42, 0x0a, 0x08, 0x50, 0x6f,
	0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x72,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x08,
	0x5a, 0x06, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_port_service_proto_rawDescOnce sync.Once
	file_port_service_proto_rawDescData = file_port_service_proto_rawDesc
)

func file_port_service_proto_rawDescGZIP() []byte {
	file_port_service_proto_rawDescOnce.Do(func() {
		file_port_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_port_service_proto_rawDescData)
	})
	return file_port_service_proto_rawDescData
}

var file_port_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_port_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_port_service_proto_goTypes = []interface{}{
	(Protocol)(0),            // 0: Protocol
	(*PortInfoRequest)(nil),  // 1: PortInfoRequest
	(*IpPort)(nil),           // 2: IpPort
	(*PortInfoResponse)(nil), // 3: PortInfoResponse
}
var file_port_service_proto_depIdxs = []int32{
	0, // 0: PortInfoRequest.Protocol:type_name -> Protocol
	2, // 1: PortInfoResponse.Source:type_name -> IpPort
	2, // 2: PortInfoResponse.Destination:type_name -> IpPort
	0, // 3: PortInfoResponse.Protocol:type_name -> Protocol
	1, // 4: PortInfo.GetPortInfo:input_type -> PortInfoRequest
	3, // 5: PortInfo.GetPortInfo:output_type -> PortInfoResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_port_service_proto_init() }
func file_port_service_proto_init() {
	if File_port_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_port_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortInfoRequest); i {
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
		file_port_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpPort); i {
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
		file_port_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortInfoResponse); i {
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
			RawDescriptor: file_port_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_port_service_proto_goTypes,
		DependencyIndexes: file_port_service_proto_depIdxs,
		EnumInfos:         file_port_service_proto_enumTypes,
		MessageInfos:      file_port_service_proto_msgTypes,
	}.Build()
	File_port_service_proto = out.File
	file_port_service_proto_rawDesc = nil
	file_port_service_proto_goTypes = nil
	file_port_service_proto_depIdxs = nil
}