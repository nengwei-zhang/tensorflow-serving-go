// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/apis/classification.proto

package tensorflow_serving

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A single class.
type Class struct {
	// Label or name of the class.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// Score for this class (e.g., the probability the item belongs to this
	// class). As per the proto3 default-value semantics, if the score is missing,
	// it should be treated as 0.
	Score                float32  `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Class) Reset()         { *m = Class{} }
func (m *Class) String() string { return proto.CompactTextString(m) }
func (*Class) ProtoMessage()    {}
func (*Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_classification_103eb944e683585f, []int{0}
}
func (m *Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Class.Unmarshal(m, b)
}
func (m *Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Class.Marshal(b, m, deterministic)
}
func (dst *Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Class.Merge(dst, src)
}
func (m *Class) XXX_Size() int {
	return xxx_messageInfo_Class.Size(m)
}
func (m *Class) XXX_DiscardUnknown() {
	xxx_messageInfo_Class.DiscardUnknown(m)
}

var xxx_messageInfo_Class proto.InternalMessageInfo

func (m *Class) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Class) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// List of classes for a single item (tensorflow.Example).
type Classifications struct {
	Classes              []*Class `protobuf:"bytes,1,rep,name=classes,proto3" json:"classes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Classifications) Reset()         { *m = Classifications{} }
func (m *Classifications) String() string { return proto.CompactTextString(m) }
func (*Classifications) ProtoMessage()    {}
func (*Classifications) Descriptor() ([]byte, []int) {
	return fileDescriptor_classification_103eb944e683585f, []int{1}
}
func (m *Classifications) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Classifications.Unmarshal(m, b)
}
func (m *Classifications) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Classifications.Marshal(b, m, deterministic)
}
func (dst *Classifications) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Classifications.Merge(dst, src)
}
func (m *Classifications) XXX_Size() int {
	return xxx_messageInfo_Classifications.Size(m)
}
func (m *Classifications) XXX_DiscardUnknown() {
	xxx_messageInfo_Classifications.DiscardUnknown(m)
}

var xxx_messageInfo_Classifications proto.InternalMessageInfo

func (m *Classifications) GetClasses() []*Class {
	if m != nil {
		return m.Classes
	}
	return nil
}

