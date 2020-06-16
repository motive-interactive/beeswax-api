// Copyright 2020, Beeswax.IO Inc.
// Protocol buffer defining ghost bidding attributes of an adgroup

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: beeswax/adgroup/ghost_bidding.proto

package adgroup

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Next Tag: 2
type GhostBidding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//if the bid is a ghost bid. If true, the bid will be logged but will NOT be submitted to exchanges.
	IsGhostBid *bool `protobuf:"varint,1,req,name=is_ghost_bid,json=isGhostBid" json:"is_ghost_bid,omitempty"`
}

func (x *GhostBidding) Reset() {
	*x = GhostBidding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beeswax_adgroup_ghost_bidding_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GhostBidding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GhostBidding) ProtoMessage() {}

func (x *GhostBidding) ProtoReflect() protoreflect.Message {
	mi := &file_beeswax_adgroup_ghost_bidding_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GhostBidding.ProtoReflect.Descriptor instead.
func (*GhostBidding) Descriptor() ([]byte, []int) {
	return file_beeswax_adgroup_ghost_bidding_proto_rawDescGZIP(), []int{0}
}

func (x *GhostBidding) GetIsGhostBid() bool {
	if x != nil && x.IsGhostBid != nil {
		return *x.IsGhostBid
	}
	return false
}

var File_beeswax_adgroup_ghost_bidding_proto protoreflect.FileDescriptor

var file_beeswax_adgroup_ghost_bidding_proto_rawDesc = []byte{
	0x0a, 0x23, 0x62, 0x65, 0x65, 0x73, 0x77, 0x61, 0x78, 0x2f, 0x61, 0x64, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2f, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x64, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x30,
	0x0a, 0x0c, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x20,
	0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x67, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x47, 0x68, 0x6f, 0x73, 0x74, 0x42, 0x69, 0x64,
	0x42, 0x18, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x65, 0x65, 0x73, 0x77, 0x61, 0x78, 0x2e,
	0x61, 0x64, 0x67, 0x72, 0x6f, 0x75, 0x70, 0xf8, 0x01, 0x01,
}

var (
	file_beeswax_adgroup_ghost_bidding_proto_rawDescOnce sync.Once
	file_beeswax_adgroup_ghost_bidding_proto_rawDescData = file_beeswax_adgroup_ghost_bidding_proto_rawDesc
)

func file_beeswax_adgroup_ghost_bidding_proto_rawDescGZIP() []byte {
	file_beeswax_adgroup_ghost_bidding_proto_rawDescOnce.Do(func() {
		file_beeswax_adgroup_ghost_bidding_proto_rawDescData = protoimpl.X.CompressGZIP(file_beeswax_adgroup_ghost_bidding_proto_rawDescData)
	})
	return file_beeswax_adgroup_ghost_bidding_proto_rawDescData
}

var file_beeswax_adgroup_ghost_bidding_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_beeswax_adgroup_ghost_bidding_proto_goTypes = []interface{}{
	(*GhostBidding)(nil), // 0: adgroup.GhostBidding
}
var file_beeswax_adgroup_ghost_bidding_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_beeswax_adgroup_ghost_bidding_proto_init() }
func file_beeswax_adgroup_ghost_bidding_proto_init() {
	if File_beeswax_adgroup_ghost_bidding_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beeswax_adgroup_ghost_bidding_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GhostBidding); i {
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
			RawDescriptor: file_beeswax_adgroup_ghost_bidding_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_beeswax_adgroup_ghost_bidding_proto_goTypes,
		DependencyIndexes: file_beeswax_adgroup_ghost_bidding_proto_depIdxs,
		MessageInfos:      file_beeswax_adgroup_ghost_bidding_proto_msgTypes,
	}.Build()
	File_beeswax_adgroup_ghost_bidding_proto = out.File
	file_beeswax_adgroup_ghost_bidding_proto_rawDesc = nil
	file_beeswax_adgroup_ghost_bidding_proto_goTypes = nil
	file_beeswax_adgroup_ghost_bidding_proto_depIdxs = nil
}
