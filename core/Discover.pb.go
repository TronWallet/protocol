// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/Discover.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Endpoint struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	NodeId               []byte   `protobuf:"bytes,3,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

type PingMessage struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Endpoint `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Version              int32     `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Timestamp            int64     `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{1}
}

func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (m *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(m, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PingMessage) GetTo() *Endpoint {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PingMessage) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PingMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PongMessage struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Echo                 int32     `protobuf:"varint,2,opt,name=echo,proto3" json:"echo,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PongMessage) Reset()         { *m = PongMessage{} }
func (m *PongMessage) String() string { return proto.CompactTextString(m) }
func (*PongMessage) ProtoMessage()    {}
func (*PongMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{2}
}

func (m *PongMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongMessage.Unmarshal(m, b)
}
func (m *PongMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongMessage.Marshal(b, m, deterministic)
}
func (m *PongMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongMessage.Merge(m, src)
}
func (m *PongMessage) XXX_Size() int {
	return xxx_messageInfo_PongMessage.Size(m)
}
func (m *PongMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PongMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PongMessage proto.InternalMessageInfo

func (m *PongMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PongMessage) GetEcho() int32 {
	if m != nil {
		return m.Echo
	}
	return 0
}

func (m *PongMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type FindNeighbours struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	TargetId             []byte    `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FindNeighbours) Reset()         { *m = FindNeighbours{} }
func (m *FindNeighbours) String() string { return proto.CompactTextString(m) }
func (*FindNeighbours) ProtoMessage()    {}
func (*FindNeighbours) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{3}
}

func (m *FindNeighbours) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNeighbours.Unmarshal(m, b)
}
func (m *FindNeighbours) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNeighbours.Marshal(b, m, deterministic)
}
func (m *FindNeighbours) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNeighbours.Merge(m, src)
}
func (m *FindNeighbours) XXX_Size() int {
	return xxx_messageInfo_FindNeighbours.Size(m)
}
func (m *FindNeighbours) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNeighbours.DiscardUnknown(m)
}

var xxx_messageInfo_FindNeighbours proto.InternalMessageInfo

func (m *FindNeighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *FindNeighbours) GetTargetId() []byte {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *FindNeighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Neighbours struct {
	From                 *Endpoint   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Neighbours           []*Endpoint `protobuf:"bytes,2,rep,name=neighbours,proto3" json:"neighbours,omitempty"`
	Timestamp            int64       `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Neighbours) Reset()         { *m = Neighbours{} }
func (m *Neighbours) String() string { return proto.CompactTextString(m) }
func (*Neighbours) ProtoMessage()    {}
func (*Neighbours) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{4}
}

func (m *Neighbours) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Neighbours.Unmarshal(m, b)
}
func (m *Neighbours) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Neighbours.Marshal(b, m, deterministic)
}
func (m *Neighbours) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Neighbours.Merge(m, src)
}
func (m *Neighbours) XXX_Size() int {
	return xxx_messageInfo_Neighbours.Size(m)
}
func (m *Neighbours) XXX_DiscardUnknown() {
	xxx_messageInfo_Neighbours.DiscardUnknown(m)
}

var xxx_messageInfo_Neighbours proto.InternalMessageInfo

func (m *Neighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Neighbours) GetNeighbours() []*Endpoint {
	if m != nil {
		return m.Neighbours
	}
	return nil
}

func (m *Neighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type BackupMessage struct {
	Flag                 bool     `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Priority             int32    `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackupMessage) Reset()         { *m = BackupMessage{} }
func (m *BackupMessage) String() string { return proto.CompactTextString(m) }
func (*BackupMessage) ProtoMessage()    {}
func (*BackupMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{5}
}

func (m *BackupMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupMessage.Unmarshal(m, b)
}
func (m *BackupMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupMessage.Marshal(b, m, deterministic)
}
func (m *BackupMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupMessage.Merge(m, src)
}
func (m *BackupMessage) XXX_Size() int {
	return xxx_messageInfo_BackupMessage.Size(m)
}
func (m *BackupMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BackupMessage proto.InternalMessageInfo

func (m *BackupMessage) GetFlag() bool {
	if m != nil {
		return m.Flag
	}
	return false
}

func (m *BackupMessage) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "protocol.Endpoint")
	proto.RegisterType((*PingMessage)(nil), "protocol.PingMessage")
	proto.RegisterType((*PongMessage)(nil), "protocol.PongMessage")
	proto.RegisterType((*FindNeighbours)(nil), "protocol.FindNeighbours")
	proto.RegisterType((*Neighbours)(nil), "protocol.Neighbours")
	proto.RegisterType((*BackupMessage)(nil), "protocol.BackupMessage")
}

func init() { proto.RegisterFile("core/Discover.proto", fileDescriptor_0471e8a4adb61f9e) }

