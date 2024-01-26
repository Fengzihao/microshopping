// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.2
// source: proto/productcatalogservice.proto

package microshopping

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 当前币种描述
type Money struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 货币code 例如：EUR 欧元 USD 美元
	CurrencyCode string `protobuf:"bytes,1,opt,name=currency_code,json=currencyCode,proto3" json:"currency_code,omitempty"`
	// 货币单位
	//
	//	例如，如果currencyCode是USD，则1个单位是1美元。
	Units int64 `protobuf:"varint,2,opt,name=units,proto3" json:"units,omitempty"`
	// 例如，$-1.75表示为 单位＝-1 和 纳米＝-75000000。
	Nanos int32 `protobuf:"varint,3,opt,name=nanos,proto3" json:"nanos,omitempty"`
}

func (x *Money) Reset() {
	*x = Money{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_productcatalogservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Money) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Money) ProtoMessage() {}

func (x *Money) ProtoReflect() protoreflect.Message {
	mi := &file_proto_productcatalogservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Money.ProtoReflect.Descriptor instead.
func (*Money) Descriptor() ([]byte, []int) {
	return file_proto_productcatalogservice_proto_rawDescGZIP(), []int{0}
}

func (x *Money) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

func (x *Money) GetUnits() int64 {
	if x != nil {
		return x.Units
	}
	return 0
}

func (x *Money) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_productcatalogservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_productcatalogservice_proto_msgTypes[1]
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
	return file_proto_productcatalogservice_proto_rawDescGZIP(), []int{1}
}

// 商品
type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 描述
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// 图片
	Picture string `protobuf:"bytes,4,opt,name=picture,proto3" json:"picture,omitempty"`
	// 价格
	PriceUsd *Money `protobuf:"bytes,5,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	// 分类
	Categories []string `protobuf:"bytes,6,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_productcatalogservice_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_productcatalogservice_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_productcatalogservice_proto_rawDescGZIP(), []int{2}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Product) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Product) GetPriceUsd() *Money {
	if x != nil {
		return x.PriceUsd
	}
	return nil
}

func (x *Product) GetCategories() []string {
	if x != nil {
		return x.Categories
	}
	return nil
}

// 商品列表响应
type ListProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 返回商品切片
	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *ListProductsResponse) Reset() {
	*x = ListProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_productcatalogservice_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProductsResponse) ProtoMessage() {}

