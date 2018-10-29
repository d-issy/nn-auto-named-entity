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
	NamedEntities        map[uint32]*NamedEntity `protobuf:"bytes,1,rep,name=namedEntities,proto3" json:"namedEntities,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
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

func init() {
	proto.RegisterType((*Wiki)(nil), "protobuf.Wiki")
	proto.RegisterMapType((map[string]*Page)(nil), "protobuf.Wiki.PagesEntry")
	proto.RegisterType((*Page)(nil), "protobuf.Page")
	proto.RegisterType((*Data)(nil), "protobuf.Data")
	proto.RegisterMapType((map[uint32]*NamedEntity)(nil), "protobuf.Data.NamedEntitiesEntry")
	proto.RegisterType((*NamedEntity)(nil), "protobuf.NamedEntity")
}

func init() { proto.RegisterFile("nn_auto.proto", fileDescriptor_f56b904611e84f2b) }

var fileDescriptor_f56b904611e84f2b = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcb, 0x4a, 0x33, 0x31,
	0x14, 0xc7, 0xc9, 0x5c, 0xfa, 0x7d, 0x3d, 0x65, 0x8a, 0x04, 0x95, 0x58, 0x44, 0xc6, 0x22, 0x32,
	0x20, 0x8c, 0x50, 0x37, 0xe2, 0xba, 0x45, 0x5d, 0x28, 0x92, 0x4d, 0x97, 0x92, 0xb6, 0xb1, 0xc4,
	0x4e, 0x33, 0x43, 0x27, 0x23, 0xcc, 0xd2, 0xb7, 0xf1, 0x49, 0x7c, 0x2e, 0x49, 0x52, 0x4d, 0xaa,
	0xab, 0xe4, 0xfc, 0x2f, 0xe7, 0xc0, 0x0f, 0x12, 0x29, 0x9f, 0x59, 0xa3, 0xca, 0xbc, 0xda, 0x94,
	0xaa, 0xc4, 0xff, 0xcd, 0x33, 0x6b, 0x5e, 0x86, 0xef, 0x08, 0xa2, 0xa9, 0x58, 0x09, 0x7c, 0x09,
	0x71, 0xc5, 0x96, 0xbc, 0x26, 0x28, 0x0d, 0xb3, 0xde, 0xe8, 0x28, 0xff, 0x8e, 0xe4, 0xda, 0xce,
	0x9f, 0xb4, 0x37, 0x91, 0x6a, 0xd3, 0x52, 0x9b, 0x1b, 0xdc, 0x01, 0x38, 0x11, 0xef, 0x41, 0xb8,
	0xe2, 0x2d, 0x41, 0x29, 0xca, 0xba, 0x54, 0x7f, 0xf1, 0x19, 0xc4, 0x6f, 0xac, 0x68, 0x38, 0x09,
	0x52, 0x94, 0xf5, 0x46, 0x7d, 0xb7, 0x50, 0xd7, 0xa8, 0x35, 0x6f, 0x82, 0x6b, 0x34, 0xfc, 0x44,
	0x10, 0x69, 0x0d, 0xf7, 0x21, 0x10, 0x0b, 0xb3, 0x23, 0xa2, 0x81, 0x58, 0xe0, 0x7d, 0x88, 0x95,
	0x50, 0x85, 0x5d, 0xd1, 0xa5, 0x76, 0xc0, 0xc7, 0xd0, 0x95, 0x6c, 0xcd, 0xeb, 0x8a, 0xcd, 0x39,
	0x09, 0x53, 0x94, 0x25, 0xd4, 0x09, 0xba, 0x53, 0x08, 0xb9, 0xaa, 0x49, 0x94, 0x86, 0xba, 0x63,
	0x06, 0x7c, 0x02, 0x30, 0x67, 0x8a, 0x2f, 0xcb, 0x8d, 0xe0, 0x35, 0x89, 0x8d, 0xe5, 0x29, 0x7a,
	0xa7, 0xe2, 0xeb, 0xaa, 0x60, 0x8a, 0xd7, 0xa4, 0x63, 0x6c, 0x27, 0xe0, 0x73, 0xe8, 0xeb, 0x03,
	0x8b, 0x89, 0x54, 0x42, 0xb5, 0xf7, 0xe3, 0x9a, 0xfc, 0x4b, 0xc3, 0x2c, 0xa1, 0xbf, 0xd4, 0xe1,
	0x07, 0x82, 0x68, 0xcc, 0x14, 0xc3, 0xb7, 0x90, 0x38, 0x4b, 0xfc, 0x40, 0x3d, 0x75, 0x0c, 0x74,
	0x2c, 0x7f, 0xf4, 0x33, 0x16, 0xee, 0x6e, 0x6f, 0x30, 0x05, 0xfc, 0x37, 0xe4, 0xc3, 0x4e, 0x2c,
	0xec, 0x8b, 0x5d, 0xd8, 0x07, 0xee, 0x90, 0xab, 0xb7, 0x3e, 0xf3, 0x07, 0xe8, 0x79, 0x8e, 0x47,
	0x3e, 0x31, 0xe4, 0x0f, 0xa1, 0xf3, 0x5a, 0xe9, 0xc0, 0x16, 0xfd, 0x76, 0xd2, 0x3a, 0x97, 0x46,
	0x0f, 0xad, 0x6e, 0xa7, 0x59, 0xc7, 0xdc, 0xbb, 0xfa, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x19, 0x31,
	0x27, 0x95, 0x68, 0x02, 0x00, 0x00,
}