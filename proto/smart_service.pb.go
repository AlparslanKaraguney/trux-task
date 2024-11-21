// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0--rc2
// source: proto/smart_service.proto

package smartservice

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

type SmartModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Identifier string          `protobuf:"bytes,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Type       string          `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Category   string          `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	Features   []*SmartFeature `protobuf:"bytes,6,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *SmartModel) Reset() {
	*x = SmartModel{}
	mi := &file_proto_smart_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmartModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartModel) ProtoMessage() {}

func (x *SmartModel) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smart_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartModel.ProtoReflect.Descriptor instead.
func (*SmartModel) Descriptor() ([]byte, []int) {
	return file_proto_smart_service_proto_rawDescGZIP(), []int{0}
}

func (x *SmartModel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SmartModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SmartModel) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *SmartModel) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SmartModel) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *SmartModel) GetFeatures() []*SmartFeature {
	if x != nil {
		return x.Features
	}
	return nil
}

type SmartFeature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Identifier    string `protobuf:"bytes,3,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Functionality string `protobuf:"bytes,4,opt,name=functionality,proto3" json:"functionality,omitempty"`
	ModelId       int32  `protobuf:"varint,5,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
}

func (x *SmartFeature) Reset() {
	*x = SmartFeature{}
	mi := &file_proto_smart_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmartFeature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartFeature) ProtoMessage() {}

func (x *SmartFeature) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smart_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartFeature.ProtoReflect.Descriptor instead.
func (*SmartFeature) Descriptor() ([]byte, []int) {
	return file_proto_smart_service_proto_rawDescGZIP(), []int{1}
}

func (x *SmartFeature) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SmartFeature) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SmartFeature) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

func (x *SmartFeature) GetFunctionality() string {
	if x != nil {
		return x.Functionality
	}
	return ""
}

func (x *SmartFeature) GetModelId() int32 {
	if x != nil {
		return x.ModelId
	}
	return 0
}

type SmartModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model *SmartModel `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *SmartModelRequest) Reset() {
	*x = SmartModelRequest{}
	mi := &file_proto_smart_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmartModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartModelRequest) ProtoMessage() {}

func (x *SmartModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smart_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartModelRequest.ProtoReflect.Descriptor instead.
func (*SmartModelRequest) Descriptor() ([]byte, []int) {
	return file_proto_smart_service_proto_rawDescGZIP(), []int{2}
}

func (x *SmartModelRequest) GetModel() *SmartModel {
	if x != nil {
		return x.Model
	}
	return nil
}

type SmartModelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model *SmartModel `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *SmartModelResponse) Reset() {
	*x = SmartModelResponse{}
	mi := &file_proto_smart_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmartModelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartModelResponse) ProtoMessage() {}

func (x *SmartModelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smart_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartModelResponse.ProtoReflect.Descriptor instead.
func (*SmartModelResponse) Descriptor() ([]byte, []int) {
	return file_proto_smart_service_proto_rawDescGZIP(), []int{3}
}

func (x *SmartModelResponse) GetModel() *SmartModel {
	if x != nil {
		return x.Model
	}
	return nil
}

type SmartModelQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *SmartModelQuery) Reset() {
	*x = SmartModelQuery{}
	mi := &file_proto_smart_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmartModelQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartModelQuery) ProtoMessage() {}

func (x *SmartModelQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smart_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartModelQuery.ProtoReflect.Descriptor instead.
func (*SmartModelQuery) Descriptor() ([]byte, []int) {
	return file_proto_smart_service_proto_rawDescGZIP(), []int{4}
}

func (x *SmartModelQuery) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

type SmartFeatureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature *SmartFeature `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *SmartFeatureRequest) Reset() {
	*x = SmartFeatureRequest{}
	mi := &file_proto_smart_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmartFeatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartFeatureRequest) ProtoMessage() {}

func (x *SmartFeatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smart_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartFeatureRequest.ProtoReflect.Descriptor instead.
func (*SmartFeatureRequest) Descriptor() ([]byte, []int) {
	return file_proto_smart_service_proto_rawDescGZIP(), []int{5}
}

func (x *SmartFeatureRequest) GetFeature() *SmartFeature {
	if x != nil {
		return x.Feature
	}
	return nil
}

type SmartFeatureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Feature *SmartFeature `protobuf:"bytes,1,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *SmartFeatureResponse) Reset() {
	*x = SmartFeatureResponse{}
	mi := &file_proto_smart_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmartFeatureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartFeatureResponse) ProtoMessage() {}

func (x *SmartFeatureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smart_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartFeatureResponse.ProtoReflect.Descriptor instead.
func (*SmartFeatureResponse) Descriptor() ([]byte, []int) {
	return file_proto_smart_service_proto_rawDescGZIP(), []int{6}
}

func (x *SmartFeatureResponse) GetFeature() *SmartFeature {
	if x != nil {
		return x.Feature
	}
	return nil
}

type SmartFeatureQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
}

