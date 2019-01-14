// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/device_attributes.proto

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

type InterconnectLink struct {
	DeviceId             int32    `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Strength             int32    `protobuf:"varint,3,opt,name=strength,proto3" json:"strength,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterconnectLink) Reset()         { *m = InterconnectLink{} }
func (m *InterconnectLink) String() string { return proto.CompactTextString(m) }
func (*InterconnectLink) ProtoMessage()    {}
func (*InterconnectLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_attributes_2592e169d0ddd9a2, []int{0}
}
func (m *InterconnectLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterconnectLink.Unmarshal(m, b)
}
func (m *InterconnectLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterconnectLink.Marshal(b, m, deterministic)
}
func (dst *InterconnectLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterconnectLink.Merge(dst, src)
}
func (m *InterconnectLink) XXX_Size() int {
	return xxx_messageInfo_InterconnectLink.Size(m)
}
func (m *InterconnectLink) XXX_DiscardUnknown() {
	xxx_messageInfo_InterconnectLink.DiscardUnknown(m)
}

var xxx_messageInfo_InterconnectLink proto.InternalMessageInfo

func (m *InterconnectLink) GetDeviceId() int32 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *InterconnectLink) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *InterconnectLink) GetStrength() int32 {
	if m != nil {
		return m.Strength
	}
	return 0
}

type LocalLinks struct {
	Link                 []*InterconnectLink `protobuf:"bytes,1,rep,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LocalLinks) Reset()         { *m = LocalLinks{} }
func (m *LocalLinks) String() string { return proto.CompactTextString(m) }
func (*LocalLinks) ProtoMessage()    {}
func (*LocalLinks) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_attributes_2592e169d0ddd9a2, []int{1}
}
func (m *LocalLinks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalLinks.Unmarshal(m, b)
}
func (m *LocalLinks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalLinks.Marshal(b, m, deterministic)
}
func (dst *LocalLinks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalLinks.Merge(dst, src)
}
func (m *LocalLinks) XXX_Size() int {
	return xxx_messageInfo_LocalLinks.Size(m)
}
func (m *LocalLinks) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalLinks.DiscardUnknown(m)
}

var xxx_messageInfo_LocalLinks proto.InternalMessageInfo

func (m *LocalLinks) GetLink() []*InterconnectLink {
	if m != nil {
		return m.Link
	}
	return nil
}

type DeviceLocality struct {
	// Optional bus locality of device.  Default value of 0 means
	// no specific locality.  Specific localities are indexed from 1.
	BusId int32 `protobuf:"varint,1,opt,name=bus_id,json=busId,proto3" json:"bus_id,omitempty"`
	// Optional NUMA locality of device.
	NumaNode int32 `protobuf:"varint,2,opt,name=numa_node,json=numaNode,proto3" json:"numa_node,omitempty"`
	// Optional local interconnect links to other devices.
	Links                *LocalLinks `protobuf:"bytes,3,opt,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DeviceLocality) Reset()         { *m = DeviceLocality{} }
func (m *DeviceLocality) String() string { return proto.CompactTextString(m) }
func (*DeviceLocality) ProtoMessage()    {}
func (*DeviceLocality) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_attributes_2592e169d0ddd9a2, []int{2}
}
func (m *DeviceLocality) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceLocality.Unmarshal(m, b)
}
func (m *DeviceLocality) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceLocality.Marshal(b, m, deterministic)
}
func (dst *DeviceLocality) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceLocality.Merge(dst, src)
}
func (m *DeviceLocality) XXX_Size() int {
	return xxx_messageInfo_DeviceLocality.Size(m)
}
func (m *DeviceLocality) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceLocality.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceLocality proto.InternalMessageInfo

func (m *DeviceLocality) GetBusId() int32 {
	if m != nil {
		return m.BusId
	}
	return 0
}

func (m *DeviceLocality) GetNumaNode() int32 {
	if m != nil {
		return m.NumaNode
	}
	return 0
}

func (m *DeviceLocality) GetLinks() *LocalLinks {
	if m != nil {
		return m.Links
	}
	return nil
}

