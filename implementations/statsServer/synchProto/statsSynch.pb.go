// Code generated by protoc-gen-go. DO NOT EDIT.
// source: statsSynch.proto

package synchProto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Device struct {
	Building             int32    `protobuf:"varint,1,opt,name=building,proto3" json:"building,omitempty"`
	Room                 int32    `protobuf:"varint,2,opt,name=room,proto3" json:"room,omitempty"`
	LedOn                bool     `protobuf:"varint,3,opt,name=ledOn,proto3" json:"ledOn,omitempty"`
	Label                string   `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	OnTime               int32    `protobuf:"varint,5,opt,name=onTime,proto3" json:"onTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3ab7bbe4f463e27, []int{0}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
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

func (m *Device) GetOnTime() int32 {
	if m != nil {
		return m.OnTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Device)(nil), "synchProto.Device")
}

func init() { proto.RegisterFile("statsSynch.proto", fileDescriptor_e3ab7bbe4f463e27) }

var fileDescriptor_e3ab7bbe4f463e27 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x2e, 0x49, 0x2c,
	0x29, 0x0e, 0xae, 0xcc, 0x4b, 0xce, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2a, 0x06,
	0x71, 0x02, 0x40, 0x6c, 0xa5, 0x1a, 0x2e, 0x36, 0x97, 0xd4, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x29,
	0x2e, 0x8e, 0xa4, 0xd2, 0xcc, 0x9c, 0x94, 0xcc, 0xbc, 0x74, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6,
	0x20, 0x38, 0x5f, 0x48, 0x88, 0x8b, 0xa5, 0x28, 0x3f, 0x3f, 0x57, 0x82, 0x09, 0x2c, 0x0e, 0x66,
	0x0b, 0x89, 0x70, 0xb1, 0xe6, 0xa4, 0xa6, 0xf8, 0xe7, 0x49, 0x30, 0x03, 0x05, 0x39, 0x82, 0x20,
	0x1c, 0xb0, 0x68, 0x62, 0x52, 0x6a, 0x8e, 0x04, 0x0b, 0x50, 0x94, 0x33, 0x08, 0xc2, 0x11, 0x12,
	0xe3, 0x62, 0xcb, 0xcf, 0x0b, 0xc9, 0xcc, 0x4d, 0x95, 0x60, 0x05, 0x9b, 0x00, 0xe5, 0x19, 0x39,
	0x73, 0xf1, 0x80, 0x1d, 0x56, 0x94, 0x9f, 0x97, 0x59, 0x95, 0x5a, 0x24, 0x64, 0xcc, 0xc5, 0x0a,
	0xe6, 0x0b, 0x09, 0xe9, 0x21, 0xdc, 0xa8, 0x07, 0x71, 0xa0, 0x14, 0x16, 0x31, 0x25, 0x06, 0x03,
	0xc6, 0x24, 0x36, 0xb0, 0xaf, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x36, 0x71, 0x7e,
	0xe9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SynchronizerClient is the client API for Synchronizer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SynchronizerClient interface {
	Synch(ctx context.Context, in *Device, opts ...grpc.CallOption) (Synchronizer_SynchClient, error)
}

type synchronizerClient struct {
	cc *grpc.ClientConn
}

func NewSynchronizerClient(cc *grpc.ClientConn) SynchronizerClient {
	return &synchronizerClient{cc}
}

func (c *synchronizerClient) Synch(ctx context.Context, in *Device, opts ...grpc.CallOption) (Synchronizer_SynchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Synchronizer_serviceDesc.Streams[0], "/synchProto.Synchronizer/Synch", opts...)
	if err != nil {
		return nil, err
	}
	x := &synchronizerSynchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Synchronizer_SynchClient interface {
	Recv() (*Device, error)
	grpc.ClientStream
}

type synchronizerSynchClient struct {
	grpc.ClientStream
}

func (x *synchronizerSynchClient) Recv() (*Device, error) {
	m := new(Device)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SynchronizerServer is the server API for Synchronizer service.
type SynchronizerServer interface {
	Synch(*Device, Synchronizer_SynchServer) error
}

// UnimplementedSynchronizerServer can be embedded to have forward compatible implementations.
type UnimplementedSynchronizerServer struct {
}

func (*UnimplementedSynchronizerServer) Synch(req *Device, srv Synchronizer_SynchServer) error {
	return status.Errorf(codes.Unimplemented, "method Synch not implemented")
}

func RegisterSynchronizerServer(s *grpc.Server, srv SynchronizerServer) {
	s.RegisterService(&_Synchronizer_serviceDesc, srv)
}

func _Synchronizer_Synch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Device)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SynchronizerServer).Synch(m, &synchronizerSynchServer{stream})
}

type Synchronizer_SynchServer interface {
	Send(*Device) error
	grpc.ServerStream
}

type synchronizerSynchServer struct {
	grpc.ServerStream
}

func (x *synchronizerSynchServer) Send(m *Device) error {
	return x.ServerStream.SendMsg(m)
}

var _Synchronizer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "synchProto.Synchronizer",
	HandlerType: (*SynchronizerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Synch",
			Handler:       _Synchronizer_Synch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "statsSynch.proto",
}
