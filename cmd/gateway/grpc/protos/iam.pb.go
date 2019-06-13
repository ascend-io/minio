// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/gateway/grpc/protos/iam.proto

package miniogrpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type GetUserRequest struct {
	AccessKey            string   `protobuf:"bytes,1,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6936646026dc912, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

type Credentials struct {
	AccessKey            string               `protobuf:"bytes,1,opt,name=accessKey,proto3" json:"accessKey,omitempty"`
	SecretKey            string               `protobuf:"bytes,2,opt,name=secretKey,proto3" json:"secretKey,omitempty"`
	Expiration           *timestamp.Timestamp `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	SessionToken         string               `protobuf:"bytes,4,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
	Status               string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Credentials) Reset()         { *m = Credentials{} }
func (m *Credentials) String() string { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()    {}
func (*Credentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6936646026dc912, []int{1}
}

func (m *Credentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Credentials.Unmarshal(m, b)
}
func (m *Credentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Credentials.Marshal(b, m, deterministic)
}
func (m *Credentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credentials.Merge(m, src)
}
func (m *Credentials) XXX_Size() int {
	return xxx_messageInfo_Credentials.Size(m)
}
func (m *Credentials) XXX_DiscardUnknown() {
	xxx_messageInfo_Credentials.DiscardUnknown(m)
}

var xxx_messageInfo_Credentials proto.InternalMessageInfo

func (m *Credentials) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *Credentials) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

func (m *Credentials) GetExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func (m *Credentials) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

func (m *Credentials) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type GetUserResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6936646026dc912, []int{2}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type IsAllowedRequest struct {
	AccountName string `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	Action      string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	BucketName  string `protobuf:"bytes,3,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	// TODO(cmaloney): ConditionValues conditions = 4;
	IsOwner              bool     `protobuf:"varint,5,opt,name=isOwner,proto3" json:"isOwner,omitempty"`
	ObjectName           string   `protobuf:"bytes,6,opt,name=objectName,proto3" json:"objectName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAllowedRequest) Reset()         { *m = IsAllowedRequest{} }
func (m *IsAllowedRequest) String() string { return proto.CompactTextString(m) }
func (*IsAllowedRequest) ProtoMessage()    {}
func (*IsAllowedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6936646026dc912, []int{3}
}

func (m *IsAllowedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAllowedRequest.Unmarshal(m, b)
}
func (m *IsAllowedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAllowedRequest.Marshal(b, m, deterministic)
}
func (m *IsAllowedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAllowedRequest.Merge(m, src)
}
func (m *IsAllowedRequest) XXX_Size() int {
	return xxx_messageInfo_IsAllowedRequest.Size(m)
}
func (m *IsAllowedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAllowedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAllowedRequest proto.InternalMessageInfo

func (m *IsAllowedRequest) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *IsAllowedRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *IsAllowedRequest) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *IsAllowedRequest) GetIsOwner() bool {
	if m != nil {
		return m.IsOwner
	}
	return false
}

func (m *IsAllowedRequest) GetObjectName() string {
	if m != nil {
		return m.ObjectName
	}
	return ""
}

type IsAllowedResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	IsAllowed            bool     `protobuf:"varint,2,opt,name=isAllowed,proto3" json:"isAllowed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAllowedResponse) Reset()         { *m = IsAllowedResponse{} }
func (m *IsAllowedResponse) String() string { return proto.CompactTextString(m) }
func (*IsAllowedResponse) ProtoMessage()    {}
func (*IsAllowedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6936646026dc912, []int{4}
}

func (m *IsAllowedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAllowedResponse.Unmarshal(m, b)
}
func (m *IsAllowedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAllowedResponse.Marshal(b, m, deterministic)
}
func (m *IsAllowedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAllowedResponse.Merge(m, src)
}
func (m *IsAllowedResponse) XXX_Size() int {
	return xxx_messageInfo_IsAllowedResponse.Size(m)
}
func (m *IsAllowedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAllowedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsAllowedResponse proto.InternalMessageInfo

func (m *IsAllowedResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *IsAllowedResponse) GetIsAllowed() bool {
	if m != nil {
		return m.IsAllowed
	}
	return false
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "miniogrpc.GetUserRequest")
	proto.RegisterType((*Credentials)(nil), "miniogrpc.Credentials")
	proto.RegisterType((*GetUserResponse)(nil), "miniogrpc.GetUserResponse")
	proto.RegisterType((*IsAllowedRequest)(nil), "miniogrpc.IsAllowedRequest")
	proto.RegisterType((*IsAllowedResponse)(nil), "miniogrpc.IsAllowedResponse")
}

func init() { proto.RegisterFile("cmd/gateway/grpc/protos/iam.proto", fileDescriptor_f6936646026dc912) }

var fileDescriptor_f6936646026dc912 = []byte{
	// 414 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x26, 0x2c, 0x74, 0x9b, 0x29, 0x02, 0xd6, 0x07, 0x14, 0xc2, 0x0a, 0x4a, 0xc4, 0x61, 0xb9,
	0x24, 0xd2, 0x72, 0x43, 0x5c, 0x16, 0x0e, 0x50, 0xa1, 0xa5, 0x52, 0x28, 0x37, 0x2e, 0x8e, 0x3b,
	0x44, 0xa6, 0x89, 0x1d, 0x32, 0x8e, 0x4a, 0xdf, 0x84, 0x27, 0xe0, 0x3d, 0x78, 0x33, 0x14, 0x3b,
	0x49, 0x53, 0x54, 0x90, 0xb8, 0x65, 0xbe, 0x9f, 0xd1, 0x37, 0x5f, 0x0c, 0x4f, 0x45, 0xb9, 0x4e,
	0x72, 0x6e, 0x70, 0xcb, 0x77, 0x49, 0x5e, 0x57, 0x22, 0xa9, 0x6a, 0x6d, 0x34, 0x25, 0x92, 0x97,
	0xb1, 0xfd, 0x64, 0x7e, 0x29, 0x95, 0xd4, 0x2d, 0x17, 0x3e, 0xfb, 0x9b, 0x9a, 0x0c, 0x37, 0x0d,
	0x39, 0x43, 0xf8, 0x24, 0xd7, 0x3a, 0x2f, 0xd0, 0x71, 0x59, 0xf3, 0x25, 0x31, 0xb2, 0x44, 0x32,
	0xbc, 0xac, 0x9c, 0x20, 0x8a, 0xe1, 0xee, 0x5b, 0x34, 0x9f, 0x08, 0xeb, 0x14, 0xbf, 0x35, 0x48,
	0x86, 0x9d, 0x83, 0xcf, 0x85, 0x40, 0xa2, 0xf7, 0xb8, 0x0b, 0xbc, 0xb9, 0x77, 0xe1, 0xa7, 0x7b,
	0x20, 0xfa, 0xe5, 0xc1, 0xec, 0x4d, 0x8d, 0x6b, 0x54, 0x46, 0xf2, 0x82, 0xfe, 0xad, 0x6e, 0x59,
	0x42, 0x51, 0xa3, 0x69, 0xd9, 0x9b, 0x8e, 0x1d, 0x00, 0xf6, 0x12, 0x00, 0xbf, 0x57, 0xb2, 0xe6,
	0x46, 0x6a, 0x15, 0x9c, 0xcc, 0xbd, 0x8b, 0xd9, 0x65, 0x18, 0xbb, 0xc4, 0x71, 0x9f, 0x38, 0x5e,
	0xf5, 0x89, 0xd3, 0x91, 0x9a, 0x45, 0x70, 0x87, 0x90, 0x48, 0x6a, 0xb5, 0xd2, 0x1b, 0x54, 0xc1,
	0x2d, 0xbb, 0xfc, 0x00, 0x63, 0x0f, 0x60, 0xe2, 0xca, 0x08, 0x6e, 0x5b, 0xb6, 0x9b, 0xa2, 0x57,
	0x70, 0x6f, 0xb8, 0x99, 0x2a, 0xad, 0x08, 0xd9, 0xf3, 0x41, 0xea, 0xd9, 0x18, 0x67, 0xf1, 0xd0,
	0x74, 0xfc, 0xd1, 0x12, 0x83, 0xfb, 0xa7, 0x07, 0xf7, 0x17, 0x74, 0x55, 0x14, 0x7a, 0x8b, 0xeb,
	0xbe, 0xb4, 0x39, 0xcc, 0xb8, 0x10, 0xba, 0x51, 0xe6, 0x03, 0x2f, 0xb1, 0x2b, 0x62, 0x0c, 0xb5,
	0x61, 0xb8, 0xb0, 0x87, 0xba, 0x1e, 0xba, 0x89, 0x3d, 0x06, 0xc8, 0x1a, 0xb1, 0x41, 0x67, 0x3c,
	0xb1, 0xdc, 0x08, 0x61, 0x01, 0x9c, 0x4a, 0x5a, 0x6e, 0x15, 0xd6, 0xf6, 0x8a, 0x69, 0xda, 0x8f,
	0xad, 0x53, 0x67, 0x5f, 0x51, 0x38, 0xe7, 0xc4, 0x39, 0xf7, 0x48, 0xf4, 0x19, 0xce, 0x46, 0x39,
	0xff, 0xfb, 0xd0, 0xf6, 0xe7, 0xc9, 0xde, 0x6f, 0x43, 0x4f, 0xd3, 0x3d, 0x70, 0xf9, 0xc3, 0x83,
	0xe9, 0xb5, 0x54, 0x8b, 0xe5, 0xe2, 0xea, 0x9a, 0xbd, 0x86, 0xd3, 0xae, 0x51, 0xf6, 0x70, 0xb4,
	0xf0, 0xf0, 0x65, 0x85, 0xe1, 0x31, 0xca, 0xe5, 0x8a, 0x6e, 0xb0, 0x77, 0xe0, 0x0f, 0x71, 0xd9,
	0xa3, 0x91, 0xf4, 0xcf, 0xb2, 0xc3, 0xf3, 0xe3, 0x64, 0xbf, 0x29, 0x9b, 0xd8, 0xb7, 0xf3, 0xe2,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xf9, 0xa4, 0xc2, 0x51, 0x03, 0x00, 0x00,
}
