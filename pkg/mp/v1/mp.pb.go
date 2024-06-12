// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: v1/mp.proto

package mp

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[0]
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
	return file_v1_mp_proto_rawDescGZIP(), []int{0}
}

type AdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductName string `protobuf:"bytes,1,opt,name=product_name,json=productName,proto3" json:"product_name,omitempty"`
	Price       uint64 `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	SellerId    uint64 `protobuf:"varint,3,opt,name=seller_id,json=sellerId,proto3" json:"seller_id,omitempty"`
}

func (x *AdRequest) Reset() {
	*x = AdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdRequest) ProtoMessage() {}

func (x *AdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdRequest.ProtoReflect.Descriptor instead.
func (*AdRequest) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{1}
}

func (x *AdRequest) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *AdRequest) GetPrice() uint64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AdRequest) GetSellerId() uint64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

type ListAdsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Substring *string `protobuf:"bytes,1,opt,name=substring,proto3,oneof" json:"substring,omitempty"`
	SellerId  *uint64 `protobuf:"varint,2,opt,name=seller_id,json=sellerId,proto3,oneof" json:"seller_id,omitempty"`
}

func (x *ListAdsRequest) Reset() {
	*x = ListAdsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAdsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAdsRequest) ProtoMessage() {}

func (x *ListAdsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAdsRequest.ProtoReflect.Descriptor instead.
func (*ListAdsRequest) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{2}
}

func (x *ListAdsRequest) GetSubstring() string {
	if x != nil && x.Substring != nil {
		return *x.Substring
	}
	return ""
}

func (x *ListAdsRequest) GetSellerId() uint64 {
	if x != nil && x.SellerId != nil {
		return *x.SellerId
	}
	return 0
}

type AdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId   uint64     `protobuf:"varint,1,opt,name=ad_id,json=adId,proto3" json:"ad_id,omitempty"`
	Status uint32     `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Props  *AdRequest `protobuf:"bytes,3,opt,name=props,proto3" json:"props,omitempty"`
}

func (x *AdResponse) Reset() {
	*x = AdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdResponse) ProtoMessage() {}

func (x *AdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdResponse.ProtoReflect.Descriptor instead.
func (*AdResponse) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{3}
}

func (x *AdResponse) GetAdId() uint64 {
	if x != nil {
		return x.AdId
	}
	return 0
}

func (x *AdResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AdResponse) GetProps() *AdRequest {
	if x != nil {
		return x.Props
	}
	return nil
}

type OrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdId    uint64 `protobuf:"varint,1,opt,name=ad_id,json=adId,proto3" json:"ad_id,omitempty"`
	BuyerId uint64 `protobuf:"varint,2,opt,name=buyer_id,json=buyerId,proto3" json:"buyer_id,omitempty"`
}

func (x *OrderRequest) Reset() {
	*x = OrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderRequest) ProtoMessage() {}

func (x *OrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderRequest.ProtoReflect.Descriptor instead.
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{4}
}

func (x *OrderRequest) GetAdId() uint64 {
	if x != nil {
		return x.AdId
	}
	return 0
}

func (x *OrderRequest) GetBuyerId() uint64 {
	if x != nil {
		return x.BuyerId
	}
	return 0
}

type ListOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UserId:
	//
	//	*ListOrdersRequest_SellerId
	//	*ListOrdersRequest_BuyerId
	UserId isListOrdersRequest_UserId `protobuf_oneof:"user_id"`
}

func (x *ListOrdersRequest) Reset() {
	*x = ListOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOrdersRequest) ProtoMessage() {}

func (x *ListOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOrdersRequest.ProtoReflect.Descriptor instead.
func (*ListOrdersRequest) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{5}
}

func (m *ListOrdersRequest) GetUserId() isListOrdersRequest_UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (x *ListOrdersRequest) GetSellerId() uint64 {
	if x, ok := x.GetUserId().(*ListOrdersRequest_SellerId); ok {
		return x.SellerId
	}
	return 0
}

func (x *ListOrdersRequest) GetBuyerId() uint64 {
	if x, ok := x.GetUserId().(*ListOrdersRequest_BuyerId); ok {
		return x.BuyerId
	}
	return 0
}

type isListOrdersRequest_UserId interface {
	isListOrdersRequest_UserId()
}

