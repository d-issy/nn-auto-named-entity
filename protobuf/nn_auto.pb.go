// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nn_auto.proto

package protobuf

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Wiki struct {
	Pages                map[string]*Page `protobuf:"bytes,1,rep,name=pages,proto3" json:"pages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Wiki) Reset()         { *m = Wiki{} }
func (m *Wiki) String() string { return proto.CompactTextString(m) }
func (*Wiki) ProtoMessage()    {}
func (*Wiki) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56b904611e84f2b, []int{0}
}

func (m *Wiki) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wiki.Unmarshal(m, b)
}
func (m *Wiki) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wiki.Marshal(b, m, deterministic)
}
func (m *Wiki) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wiki.Merge(m, src)
}
func (m *Wiki) XXX_Size() int {
	return xxx_messageInfo_Wiki.Size(m)
}
func (m *Wiki) XXX_DiscardUnknown() {
	xxx_messageInfo_Wiki.DiscardUnknown(m)
}

var xxx_messageInfo_Wiki proto.InternalMessageInfo

func (m *Wiki) GetPages() map[string]*Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

type Page struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Namespace            uint32   `protobuf:"varint,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Links                []string `protobuf:"bytes,4,rep,name=links,proto3" json:"links,omitempty"`
	Categories           []string `protobuf:"bytes,5,rep,name=categories,proto3" json:"categories,omitempty"`
	Templates            []string `protobuf:"bytes,6,rep,name=templates,proto3" json:"templates,omitempty"`
	NamedEntityIDs       []uint32 `protobuf:"varint,7,rep,packed,name=namedEntityIDs,proto3" json:"namedEntityIDs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56b904611e84f2b, []int{1}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Page) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Page) GetNamespace() uint32 {
	if m != nil {
		return m.Namespace
	}
	return 0
}

func (m *Page) GetLinks() []string {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Page) GetCategories() []string {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Page) GetTemplates() []string {
	if m != nil {
		return m.Templates
	}
	return nil
}

func (m *Page) GetNamedEntityIDs() []uint32 {
	if m != nil {
		return m.NamedEntityIDs
	}
	return nil
}

type Data struct {
	NamedEntities        map[uint32]*NamedEntity         `protobuf:"bytes,1,rep,name=namedEntities,proto3" json:"namedEntities,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TaggedPages          map[string]*TaggedNamedEntities `protobuf:"bytes,2,rep,name=taggedPages,proto3" json:"taggedPages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *Data) Reset()         { *m = Data{} }
func (m *Data) String() string { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()    {}
func (*Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56b904611e84f2b, []int{2}
}

func (m *Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Data.Unmarshal(m, b)
}
func (m *Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Data.Marshal(b, m, deterministic)
}
func (m *Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Data.Merge(m, src)
}
func (m *Data) XXX_Size() int {
	return xxx_messageInfo_Data.Size(m)
}
func (m *Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Data proto.InternalMessageInfo

func (m *Data) GetNamedEntities() map[uint32]*NamedEntity {
	if m != nil {
		return m.NamedEntities
	}
	return nil
}

func (m *Data) GetTaggedPages() map[string]*TaggedNamedEntities {
	if m != nil {
		return m.TaggedPages
	}
	return nil
}

type NamedEntity struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	JpName               string   `protobuf:"bytes,2,opt,name=jpName,proto3" json:"jpName,omitempty"`
	EnName               string   `protobuf:"bytes,3,opt,name=enName,proto3" json:"enName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NamedEntity) Reset()         { *m = NamedEntity{} }
func (m *NamedEntity) String() string { return proto.CompactTextString(m) }
func (*NamedEntity) ProtoMessage()    {}
func (*NamedEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56b904611e84f2b, []int{3}
}

func (m *NamedEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NamedEntity.Unmarshal(m, b)
}
func (m *NamedEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NamedEntity.Marshal(b, m, deterministic)
}
func (m *NamedEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamedEntity.Merge(m, src)
}
func (m *NamedEntity) XXX_Size() int {
	return xxx_messageInfo_NamedEntity.Size(m)
}
func (m *NamedEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_NamedEntity.DiscardUnknown(m)
}

var xxx_messageInfo_NamedEntity proto.InternalMessageInfo

func (m *NamedEntity) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *NamedEntity) GetJpName() string {
	if m != nil {
		return m.JpName
	}
	return ""
}

func (m *NamedEntity) GetEnName() string {
	if m != nil {
		return m.EnName
	}
	return ""
}

type TaggedNamedEntities struct {
	Attrs                []uint32 `protobuf:"varint,1,rep,packed,name=attrs,proto3" json:"attrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaggedNamedEntities) Reset()         { *m = TaggedNamedEntities{} }
func (m *TaggedNamedEntities) String() string { return proto.CompactTextString(m) }
func (*TaggedNamedEntities) ProtoMessage()    {}
func (*TaggedNamedEntities) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56b904611e84f2b, []int{4}
}

func (m *TaggedNamedEntities) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaggedNamedEntities.Unmarshal(m, b)
}
func (m *TaggedNamedEntities) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaggedNamedEntities.Marshal(b, m, deterministic)
}
func (m *TaggedNamedEntities) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaggedNamedEntities.Merge(m, src)
}
func (m *TaggedNamedEntities) XXX_Size() int {
	return xxx_messageInfo_TaggedNamedEntities.Size(m)
}
func (m *TaggedNamedEntities) XXX_DiscardUnknown() {
	xxx_messageInfo_TaggedNamedEntities.DiscardUnknown(m)
}

