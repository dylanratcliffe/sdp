// Code generated by protoc-gen-go. DO NOT EDIT.
// source: items.proto

package sdp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// RequestMethod represents the available request methods. The details of these
// methods are:
//
// GET: This takes a single unique query and should only return a single item.
//      If an item matching th parameter passed doesn't exist the server should
//      fail
//
// FIND: This takes no query (or ignores it) and should return all items that it
//       can find
//
// SEARCH: This takes a non-unique query which is designed to be used as a
//         search term. It should return some number of items (or zero) which
//         match the query
type RequestMethod int32

const (
	RequestMethod_GET    RequestMethod = 0
	RequestMethod_FIND   RequestMethod = 1
	RequestMethod_SEARCH RequestMethod = 2
)

var RequestMethod_name = map[int32]string{
	0: "GET",
	1: "FIND",
	2: "SEARCH",
}

var RequestMethod_value = map[string]int32{
	"GET":    0,
	"FIND":   1,
	"SEARCH": 2,
}

func (x RequestMethod) String() string {
	return proto.EnumName(RequestMethod_name, int32(x))
}

func (RequestMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f8bad7296e45e12f, []int{0}
}

// The error type. Any types in here will be gracefully handled unless the
// type os "OTHER"
type ItemRequestError_ErrorType int32

const (
	// NOTFOUND means that the item was not found. This is only returned as the
	// result of a GET request since all other requests would return an empty
	// list instead
	ItemRequestError_NOTFOUND ItemRequestError_ErrorType = 0
	// This should be used of all other failure modes
	ItemRequestError_OTHER ItemRequestError_ErrorType = 1
)

var ItemRequestError_ErrorType_name = map[int32]string{
	0: "NOTFOUND",
	1: "OTHER",
}

var ItemRequestError_ErrorType_value = map[string]int32{
	"NOTFOUND": 0,
	"OTHER":    1,
}

func (x ItemRequestError_ErrorType) String() string {
	return proto.EnumName(ItemRequestError_ErrorType_name, int32(x))
}

func (ItemRequestError_ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f8bad7296e45e12f, []int{1, 0}
}

// ItemRequest represents a request for an item.
//
// Type: (Optional) The type of item that you are looking for, is this is not
// provided then the request will be for all types that the receiver knows
// about.
//
// Method: (Required) The request method to use
//
// Query: (Optional) The query to pass
type ItemRequest struct {
	// The type of item to search for
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Which method to use when looking for it
	Method RequestMethod `protobuf:"varint,2,opt,name=method,proto3,enum=RequestMethod" json:"method,omitempty"`
	// What query should be passed to that method
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// Whether or not to resolve linked items
	// TODO: This will require a fair amount of refactoring in redacted_cli
	Link bool `protobuf:"varint,4,opt,name=link,proto3" json:"link,omitempty"`
	// Subject that items resulting from the request should be sent to
	ItemSubject string `protobuf:"bytes,16,opt,name=itemSubject,proto3" json:"itemSubject,omitempty"`
	// Subject that both interim and final responses should be sent to
	ResponseSubject string `protobuf:"bytes,17,opt,name=responseSubject,proto3" json:"responseSubject,omitempty"`
	// Subject that errors should be sent to
	ErrorSubject         string   `protobuf:"bytes,18,opt,name=errorSubject,proto3" json:"errorSubject,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemRequest) Reset()         { *m = ItemRequest{} }
func (m *ItemRequest) String() string { return proto.CompactTextString(m) }
func (*ItemRequest) ProtoMessage()    {}
func (*ItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bad7296e45e12f, []int{0}
}

func (m *ItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemRequest.Unmarshal(m, b)
}
func (m *ItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemRequest.Marshal(b, m, deterministic)
}
func (m *ItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRequest.Merge(m, src)
}
func (m *ItemRequest) XXX_Size() int {
	return xxx_messageInfo_ItemRequest.Size(m)
}
func (m *ItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRequest proto.InternalMessageInfo

func (m *ItemRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ItemRequest) GetMethod() RequestMethod {
	if m != nil {
		return m.Method
	}
	return RequestMethod_GET
}

func (m *ItemRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *ItemRequest) GetLink() bool {
	if m != nil {
		return m.Link
	}
	return false
}

func (m *ItemRequest) GetItemSubject() string {
	if m != nil {
		return m.ItemSubject
	}
	return ""
}

func (m *ItemRequest) GetResponseSubject() string {
	if m != nil {
		return m.ResponseSubject
	}
	return ""
}

func (m *ItemRequest) GetErrorSubject() string {
	if m != nil {
		return m.ErrorSubject
	}
	return ""
}

// ItemRequestError is sent back when an item request fails
type ItemRequestError struct {
	ErrorType ItemRequestError_ErrorType `protobuf:"varint,2,opt,name=errorType,proto3,enum=ItemRequestError_ErrorType" json:"errorType,omitempty"`
	// The string contents of the error
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// The context from whihc the error was raised
	Context              string   `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ItemRequestError) Reset()         { *m = ItemRequestError{} }