type DeviceAttributes struct {
	// Fully specified name of the device within a cluster.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// String representation of device_type.
	DeviceType string `protobuf:"bytes,2,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	// Memory capacity of device in bytes.
	MemoryLimit int64 `protobuf:"varint,4,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	// Platform-specific data about device that may be useful
	// for supporting efficient data transfers.
	Locality *DeviceLocality `protobuf:"bytes,5,opt,name=locality,proto3" json:"locality,omitempty"`
	// A device is assigned a global unique number each time it is
	// initialized. "incarnation" should never be 0.
	Incarnation uint64 `protobuf:"fixed64,6,opt,name=incarnation,proto3" json:"incarnation,omitempty"`
	// String representation of the physical device that this device maps to.
	PhysicalDeviceDesc   string   `protobuf:"bytes,7,opt,name=physical_device_desc,json=physicalDeviceDesc,proto3" json:"physical_device_desc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceAttributes) Reset()         { *m = DeviceAttributes{} }
func (m *DeviceAttributes) String() string { return proto.CompactTextString(m) }
func (*DeviceAttributes) ProtoMessage()    {}
func (*DeviceAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_device_attributes_2592e169d0ddd9a2, []int{3}
}
func (m *DeviceAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceAttributes.Unmarshal(m, b)
}
func (m *DeviceAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceAttributes.Marshal(b, m, deterministic)
}
func (dst *DeviceAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceAttributes.Merge(dst, src)
}
func (m *DeviceAttributes) XXX_Size() int {
	return xxx_messageInfo_DeviceAttributes.Size(m)
}
func (m *DeviceAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceAttributes proto.InternalMessageInfo

func (m *DeviceAttributes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceAttributes) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *DeviceAttributes) GetMemoryLimit() int64 {
	if m != nil {
		return m.MemoryLimit
	}
	return 0
}

func (m *DeviceAttributes) GetLocality() *DeviceLocality {
	if m != nil {
		return m.Locality
	}
	return nil
}

func (m *DeviceAttributes) GetIncarnation() uint64 {
	if m != nil {
		return m.Incarnation
	}
	return 0
}

func (m *DeviceAttributes) GetPhysicalDeviceDesc() string {
	if m != nil {
		return m.PhysicalDeviceDesc
	}
	return ""
}

func init() {
	proto.RegisterType((*InterconnectLink)(nil), "tensorflow.InterconnectLink")
	proto.RegisterType((*LocalLinks)(nil), "tensorflow.LocalLinks")
	proto.RegisterType((*DeviceLocality)(nil), "tensorflow.DeviceLocality")
	proto.RegisterType((*DeviceAttributes)(nil), "tensorflow.DeviceAttributes")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/device_attributes.proto", fileDescriptor_device_attributes_2592e169d0ddd9a2)
}

var fileDescriptor_device_attributes_2592e169d0ddd9a2 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x6b, 0xdb, 0x30,
	0x14, 0xc6, 0xd1, 0x12, 0x67, 0xc9, 0xf3, 0x18, 0x45, 0x6c, 0x45, 0x74, 0x83, 0x79, 0x39, 0xf9,
	0x30, 0x92, 0xae, 0x83, 0xdd, 0x36, 0x58, 0xe9, 0x25, 0x10, 0x46, 0x11, 0x3b, 0xed, 0x62, 0x64,
	0x59, 0x4d, 0x44, 0x6c, 0x29, 0x48, 0x72, 0x4b, 0xfe, 0xf1, 0xb1, 0xe3, 0x90, 0xe4, 0xc4, 0x5e,
	0xe8, 0xed, 0xf9, 0xf3, 0xd3, 0xfb, 0x7e, 0xef, 0xf1, 0xc1, 0x67, 0x27, 0x94, 0xd5, 0xe6, 0xa1,
	0xd6, 0x4f, 0x4b, 0xae, 0x8d, 0x58, 0x3e, 0x18, 0xd6, 0x88, 0x27, 0x6d, 0x76, 0xcb, 0x4a, 0x3c,
	0x4a, 0x2e, 0x0a, 0xe6, 0x9c, 0x91, 0x65, 0xeb, 0x84, 0x5d, 0xec, 0x8d, 0x76, 0x1a, 0x43, 0xff,
	0x64, 0x5e, 0xc0, 0xc5, 0x4a, 0x39, 0x61, 0xb8, 0x56, 0x4a, 0x70, 0xb7, 0x96, 0x6a, 0x87, 0xdf,
	0xc1, 0xac, 0x7b, 0x2a, 0x2b, 0x82, 0x32, 0x94, 0x27, 0x74, 0x1a, 0x85, 0x55, 0x85, 0x31, 0x8c,
	0xdd, 0x61, 0x2f, 0xc8, 0x8b, 0x0c, 0xe5, 0x33, 0x1a, 0x6a, 0x7c, 0x05, 0x53, 0xeb, 0x8c, 0x50,
	0x1b, 0xb7, 0x25, 0xa3, 0xd8, 0x7f, 0xfc, 0x9e, 0x7f, 0x07, 0x58, 0x6b, 0xce, 0x6a, 0x3f, 0xd9,
	0xe2, 0x6b, 0x18, 0xd7, 0x52, 0xed, 0x08, 0xca, 0x46, 0x79, 0x7a, 0xf3, 0x7e, 0xd1, 0x93, 0x2c,
	0xce, 0x31, 0x68, 0xe8, 0x9c, 0x1b, 0x78, 0x7d, 0x17, 0xbc, 0xc3, 0x14, 0xe9, 0x0e, 0xf8, 0x2d,
	0x4c, 0xca, 0xd6, 0xf6, 0x6c, 0x49, 0xd9, 0xda, 0x55, 0xe5, 0xa9, 0x55, 0xdb, 0xb0, 0x42, 0xe9,
	0x2a, 0xd2, 0x25, 0x74, 0xea, 0x85, 0x9f, 0xba, 0x12, 0xf8, 0x13, 0x24, 0x7e, 0x9a, 0x0d, 0x78,
	0xe9, 0xcd, 0xe5, 0xd0, 0xb8, 0xc7, 0xa3, 0xb1, 0x69, 0xfe, 0x07, 0xc1, 0x45, 0x34, 0xfd, 0x71,
	0xba, 0x9d, 0x5f, 0x5c, 0xb1, 0x46, 0x04, 0xd3, 0x19, 0x0d, 0x35, 0xfe, 0x00, 0x69, 0x77, 0xa9,
	0xc1, 0x4d, 0x20, 0x4a, 0xbf, 0xfc, 0x65, 0x3e, 0xc2, 0xab, 0x46, 0x34, 0xda, 0x1c, 0x8a, 0x5a,
	0x36, 0xd2, 0x91, 0x71, 0x86, 0xf2, 0x11, 0x4d, 0xa3, 0xb6, 0xf6, 0x12, 0xfe, 0x0a, 0xd3, 0xba,
	0x5b, 0x8d, 0x24, 0x81, 0xee, 0x6a, 0x48, 0xf7, 0xff, 0xf2, 0xf4, 0xd4, 0x8b, 0x33, 0x48, 0xa5,
	0xe2, 0xcc, 0x28, 0xe6, 0xa4, 0x56, 0x64, 0x92, 0xa1, 0x7c, 0x42, 0x87, 0x12, 0xbe, 0x86, 0x37,
	0xfb, 0xed, 0xc1, 0x4a, 0xce, 0xea, 0xa2, 0xc3, 0xac, 0x84, 0xe5, 0xe4, 0x65, 0xc0, 0xc4, 0xc7,
	0x7f, 0xd1, 0xe1, 0x4e, 0x58, 0x7e, 0xfb, 0x08, 0x44, 0x9b, 0xcd, 0xd0, 0xfe, 0x94, 0xa6, 0xdb,
	0xcb, 0xf3, 0x8b, 0xdc, 0xfb, 0x30, 0xd9, 0x7b, 0xf4, 0xfb, 0xdb, 0x46, 0xba, 0x6d, 0x5b, 0x2e,
	0xb8, 0x6e, 0x96, 0x83, 0x34, 0x3e, 0x5f, 0x6e, 0xf4, 0x59, 0x4c, 0xff, 0x22, 0x54, 0x4e, 0x42,
	0x30, 0xbf, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x5f, 0x72, 0x19, 0xcd, 0x02, 0x00, 0x00,
}