var xxx_messageInfo_TaggedNamedEntities proto.InternalMessageInfo

func (m *TaggedNamedEntities) GetAttrs() []uint32 {
	if m != nil {
		return m.Attrs
	}
	return nil
}

type Statistics struct {
	Categories           []*Counter `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
	Links                []*Counter `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	Templates            []*Counter `protobuf:"bytes,3,rep,name=templates,proto3" json:"templates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Statistics) Reset()         { *m = Statistics{} }
func (m *Statistics) String() string { return proto.CompactTextString(m) }
func (*Statistics) ProtoMessage()    {}
func (*Statistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56b904611e84f2b, []int{5}
}

func (m *Statistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Statistics.Unmarshal(m, b)
}
func (m *Statistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Statistics.Marshal(b, m, deterministic)
}
func (m *Statistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Statistics.Merge(m, src)
}
func (m *Statistics) XXX_Size() int {
	return xxx_messageInfo_Statistics.Size(m)
}
func (m *Statistics) XXX_DiscardUnknown() {
	xxx_messageInfo_Statistics.DiscardUnknown(m)
}

var xxx_messageInfo_Statistics proto.InternalMessageInfo

func (m *Statistics) GetCategories() []*Counter {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Statistics) GetLinks() []*Counter {
	if m != nil {
		return m.Links
	}
	return nil
}

func (m *Statistics) GetTemplates() []*Counter {
	if m != nil {
		return m.Templates
	}
	return nil
}

type Counter struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Count                uint64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Counter) Reset()         { *m = Counter{} }
func (m *Counter) String() string { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()    {}
func (*Counter) Descriptor() ([]byte, []int) {
	return fileDescriptor_f56b904611e84f2b, []int{6}
}

func (m *Counter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Counter.Unmarshal(m, b)
}
func (m *Counter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Counter.Marshal(b, m, deterministic)
}
func (m *Counter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Counter.Merge(m, src)
}
func (m *Counter) XXX_Size() int {
	return xxx_messageInfo_Counter.Size(m)
}
func (m *Counter) XXX_DiscardUnknown() {
	xxx_messageInfo_Counter.DiscardUnknown(m)
}

var xxx_messageInfo_Counter proto.InternalMessageInfo

func (m *Counter) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Counter) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*Wiki)(nil), "protobuf.Wiki")
	proto.RegisterMapType((map[string]*Page)(nil), "protobuf.Wiki.PagesEntry")
	proto.RegisterType((*Page)(nil), "protobuf.Page")
	proto.RegisterType((*Data)(nil), "protobuf.Data")
	proto.RegisterMapType((map[uint32]*NamedEntity)(nil), "protobuf.Data.NamedEntitiesEntry")
	proto.RegisterMapType((map[string]*TaggedNamedEntities)(nil), "protobuf.Data.TaggedPagesEntry")
	proto.RegisterType((*NamedEntity)(nil), "protobuf.NamedEntity")
	proto.RegisterType((*TaggedNamedEntities)(nil), "protobuf.TaggedNamedEntities")
	proto.RegisterType((*Statistics)(nil), "protobuf.Statistics")
	proto.RegisterType((*Counter)(nil), "protobuf.Counter")
}

func init() { proto.RegisterFile("nn_auto.proto", fileDescriptor_f56b904611e84f2b) }

var fileDescriptor_f56b904611e84f2b = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xd5, 0xfa, 0x23, 0x25, 0x13, 0x39, 0x2a, 0x4b, 0x41, 0x4b, 0xc4, 0x87, 0xb1, 0x10, 0x58,
	0xaa, 0xe4, 0xaa, 0xed, 0x05, 0x71, 0x43, 0xb4, 0x02, 0x0e, 0x20, 0xb4, 0x20, 0xf5, 0x84, 0xd0,
	0x36, 0x59, 0xa2, 0x25, 0xa9, 0x6d, 0x79, 0x27, 0x48, 0x39, 0xf2, 0x33, 0xf8, 0x43, 0x1c, 0xf8,
	0x55, 0x68, 0x77, 0x53, 0x76, 0x1d, 0xd2, 0x93, 0x3d, 0x6f, 0xde, 0x7b, 0x63, 0xcf, 0x3c, 0xc8,
	0xea, 0xfa, 0xab, 0x58, 0x61, 0x53, 0xb5, 0x5d, 0x83, 0x0d, 0xbd, 0x65, 0x1f, 0x97, 0xab, 0x6f,
	0xc5, 0x4f, 0x02, 0xc9, 0x85, 0x5a, 0x28, 0x7a, 0x04, 0x69, 0x2b, 0xe6, 0x52, 0x33, 0x92, 0xc7,
	0xe5, 0xe8, 0xe4, 0x7e, 0x75, 0x4d, 0xa9, 0x4c, 0xbb, 0xfa, 0x68, 0x7a, 0xe7, 0x35, 0x76, 0x6b,
	0xee, 0x78, 0x93, 0xb7, 0x00, 0x1e, 0xa4, 0xfb, 0x10, 0x2f, 0xe4, 0x9a, 0x91, 0x9c, 0x94, 0x43,
	0x6e, 0x5e, 0xe9, 0x53, 0x48, 0x7f, 0x88, 0xe5, 0x4a, 0xb2, 0x28, 0x27, 0xe5, 0xe8, 0x64, 0xec,
	0x0d, 0x8d, 0x8c, 0xbb, 0xe6, 0xcb, 0xe8, 0x05, 0x29, 0x7e, 0x13, 0x48, 0x0c, 0x46, 0xc7, 0x10,
	0xa9, 0x99, 0xf5, 0x48, 0x78, 0xa4, 0x66, 0xf4, 0x00, 0x52, 0x54, 0xb8, 0x74, 0x16, 0x43, 0xee,
	0x0a, 0xfa, 0x00, 0x86, 0xb5, 0xb8, 0x92, 0xba, 0x15, 0x53, 0xc9, 0xe2, 0x9c, 0x94, 0x19, 0xf7,
	0x80, 0xd1, 0x2c, 0x55, 0xbd, 0xd0, 0x2c, 0xc9, 0x63, 0xa3, 0xb1, 0x05, 0x7d, 0x04, 0x30, 0x15,
	0x28, 0xe7, 0x4d, 0xa7, 0xa4, 0x66, 0xa9, 0x6d, 0x05, 0x88, 0xf1, 0x44, 0x79, 0xd5, 0x2e, 0x05,
	0x4a, 0xcd, 0x06, 0xb6, 0xed, 0x01, 0xfa, 0x0c, 0xc6, 0x66, 0xc0, 0xec, 0xbc, 0x46, 0x85, 0xeb,
	0x77, 0x67, 0x9a, 0xed, 0xe5, 0x71, 0x99, 0xf1, 0x2d, 0xb4, 0xf8, 0x13, 0x41, 0x72, 0x26, 0x50,
	0xd0, 0x37, 0x90, 0xf9, 0x96, 0xfa, 0xb7, 0xd4, 0x27, 0x7e, 0x07, 0x86, 0x56, 0x7d, 0x08, 0x39,
	0x6e, 0xb9, 0x7d, 0x1d, 0x7d, 0x05, 0x23, 0x14, 0xf3, 0xb9, 0x9c, 0xd9, 0x55, 0xb3, 0xc8, 0xda,
	0x3c, 0xde, 0xb2, 0xf9, 0xec, 0x19, 0xce, 0x24, 0xd4, 0x4c, 0x2e, 0x80, 0xfe, 0x3f, 0x27, 0xbc,
	0x57, 0xe6, 0xee, 0x75, 0xd8, 0xbf, 0xd7, 0x5d, 0x3f, 0xc4, 0xcb, 0xd7, 0xc1, 0xd9, 0x26, 0x5f,
	0x60, 0x7f, 0x7b, 0xf2, 0x8e, 0x18, 0x9c, 0xf6, 0x6d, 0x1f, 0x7a, 0x5b, 0x27, 0xee, 0x7d, 0x5b,
	0x98, 0x8a, 0xf7, 0x30, 0x0a, 0x06, 0x07, 0xd9, 0xc8, 0x6c, 0x36, 0xee, 0xc1, 0xe0, 0x7b, 0x6b,
	0x08, 0x9b, 0x70, 0x6c, 0x2a, 0x83, 0xcb, 0xda, 0xe2, 0xb1, 0xc3, 0x5d, 0x55, 0x1c, 0xc2, 0x9d,
	0x1d, 0x03, 0x4d, 0x5c, 0x04, 0x62, 0xe7, 0x2e, 0x94, 0x71, 0x57, 0x14, 0xbf, 0x08, 0xc0, 0x27,
	0x14, 0xa8, 0x34, 0xaa, 0xa9, 0xa6, 0xc7, 0xbd, 0xf4, 0xb8, 0x5b, 0xde, 0xf6, 0x3f, 0xf2, 0xba,
	0x59, 0xd5, 0x28, 0xbb, 0x5e, 0xa0, 0x9e, 0x5f, 0xc7, 0x30, 0xba, 0x89, 0xbd, 0x49, 0xe6, 0x51,
	0x98, 0xbc, 0xf8, 0x26, 0xb2, 0xe7, 0x14, 0xc7, 0xb0, 0xb7, 0x41, 0x77, 0x6c, 0xfb, 0x00, 0xd2,
	0xa9, 0x69, 0xda, 0xa5, 0x24, 0xdc, 0x15, 0x97, 0x03, 0xeb, 0x77, 0xfa, 0x37, 0x00, 0x00, 0xff,
	0xff, 0xe4, 0x3d, 0x4c, 0xa4, 0x06, 0x04, 0x00, 0x00,
}
