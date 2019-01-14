// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow_serving/util/status.proto

package tensorflow_serving

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/maxhawkins/tensorflow-serving-go/tensorflow/core/lib/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Status that corresponds to Status in
// third_party/tensorflow/core/lib/core/status.h.
type StatusProto struct {
	// Error code.
	ErrorCode core.Code `protobuf:"varint,1,opt,name=error_code,proto3,enum=tensorflow.error.Code" json:"error_code,omitempty"`
	// Error message. Will only be set if an error was encountered.
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=error_message,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusProto) Reset()         { *m = StatusProto{} }
func (m *StatusProto) String() string { return proto.CompactTextString(m) }
func (*StatusProto) ProtoMessage()    {}
func (*StatusProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_status_b3bbb28628a980f2, []int{0}
}
func (m *StatusProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusProto.Unmarshal(m, b)
}
func (m *StatusProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusProto.Marshal(b, m, deterministic)
}
func (dst *StatusProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusProto.Merge(dst, src)
}
func (m *StatusProto) XXX_Size() int {
	return xxx_messageInfo_StatusProto.Size(m)
}
func (m *StatusProto) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusProto.DiscardUnknown(m)
}

var xxx_messageInfo_StatusProto proto.InternalMessageInfo

func (m *StatusProto) GetErrorCode() core.Code {
	if m != nil {
		return m.ErrorCode
	}
	return core.Code_OK
}

func (m *StatusProto) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusProto)(nil), "tensorflow.serving.StatusProto")
}

func init() {
	proto.RegisterFile("tensorflow_serving/util/status.proto", fileDescriptor_status_b3bbb28628a980f2)
}

var fileDescriptor_status_b3bbb28628a980f2 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x49, 0xcd, 0x2b,
	0xce, 0x2f, 0x4a, 0xcb, 0xc9, 0x2f, 0x8f, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0xcc, 0x4b, 0xd7, 0x2f,
	0x2d, 0xc9, 0xcc, 0xd1, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x42, 0xa8, 0xd2, 0x83, 0xaa, 0x92, 0xd2, 0x42, 0x88, 0xe9, 0x27, 0xe7, 0x17, 0xa5,
	0xea, 0xe7, 0x64, 0x26, 0x41, 0x18, 0xa9, 0x45, 0x45, 0xf9, 0x45, 0xf1, 0xc9, 0xf9, 0x29, 0xa9,
	0x50, 0xfd, 0x4a, 0xd9, 0x5c, 0xdc, 0xc1, 0x60, 0xf3, 0x02, 0xc0, 0xc6, 0x99, 0x71, 0x71, 0x21,
	0xd4, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x89, 0xe9, 0x21, 0xd9, 0x01, 0x96, 0xd5, 0x73,
	0xce, 0x4f, 0x49, 0x0d, 0x42, 0x52, 0x29, 0xa4, 0xc2, 0xc5, 0x0b, 0xe1, 0xe5, 0xa6, 0x16, 0x17,
	0x27, 0xa6, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xa1, 0x0a, 0x3a, 0x31, 0xff, 0x60,
	0x64, 0x4c, 0x62, 0x03, 0x5b, 0x6c, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x6e, 0x2f, 0xb5,
	0xe0, 0x00, 0x00, 0x00,
}
