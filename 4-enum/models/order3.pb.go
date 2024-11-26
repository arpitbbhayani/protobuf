// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v4.25.1
// source: protos/order3.proto

package models

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

type OrderStatus3 int32

const (
	// remove ORDER_STATUS3 prefixes and it gives a compile-time error
	OrderStatus3_ORDER_STATUS3_UNSPECIFIDED OrderStatus3 = 0 // remove this line and it gives a compile-time error
	OrderStatus3_ORDER_STATUS3_PENDING      OrderStatus3 = 1
	OrderStatus3_ORDER_STATUS3_PROCESSING   OrderStatus3 = 2
	OrderStatus3_ORDER_STATUS3_SHIPPED      OrderStatus3 = 3
	OrderStatus3_ORDER_STATUS3_DELIVERED    OrderStatus3 = 4
)

// Enum value maps for OrderStatus3.
var (
	OrderStatus3_name = map[int32]string{
		0: "ORDER_STATUS3_UNSPECIFIDED",
		1: "ORDER_STATUS3_PENDING",
		2: "ORDER_STATUS3_PROCESSING",
		3: "ORDER_STATUS3_SHIPPED",
		4: "ORDER_STATUS3_DELIVERED",
	}
	OrderStatus3_value = map[string]int32{
		"ORDER_STATUS3_UNSPECIFIDED": 0,
		"ORDER_STATUS3_PENDING":      1,
		"ORDER_STATUS3_PROCESSING":   2,
		"ORDER_STATUS3_SHIPPED":      3,
		"ORDER_STATUS3_DELIVERED":    4,
	}
)

func (x OrderStatus3) Enum() *OrderStatus3 {
	p := new(OrderStatus3)
	*p = x
	return p
}

func (x OrderStatus3) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderStatus3) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_order3_proto_enumTypes[0].Descriptor()
}

func (OrderStatus3) Type() protoreflect.EnumType {
	return &file_protos_order3_proto_enumTypes[0]
}

func (x OrderStatus3) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderStatus3.Descriptor instead.
func (OrderStatus3) EnumDescriptor() ([]byte, []int) {
	return file_protos_order3_proto_rawDescGZIP(), []int{0}
}

type Order3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string       `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Amount  float64      `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Status  OrderStatus3 `protobuf:"varint,3,opt,name=status,proto3,enum=models.OrderStatus3" json:"status,omitempty"`
}

func (x *Order3) Reset() {
	*x = Order3{}
	mi := &file_protos_order3_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Order3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order3) ProtoMessage() {}

func (x *Order3) ProtoReflect() protoreflect.Message {
	mi := &file_protos_order3_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order3.ProtoReflect.Descriptor instead.
func (*Order3) Descriptor() ([]byte, []int) {
	return file_protos_order3_proto_rawDescGZIP(), []int{0}
}

func (x *Order3) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *Order3) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Order3) GetStatus() OrderStatus3 {
	if x != nil {
		return x.Status
	}
	return OrderStatus3_ORDER_STATUS3_UNSPECIFIDED
}

var File_protos_order3_proto protoreflect.FileDescriptor

var file_protos_order3_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x33, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x22, 0x69, 0x0a,
	0x06, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x33, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x33,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x9f, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x33, 0x12, 0x1e, 0x0a, 0x1a, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x33, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x44, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x52, 0x44,
	0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x33, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x33, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47,
	0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x33, 0x5f, 0x53, 0x48, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x03, 0x12, 0x1b, 0x0a,
	0x17, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x33, 0x5f, 0x44,
	0x45, 0x4c, 0x49, 0x56, 0x45, 0x52, 0x45, 0x44, 0x10, 0x04, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_order3_proto_rawDescOnce sync.Once
	file_protos_order3_proto_rawDescData = file_protos_order3_proto_rawDesc
)

func file_protos_order3_proto_rawDescGZIP() []byte {
	file_protos_order3_proto_rawDescOnce.Do(func() {
		file_protos_order3_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_order3_proto_rawDescData)
	})
	return file_protos_order3_proto_rawDescData
}

var file_protos_order3_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_order3_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protos_order3_proto_goTypes = []any{
	(OrderStatus3)(0), // 0: models.OrderStatus3
	(*Order3)(nil),    // 1: models.Order3
}
var file_protos_order3_proto_depIdxs = []int32{
	0, // 0: models.Order3.status:type_name -> models.OrderStatus3
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_order3_proto_init() }
func file_protos_order3_proto_init() {
	if File_protos_order3_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_order3_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_order3_proto_goTypes,
		DependencyIndexes: file_protos_order3_proto_depIdxs,
		EnumInfos:         file_protos_order3_proto_enumTypes,
		MessageInfos:      file_protos_order3_proto_msgTypes,
	}.Build()
	File_protos_order3_proto = out.File
	file_protos_order3_proto_rawDesc = nil
	file_protos_order3_proto_goTypes = nil
	file_protos_order3_proto_depIdxs = nil
}