var fileDescriptor_0471e8a4adb61f9e = []byte{
	// 359 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xd1, 0x1f, 0xbb, 0xea, 0xd8, 0x6d, 0x61, 0x0b, 0x45, 0x94, 0x1e, 0x84, 0x0a, 0xc5,
	0x27, 0x09, 0xdc, 0x07, 0x68, 0x31, 0x49, 0xc0, 0x87, 0x04, 0x23, 0x02, 0x81, 0xdc, 0x64, 0x69,
	0xbd, 0x5e, 0x22, 0xed, 0x88, 0xd9, 0xb5, 0x21, 0x2f, 0x90, 0x7b, 0xde, 0x38, 0x68, 0xe3, 0x75,
	0xfe, 0x90, 0xe0, 0x24, 0x27, 0xcd, 0xa7, 0xfd, 0x76, 0x7f, 0xdf, 0x0c, 0x03, 0xdf, 0x2b, 0x24,
	0x9e, 0x1f, 0x49, 0x5d, 0xe1, 0x96, 0x53, 0xd6, 0x11, 0x1a, 0x64, 0x91, 0xfd, 0x54, 0xd8, 0xa4,
	0x0b, 0x88, 0x8e, 0x55, 0xdd, 0xa1, 0x54, 0x86, 0xc5, 0xf0, 0xa9, 0xac, 0x6b, 0xe2, 0x5a, 0xc7,
	0x5e, 0xe2, 0x4d, 0xc6, 0x85, 0x93, 0x8c, 0x41, 0xd8, 0x21, 0x99, 0xd8, 0x4f, 0xbc, 0xc9, 0xa0,
	0xb0, 0x35, 0xfb, 0x01, 0x43, 0x85, 0x35, 0x9f, 0xd7, 0x71, 0x60, 0xcd, 0x3b, 0x95, 0xde, 0x7a,
	0x30, 0x5a, 0x48, 0x25, 0x4e, 0xb9, 0xd6, 0xa5, 0xe0, 0xec, 0x0f, 0x84, 0x2b, 0xc2, 0xd6, 0x3e,
	0x39, 0x9a, 0xb2, 0xcc, 0xa1, 0x33, 0xc7, 0x2d, 0xec, 0x39, 0x4b, 0xc1, 0x37, 0x68, 0x09, 0x2f,
	0xbb, 0x7c, 0x83, 0x7d, 0xc2, 0x2d, 0x27, 0x2d, 0x51, 0x59, 0xe8, 0xa0, 0x70, 0x92, 0xfd, 0x82,
	0xcf, 0x46, 0xb6, 0x5c, 0x9b, 0xb2, 0xed, 0xe2, 0x30, 0xf1, 0x26, 0x41, 0xf1, 0xf0, 0x23, 0x15,
	0x30, 0x5a, 0xe0, 0xfb, 0x23, 0x31, 0x08, 0x79, 0xb5, 0x46, 0xd7, 0x76, 0x5f, 0x3f, 0x05, 0x05,
	0xcf, 0x41, 0x04, 0x5f, 0x4f, 0xa4, 0xaa, 0xcf, 0xb8, 0x14, 0xeb, 0x25, 0x6e, 0x48, 0xbf, 0x99,
	0xf5, 0x13, 0x22, 0x53, 0x92, 0xe0, 0x66, 0x5e, 0x5b, 0xde, 0xb8, 0xd8, 0xeb, 0x03, 0xcc, 0x1b,
	0x0f, 0xe0, 0x03, 0xc0, 0x29, 0x80, 0xda, 0xdf, 0x8a, 0xfd, 0x24, 0x78, 0xc5, 0xfd, 0xc8, 0x75,
	0x20, 0xc8, 0x3f, 0xf8, 0x32, 0x2b, 0xab, 0xab, 0x4d, 0xe7, 0xe6, 0xcc, 0x20, 0x5c, 0x35, 0xa5,
	0xb0, 0x51, 0xa2, 0xc2, 0xd6, 0x7d, 0x9f, 0x1d, 0x49, 0x24, 0x69, 0xae, 0x77, 0x73, 0xdd, 0xeb,
	0xd9, 0x7f, 0xf8, 0x86, 0x24, 0x32, 0x43, 0xa8, 0xee, 0x83, 0xe8, 0x59, 0xe4, 0x36, 0xf7, 0xf2,
	0xb7, 0x90, 0x66, 0xbd, 0x59, 0x66, 0x15, 0xb6, 0xf9, 0x39, 0xa1, 0xba, 0x28, 0x9b, 0x86, 0x9b,
	0xdc, 0x05, 0xce, 0xfb, 0x3d, 0x5f, 0x0e, 0xad, 0xfc, 0x7b, 0x17, 0x00, 0x00, 0xff, 0xff, 0xaa,
	0xce, 0x1f, 0xc8, 0xf6, 0x02, 0x00, 0x00,
}