func (m *ItemRequestError) String() string { return proto.CompactTextString(m) }
func (*ItemRequestError) ProtoMessage()    {}
func (*ItemRequestError) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bad7296e45e12f, []int{1}
}

func (m *ItemRequestError) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemRequestError.Unmarshal(m, b)
}
func (m *ItemRequestError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemRequestError.Marshal(b, m, deterministic)
}
func (m *ItemRequestError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemRequestError.Merge(m, src)
}
func (m *ItemRequestError) XXX_Size() int {
	return xxx_messageInfo_ItemRequestError.Size(m)
}
func (m *ItemRequestError) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemRequestError.DiscardUnknown(m)
}

var xxx_messageInfo_ItemRequestError proto.InternalMessageInfo

func (m *ItemRequestError) GetErrorType() ItemRequestError_ErrorType {
	if m != nil {
		return m.ErrorType
	}
	return ItemRequestError_NOTFOUND
}

func (m *ItemRequestError) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *ItemRequestError) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

// ItemAttributes represents the known attributes for an item. These are likely
// to be common to a given type, but even this is not guaranteed. All items must
// have at least one attribute however as it needs something to uniquely
// identify it
type ItemAttributes struct {
	AttrStruct           *_struct.Struct `protobuf:"bytes,1,opt,name=attrStruct,proto3" json:"attrStruct,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ItemAttributes) Reset()         { *m = ItemAttributes{} }
func (m *ItemAttributes) String() string { return proto.CompactTextString(m) }
func (*ItemAttributes) ProtoMessage()    {}
func (*ItemAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bad7296e45e12f, []int{2}
}

func (m *ItemAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ItemAttributes.Unmarshal(m, b)
}
func (m *ItemAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ItemAttributes.Marshal(b, m, deterministic)
}
func (m *ItemAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ItemAttributes.Merge(m, src)
}
func (m *ItemAttributes) XXX_Size() int {
	return xxx_messageInfo_ItemAttributes.Size(m)
}
func (m *ItemAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ItemAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ItemAttributes proto.InternalMessageInfo

func (m *ItemAttributes) GetAttrStruct() *_struct.Struct {
	if m != nil {
		return m.AttrStruct
	}
	return nil
}

// This is the same as Item within the package with a couple of exceptions, no
// real reason why this whole thing couldn't be modelled in protobuf though if
// required. Just need to decide what if anything should remain private
type Item struct {
	Type            string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	UniqueAttribute string          `protobuf:"bytes,2,opt,name=uniqueAttribute,proto3" json:"uniqueAttribute,omitempty"`
	Attributes      *ItemAttributes `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
	Metadata        *Metadata       `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Not all items will have relatedItems we are are using a two byte
	// integer to save one byte integers for more common things
	LinkedItemRequests []*ItemRequest `protobuf:"bytes,16,rep,name=linkedItemRequests,proto3" json:"linkedItemRequests,omitempty"`
	// Linked items
	LinkedItems          []*Item  `protobuf:"bytes,17,rep,name=linkedItems,proto3" json:"linkedItems,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bad7296e45e12f, []int{3}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Item) GetUniqueAttribute() string {
	if m != nil {
		return m.UniqueAttribute
	}
	return ""
}

