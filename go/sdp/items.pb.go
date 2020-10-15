// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: items.proto

package sdp

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

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

// Enum value maps for RequestMethod.
var (
	RequestMethod_name = map[int32]string{
		0: "GET",
		1: "FIND",
		2: "SEARCH",
	}
	RequestMethod_value = map[string]int32{
		"GET":    0,
		"FIND":   1,
		"SEARCH": 2,
	}
)

func (x RequestMethod) Enum() *RequestMethod {
	p := new(RequestMethod)
	*p = x
	return p
}

func (x RequestMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_items_proto_enumTypes[0].Descriptor()
}

func (RequestMethod) Type() protoreflect.EnumType {
	return &file_items_proto_enumTypes[0]
}

func (x RequestMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestMethod.Descriptor instead.
func (RequestMethod) EnumDescriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{0}
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

// Enum value maps for ItemRequestError_ErrorType.
var (
	ItemRequestError_ErrorType_name = map[int32]string{
		0: "NOTFOUND",
		1: "OTHER",
	}
	ItemRequestError_ErrorType_value = map[string]int32{
		"NOTFOUND": 0,
		"OTHER":    1,
	}
)

func (x ItemRequestError_ErrorType) Enum() *ItemRequestError_ErrorType {
	p := new(ItemRequestError_ErrorType)
	*p = x
	return p
}

func (x ItemRequestError_ErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ItemRequestError_ErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_items_proto_enumTypes[1].Descriptor()
}

func (ItemRequestError_ErrorType) Type() protoreflect.EnumType {
	return &file_items_proto_enumTypes[1]
}

func (x ItemRequestError_ErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ItemRequestError_ErrorType.Descriptor instead.
func (ItemRequestError_ErrorType) EnumDescriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{1, 0}
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of item to search for. If blank we assume all types
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Which method to use when looking for it
	Method RequestMethod `protobuf:"varint,2,opt,name=method,proto3,enum=RequestMethod" json:"method,omitempty"`
	// What query should be passed to that method
	Query string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	// Whether or not to resolve linked items
	// TODO: This will require a fair amount of refactoring in redacted_cli
	Link bool `protobuf:"varint,4,opt,name=link,proto3" json:"link,omitempty"`
	// The context for which we are requesting. To query all contexts use the the
	// wildcard '*'
	Context string `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
	// Subject that items resulting from the request should be sent to
	ItemSubject string `protobuf:"bytes,16,opt,name=itemSubject,proto3" json:"itemSubject,omitempty"`
	// Subject that both interim and final responses should be sent to
	ResponseSubject string `protobuf:"bytes,17,opt,name=responseSubject,proto3" json:"responseSubject,omitempty"`
	// Subject that errors should be sent to
	ErrorSubject string `protobuf:"bytes,18,opt,name=errorSubject,proto3" json:"errorSubject,omitempty"`
}

func (x *ItemRequest) Reset() {
	*x = ItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemRequest) ProtoMessage() {}

func (x *ItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemRequest.ProtoReflect.Descriptor instead.
func (*ItemRequest) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{0}
}

func (x *ItemRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ItemRequest) GetMethod() RequestMethod {
	if x != nil {
		return x.Method
	}
	return RequestMethod_GET
}

func (x *ItemRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ItemRequest) GetLink() bool {
	if x != nil {
		return x.Link
	}
	return false
}

func (x *ItemRequest) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *ItemRequest) GetItemSubject() string {
	if x != nil {
		return x.ItemSubject
	}
	return ""
}

func (x *ItemRequest) GetResponseSubject() string {
	if x != nil {
		return x.ResponseSubject
	}
	return ""
}

func (x *ItemRequest) GetErrorSubject() string {
	if x != nil {
		return x.ErrorSubject
	}
	return ""
}

// ItemRequestError is sent back when an item request fails
type ItemRequestError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorType ItemRequestError_ErrorType `protobuf:"varint,2,opt,name=errorType,proto3,enum=ItemRequestError_ErrorType" json:"errorType,omitempty"`
	// The string contents of the error
	Error string `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	// The context from whihc the error was raised
	Context string `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
}

func (x *ItemRequestError) Reset() {
	*x = ItemRequestError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemRequestError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemRequestError) ProtoMessage() {}

func (x *ItemRequestError) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemRequestError.ProtoReflect.Descriptor instead.
func (*ItemRequestError) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{1}
}

func (x *ItemRequestError) GetErrorType() ItemRequestError_ErrorType {
	if x != nil {
		return x.ErrorType
	}
	return ItemRequestError_NOTFOUND
}

func (x *ItemRequestError) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ItemRequestError) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

// ItemAttributes represents the known attributes for an item. These are likely
// to be common to a given type, but even this is not guaranteed. All items must
// have at least one attribute however as it needs something to uniquely
// identify it
type ItemAttributes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttrStruct *_struct.Struct `protobuf:"bytes,1,opt,name=attrStruct,proto3" json:"attrStruct,omitempty"`
}

func (x *ItemAttributes) Reset() {
	*x = ItemAttributes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemAttributes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemAttributes) ProtoMessage() {}

func (x *ItemAttributes) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemAttributes.ProtoReflect.Descriptor instead.
func (*ItemAttributes) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{2}
}

func (x *ItemAttributes) GetAttrStruct() *_struct.Struct {
	if x != nil {
		return x.AttrStruct
	}
	return nil
}

// This is the same as Item within the package with a couple of exceptions, no
// real reason why this whole thing couldn't be modelled in protobuf though if
// required. Just need to decide what if anything should remain private
type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type            string          `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	UniqueAttribute string          `protobuf:"bytes,2,opt,name=uniqueAttribute,proto3" json:"uniqueAttribute,omitempty"`
	Attributes      *ItemAttributes `protobuf:"bytes,3,opt,name=attributes,proto3" json:"attributes,omitempty"`
	Metadata        *Metadata       `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The context within which the item is unique. Item uniqueness is determined
	// by the combination of type and uniqueAttribute value. However it is
	// possible for the same item to exist in many contexts. There is not formal
	// definition for what a context should be other than the fact that it should
	// be somewhat descriptive and should ensure item uniqueness
	Context string `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
	// Not all items will have relatedItems we are are using a two byte
	// integer to save one byte integers for more common things
	LinkedItemRequests []*ItemRequest `protobuf:"bytes,16,rep,name=linkedItemRequests,proto3" json:"linkedItemRequests,omitempty"`
	// Linked items
	LinkedItems []*Item `protobuf:"bytes,17,rep,name=linkedItems,proto3" json:"linkedItems,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{3}
}

func (x *Item) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Item) GetUniqueAttribute() string {
	if x != nil {
		return x.UniqueAttribute
	}
	return ""
}

func (x *Item) GetAttributes() *ItemAttributes {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Item) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Item) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *Item) GetLinkedItemRequests() []*ItemRequest {
	if x != nil {
		return x.LinkedItemRequests
	}
	return nil
}

func (x *Item) GetLinkedItems() []*Item {
	if x != nil {
		return x.LinkedItems
	}
	return nil
}

// This is a list of items, like a Find() would return
type Items struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Items) Reset() {
	*x = Items{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Items) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Items) ProtoMessage() {}

func (x *Items) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Items.ProtoReflect.Descriptor instead.
func (*Items) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{4}
}

func (x *Items) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

// Metadata about the item. Where it came from, how long it took, etc.
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_items_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_items_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_items_proto_rawDescGZIP(), []int{5}
}

func (x *Metadata) GetBackendName() string {
	if x != nil {
		return x.BackendName
	}
	return ""
}

func (x *Metadata) GetRequestMethod() RequestMethod {
	if x != nil {
		return x.RequestMethod
	}
	return RequestMethod_GET
}

func (x *Metadata) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Metadata) GetBackendDuration() *duration.Duration {
	if x != nil {
		return x.BackendDuration
	}
	return nil
}

func (x *Metadata) GetBackendDurationPerItem() *duration.Duration {
	if x != nil {
		return x.BackendDurationPerItem
	}
	return nil
}

var File_items_proto protoreflect.FileDescriptor

var file_items_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x01, 0x0a,
	0x0b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x26, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6c, 0x69,
	0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x69, 0x74, 0x65, 0x6d, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x28,
	0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0xa3, 0x01, 0x0a,
	0x10, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x39, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x24, 0x0a, 0x09,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x54,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52,
	0x10, 0x01, 0x22, 0x49, 0x0a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x22, 0x9d, 0x02,
	0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3c, 0x0a, 0x12, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x10, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x12, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x12, 0x27, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x0b, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x24, 0x0a,
	0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x22, 0xb4, 0x02, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x43, 0x0a, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x16, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x49, 0x74, 0x65,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x16, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x2a, 0x2e, 0x0a, 0x0d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x47,
	0x45, 0x54, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x10, 0x02, 0x42, 0x05, 0x5a, 0x03, 0x73, 0x64,
	0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_items_proto_rawDescOnce sync.Once
	file_items_proto_rawDescData = file_items_proto_rawDesc
)

func file_items_proto_rawDescGZIP() []byte {
	file_items_proto_rawDescOnce.Do(func() {
		file_items_proto_rawDescData = protoimpl.X.CompressGZIP(file_items_proto_rawDescData)
	})
	return file_items_proto_rawDescData
}

var file_items_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_items_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_items_proto_goTypes = []interface{}{
	(RequestMethod)(0),              // 0: RequestMethod
	(ItemRequestError_ErrorType)(0), // 1: ItemRequestError.ErrorType
	(*ItemRequest)(nil),             // 2: ItemRequest
	(*ItemRequestError)(nil),        // 3: ItemRequestError
	(*ItemAttributes)(nil),          // 4: ItemAttributes
	(*Item)(nil),                    // 5: Item
	(*Items)(nil),                   // 6: Items
	(*Metadata)(nil),                // 7: Metadata
	(*_struct.Struct)(nil),          // 8: google.protobuf.Struct
	(*timestamp.Timestamp)(nil),     // 9: google.protobuf.Timestamp
	(*duration.Duration)(nil),       // 10: google.protobuf.Duration
}
var file_items_proto_depIdxs = []int32{
	0,  // 0: ItemRequest.method:type_name -> RequestMethod
	1,  // 1: ItemRequestError.errorType:type_name -> ItemRequestError.ErrorType
	8,  // 2: ItemAttributes.attrStruct:type_name -> google.protobuf.Struct
	4,  // 3: Item.attributes:type_name -> ItemAttributes
	7,  // 4: Item.metadata:type_name -> Metadata
	2,  // 5: Item.linkedItemRequests:type_name -> ItemRequest
	5,  // 6: Item.linkedItems:type_name -> Item
	5,  // 7: Items.items:type_name -> Item
	0,  // 8: Metadata.requestMethod:type_name -> RequestMethod
	9,  // 9: Metadata.timestamp:type_name -> google.protobuf.Timestamp
	10, // 10: Metadata.backendDuration:type_name -> google.protobuf.Duration
	10, // 11: Metadata.backendDurationPerItem:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_items_proto_init() }
func file_items_proto_init() {
	if File_items_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_items_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_items_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemRequestError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_items_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemAttributes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_items_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_items_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Items); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_items_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_items_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_items_proto_goTypes,
		DependencyIndexes: file_items_proto_depIdxs,
		EnumInfos:         file_items_proto_enumTypes,
		MessageInfos:      file_items_proto_msgTypes,
	}.Build()
	File_items_proto = out.File
	file_items_proto_rawDesc = nil
	file_items_proto_goTypes = nil
	file_items_proto_depIdxs = nil
}