type ListOrdersRequest_SellerId struct {
	SellerId uint64 `protobuf:"varint,1,opt,name=seller_id,json=sellerId,proto3,oneof"`
}

type ListOrdersRequest_BuyerId struct {
	BuyerId uint64 `protobuf:"varint,2,opt,name=buyer_id,json=buyerId,proto3,oneof"`
}

func (*ListOrdersRequest_SellerId) isListOrdersRequest_UserId() {}

func (*ListOrdersRequest_BuyerId) isListOrdersRequest_UserId() {}

type OrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64        `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status  uint32        `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Props   *OrderRequest `protobuf:"bytes,3,opt,name=props,proto3" json:"props,omitempty"`
}

func (x *OrderResponse) Reset() {
	*x = OrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResponse) ProtoMessage() {}

func (x *OrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResponse.ProtoReflect.Descriptor instead.
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{6}
}

func (x *OrderResponse) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *OrderResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *OrderResponse) GetProps() *OrderRequest {
	if x != nil {
		return x.Props
	}
	return nil
}

type ShipmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	SellerId uint64 `protobuf:"varint,2,opt,name=seller_id,json=sellerId,proto3" json:"seller_id,omitempty"`
}

func (x *ShipmentRequest) Reset() {
	*x = ShipmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipmentRequest) ProtoMessage() {}

func (x *ShipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipmentRequest.ProtoReflect.Descriptor instead.
func (*ShipmentRequest) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{7}
}

func (x *ShipmentRequest) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *ShipmentRequest) GetSellerId() uint64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

type ListShipmentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UserId:
	//
	//	*ListShipmentsRequest_SellerId
	//	*ListShipmentsRequest_BuyerId
	UserId isListShipmentsRequest_UserId `protobuf_oneof:"user_id"`
}

func (x *ListShipmentsRequest) Reset() {
	*x = ListShipmentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShipmentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShipmentsRequest) ProtoMessage() {}

func (x *ListShipmentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShipmentsRequest.ProtoReflect.Descriptor instead.
func (*ListShipmentsRequest) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{8}
}

func (m *ListShipmentsRequest) GetUserId() isListShipmentsRequest_UserId {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (x *ListShipmentsRequest) GetSellerId() uint64 {
	if x, ok := x.GetUserId().(*ListShipmentsRequest_SellerId); ok {
		return x.SellerId
	}
	return 0
}

func (x *ListShipmentsRequest) GetBuyerId() uint64 {
	if x, ok := x.GetUserId().(*ListShipmentsRequest_BuyerId); ok {
		return x.BuyerId
	}
	return 0
}

type isListShipmentsRequest_UserId interface {
	isListShipmentsRequest_UserId()
}

type ListShipmentsRequest_SellerId struct {
	SellerId uint64 `protobuf:"varint,1,opt,name=seller_id,json=sellerId,proto3,oneof"`
}

type ListShipmentsRequest_BuyerId struct {
	BuyerId uint64 `protobuf:"varint,2,opt,name=buyer_id,json=buyerId,proto3,oneof"`
}

func (*ListShipmentsRequest_SellerId) isListShipmentsRequest_UserId() {}

func (*ListShipmentsRequest_BuyerId) isListShipmentsRequest_UserId() {}

type ShipmentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipmentId uint64 `protobuf:"varint,1,opt,name=shipment_id,json=shipmentId,proto3" json:"shipment_id,omitempty"`
	Status     uint32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	OrderId    uint64 `protobuf:"varint,3,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *ShipmentResponse) Reset() {
	*x = ShipmentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShipmentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShipmentResponse) ProtoMessage() {}

func (x *ShipmentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShipmentResponse.ProtoReflect.Descriptor instead.
func (*ShipmentResponse) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{9}
}

func (x *ShipmentResponse) GetShipmentId() uint64 {
	if x != nil {
		return x.ShipmentId
	}
	return 0
}

func (x *ShipmentResponse) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ShipmentResponse) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type DeliveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShipmentId uint64 `protobuf:"varint,1,opt,name=shipment_id,json=shipmentId,proto3" json:"shipment_id,omitempty"`
	BuyerId    uint64 `protobuf:"varint,2,opt,name=buyer_id,json=buyerId,proto3" json:"buyer_id,omitempty"`
}

func (x *DeliveryRequest) Reset() {
	*x = DeliveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryRequest) ProtoMessage() {}

func (x *DeliveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryRequest.ProtoReflect.Descriptor instead.
func (*DeliveryRequest) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{10}
}

func (x *DeliveryRequest) GetShipmentId() uint64 {
	if x != nil {
		return x.ShipmentId
	}
	return 0
}

func (x *DeliveryRequest) GetBuyerId() uint64 {
	if x != nil {
		return x.BuyerId
	}
	return 0
}

type Dictionary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*Dictionary_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *Dictionary) Reset() {
	*x = Dictionary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dictionary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dictionary) ProtoMessage() {}

func (x *Dictionary) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dictionary.ProtoReflect.Descriptor instead.
func (*Dictionary) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{11}
}

func (x *Dictionary) GetEntries() []*Dictionary_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type Dictionary_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`      // 0
	Code  string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`   // UNKNOWN
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"` // Неизвестно
}