func (m *Item) GetAttributes() *ItemAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Item) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Item) GetLinkedItemRequests() []*ItemRequest {
	if m != nil {
		return m.LinkedItemRequests
	}
	return nil
}

func (m *Item) GetLinkedItems() []*Item {
	if m != nil {
		return m.LinkedItems
	}
	return nil
}

// This is a list of items, like a Find() would return
type Items struct {
	Items                []*Item  `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Items) Reset()         { *m = Items{} }
func (m *Items) String() string { return proto.CompactTextString(m) }
func (*Items) ProtoMessage()    {}
func (*Items) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bad7296e45e12f, []int{4}
}

func (m *Items) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Items.Unmarshal(m, b)
}
func (m *Items) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Items.Marshal(b, m, deterministic)
}
func (m *Items) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Items.Merge(m, src)
}
func (m *Items) XXX_Size() int {
	return xxx_messageInfo_Items.Size(m)
}
func (m *Items) XXX_DiscardUnknown() {
	xxx_messageInfo_Items.DiscardUnknown(m)
}

var xxx_messageInfo_Items proto.InternalMessageInfo

func (m *Items) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

// Metadata about the item. Where it came from, how long it took, etc.
type Metadata struct {
	// The context within which the item is unique. Item uniqueness is determined
	// by the combination of type and uniqueAttribute value. However it is
	// possible for the same item to exist in many contexts. There is not formal
	// definition for what a context should be other than the fact that it should
	// be somewhat descriptive and should ensure item uniqueness
	Context string `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	// This is the name of the backend that was used to find the item.
	BackendName   string        `protobuf:"bytes,2,opt,name=backendName,proto3" json:"backendName,omitempty"`
	RequestMethod RequestMethod `protobuf:"varint,3,opt,name=requestMethod,proto3,enum=RequestMethod" json:"requestMethod,omitempty"`
	// The time that the item was found
	Timestamp *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// How long the backend took to execute in total when processing the
	// ItemRequest
	BackendDuration *duration.Duration `protobuf:"bytes,5,opt,name=backendDuration,proto3" json:"backendDuration,omitempty"`
	// How long the backend took to execute per item when processing the
	// ItemRequest
	BackendDurationPerItem *duration.Duration `protobuf:"bytes,6,opt,name=backendDurationPerItem,proto3" json:"backendDurationPerItem,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8bad7296e45e12f, []int{5}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetContext() string {
	if m != nil {
		return m.Context
	}
	return ""
}

func (m *Metadata) GetBackendName() string {
	if m != nil {
		return m.BackendName
	}
	return ""
}

func (m *Metadata) GetRequestMethod() RequestMethod {
	if m != nil {
		return m.RequestMethod
	}
	return RequestMethod_GET
}

func (m *Metadata) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Metadata) GetBackendDuration() *duration.Duration {
	if m != nil {
		return m.BackendDuration
	}
	return nil
}

func (m *Metadata) GetBackendDurationPerItem() *duration.Duration {
	if m != nil {
		return m.BackendDurationPerItem
	}
	return nil
}

func init() {
	proto.RegisterEnum("RequestMethod", RequestMethod_name, RequestMethod_value)
	proto.RegisterEnum("ItemRequestError_ErrorType", ItemRequestError_ErrorType_name, ItemRequestError_ErrorType_value)
	proto.RegisterType((*ItemRequest)(nil), "ItemRequest")
	proto.RegisterType((*ItemRequestError)(nil), "ItemRequestError")
	proto.RegisterType((*ItemAttributes)(nil), "ItemAttributes")
	proto.RegisterType((*Item)(nil), "Item")
	proto.RegisterType((*Items)(nil), "Items")
	proto.RegisterType((*Metadata)(nil), "Metadata")
}

func init() {
	proto.RegisterFile("items.proto", fileDescriptor_f8bad7296e45e12f)
}

var fileDescriptor_f8bad7296e45e12f = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xeb, 0x24, 0x4e, 0xe3, 0x71, 0xdb, 0xb8, 0x23, 0x04, 0xa6, 0x45, 0x10, 0x59, 0x05,
	0x2c, 0x0e, 0x8e, 0x14, 0x90, 0x00, 0x89, 0x4b, 0x69, 0x52, 0x9a, 0x43, 0x13, 0xd8, 0x84, 0x0b,
	0x37, 0x27, 0x5e, 0x8a, 0x69, 0x6d, 0x27, 0xeb, 0xb5, 0x44, 0xcf, 0x3c, 0x0a, 0xef, 0xc3, 0x4b,
	0xf0, 0x22, 0xc8, 0x63, 0x3b, 0x71, 0x9c, 0x22, 0x6e, 0xde, 0x7f, 0xbe, 0xd9, 0xd9, 0x7f, 0xc6,
	0x03, 0xba, 0x2f, 0x79, 0x10, 0x3b, 0x0b, 0x11, 0xc9, 0xe8, 0xe8, 0xd1, 0x55, 0x14, 0x5d, 0xdd,
	0xf0, 0x2e, 0x9d, 0x66, 0xc9, 0xd7, 0x6e, 0x2c, 0x45, 0x32, 0x97, 0x79, 0xf4, 0x49, 0x35, 0x2a,
	0xfd, 0x80, 0xc7, 0xd2, 0x0d, 0x16, 0x39, 0xf0, 0xb8, 0x0a, 0x78, 0x89, 0x70, 0xa5, 0x1f, 0x85,
	0x59, 0xdc, 0xfa, 0xa3, 0x80, 0x3e, 0x94, 0x3c, 0x60, 0x7c, 0x99, 0xf0, 0x58, 0x22, 0x42, 0x43,
	0xde, 0x2e, 0xb8, 0xa9, 0x74, 0x14, 0x5b, 0x63, 0xf4, 0x8d, 0xcf, 0xa0, 0x19, 0x70, 0xf9, 0x2d,
	0xf2, 0xcc, 0x5a, 0x47, 0xb1, 0x0f, 0x7a, 0x07, 0x4e, 0x4e, 0x5f, 0x92, 0xca, 0xf2, 0x28, 0xde,
	0x03, 0x75, 0x99, 0x70, 0x71, 0x6b, 0xd6, 0x29, 0x39, 0x3b, 0xa4, 0x37, 0xde, 0xf8, 0xe1, 0xb5,
	0xd9, 0xe8, 0x28, 0x76, 0x8b, 0xd1, 0x37, 0x76, 0x32, 0x8f, 0x93, 0x64, 0xf6, 0x9d, 0xcf, 0xa5,
	0x69, 0x10, 0x5f, 0x96, 0xd0, 0x86, 0xb6, 0xe0, 0xf1, 0x22, 0x0a, 0x63, 0x5e, 0x50, 0x87, 0x44,
	0x55, 0x65, 0xb4, 0x60, 0x8f, 0x0b, 0x11, 0x89, 0x02, 0x43, 0xc2, 0x36, 0x34, 0xeb, 0x97, 0x02,
	0x46, 0xc9, 0xe5, 0x20, 0x8d, 0xe1, 0x5b, 0xd0, 0x08, 0x9a, 0xa6, 0x7e, 0x33, 0x67, 0xc7, 0x4e,
	0x95, 0x72, 0x06, 0x05, 0xc2, 0xd6, 0x74, 0xea, 0x94, 0x0e, 0x85, 0x53, 0x3a, 0xa0, 0x09, 0xbb,
	0xf3, 0x28, 0x94, 0xfc, 0x87, 0x24, 0xb3, 0x1a, 0x2b, 0x8e, 0xd6, 0x09, 0x68, 0xab, 0x7b, 0x70,
	0x0f, 0x5a, 0xa3, 0xf1, 0xf4, 0x7c, 0xfc, 0x79, 0xd4, 0x37, 0x76, 0x50, 0x03, 0x75, 0x3c, 0xbd,
	0x18, 0x30, 0x43, 0xb1, 0x86, 0x70, 0x90, 0x96, 0x3f, 0x95, 0x52, 0xf8, 0xb3, 0x44, 0xf2, 0x18,
	0x5f, 0x03, 0xb8, 0x52, 0x8a, 0x09, 0x8d, 0x9c, 0x66, 0xa2, 0xf7, 0x1e, 0x38, 0xd9, 0x48, 0x9d,
	0x62, 0xa4, 0x4e, 0x16, 0x66, 0x25, 0xd4, 0xfa, 0x59, 0x83, 0x46, 0x7a, 0xd7, 0x9d, 0xf3, 0xb4,
	0xa1, 0x9d, 0x84, 0xfe, 0x32, 0xe1, 0xab, 0x4a, 0x64, 0x5f, 0x63, 0x55, 0x19, 0xbb, 0x59, 0xfd,
	0xec, 0x35, 0x64, 0x56, 0xef, 0xb5, 0x9d, 0xcd, 0x47, 0xb2, 0x12, 0x82, 0x4f, 0xa1, 0x15, 0x70,
	0xe9, 0x7a, 0xae, 0x74, 0xa9, 0x07, 0x7a, 0x4f, 0x73, 0x2e, 0x73, 0x81, 0xad, 0x42, 0xf8, 0x0e,
	0x30, 0xfd, 0x0f, 0xb8, 0x57, 0x6a, 0x77, 0x6c, 0x1a, 0x9d, 0xba, 0xad, 0xf7, 0xf6, 0xca, 0x33,
	0x60, 0x77, 0x70, 0xf8, 0x1c, 0xf4, 0xb5, 0x1a, 0x9b, 0x87, 0x94, 0xa6, 0x66, 0x69, 0xe5, 0x88,
	0x75, 0x02, 0x2a, 0x7d, 0xe0, 0x31, 0xa8, 0xb4, 0x53, 0xa6, 0x52, 0x66, 0x33, 0xcd, 0xfa, 0x5d,
	0x83, 0x56, 0xf1, 0xc6, 0xf2, 0x0c, 0x95, 0x8d, 0x19, 0xa6, 0xff, 0xec, 0xcc, 0x9d, 0x5f, 0xf3,
	0xd0, 0x1b, 0xb9, 0x41, 0xd1, 0xb1, 0xb2, 0x84, 0xaf, 0x60, 0x5f, 0x94, 0x17, 0x83, 0x1a, 0xb6,
	0xbd, 0x2e, 0x9b, 0x10, 0xbe, 0x01, 0x6d, 0xb5, 0xb4, 0x79, 0xcf, 0x8e, 0xb6, 0x46, 0x3c, 0x2d,
	0x08, 0xb6, 0x86, 0xf1, 0x0c, 0xda, 0x79, 0xf9, 0x7e, 0xbe, 0xd4, 0xa6, 0x4a, 0xf9, 0x0f, 0xb7,
	0xf2, 0x0b, 0x80, 0x55, 0x33, 0xf0, 0x13, 0xdc, 0xaf, 0x48, 0x1f, 0xb9, 0x48, 0xdb, 0x63, 0x36,
	0xff, 0x77, 0xd7, 0x3f, 0x12, 0x5f, 0x38, 0xb0, 0xbf, 0xe1, 0x18, 0x77, 0xa1, 0xfe, 0x61, 0x30,
	0x35, 0x76, 0xb0, 0x05, 0x8d, 0xf3, 0xe1, 0xa8, 0x6f, 0x28, 0x08, 0xd0, 0x9c, 0x0c, 0x4e, 0xd9,
	0xd9, 0x85, 0x51, 0x7b, 0xaf, 0x7e, 0xa9, 0xc7, 0xde, 0x62, 0xd6, 0xa4, 0x0a, 0x2f, 0xff, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x27, 0xaa, 0xde, 0x59, 0xff, 0x04, 0x00, 0x00,
}
