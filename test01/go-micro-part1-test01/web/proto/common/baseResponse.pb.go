// Code generated by protoc-gen-go. DO NOT EDIT.
// source: baseResponse.proto

package common

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

type BaseResponse struct {
	IsSuccess            bool     `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a4673f0ea6d60b3, []int{0}
}

func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (m *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(m, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func (m *BaseResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseResponse)(nil), "common.baseResponse")
}

func init() {
	proto.RegisterFile("baseResponse.proto", fileDescriptor_6a4673f0ea6d60b3)
}

var fileDescriptor_6a4673f0ea6d60b3 = []byte{
	// 110 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x4a, 0x2c, 0x4e,
	0x0d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4b, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x53, 0x0a, 0xe0, 0xe2, 0x41, 0x96, 0x15, 0x92, 0xe1, 0xe2,
	0xcc, 0x2c, 0x0e, 0x2e, 0x4d, 0x4e, 0x4e, 0x2d, 0x2e, 0x96, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08,
	0x42, 0x08, 0x08, 0x29, 0x71, 0xf1, 0xa4, 0x16, 0x15, 0xe5, 0x17, 0xf9, 0xa6, 0x16, 0x17, 0x27,
	0xa6, 0xa7, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xa1, 0x88, 0x25, 0xb1, 0x81, 0x2d, 0x30,
	0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xad, 0x05, 0xcf, 0xaf, 0x76, 0x00, 0x00, 0x00,
}