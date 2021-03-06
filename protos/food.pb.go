//新增菜品
//修改菜品
//按条件查询菜品
//删除菜品

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: food.proto

// user 包

package protos

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

type Food struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodId      int32   `protobuf:"varint,1,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
	FoodName    string  `protobuf:"bytes,2,opt,name=food_name,json=foodName,proto3" json:"food_name,omitempty"`
	FoodCw      float32 `protobuf:"fixed32,3,opt,name=food_cw,json=foodCw,proto3" json:"food_cw,omitempty"`
	FoodOil     float32 `protobuf:"fixed32,4,opt,name=food_oil,json=foodOil,proto3" json:"food_oil,omitempty"`
	FoodEnergy  float32 `protobuf:"fixed32,5,opt,name=food_energy,json=foodEnergy,proto3" json:"food_energy,omitempty"`
	FoodProtein float32 `protobuf:"fixed32,6,opt,name=food_protein,json=foodProtein,proto3" json:"food_protein,omitempty"`
	FoodType    string  `protobuf:"bytes,7,opt,name=food_type,json=foodType,proto3" json:"food_type,omitempty"`
}

func (x *Food) Reset() {
	*x = Food{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Food) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Food) ProtoMessage() {}

func (x *Food) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Food.ProtoReflect.Descriptor instead.
func (*Food) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{0}
}

func (x *Food) GetFoodId() int32 {
	if x != nil {
		return x.FoodId
	}
	return 0
}

func (x *Food) GetFoodName() string {
	if x != nil {
		return x.FoodName
	}
	return ""
}

func (x *Food) GetFoodCw() float32 {
	if x != nil {
		return x.FoodCw
	}
	return 0
}

func (x *Food) GetFoodOil() float32 {
	if x != nil {
		return x.FoodOil
	}
	return 0
}

func (x *Food) GetFoodEnergy() float32 {
	if x != nil {
		return x.FoodEnergy
	}
	return 0
}

func (x *Food) GetFoodProtein() float32 {
	if x != nil {
		return x.FoodProtein
	}
	return 0
}

func (x *Food) GetFoodType() string {
	if x != nil {
		return x.FoodType
	}
	return ""
}

//今天需要写完dao层和grpc，后台controller和错误码分类
type CreateFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodData []*Food `protobuf:"bytes,1,rep,name=food_data,json=foodData,proto3" json:"food_data,omitempty"`
}

func (x *CreateFoodRequest) Reset() {
	*x = CreateFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFoodRequest) ProtoMessage() {}

func (x *CreateFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFoodRequest.ProtoReflect.Descriptor instead.
func (*CreateFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFoodRequest) GetFoodData() []*Food {
	if x != nil {
		return x.FoodData
	}
	return nil
}

type CreateFoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err int32  `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateFoodResponse) Reset() {
	*x = CreateFoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFoodResponse) ProtoMessage() {}

