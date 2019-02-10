// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/ponger/service.proto

package ponger

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

type PongRequest struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Delay                int64    `protobuf:"varint,2,opt,name=delay,proto3" json:"delay,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongRequest) Reset()         { *m = PongRequest{} }
func (m *PongRequest) String() string { return proto.CompactTextString(m) }
func (*PongRequest) ProtoMessage()    {}
func (*PongRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3b54b7b531281c, []int{0}
}

func (m *PongRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongRequest.Unmarshal(m, b)
}
func (m *PongRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongRequest.Marshal(b, m, deterministic)
}
func (m *PongRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongRequest.Merge(m, src)
}
func (m *PongRequest) XXX_Size() int {
	return xxx_messageInfo_PongRequest.Size(m)
}
func (m *PongRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PongRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PongRequest proto.InternalMessageInfo

func (m *PongRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PongRequest) GetDelay() int64 {
	if m != nil {
		return m.Delay
	}
	return 0
}

type PongResponse struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongResponse) Reset()         { *m = PongResponse{} }
func (m *PongResponse) String() string { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()    {}
func (*PongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b3b54b7b531281c, []int{1}
}

func (m *PongResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongResponse.Unmarshal(m, b)
}
func (m *PongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongResponse.Marshal(b, m, deterministic)
}
func (m *PongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongResponse.Merge(m, src)
}
func (m *PongResponse) XXX_Size() int {
	return xxx_messageInfo_PongResponse.Size(m)
}
func (m *PongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PongResponse proto.InternalMessageInfo

func (m *PongResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*PongRequest)(nil), "iheanyi.goopentracingexample.ponger.PongRequest")
	proto.RegisterType((*PongResponse)(nil), "iheanyi.goopentracingexample.ponger.PongResponse")
}

func init() { proto.RegisterFile("rpc/ponger/service.proto", fileDescriptor_4b3b54b7b531281c) }

var fileDescriptor_4b3b54b7b531281c = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x31, 0x6f, 0xc2, 0x30,
	0x10, 0x85, 0x95, 0x36, 0x4d, 0xdb, 0x6b, 0x27, 0xab, 0x83, 0xd5, 0x29, 0x4a, 0x97, 0x4c, 0x4e,
	0x81, 0x99, 0x85, 0x5f, 0x10, 0x65, 0x64, 0x73, 0x92, 0x93, 0xb1, 0x48, 0x7c, 0xc6, 0x4e, 0x10,
	0xf9, 0xf7, 0x88, 0x18, 0x24, 0x46, 0xd8, 0xde, 0x3b, 0xe9, 0xd3, 0x7d, 0x7a, 0xc0, 0x9d, 0x6d,
	0x0a, 0x4b, 0x46, 0xa1, 0x2b, 0x3c, 0xba, 0xa3, 0x6e, 0x50, 0x58, 0x47, 0x03, 0xb1, 0x3f, 0xbd,
	0x43, 0x69, 0x26, 0x2d, 0x14, 0x91, 0x45, 0x33, 0x38, 0xd9, 0x68, 0xa3, 0xf0, 0x24, 0x7b, 0xdb,
	0xa1, 0x08, 0x48, 0xb6, 0x86, 0xaf, 0x92, 0x8c, 0xaa, 0xf0, 0x30, 0xa2, 0x1f, 0x18, 0x87, 0xf7,
	0x1e, 0xbd, 0x97, 0x0a, 0x79, 0x94, 0x46, 0xf9, 0x67, 0x75, 0xab, 0xec, 0x07, 0xde, 0x5a, 0xec,
	0xe4, 0xc4, 0x5f, 0xd2, 0x28, 0x7f, 0xad, 0x42, 0xc9, 0x32, 0xf8, 0x0e, 0xb8, 0xb7, 0x64, 0x3c,
	0x32, 0x06, 0x71, 0x4d, 0xed, 0x74, 0x85, 0xe7, 0xbc, 0x1c, 0x21, 0x29, 0xe7, 0x67, 0x6c, 0x0f,
	0xf1, 0x25, 0xb1, 0x7f, 0xf1, 0x80, 0x9a, 0xb8, 0xf3, 0xfa, 0x5d, 0x3c, 0x41, 0x04, 0x95, 0xcd,
	0xc7, 0x36, 0x09, 0xe7, 0x3a, 0x99, 0xf7, 0x58, 0x9d, 0x03, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x73,
	0x6a, 0x05, 0x2b, 0x01, 0x00, 0x00,
}