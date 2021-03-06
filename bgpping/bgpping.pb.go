// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: bgpping/bgpping.proto

package bgpping

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

type BgpPingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerIp                        string `protobuf:"bytes,1,opt,name=peer_ip,json=peerIp,proto3" json:"peer_ip,omitempty"`
	StopAfterSendingCountRequests int64  `protobuf:"varint,2,opt,name=stop_after_sending_count_requests,json=stopAfterSendingCountRequests,proto3" json:"stop_after_sending_count_requests,omitempty"`
}

func (x *BgpPingRequest) Reset() {
	*x = BgpPingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgpping_bgpping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BgpPingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BgpPingRequest) ProtoMessage() {}

func (x *BgpPingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bgpping_bgpping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BgpPingRequest.ProtoReflect.Descriptor instead.
func (*BgpPingRequest) Descriptor() ([]byte, []int) {
	return file_bgpping_bgpping_proto_rawDescGZIP(), []int{0}
}

func (x *BgpPingRequest) GetPeerIp() string {
	if x != nil {
		return x.PeerIp
	}
	return ""
}

func (x *BgpPingRequest) GetStopAfterSendingCountRequests() int64 {
	if x != nil {
		return x.StopAfterSendingCountRequests
	}
	return 0
}

type BgpPingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProbesSent       int64 `protobuf:"varint,1,opt,name=probes_sent,json=probesSent,proto3" json:"probes_sent,omitempty"`
	ProbesSuccessful int64 `protobuf:"varint,2,opt,name=probes_successful,json=probesSuccessful,proto3" json:"probes_successful,omitempty"`
	ProbesFailed     int64 `protobuf:"varint,3,opt,name=probes_failed,json=probesFailed,proto3" json:"probes_failed,omitempty"`
}

func (x *BgpPingResponse) Reset() {
	*x = BgpPingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bgpping_bgpping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BgpPingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BgpPingResponse) ProtoMessage() {}

func (x *BgpPingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bgpping_bgpping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BgpPingResponse.ProtoReflect.Descriptor instead.
func (*BgpPingResponse) Descriptor() ([]byte, []int) {
	return file_bgpping_bgpping_proto_rawDescGZIP(), []int{1}
}

func (x *BgpPingResponse) GetProbesSent() int64 {
	if x != nil {
		return x.ProbesSent
	}
	return 0
}

func (x *BgpPingResponse) GetProbesSuccessful() int64 {
	if x != nil {
		return x.ProbesSuccessful
	}
	return 0
}

func (x *BgpPingResponse) GetProbesFailed() int64 {
	if x != nil {
		return x.ProbesFailed
	}
	return 0
}

var File_bgpping_bgpping_proto protoreflect.FileDescriptor

var file_bgpping_bgpping_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x67, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x67, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x67, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x22, 0x73, 0x0a, 0x0e, 0x42, 0x67, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x70, 0x12, 0x48, 0x0a, 0x21, 0x73,
	0x74, 0x6f, 0x70, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1d, 0x73, 0x74, 0x6f, 0x70, 0x41, 0x66, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x84, 0x01, 0x0a, 0x0f, 0x42, 0x67, 0x70, 0x50, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x62, 0x65, 0x73, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x70, 0x72,
	0x6f, 0x62, 0x65, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x62, 0x65,
	0x73, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x70, 0x72, 0x6f, 0x62, 0x65, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x32, 0x9d, 0x01, 0x0a,
	0x07, 0x42, 0x67, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x42, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67,
	0x42, 0x67, 0x70, 0x50, 0x65, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x62, 0x67, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x67, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x62, 0x67, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x67, 0x70, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x17,
	0x50, 0x69, 0x6e, 0x67, 0x42, 0x67, 0x70, 0x50, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x69,
	0x6e, 0x75, 0x6f, 0x75, 0x73, 0x6c, 0x79, 0x12, 0x17, 0x2e, 0x62, 0x67, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x67, 0x70, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x18, 0x2e, 0x62, 0x67, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x67, 0x70, 0x50, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x24, 0x5a, 0x22,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x2f, 0x62, 0x67, 0x70, 0x2d, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x67, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bgpping_bgpping_proto_rawDescOnce sync.Once
	file_bgpping_bgpping_proto_rawDescData = file_bgpping_bgpping_proto_rawDesc
)

func file_bgpping_bgpping_proto_rawDescGZIP() []byte {
	file_bgpping_bgpping_proto_rawDescOnce.Do(func() {
		file_bgpping_bgpping_proto_rawDescData = protoimpl.X.CompressGZIP(file_bgpping_bgpping_proto_rawDescData)
	})
	return file_bgpping_bgpping_proto_rawDescData
}

var file_bgpping_bgpping_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_bgpping_bgpping_proto_goTypes = []interface{}{
	(*BgpPingRequest)(nil),  // 0: bgpping.BgpPingRequest
	(*BgpPingResponse)(nil), // 1: bgpping.BgpPingResponse
}
var file_bgpping_bgpping_proto_depIdxs = []int32{
	0, // 0: bgpping.BgpPing.PingBgpPeer:input_type -> bgpping.BgpPingRequest
	0, // 1: bgpping.BgpPing.PingBgpPeerContinuously:input_type -> bgpping.BgpPingRequest
	1, // 2: bgpping.BgpPing.PingBgpPeer:output_type -> bgpping.BgpPingResponse
	1, // 3: bgpping.BgpPing.PingBgpPeerContinuously:output_type -> bgpping.BgpPingResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bgpping_bgpping_proto_init() }
func file_bgpping_bgpping_proto_init() {
	if File_bgpping_bgpping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bgpping_bgpping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BgpPingRequest); i {
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
		file_bgpping_bgpping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BgpPingResponse); i {
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
			RawDescriptor: file_bgpping_bgpping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bgpping_bgpping_proto_goTypes,
		DependencyIndexes: file_bgpping_bgpping_proto_depIdxs,
		MessageInfos:      file_bgpping_bgpping_proto_msgTypes,
	}.Build()
	File_bgpping_bgpping_proto = out.File
	file_bgpping_bgpping_proto_rawDesc = nil
	file_bgpping_bgpping_proto_goTypes = nil
	file_bgpping_bgpping_proto_depIdxs = nil
}
