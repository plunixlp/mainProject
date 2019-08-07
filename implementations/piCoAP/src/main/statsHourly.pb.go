// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statsHourly.proto

package proto

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

type Device struct {
	Building             int32    `protobuf:"varint,1,opt,name=building,proto3" json:"building,omitempty"`
	Room                 int32    `protobuf:"varint,2,opt,name=room,proto3" json:"room,omitempty"`
	LedOn                bool     `protobuf:"varint,3,opt,name=ledOn,proto3" json:"ledOn,omitempty"`
	Label                string   `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_statsHourly_7e027e23c6a923b6, []int{0}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (dst *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(dst, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetBuilding() int32 {
	if m != nil {
		return m.Building
	}
	return 0
}

func (m *Device) GetRoom() int32 {
	if m != nil {
		return m.Room
	}
	return 0
}

func (m *Device) GetLedOn() bool {
	if m != nil {
		return m.LedOn
	}
	return false
}

func (m *Device) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func init() {
	proto.RegisterType((*Device)(nil), "proto.Device")
}

func init() { proto.RegisterFile("statsHourly.proto", fileDescriptor_statsHourly_7e027e23c6a923b6) }

var fileDescriptor_statsHourly_7e027e23c6a923b6 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2e, 0x49, 0x2c,
	0x29, 0xf6, 0xc8, 0x2f, 0x2d, 0xca, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05,
	0x53, 0x4a, 0x29, 0x5c, 0x6c, 0x2e, 0xa9, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0x49,
	0xa5, 0x99, 0x39, 0x29, 0x99, 0x79, 0xe9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x70, 0xbe,
	0x90, 0x10, 0x17, 0x4b, 0x51, 0x7e, 0x7e, 0xae, 0x04, 0x13, 0x58, 0x1c, 0xcc, 0x16, 0x12, 0xe1,
	0x62, 0xcd, 0x49, 0x4d, 0xf1, 0xcf, 0x93, 0x60, 0x06, 0x0a, 0x72, 0x04, 0x41, 0x38, 0x60, 0xd1,
	0xc4, 0xa4, 0xd4, 0x1c, 0x09, 0x16, 0xa0, 0x28, 0x67, 0x10, 0x84, 0x63, 0x64, 0xc9, 0xc5, 0x13,
	0x5c, 0x99, 0x97, 0x9c, 0x51, 0x94, 0x9f, 0x97, 0x59, 0x95, 0x5a, 0x24, 0xa4, 0xc9, 0xc5, 0x0a,
	0xe6, 0x0b, 0xf1, 0x42, 0x5c, 0xa3, 0x07, 0x71, 0x83, 0x14, 0x2a, 0x57, 0x89, 0xc1, 0x80, 0x31,
	0x89, 0x0d, 0x2c, 0x62, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xb8, 0x16, 0xa7, 0xc3, 0x00,
	0x00, 0x00,
}