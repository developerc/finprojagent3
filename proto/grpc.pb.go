// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: proto/grpc.proto

package proto

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

// Сообщение, описывающее параметры подключаемого агента
type AgentParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // идентификатор агента
	Ip   string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *AgentParams) Reset() {
	*x = AgentParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentParams) ProtoMessage() {}

func (x *AgentParams) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentParams.ProtoReflect.Descriptor instead.
func (*AgentParams) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *AgentParams) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AgentParams) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *AgentParams) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// Сообщаем агенту его идентификатор
type AgentParamsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // идентификатор агента
}

func (x *AgentParamsResponse) Reset() {
	*x = AgentParamsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentParamsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentParamsResponse) ProtoMessage() {}

func (x *AgentParamsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentParamsResponse.ProtoReflect.Descriptor instead.
func (*AgentParamsResponse) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *AgentParamsResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Сообщение описывающее задачу
type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Agentid   int32   `protobuf:"varint,2,opt,name=agentid,proto3" json:"agentid,omitempty"`
	Status    string  `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Expr      string  `protobuf:"bytes,4,opt,name=expr,proto3" json:"expr,omitempty"`
	Result    float32 `protobuf:"fixed32,5,opt,name=result,proto3" json:"result,omitempty"`
	Begindate int64   `protobuf:"varint,6,opt,name=begindate,proto3" json:"begindate,omitempty"`
	Enddate   int64   `protobuf:"varint,7,opt,name=enddate,proto3" json:"enddate,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *Task) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Task) GetAgentid() int32 {
	if x != nil {
		return x.Agentid
	}
	return 0
}

func (x *Task) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Task) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

func (x *Task) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

func (x *Task) GetBegindate() int64 {
	if x != nil {
		return x.Begindate
	}
	return 0
}

func (x *Task) GetEnddate() int64 {
	if x != nil {
		return x.Enddate
	}
	return 0
}

// Сообщение HeartBit request
type HeartBit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //id агента
}

func (x *HeartBit) Reset() {
	*x = HeartBit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBit) ProtoMessage() {}

func (x *HeartBit) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBit.ProtoReflect.Descriptor instead.
func (*HeartBit) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *HeartBit) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Сообщение HeartBit response
type HeartBitResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` //id агента
}

func (x *HeartBitResp) Reset() {
	*x = HeartBitResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartBitResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartBitResp) ProtoMessage() {}

func (x *HeartBitResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartBitResp.ProtoReflect.Descriptor instead.
func (*HeartBitResp) Descriptor() ([]byte, []int) {
	return file_proto_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *HeartBitResp) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_grpc_proto protoreflect.FileDescriptor

var file_proto_grpc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x66, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x6a, 0x22, 0x41, 0x0a, 0x0b, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x25,
	0x0a, 0x13, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0xac, 0x01, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x65, 0x78, 0x70, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x62, 0x65, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e,
	0x64, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64,
	0x64, 0x61, 0x74, 0x65, 0x22, 0x1a, 0x0a, 0x08, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x69, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x1e, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x32, 0x98, 0x02, 0x0a, 0x11, 0x4f, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x65, 0x77, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x2e, 0x66, 0x69, 0x6e,
	0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x1a, 0x1c, 0x2e, 0x66, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0d, 0x2e, 0x66, 0x69, 0x6e,
	0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x0d, 0x2e, 0x66, 0x69, 0x6e, 0x70,
	0x72, 0x6f, 0x6a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x31, 0x0a, 0x05, 0x48, 0x42, 0x72, 0x65,
	0x71, 0x12, 0x11, 0x2e, 0x66, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x42, 0x69, 0x74, 0x1a, 0x15, 0x2e, 0x66, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x42, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0e, 0x50,
	0x75, 0x6c, 0x6c, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0d, 0x2e,
	0x66, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x0d, 0x2e, 0x66,
	0x69, 0x6e, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x2e, 0x0a, 0x0e, 0x50,
	0x75, 0x73, 0x68, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0d, 0x2e,
	0x66, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x0d, 0x2e, 0x66,
	0x69, 0x6e, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x42, 0x2b, 0x5a, 0x29, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x72, 0x63, 0x2f, 0x66, 0x69, 0x6e, 0x70, 0x72, 0x6f, 0x6a, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x33, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_grpc_proto_rawDescOnce sync.Once
	file_proto_grpc_proto_rawDescData = file_proto_grpc_proto_rawDesc
)

func file_proto_grpc_proto_rawDescGZIP() []byte {
	file_proto_grpc_proto_rawDescOnce.Do(func() {
		file_proto_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_grpc_proto_rawDescData)
	})
	return file_proto_grpc_proto_rawDescData
}

var file_proto_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_grpc_proto_goTypes = []interface{}{
	(*AgentParams)(nil),         // 0: finproj.AgentParams
	(*AgentParamsResponse)(nil), // 1: finproj.AgentParamsResponse
	(*Task)(nil),                // 2: finproj.Task
	(*HeartBit)(nil),            // 3: finproj.HeartBit
	(*HeartBitResp)(nil),        // 4: finproj.HeartBitResp
}
var file_proto_grpc_proto_depIdxs = []int32{
	0, // 0: finproj.OrchServerService.RegisterNewAgent:input_type -> finproj.AgentParams
	2, // 1: finproj.OrchServerService.SendTask:input_type -> finproj.Task
	3, // 2: finproj.OrchServerService.HBreq:input_type -> finproj.HeartBit
	2, // 3: finproj.OrchServerService.PullFinishTask:input_type -> finproj.Task
	2, // 4: finproj.OrchServerService.PushFinishTask:input_type -> finproj.Task
	1, // 5: finproj.OrchServerService.RegisterNewAgent:output_type -> finproj.AgentParamsResponse
	2, // 6: finproj.OrchServerService.SendTask:output_type -> finproj.Task
	4, // 7: finproj.OrchServerService.HBreq:output_type -> finproj.HeartBitResp
	2, // 8: finproj.OrchServerService.PullFinishTask:output_type -> finproj.Task
	2, // 9: finproj.OrchServerService.PushFinishTask:output_type -> finproj.Task
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_grpc_proto_init() }
func file_proto_grpc_proto_init() {
	if File_proto_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentParams); i {
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
		file_proto_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentParamsResponse); i {
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
		file_proto_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_proto_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBit); i {
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
		file_proto_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartBitResp); i {
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
			RawDescriptor: file_proto_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_grpc_proto_goTypes,
		DependencyIndexes: file_proto_grpc_proto_depIdxs,
		MessageInfos:      file_proto_grpc_proto_msgTypes,
	}.Build()
	File_proto_grpc_proto = out.File
	file_proto_grpc_proto_rawDesc = nil
	file_proto_grpc_proto_goTypes = nil
	file_proto_grpc_proto_depIdxs = nil
}