func (x *CreateFoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFoodResponse.ProtoReflect.Descriptor instead.
func (*CreateFoodResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFoodResponse) GetErr() int32 {
	if x != nil {
		return x.Err
	}
	return 0
}

func (x *CreateFoodResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DeleteFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodId int32 `protobuf:"varint,1,opt,name=food_id,json=foodId,proto3" json:"food_id,omitempty"`
}

func (x *DeleteFoodRequest) Reset() {
	*x = DeleteFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFoodRequest) ProtoMessage() {}

func (x *DeleteFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFoodRequest.ProtoReflect.Descriptor instead.
func (*DeleteFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteFoodRequest) GetFoodId() int32 {
	if x != nil {
		return x.FoodId
	}
	return 0
}

type DeleteFoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err int32  `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteFoodResponse) Reset() {
	*x = DeleteFoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFoodResponse) ProtoMessage() {}

func (x *DeleteFoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFoodResponse.ProtoReflect.Descriptor instead.
func (*DeleteFoodResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFoodResponse) GetErr() int32 {
	if x != nil {
		return x.Err
	}
	return 0
}

func (x *DeleteFoodResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UpdateFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Food *Food `protobuf:"bytes,1,opt,name=food,proto3" json:"food,omitempty"`
}

func (x *UpdateFoodRequest) Reset() {
	*x = UpdateFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFoodRequest) ProtoMessage() {}

func (x *UpdateFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFoodRequest.ProtoReflect.Descriptor instead.
func (*UpdateFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFoodRequest) GetFood() *Food {
	if x != nil {
		return x.Food
	}
	return nil
}

type UpdateFoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err int32  `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UpdateFoodResponse) Reset() {
	*x = UpdateFoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFoodResponse) ProtoMessage() {}

func (x *UpdateFoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFoodResponse.ProtoReflect.Descriptor instead.
func (*UpdateFoodResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFoodResponse) GetErr() int32 {
	if x != nil {
		return x.Err
	}
	return 0
}

func (x *UpdateFoodResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit    int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`                       //每页多少
	Page     int32 `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`                         //页数
	HavingId int32 `protobuf:"varint,3,opt,name=having_id,json=havingId,proto3" json:"having_id,omitempty"` //菜品id
	Food     *Food `protobuf:"bytes,4,opt,name=food,proto3" json:"food,omitempty"`
}

func (x *GetFoodRequest) Reset() {
	*x = GetFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFoodRequest) ProtoMessage() {}

func (x *GetFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFoodRequest.ProtoReflect.Descriptor instead.
func (*GetFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{7}
}

func (x *GetFoodRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetFoodRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetFoodRequest) GetHavingId() int32 {
	if x != nil {
		return x.HavingId
	}
	return 0
}

func (x *GetFoodRequest) GetFood() *Food {
	if x != nil {
		return x.Food
	}
	return nil
}

type GetFoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err  int32   `protobuf:"varint,1,opt,name=err,proto3" json:"err,omitempty"`
	Msg  string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Food []*Food `protobuf:"bytes,3,rep,name=food,proto3" json:"food,omitempty"`
}

func (x *GetFoodResponse) Reset() {
	*x = GetFoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFoodResponse) ProtoMessage() {}

func (x *GetFoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFoodResponse.ProtoReflect.Descriptor instead.
func (*GetFoodResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{8}
}

func (x *GetFoodResponse) GetErr() int32 {
	if x != nil {
		return x.Err
	}
	return 0
}

func (x *GetFoodResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetFoodResponse) GetFood() []*Food {
	if x != nil {
		return x.Food
	}
	return nil
}

var File_food_proto protoreflect.FileDescriptor

var file_food_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x22, 0xd1, 0x01, 0x0a, 0x04, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x6f,
	0x6f, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6f, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x63, 0x77, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x43, 0x77, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x6f,
	0x6f, 0x64, 0x5f, 0x6f, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x66, 0x6f,
	0x6f, 0x64, 0x4f, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x65, 0x6e,
	0x65, 0x72, 0x67, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x6f, 0x6f, 0x64,
	0x45, 0x6e, 0x65, 0x72, 0x67, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x66, 0x6f,
	0x6f, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x69, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6f,
	0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f,
	0x6f, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x3c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x09, 0x66,
	0x6f, 0x6f, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x08, 0x66, 0x6f, 0x6f, 0x64,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x38, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2c,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x33, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x66,
	0x6f, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x22, 0x38, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x77, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x61, 0x76, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x1e,
	0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x22, 0x55,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x65, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1e, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52,
	0x04, 0x66, 0x6f, 0x6f, 0x64, 0x32, 0x89, 0x02, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x64, 0x12, 0x41,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12,
	0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f,
	0x6f, 0x64, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x6f, 0x64, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x64, 0x72, 0x69, 0x6e, 0x6b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_food_proto_rawDescOnce sync.Once
	file_food_proto_rawDescData = file_food_proto_rawDesc
)

func file_food_proto_rawDescGZIP() []byte {
	file_food_proto_rawDescOnce.Do(func() {
		file_food_proto_rawDescData = protoimpl.X.CompressGZIP(file_food_proto_rawDescData)
	})
	return file_food_proto_rawDescData
}

var file_food_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_food_proto_goTypes = []interface{}{
	(*Food)(nil),               // 0: user.Food
	(*CreateFoodRequest)(nil),  // 1: user.CreateFoodRequest
	(*CreateFoodResponse)(nil), // 2: user.CreateFoodResponse
	(*DeleteFoodRequest)(nil),  // 3: user.DeleteFoodRequest
	(*DeleteFoodResponse)(nil), // 4: user.DeleteFoodResponse
	(*UpdateFoodRequest)(nil),  // 5: user.UpdateFoodRequest
	(*UpdateFoodResponse)(nil), // 6: user.UpdateFoodResponse
	(*GetFoodRequest)(nil),     // 7: user.GetFoodRequest
	(*GetFoodResponse)(nil),    // 8: user.GetFoodResponse
}
var file_food_proto_depIdxs = []int32{
	0, // 0: user.CreateFoodRequest.food_data:type_name -> user.Food
	0, // 1: user.UpdateFoodRequest.food:type_name -> user.Food
	0, // 2: user.GetFoodRequest.food:type_name -> user.Food
	0, // 3: user.GetFoodResponse.food:type_name -> user.Food
	1, // 4: user.food.CreateFood:input_type -> user.CreateFoodRequest
	3, // 5: user.food.DeleteFood:input_type -> user.DeleteFoodRequest
	5, // 6: user.food.UpdateFood:input_type -> user.UpdateFoodRequest
	7, // 7: user.food.GetFood:input_type -> user.GetFoodRequest
	2, // 8: user.food.CreateFood:output_type -> user.CreateFoodResponse
	4, // 9: user.food.DeleteFood:output_type -> user.DeleteFoodResponse
	6, // 10: user.food.UpdateFood:output_type -> user.UpdateFoodResponse
	8, // 11: user.food.GetFood:output_type -> user.GetFoodResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_food_proto_init() }
func file_food_proto_init() {
	if File_food_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_food_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Food); i {
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
		file_food_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFoodRequest); i {
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
		file_food_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFoodResponse); i {
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
		file_food_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFoodRequest); i {
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
		file_food_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFoodResponse); i {
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
		file_food_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFoodRequest); i {
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
		file_food_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFoodResponse); i {
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
		file_food_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFoodRequest); i {
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
		file_food_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFoodResponse); i {
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
			RawDescriptor: file_food_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_food_proto_goTypes,
		DependencyIndexes: file_food_proto_depIdxs,
		MessageInfos:      file_food_proto_msgTypes,
	}.Build()
	File_food_proto = out.File
	file_food_proto_rawDesc = nil
	file_food_proto_goTypes = nil
	file_food_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FoodClient is the client API for Food service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FoodClient interface {
	CreateFood(ctx context.Context, in *CreateFoodRequest, opts ...grpc.CallOption) (*CreateFoodResponse, error)
	DeleteFood(ctx context.Context, in *DeleteFoodRequest, opts ...grpc.CallOption) (*DeleteFoodResponse, error)
	UpdateFood(ctx context.Context, in *UpdateFoodRequest, opts ...grpc.CallOption) (*UpdateFoodResponse, error)
	GetFood(ctx context.Context, in *GetFoodRequest, opts ...grpc.CallOption) (*GetFoodResponse, error)
}

type foodClient struct {
	cc grpc.ClientConnInterface
}

func NewFoodClient(cc grpc.ClientConnInterface) FoodClient {
	return &foodClient{cc}
}

func (c *foodClient) CreateFood(ctx context.Context, in *CreateFoodRequest, opts ...grpc.CallOption) (*CreateFoodResponse, error) {
	out := new(CreateFoodResponse)
	err := c.cc.Invoke(ctx, "/user.food/CreateFood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) DeleteFood(ctx context.Context, in *DeleteFoodRequest, opts ...grpc.CallOption) (*DeleteFoodResponse, error) {
	out := new(DeleteFoodResponse)
	err := c.cc.Invoke(ctx, "/user.food/DeleteFood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) UpdateFood(ctx context.Context, in *UpdateFoodRequest, opts ...grpc.CallOption) (*UpdateFoodResponse, error) {
	out := new(UpdateFoodResponse)
	err := c.cc.Invoke(ctx, "/user.food/UpdateFood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *foodClient) GetFood(ctx context.Context, in *GetFoodRequest, opts ...grpc.CallOption) (*GetFoodResponse, error) {
	out := new(GetFoodResponse)
	err := c.cc.Invoke(ctx, "/user.food/GetFood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FoodServer is the server API for Food service.
type FoodServer interface {
	CreateFood(context.Context, *CreateFoodRequest) (*CreateFoodResponse, error)
	DeleteFood(context.Context, *DeleteFoodRequest) (*DeleteFoodResponse, error)
	UpdateFood(context.Context, *UpdateFoodRequest) (*UpdateFoodResponse, error)
	GetFood(context.Context, *GetFoodRequest) (*GetFoodResponse, error)
}

// UnimplementedFoodServer can be embedded to have forward compatible implementations.
type UnimplementedFoodServer struct {
}

func (*UnimplementedFoodServer) CreateFood(context.Context, *CreateFoodRequest) (*CreateFoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFood not implemented")
}
func (*UnimplementedFoodServer) DeleteFood(context.Context, *DeleteFoodRequest) (*DeleteFoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFood not implemented")
}
func (*UnimplementedFoodServer) UpdateFood(context.Context, *UpdateFoodRequest) (*UpdateFoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFood not implemented")
}
func (*UnimplementedFoodServer) GetFood(context.Context, *GetFoodRequest) (*GetFoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFood not implemented")
}

func RegisterFoodServer(s *grpc.Server, srv FoodServer) {
	s.RegisterService(&_Food_serviceDesc, srv)
}

func _Food_CreateFood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).CreateFood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.food/CreateFood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).CreateFood(ctx, req.(*CreateFoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_DeleteFood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).DeleteFood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.food/DeleteFood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).DeleteFood(ctx, req.(*DeleteFoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_UpdateFood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).UpdateFood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.food/UpdateFood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).UpdateFood(ctx, req.(*UpdateFoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Food_GetFood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FoodServer).GetFood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.food/GetFood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FoodServer).GetFood(ctx, req.(*GetFoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Food_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.food",
	HandlerType: (*FoodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFood",
			Handler:    _Food_CreateFood_Handler,
		},
		{
			MethodName: "DeleteFood",
			Handler:    _Food_DeleteFood_Handler,
		},
		{
			MethodName: "UpdateFood",
			Handler:    _Food_UpdateFood_Handler,
		},
		{
			MethodName: "GetFood",
			Handler:    _Food_GetFood_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "food.proto",
}