func (x *SmartFeatureQuery) Reset() {
	*x = SmartFeatureQuery{}
	mi := &file_proto_smart_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SmartFeatureQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SmartFeatureQuery) ProtoMessage() {}

func (x *SmartFeatureQuery) ProtoReflect() protoreflect.Message {
	mi := &file_proto_smart_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SmartFeatureQuery.ProtoReflect.Descriptor instead.
func (*SmartFeatureQuery) Descriptor() ([]byte, []int) {
	return file_proto_smart_service_proto_rawDescGZIP(), []int{7}
}

func (x *SmartFeatureQuery) GetIdentifier() string {
	if x != nil {
		return x.Identifier
	}
	return ""
}

var File_proto_smart_service_proto protoreflect.FileDescriptor

var file_proto_smart_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xb8, 0x01, 0x0a, 0x0a, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x36, 0x0a, 0x08,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x08, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x0c, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x11, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22,
	0x44, 0x0a, 0x12, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x31, 0x0a, 0x0f, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x22, 0x4b, 0x0a, 0x13, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x34, 0x0a, 0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x4c, 0x0a, 0x14, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x22, 0x33, 0x0a, 0x11, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x32, 0xec, 0x02, 0x0a, 0x0c, 0x53, 0x6d, 0x61,
	0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1f, 0x2e,
	0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6d, 0x61,
	0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x12, 0x1d, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x1a, 0x20, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x5b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x21, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x22, 0x2e, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_smart_service_proto_rawDescOnce sync.Once
	file_proto_smart_service_proto_rawDescData = file_proto_smart_service_proto_rawDesc
)

func file_proto_smart_service_proto_rawDescGZIP() []byte {
	file_proto_smart_service_proto_rawDescOnce.Do(func() {
		file_proto_smart_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_smart_service_proto_rawDescData)
	})
	return file_proto_smart_service_proto_rawDescData
}

var file_proto_smart_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_smart_service_proto_goTypes = []any{
	(*SmartModel)(nil),           // 0: smartservice.SmartModel
	(*SmartFeature)(nil),         // 1: smartservice.SmartFeature
	(*SmartModelRequest)(nil),    // 2: smartservice.SmartModelRequest
	(*SmartModelResponse)(nil),   // 3: smartservice.SmartModelResponse
	(*SmartModelQuery)(nil),      // 4: smartservice.SmartModelQuery
	(*SmartFeatureRequest)(nil),  // 5: smartservice.SmartFeatureRequest
	(*SmartFeatureResponse)(nil), // 6: smartservice.SmartFeatureResponse
	(*SmartFeatureQuery)(nil),    // 7: smartservice.SmartFeatureQuery
}
var file_proto_smart_service_proto_depIdxs = []int32{
	1, // 0: smartservice.SmartModel.features:type_name -> smartservice.SmartFeature
	0, // 1: smartservice.SmartModelRequest.model:type_name -> smartservice.SmartModel
	0, // 2: smartservice.SmartModelResponse.model:type_name -> smartservice.SmartModel
	1, // 3: smartservice.SmartFeatureRequest.feature:type_name -> smartservice.SmartFeature
	1, // 4: smartservice.SmartFeatureResponse.feature:type_name -> smartservice.SmartFeature
	2, // 5: smartservice.SmartService.CreateSmartModel:input_type -> smartservice.SmartModelRequest
	4, // 6: smartservice.SmartService.GetSmartModel:input_type -> smartservice.SmartModelQuery
	5, // 7: smartservice.SmartService.CreateSmartFeature:input_type -> smartservice.SmartFeatureRequest
	7, // 8: smartservice.SmartService.GetSmartFeature:input_type -> smartservice.SmartFeatureQuery
	3, // 9: smartservice.SmartService.CreateSmartModel:output_type -> smartservice.SmartModelResponse
	3, // 10: smartservice.SmartService.GetSmartModel:output_type -> smartservice.SmartModelResponse
	6, // 11: smartservice.SmartService.CreateSmartFeature:output_type -> smartservice.SmartFeatureResponse
	6, // 12: smartservice.SmartService.GetSmartFeature:output_type -> smartservice.SmartFeatureResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_smart_service_proto_init() }
func file_proto_smart_service_proto_init() {
	if File_proto_smart_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_smart_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_smart_service_proto_goTypes,
		DependencyIndexes: file_proto_smart_service_proto_depIdxs,
		MessageInfos:      file_proto_smart_service_proto_msgTypes,
	}.Build()
	File_proto_smart_service_proto = out.File
	file_proto_smart_service_proto_rawDesc = nil
	file_proto_smart_service_proto_goTypes = nil
	file_proto_smart_service_proto_depIdxs = nil
}