func (x *Dictionary_Entry) Reset() {
	*x = Dictionary_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_mp_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dictionary_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dictionary_Entry) ProtoMessage() {}

func (x *Dictionary_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_mp_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dictionary_Entry.ProtoReflect.Descriptor instead.
func (*Dictionary_Entry) Descriptor() ([]byte, []int) {
	return file_v1_mp_proto_rawDescGZIP(), []int{11, 0}
}

func (x *Dictionary_Entry) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dictionary_Entry) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Dictionary_Entry) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_v1_mp_proto protoreflect.FileDescriptor

var file_v1_mp_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x31, 0x2f, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x61, 0x0a, 0x09, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x71, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x73, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x01, 0x52, 0x08,
	0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x73, 0x75, 0x62, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x65,
	0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x0a, 0x41, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x61, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73,
	0x22, 0x3e, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x13, 0x0a, 0x05, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x61, 0x64, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x75, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x5a, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x07, 0x62, 0x75, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x42, 0x09, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x73, 0x0a, 0x0d,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x2f, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x70,
	0x73, 0x22, 0x49, 0x0a, 0x0f, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x22, 0x5d, 0x0a, 0x14,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x08, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x07, 0x62, 0x75, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x09, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x66, 0x0a, 0x10, 0x53,
	0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x68, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x75, 0x79, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72,
	0x79, 0x12, 0x37, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x41, 0x0a, 0x05, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x32, 0xf2, 0x05,
	0x0a, 0x0b, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x49, 0x0a,
	0x08, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x64, 0x12, 0x16, 0x2e, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x3a, 0x01, 0x2a,
	0x22, 0x06, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x12, 0x52, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x64, 0x73, 0x12, 0x1b, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x41,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x09, 0x12, 0x07, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x73, 0x30, 0x01, 0x12, 0x57, 0x0a, 0x05,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22,
	0x14, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x2f, 0x7b, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x5e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x12, 0x1e, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x30, 0x01, 0x12, 0x5f, 0x0a, 0x04, 0x53, 0x68, 0x69, 0x70, 0x12, 0x1c, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x68, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x6a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x53, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x30, 0x01, 0x12, 0x6b, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x1c, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x68, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x73, 0x68, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12,
	0x51, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12,
	0x12, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x44, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x72, 0x79, 0x22, 0x14, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x65, 0x73, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x69, 0x74, 0x73, 0x76, 0x65, 0x74, 0x2f, 0x73, 0x61, 0x67, 0x61, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6d, 0x70, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_mp_proto_rawDescOnce sync.Once
	file_v1_mp_proto_rawDescData = file_v1_mp_proto_rawDesc
)

func file_v1_mp_proto_rawDescGZIP() []byte {
	file_v1_mp_proto_rawDescOnce.Do(func() {
		file_v1_mp_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_mp_proto_rawDescData)
	})
	return file_v1_mp_proto_rawDescData
}