// Contains one result per input example, in the same order as the input in
// ClassificationRequest.
type ClassificationResult struct {
	Classifications      []*Classifications `protobuf:"bytes,1,rep,name=classifications,proto3" json:"classifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClassificationResult) Reset()         { *m = ClassificationResult{} }
func (m *ClassificationResult) String() string { return proto.CompactTextString(m) }
func (*ClassificationResult) ProtoMessage()    {}
func (*ClassificationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_classification_103eb944e683585f, []int{2}
}
func (m *ClassificationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationResult.Unmarshal(m, b)
}
func (m *ClassificationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationResult.Marshal(b, m, deterministic)
}
func (dst *ClassificationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationResult.Merge(dst, src)
}
func (m *ClassificationResult) XXX_Size() int {
	return xxx_messageInfo_ClassificationResult.Size(m)
}
func (m *ClassificationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationResult.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationResult proto.InternalMessageInfo

func (m *ClassificationResult) GetClassifications() []*Classifications {
	if m != nil {
		return m.Classifications
	}
	return nil
}

type ClassificationRequest struct {
	// Model Specification. If version is not specified, will use the latest
	// (numerical) version.
	ModelSpec *ModelSpec `protobuf:"bytes,1,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Input data.
	Input                *Input   `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassificationRequest) Reset()         { *m = ClassificationRequest{} }
func (m *ClassificationRequest) String() string { return proto.CompactTextString(m) }
func (*ClassificationRequest) ProtoMessage()    {}
func (*ClassificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_classification_103eb944e683585f, []int{3}
}
func (m *ClassificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationRequest.Unmarshal(m, b)
}
func (m *ClassificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationRequest.Marshal(b, m, deterministic)
}
func (dst *ClassificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationRequest.Merge(dst, src)
}
func (m *ClassificationRequest) XXX_Size() int {
	return xxx_messageInfo_ClassificationRequest.Size(m)
}
func (m *ClassificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationRequest proto.InternalMessageInfo

func (m *ClassificationRequest) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *ClassificationRequest) GetInput() *Input {
	if m != nil {
		return m.Input
	}
	return nil
}

type ClassificationResponse struct {
	// Effective Model Specification used for classification.
	ModelSpec *ModelSpec `protobuf:"bytes,2,opt,name=model_spec,json=modelSpec,proto3" json:"model_spec,omitempty"`
	// Result of the classification.
	Result               *ClassificationResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ClassificationResponse) Reset()         { *m = ClassificationResponse{} }
func (m *ClassificationResponse) String() string { return proto.CompactTextString(m) }
func (*ClassificationResponse) ProtoMessage()    {}
func (*ClassificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_classification_103eb944e683585f, []int{4}
}
func (m *ClassificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassificationResponse.Unmarshal(m, b)
}
func (m *ClassificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassificationResponse.Marshal(b, m, deterministic)
}
func (dst *ClassificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassificationResponse.Merge(dst, src)
}
func (m *ClassificationResponse) XXX_Size() int {
	return xxx_messageInfo_ClassificationResponse.Size(m)
}
func (m *ClassificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClassificationResponse proto.InternalMessageInfo

func (m *ClassificationResponse) GetModelSpec() *ModelSpec {
	if m != nil {
		return m.ModelSpec
	}
	return nil
}

func (m *ClassificationResponse) GetResult() *ClassificationResult {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*Class)(nil), "tensorflow.serving.Class")
	proto.RegisterType((*Classifications)(nil), "tensorflow.serving.Classifications")
	proto.RegisterType((*ClassificationResult)(nil), "tensorflow.serving.ClassificationResult")
	proto.RegisterType((*ClassificationRequest)(nil), "tensorflow.serving.ClassificationRequest")
	proto.RegisterType((*ClassificationResponse)(nil), "tensorflow.serving.ClassificationResponse")
}

func init() {
	proto.RegisterFile("tensorflow_serving/apis/classification.proto", fileDescriptor_classification_103eb944e683585f)
}

var fileDescriptor_classification_103eb944e683585f = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x99, 0x96, 0xf6, 0xa7, 0xb7, 0x8b, 0xc2, 0xd0, 0x5f, 0xaa, 0x20, 0x94, 0x74, 0x93,
	0x85, 0x24, 0xd0, 0x6c, 0x5d, 0x88, 0x05, 0xc1, 0x45, 0x37, 0xe3, 0x03, 0x94, 0x34, 0xde, 0xca,
	0xc0, 0x34, 0x33, 0xe6, 0x4e, 0xf4, 0x0d, 0x7c, 0x06, 0x1f, 0xd5, 0xa5, 0x64, 0x26, 0xa1, 0x24,
	0x35, 0x88, 0xbb, 0xdc, 0xf0, 0xdd, 0x73, 0xcf, 0x39, 0x0c, 0xdc, 0x58, 0xcc, 0x49, 0x17, 0x07,
	0xa5, 0xdf, 0x77, 0x84, 0xc5, 0x9b, 0xcc, 0x5f, 0xe2, 0xd4, 0x48, 0x8a, 0x33, 0x95, 0x12, 0xc9,
	0x83, 0xcc, 0x52, 0x2b, 0x75, 0x1e, 0x99, 0x42, 0x5b, 0xcd, 0xf9, 0x89, 0x8e, 0x6a, 0xfa, 0x6a,
	0xd5, 0xa7, 0x20, 0x73, 0x53, 0x5a, 0xbf, 0xd8, 0x0f, 0x1d, 0xf5, 0x33, 0x2a, 0x0f, 0x05, 0x09,
	0x8c, 0x36, 0xd5, 0x55, 0x3e, 0x87, 0x91, 0x4a, 0xf7, 0xa8, 0x16, 0x6c, 0xc9, 0xc2, 0x89, 0xf0,
	0x43, 0xf5, 0x97, 0x32, 0x5d, 0xe0, 0x62, 0xb0, 0x64, 0xe1, 0x40, 0xf8, 0x21, 0x78, 0x80, 0xd9,
	0xa6, 0x65, 0x95, 0x78, 0x02, 0xff, 0x9c, 0x7b, 0xa4, 0x05, 0x5b, 0x0e, 0xc3, 0xe9, 0xfa, 0x32,
	0x3a, 0xf7, 0x1d, 0xb9, 0x2d, 0xd1, 0x90, 0x01, 0xc2, 0xbc, 0xad, 0x23, 0x90, 0x4a, 0x65, 0xf9,
	0x16, 0x66, 0xed, 0x2a, 0x1a, 0xd1, 0x55, 0xaf, 0xe8, 0x09, 0x15, 0xdd, 0xdd, 0xe0, 0x83, 0xc1,
	0xff, 0xee, 0x9d, 0xd7, 0x12, 0xc9, 0xf2, 0x5b, 0x00, 0x57, 0xc6, 0x8e, 0x0c, 0x66, 0x2e, 0xf9,
	0x74, 0x7d, 0xfd, 0xd3, 0x8d, 0x6d, 0x45, 0x3d, 0x19, 0xcc, 0xc4, 0xe4, 0xd8, 0x7c, 0xf2, 0x18,
	0x46, 0xae, 0x6f, 0x57, 0x4e, 0x4f, 0xe2, 0xc7, 0x0a, 0x10, 0x9e, 0x0b, 0x3e, 0x19, 0x5c, 0x9c,
	0x05, 0x36, 0x3a, 0x27, 0xec, 0x38, 0x19, 0xfc, 0xd1, 0xc9, 0x1d, 0x8c, 0x0b, 0x57, 0x5d, 0x9d,
	0x21, 0xfc, 0xbd, 0x27, 0x5f, 0xb5, 0xa8, 0xf7, 0xee, 0x87, 0x5f, 0x8c, 0xed, 0xc7, 0xee, 0x4d,
	0x24, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa9, 0x20, 0xe3, 0xfc, 0xa1, 0x02, 0x00, 0x00,
}
