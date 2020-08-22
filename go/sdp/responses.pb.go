// Code generated by protoc-gen-go. DO NOT EDIT.
// source: responses.proto

package sdp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// ResponseState represents the state of the responder, is it still working on
// its response or is it done?
type Response_ResponseState int32

const (
	Response_WORKING  Response_ResponseState = 0
	Response_COMPLETE Response_ResponseState = 1
)

var Response_ResponseState_name = map[int32]string{
	0: "WORKING",
	1: "COMPLETE",
}

var Response_ResponseState_value = map[string]int32{
	"WORKING":  0,
	"COMPLETE": 1,
}

func (x Response_ResponseState) String() string {
	return proto.EnumName(Response_ResponseState_name, int32(x))
}

func (Response_ResponseState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2f8192a35440e287, []int{0, 0}
}

// Response is returned when a query is made
type Response struct {
	// The context that is working on a response
	Context string `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	// The state of the responder, either WORKING or COMPLETE
	State Response_ResponseState `protobuf:"varint,2,opt,name=state,proto3,enum=Response_ResponseState" json:"state,omitempty"`
	// The timespan within which to expect the next update. (e.g. 10s) If no
	// further interim responses are recieved within this time the connection
	// can be considered stale and the requester may give up
	NextUpdateIn         *duration.Duration `protobuf:"bytes,3,opt,name=nextUpdateIn,proto3" json:"nextUpdateIn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f8192a35440e287, []int{0}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *Response) GetState() Response_ResponseState {
	if m != nil {
		return m.State
	}
	return Response_WORKING
}

func (m *Response) GetNextUpdateIn() *duration.Duration {
	if m != nil {
		return m.NextUpdateIn
	}
	return nil
}

func init() {
	proto.RegisterEnum("Response_ResponseState", Response_ResponseState_name, Response_ResponseState_value)
	proto.RegisterType((*Response)(nil), "Response")
}

func init() {
	proto.RegisterFile("responses.proto", fileDescriptor_2f8192a35440e287)
}

var fileDescriptor_2f8192a35440e287 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4b, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x4a, 0x8b, 0x12, 0x4b, 0x32,
	0xf3, 0xf3, 0x20, 0xf2, 0x4a, 0xfb, 0x18, 0xb9, 0x38, 0x82, 0xa0, 0x7a, 0x84, 0x24, 0xb8, 0xd8,
	0x93, 0xf3, 0xf3, 0x4a, 0x52, 0x2b, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c,
	0x21, 0x5d, 0x2e, 0xd6, 0xe2, 0x92, 0xc4, 0x92, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x3e, 0x23,
	0x71, 0x3d, 0x98, 0x1e, 0x38, 0x23, 0x18, 0x24, 0x1d, 0x04, 0x51, 0x25, 0x64, 0xcb, 0xc5, 0x93,
	0x97, 0x5a, 0x51, 0x12, 0x5a, 0x90, 0x92, 0x58, 0x92, 0xea, 0x99, 0x27, 0xc1, 0xac, 0xc0, 0xa8,
	0xc1, 0x6d, 0x24, 0xa9, 0x07, 0x71, 0x8c, 0x1e, 0xcc, 0x31, 0x7a, 0x2e, 0x50, 0xc7, 0x04, 0xa1,
	0x28, 0x57, 0xd2, 0xe2, 0xe2, 0x45, 0x31, 0x56, 0x88, 0x9b, 0x8b, 0x3d, 0xdc, 0x3f, 0xc8, 0xdb,
	0xd3, 0xcf, 0x5d, 0x80, 0x41, 0x88, 0x87, 0x8b, 0xc3, 0xd9, 0xdf, 0x37, 0xc0, 0xc7, 0x35, 0xc4,
	0x55, 0x80, 0xd1, 0x89, 0x35, 0x8a, 0xb9, 0x38, 0xa5, 0x20, 0x89, 0x0d, 0x6c, 0xa6, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x7f, 0xeb, 0xa8, 0x70, 0x01, 0x01, 0x00, 0x00,
}