func (x *ListProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_productcatalogservice_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProductsResponse.ProtoReflect.Descriptor instead.
func (*ListProductsResponse) Descriptor() ([]byte, []int) {
	return file_proto_productcatalogservice_proto_rawDescGZIP(), []int{3}
}

func (x *ListProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

// 根据id搜索商品请求
type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 根据id查询
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_productcatalogservice_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_productcatalogservice_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_proto_productcatalogservice_proto_rawDescGZIP(), []int{4}
}

func (x *GetProductRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// 搜索商品请求
type SearchProductsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *SearchProductsRequest) Reset() {
	*x = SearchProductsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_productcatalogservice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsRequest) ProtoMessage() {}

func (x *SearchProductsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_productcatalogservice_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsRequest.ProtoReflect.Descriptor instead.
func (*SearchProductsRequest) Descriptor() ([]byte, []int) {
	return file_proto_productcatalogservice_proto_rawDescGZIP(), []int{5}
}

func (x *SearchProductsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// 搜索商品响应
type SearchProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Product `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *SearchProductsResponse) Reset() {
	*x = SearchProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_productcatalogservice_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchProductsResponse) ProtoMessage() {}

func (x *SearchProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_productcatalogservice_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchProductsResponse.ProtoReflect.Descriptor instead.
func (*SearchProductsResponse) Descriptor() ([]byte, []int) {
	return file_proto_productcatalogservice_proto_rawDescGZIP(), []int{6}
}

func (x *SearchProductsResponse) GetResults() []*Product {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_proto_productcatalogservice_proto protoreflect.FileDescriptor

var file_proto_productcatalogservice_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x63,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x22, 0x58, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x12, 0x31, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x75, 0x73, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x52, 0x08, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x55, 0x73, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2d, 0x0a, 0x15, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x22, 0x4a, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x32, 0x8f, 0x02, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0c, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x23, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22,
	0x00, 0x12, 0x5f, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x12, 0x24, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x17, 0x5a, 0x15, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_productcatalogservice_proto_rawDescOnce sync.Once
	file_proto_productcatalogservice_proto_rawDescData = file_proto_productcatalogservice_proto_rawDesc
)

func file_proto_productcatalogservice_proto_rawDescGZIP() []byte {
	file_proto_productcatalogservice_proto_rawDescOnce.Do(func() {
		file_proto_productcatalogservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_productcatalogservice_proto_rawDescData)
	})
	return file_proto_productcatalogservice_proto_rawDescData
}

var file_proto_productcatalogservice_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_productcatalogservice_proto_goTypes = []interface{}{
	(*Money)(nil),                  // 0: microshopping.Money
	(*Empty)(nil),                  // 1: microshopping.Empty
	(*Product)(nil),                // 2: microshopping.Product
	(*ListProductsResponse)(nil),   // 3: microshopping.ListProductsResponse
	(*GetProductRequest)(nil),      // 4: microshopping.GetProductRequest
	(*SearchProductsRequest)(nil),  // 5: microshopping.SearchProductsRequest
	(*SearchProductsResponse)(nil), // 6: microshopping.SearchProductsResponse
}
var file_proto_productcatalogservice_proto_depIdxs = []int32{
	0, // 0: microshopping.Product.price_usd:type_name -> microshopping.Money
	2, // 1: microshopping.ListProductsResponse.products:type_name -> microshopping.Product
	2, // 2: microshopping.SearchProductsResponse.results:type_name -> microshopping.Product
	1, // 3: microshopping.ProductCatalogService.ListProducts:input_type -> microshopping.Empty
	4, // 4: microshopping.ProductCatalogService.GetProduct:input_type -> microshopping.GetProductRequest
	5, // 5: microshopping.ProductCatalogService.SearchProducts:input_type -> microshopping.SearchProductsRequest
	3, // 6: microshopping.ProductCatalogService.ListProducts:output_type -> microshopping.ListProductsResponse
	2, // 7: microshopping.ProductCatalogService.GetProduct:output_type -> microshopping.Product
	6, // 8: microshopping.ProductCatalogService.SearchProducts:output_type -> microshopping.SearchProductsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_productcatalogservice_proto_init() }
func file_proto_productcatalogservice_proto_init() {
	if File_proto_productcatalogservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_productcatalogservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Money); i {
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
		file_proto_productcatalogservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_productcatalogservice_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_proto_productcatalogservice_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProductsResponse); i {
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
		file_proto_productcatalogservice_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRequest); i {
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
		file_proto_productcatalogservice_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProductsRequest); i {
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
		file_proto_productcatalogservice_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchProductsResponse); i {
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
			RawDescriptor: file_proto_productcatalogservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_productcatalogservice_proto_goTypes,
		DependencyIndexes: file_proto_productcatalogservice_proto_depIdxs,
		MessageInfos:      file_proto_productcatalogservice_proto_msgTypes,
	}.Build()
	File_proto_productcatalogservice_proto = out.File
	file_proto_productcatalogservice_proto_rawDesc = nil
	file_proto_productcatalogservice_proto_goTypes = nil
	file_proto_productcatalogservice_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProductCatalogServiceClient is the client API for ProductCatalogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProductCatalogServiceClient interface {
	ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListProductsResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error)
	SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error)
}

type productCatalogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCatalogServiceClient(cc grpc.ClientConnInterface) ProductCatalogServiceClient {
	return &productCatalogServiceClient{cc}
}

func (c *productCatalogServiceClient) ListProducts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ListProductsResponse, error) {
	out := new(ListProductsResponse)
	err := c.cc.Invoke(ctx, "/microshopping.ProductCatalogService/ListProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/microshopping.ProductCatalogService/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogServiceClient) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*SearchProductsResponse, error) {
	out := new(SearchProductsResponse)
	err := c.cc.Invoke(ctx, "/microshopping.ProductCatalogService/SearchProducts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCatalogServiceServer is the server API for ProductCatalogService service.
type ProductCatalogServiceServer interface {
	ListProducts(context.Context, *Empty) (*ListProductsResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*Product, error)
	SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error)
}

// UnimplementedProductCatalogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProductCatalogServiceServer struct {
}

func (*UnimplementedProductCatalogServiceServer) ListProducts(context.Context, *Empty) (*ListProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (*UnimplementedProductCatalogServiceServer) GetProduct(context.Context, *GetProductRequest) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (*UnimplementedProductCatalogServiceServer) SearchProducts(context.Context, *SearchProductsRequest) (*SearchProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProducts not implemented")
}

func RegisterProductCatalogServiceServer(s *grpc.Server, srv ProductCatalogServiceServer) {
	s.RegisterService(&_ProductCatalogService_serviceDesc, srv)
}

func _ProductCatalogService_ListProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microshopping.ProductCatalogService/ListProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).ListProducts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microshopping.ProductCatalogService/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogService_SearchProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/microshopping.ProductCatalogService/SearchProducts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogServiceServer).SearchProducts(ctx, req.(*SearchProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductCatalogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "microshopping.ProductCatalogService",
	HandlerType: (*ProductCatalogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProducts",
			Handler:    _ProductCatalogService_ListProducts_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductCatalogService_GetProduct_Handler,
		},
		{
			MethodName: "SearchProducts",
			Handler:    _ProductCatalogService_SearchProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/productcatalogservice.proto",
}
