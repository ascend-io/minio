// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/gateway/grpc/protos/status.proto

package miniogrpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// TODO(cmaloney): This needs to _easily_ map into all the minio specific errors.
// + easily layer custom ones on top.
type Status struct {
	Code                 int32      `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string     `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Details              []*any.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_92133c0e4af3a570, []int{0}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Status) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Status) GetDetails() []*any.Any {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "miniogrpc.Status")
}

func init() {
	proto.RegisterFile("cmd/gateway/grpc/protos/status.proto", fileDescriptor_92133c0e4af3a570)
}

var fileDescriptor_92133c0e4af3a570 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8d, 0xc1, 0x0e, 0x82, 0x30,
	0x0c, 0x40, 0x83, 0x28, 0x84, 0x79, 0x5b, 0x3c, 0x4c, 0x4f, 0xc4, 0x78, 0xe0, 0xb4, 0x25, 0xfa,
	0x05, 0xfe, 0x02, 0x7e, 0x41, 0x81, 0xb2, 0x90, 0x00, 0x25, 0x74, 0xc4, 0xf0, 0xf7, 0xc6, 0x2d,
	0xbb, 0xf5, 0xb5, 0xaf, 0x79, 0xe2, 0xd1, 0x4e, 0x9d, 0xb1, 0xe0, 0xf0, 0x0b, 0xbb, 0xb1, 0xeb,
	0xd2, 0x9a, 0x65, 0x25, 0x47, 0x6c, 0xd8, 0x81, 0xdb, 0x58, 0x7b, 0x92, 0xc5, 0x34, 0xcc, 0x03,
	0xfd, 0xcf, 0xb7, 0xab, 0x25, 0xb2, 0x23, 0x06, 0xad, 0xd9, 0x7a, 0x03, 0xf3, 0x1e, 0xac, 0x7b,
	0x2f, 0xb2, 0x8f, 0xff, 0x92, 0x52, 0x1c, 0x5b, 0xea, 0x50, 0x25, 0x65, 0x52, 0x9d, 0x6a, 0x3f,
	0x4b, 0x25, 0xf2, 0x09, 0x99, 0xc1, 0xa2, 0x3a, 0x94, 0x49, 0x55, 0xd4, 0x11, 0xa5, 0x16, 0x79,
	0x87, 0x0e, 0x86, 0x91, 0x55, 0x5a, 0xa6, 0xd5, 0xf9, 0x79, 0xd1, 0x21, 0xa2, 0x63, 0x44, 0xbf,
	0xe7, 0xbd, 0x8e, 0x52, 0x93, 0xf9, 0xf5, 0xeb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x65, 0xca,
	0xdf, 0xbc, 0x00, 0x00, 0x00,
}