var file_v1_mp_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_v1_mp_proto_goTypes = []interface{}{
	(*Empty)(nil),                // 0: marketplace.Empty
	(*AdRequest)(nil),            // 1: marketplace.AdRequest
	(*ListAdsRequest)(nil),       // 2: marketplace.ListAdsRequest
	(*AdResponse)(nil),           // 3: marketplace.AdResponse
	(*OrderRequest)(nil),         // 4: marketplace.OrderRequest
	(*ListOrdersRequest)(nil),    // 5: marketplace.ListOrdersRequest
	(*OrderResponse)(nil),        // 6: marketplace.OrderResponse
	(*ShipmentRequest)(nil),      // 7: marketplace.ShipmentRequest
	(*ListShipmentsRequest)(nil), // 8: marketplace.ListShipmentsRequest
	(*ShipmentResponse)(nil),     // 9: marketplace.ShipmentResponse
	(*DeliveryRequest)(nil),      // 10: marketplace.DeliveryRequest
	(*Dictionary)(nil),           // 11: marketplace.Dictionary
	(*Dictionary_Entry)(nil),     // 12: marketplace.Dictionary.Entry
}
var file_v1_mp_proto_depIdxs = []int32{
	1,  // 0: marketplace.AdResponse.props:type_name -> marketplace.AdRequest
	4,  // 1: marketplace.OrderResponse.props:type_name -> marketplace.OrderRequest
	12, // 2: marketplace.Dictionary.entries:type_name -> marketplace.Dictionary.Entry
	1,  // 3: marketplace.Marketplace.CreateAd:input_type -> marketplace.AdRequest
	2,  // 4: marketplace.Marketplace.ListAds:input_type -> marketplace.ListAdsRequest
	4,  // 5: marketplace.Marketplace.Order:input_type -> marketplace.OrderRequest
	5,  // 6: marketplace.Marketplace.ListOrders:input_type -> marketplace.ListOrdersRequest
	7,  // 7: marketplace.Marketplace.Ship:input_type -> marketplace.ShipmentRequest
	8,  // 8: marketplace.Marketplace.ListShipments:input_type -> marketplace.ListShipmentsRequest
	10, // 9: marketplace.Marketplace.Receive:input_type -> marketplace.DeliveryRequest
	0,  // 10: marketplace.Marketplace.ListStatuses:input_type -> marketplace.Empty
	0,  // 11: marketplace.Marketplace.CreateAd:output_type -> marketplace.Empty
	3,  // 12: marketplace.Marketplace.ListAds:output_type -> marketplace.AdResponse
	0,  // 13: marketplace.Marketplace.Order:output_type -> marketplace.Empty
	6,  // 14: marketplace.Marketplace.ListOrders:output_type -> marketplace.OrderResponse
	0,  // 15: marketplace.Marketplace.Ship:output_type -> marketplace.Empty
	9,  // 16: marketplace.Marketplace.ListShipments:output_type -> marketplace.ShipmentResponse
	0,  // 17: marketplace.Marketplace.Receive:output_type -> marketplace.Empty
	11, // 18: marketplace.Marketplace.ListStatuses:output_type -> marketplace.Dictionary
	11, // [11:19] is the sub-list for method output_type
	3,  // [3:11] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_v1_mp_proto_init() }
func file_v1_mp_proto_init() {
	if File_v1_mp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_mp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_v1_mp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdRequest); i {
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
		file_v1_mp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAdsRequest); i {
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
		file_v1_mp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdResponse); i {
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
		file_v1_mp_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderRequest); i {
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
		file_v1_mp_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOrdersRequest); i {
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
		file_v1_mp_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResponse); i {
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
		file_v1_mp_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipmentRequest); i {
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
		file_v1_mp_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListShipmentsRequest); i {
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
		file_v1_mp_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShipmentResponse); i {
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
		file_v1_mp_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryRequest); i {
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
		file_v1_mp_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dictionary); i {
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
		file_v1_mp_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dictionary_Entry); i {
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
	file_v1_mp_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_v1_mp_proto_msgTypes[5].OneofWrappers = []interface{}{
		(*ListOrdersRequest_SellerId)(nil),
		(*ListOrdersRequest_BuyerId)(nil),
	}
	file_v1_mp_proto_msgTypes[8].OneofWrappers = []interface{}{
		(*ListShipmentsRequest_SellerId)(nil),
		(*ListShipmentsRequest_BuyerId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_mp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_mp_proto_goTypes,
		DependencyIndexes: file_v1_mp_proto_depIdxs,
		MessageInfos:      file_v1_mp_proto_msgTypes,
	}.Build()
	File_v1_mp_proto = out.File
	file_v1_mp_proto_rawDesc = nil
	file_v1_mp_proto_goTypes = nil
	file_v1_mp_proto_depIdxs = nil
}