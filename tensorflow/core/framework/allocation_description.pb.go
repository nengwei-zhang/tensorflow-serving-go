// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/allocation_description.proto

package framework // import "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"

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

type AllocationDescription struct {
	// Total number of bytes requested
	RequestedBytes int64 `protobuf:"varint,1,opt,name=requested_bytes,json=requestedBytes,proto3" json:"requested_bytes,omitempty"`
	// Total number of bytes allocated if known
	AllocatedBytes int64 `protobuf:"varint,2,opt,name=allocated_bytes,json=allocatedBytes,proto3" json:"allocated_bytes,omitempty"`
	// Name of the allocator used
	AllocatorName string `protobuf:"bytes,3,opt,name=allocator_name,json=allocatorName,proto3" json:"allocator_name,omitempty"`
	// Identifier of the allocated buffer if known
	AllocationId int64 `protobuf:"varint,4,opt,name=allocation_id,json=allocationId,proto3" json:"allocation_id,omitempty"`
	// Set if this tensor only has one remaining reference
	HasSingleReference bool `protobuf:"varint,5,opt,name=has_single_reference,json=hasSingleReference,proto3" json:"has_single_reference,omitempty"`
	// Address of the allocation.
	Ptr                  uint64   `protobuf:"varint,6,opt,name=ptr,proto3" json:"ptr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocationDescription) Reset()         { *m = AllocationDescription{} }
func (m *AllocationDescription) String() string { return proto.CompactTextString(m) }
func (*AllocationDescription) ProtoMessage()    {}
func (*AllocationDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_allocation_description_821408c038f72727, []int{0}
}
func (m *AllocationDescription) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationDescription.Unmarshal(m, b)
}
func (m *AllocationDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationDescription.Marshal(b, m, deterministic)
}
func (dst *AllocationDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationDescription.Merge(dst, src)
}
func (m *AllocationDescription) XXX_Size() int {
	return xxx_messageInfo_AllocationDescription.Size(m)
}
func (m *AllocationDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationDescription.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationDescription proto.InternalMessageInfo

func (m *AllocationDescription) GetRequestedBytes() int64 {
	if m != nil {
		return m.RequestedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatedBytes() int64 {
	if m != nil {
		return m.AllocatedBytes
	}
	return 0
}

func (m *AllocationDescription) GetAllocatorName() string {
	if m != nil {
		return m.AllocatorName
	}
	return ""
}

func (m *AllocationDescription) GetAllocationId() int64 {
	if m != nil {
		return m.AllocationId
	}
	return 0
}

func (m *AllocationDescription) GetHasSingleReference() bool {
	if m != nil {
		return m.HasSingleReference
	}
	return false
}

func (m *AllocationDescription) GetPtr() uint64 {
	if m != nil {
		return m.Ptr
	}
	return 0
}

func init() {
	proto.RegisterType((*AllocationDescription)(nil), "tensorflow.AllocationDescription")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/allocation_description.proto", fileDescriptor_allocation_description_821408c038f72727)
}

var fileDescriptor_allocation_description_821408c038f72727 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0x59, 0x5b, 0x8b, 0x2e, 0x56, 0x65, 0x51, 0x58, 0xf0, 0x12, 0x14, 0x31, 0xa7, 0x44,
	0x10, 0xbc, 0x79, 0x30, 0x78, 0xf1, 0x22, 0x25, 0xde, 0xbc, 0x84, 0x4d, 0x32, 0xf9, 0xc0, 0x24,
	0x13, 0x67, 0xb7, 0x14, 0xf1, 0x8f, 0xeb, 0x4d, 0x36, 0xb5, 0x9b, 0x22, 0xbd, 0x0d, 0xcf, 0x3c,
	0xfb, 0xc1, 0xfb, 0xf2, 0x7b, 0x03, 0x9d, 0x46, 0x2a, 0x1a, 0x5c, 0x85, 0x19, 0x12, 0x84, 0x05,
	0xa9, 0x16, 0x56, 0x48, 0xef, 0xa1, 0x6a, 0x1a, 0xcc, 0x94, 0xa9, 0xb1, 0x4b, 0x72, 0xd0, 0x19,
	0xd5, 0xbd, 0x9d, 0x83, 0x9e, 0xd0, 0xa0, 0xe0, 0xe3, 0xb9, 0xcb, 0x1f, 0xc6, 0xcf, 0x1f, 0x9d,
	0xfc, 0x34, 0xba, 0xe2, 0x86, 0x9f, 0x10, 0x7c, 0x2c, 0x41, 0x1b, 0xc8, 0x93, 0xf4, 0xd3, 0x80,
	0x96, 0xcc, 0x63, 0xfe, 0x24, 0x3e, 0x76, 0x38, 0xb2, 0xd4, 0x8a, 0x7f, 0xcf, 0x39, 0x71, 0x6f,
	0x2d, 0x3a, 0xbc, 0x16, 0xaf, 0xf9, 0x86, 0x20, 0x25, 0x9d, 0x6a, 0x41, 0x4e, 0x3c, 0xe6, 0x1f,
	0xc6, 0x73, 0x47, 0x5f, 0x54, 0x0b, 0xe2, 0x8a, 0xcf, 0xb7, 0xbe, 0x5f, 0xe7, 0x72, 0x3a, 0xdc,
	0x76, 0x34, 0xc2, 0xe7, 0x5c, 0xdc, 0xf2, 0xb3, 0x4a, 0xe9, 0x44, 0xd7, 0x5d, 0xd9, 0x40, 0x42,
	0x50, 0x00, 0x41, 0x97, 0x81, 0xdc, 0xf7, 0x98, 0x7f, 0x10, 0x8b, 0x4a, 0xe9, 0xd7, 0x61, 0x15,
	0x6f, 0x36, 0xe2, 0x94, 0x4f, 0x7a, 0x43, 0x72, 0xe6, 0x31, 0x7f, 0x1a, 0xdb, 0x31, 0xfa, 0xe2,
	0x12, 0xa9, 0x0c, 0xc6, 0x34, 0x02, 0x17, 0x60, 0x74, 0xb1, 0x33, 0x94, 0x85, 0xcd, 0x4f, 0x2f,
	0xd8, 0xdb, 0x43, 0x59, 0x9b, 0x6a, 0x99, 0x06, 0x19, 0xb6, 0xe1, 0x56, 0x0b, 0xbb, 0xc7, 0x12,
	0xff, 0xd5, 0xf3, 0xcd, 0x58, 0x3a, 0x1b, 0xba, 0xb8, 0xfb, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xca,
	0xac, 0x8d, 0x6f, 0xc5, 0x01, 0x00, 0x00,
}
