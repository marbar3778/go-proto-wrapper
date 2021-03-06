// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: examples/foo.proto

package wrapper_examples

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/marbar3778/go-proto-wrapper"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type FooMsg struct {
	// Types that are valid to be assigned to Sum:
	//	*FooMsg_Foo
	//	*FooMsg_Bar
	//	*FooMsg_Baz
	Sum                  isFooMsg_Sum `protobuf_oneof:"sum"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *FooMsg) Reset()         { *m = FooMsg{} }
func (m *FooMsg) String() string { return proto.CompactTextString(m) }
func (*FooMsg) ProtoMessage()    {}
func (*FooMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ef2189187829863, []int{0}
}
func (m *FooMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FooMsg.Unmarshal(m, b)
}
func (m *FooMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FooMsg.Marshal(b, m, deterministic)
}
func (m *FooMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FooMsg.Merge(m, src)
}
func (m *FooMsg) XXX_Size() int {
	return xxx_messageInfo_FooMsg.Size(m)
}
func (m *FooMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_FooMsg.DiscardUnknown(m)
}

var xxx_messageInfo_FooMsg proto.InternalMessageInfo

type isFooMsg_Sum interface {
	isFooMsg_Sum()
}

type FooMsg_Foo struct {
	Foo *Foo `protobuf:"bytes,1,opt,name=foo,proto3,oneof" json:"foo,omitempty"`
}
type FooMsg_Bar struct {
	Bar *Bar `protobuf:"bytes,2,opt,name=bar,proto3,oneof" json:"bar,omitempty"`
}
type FooMsg_Baz struct {
	Baz *Baz `protobuf:"bytes,3,opt,name=baz,proto3,oneof" json:"baz,omitempty"`
}

func (*FooMsg_Foo) isFooMsg_Sum() {}
func (*FooMsg_Bar) isFooMsg_Sum() {}
func (*FooMsg_Baz) isFooMsg_Sum() {}

func (m *FooMsg) GetSum() isFooMsg_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *FooMsg) GetFoo() *Foo {
	if x, ok := m.GetSum().(*FooMsg_Foo); ok {
		return x.Foo
	}
	return nil
}

func (m *FooMsg) GetBar() *Bar {
	if x, ok := m.GetSum().(*FooMsg_Bar); ok {
		return x.Bar
	}
	return nil
}

func (m *FooMsg) GetBaz() *Baz {
	if x, ok := m.GetSum().(*FooMsg_Baz); ok {
		return x.Baz
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FooMsg) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FooMsg_Foo)(nil),
		(*FooMsg_Bar)(nil),
		(*FooMsg_Baz)(nil),
	}
}

type BarMsg struct {
	// Types that are valid to be assigned to Sum:
	//	*BarMsg_Foo
	//	*BarMsg_Bar
	//	*BarMsg_Baz
	Sum                  isBarMsg_Sum `protobuf_oneof:"sum"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BarMsg) Reset()         { *m = BarMsg{} }
func (m *BarMsg) String() string { return proto.CompactTextString(m) }
func (*BarMsg) ProtoMessage()    {}
func (*BarMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ef2189187829863, []int{1}
}
func (m *BarMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BarMsg.Unmarshal(m, b)
}
func (m *BarMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BarMsg.Marshal(b, m, deterministic)
}
func (m *BarMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BarMsg.Merge(m, src)
}
func (m *BarMsg) XXX_Size() int {
	return xxx_messageInfo_BarMsg.Size(m)
}
func (m *BarMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_BarMsg.DiscardUnknown(m)
}

var xxx_messageInfo_BarMsg proto.InternalMessageInfo

type isBarMsg_Sum interface {
	isBarMsg_Sum()
}

type BarMsg_Foo struct {
	Foo *Foo `protobuf:"bytes,1,opt,name=foo,proto3,oneof" json:"foo,omitempty"`
}
type BarMsg_Bar struct {
	Bar *Bar `protobuf:"bytes,2,opt,name=bar,proto3,oneof" json:"bar,omitempty"`
}
type BarMsg_Baz struct {
	Baz *Baz `protobuf:"bytes,3,opt,name=baz,proto3,oneof" json:"baz,omitempty"`
}

func (*BarMsg_Foo) isBarMsg_Sum() {}
func (*BarMsg_Bar) isBarMsg_Sum() {}
func (*BarMsg_Baz) isBarMsg_Sum() {}

func (m *BarMsg) GetSum() isBarMsg_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *BarMsg) GetFoo() *Foo {
	if x, ok := m.GetSum().(*BarMsg_Foo); ok {
		return x.Foo
	}
	return nil
}

func (m *BarMsg) GetBar() *Bar {
	if x, ok := m.GetSum().(*BarMsg_Bar); ok {
		return x.Bar
	}
	return nil
}

func (m *BarMsg) GetBaz() *Baz {
	if x, ok := m.GetSum().(*BarMsg_Baz); ok {
		return x.Baz
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BarMsg) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BarMsg_Foo)(nil),
		(*BarMsg_Bar)(nil),
		(*BarMsg_Baz)(nil),
	}
}

type Foo struct {
	A                    uint64   `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()         { *m = Foo{} }
func (m *Foo) String() string { return proto.CompactTextString(m) }
func (*Foo) ProtoMessage()    {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ef2189187829863, []int{2}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return xxx_messageInfo_Foo.Size(m)
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

func (m *Foo) GetA() uint64 {
	if m != nil {
		return m.A
	}
	return 0
}

type Bar struct {
	A                    uint64   `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Bar) Reset()         { *m = Bar{} }
func (m *Bar) String() string { return proto.CompactTextString(m) }
func (*Bar) ProtoMessage()    {}
func (*Bar) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ef2189187829863, []int{3}
}
func (m *Bar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bar.Unmarshal(m, b)
}
func (m *Bar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bar.Marshal(b, m, deterministic)
}
func (m *Bar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bar.Merge(m, src)
}
func (m *Bar) XXX_Size() int {
	return xxx_messageInfo_Bar.Size(m)
}
func (m *Bar) XXX_DiscardUnknown() {
	xxx_messageInfo_Bar.DiscardUnknown(m)
}

var xxx_messageInfo_Bar proto.InternalMessageInfo

func (m *Bar) GetA() uint64 {
	if m != nil {
		return m.A
	}
	return 0
}

type Baz struct {
	A                    uint64   `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Baz) Reset()         { *m = Baz{} }
func (m *Baz) String() string { return proto.CompactTextString(m) }
func (*Baz) ProtoMessage()    {}
func (*Baz) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ef2189187829863, []int{4}
}
func (m *Baz) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Baz.Unmarshal(m, b)
}
func (m *Baz) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Baz.Marshal(b, m, deterministic)
}
func (m *Baz) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Baz.Merge(m, src)
}
func (m *Baz) XXX_Size() int {
	return xxx_messageInfo_Baz.Size(m)
}
func (m *Baz) XXX_DiscardUnknown() {
	xxx_messageInfo_Baz.DiscardUnknown(m)
}

var xxx_messageInfo_Baz proto.InternalMessageInfo

func (m *Baz) GetA() uint64 {
	if m != nil {
		return m.A
	}
	return 0
}

func init() {
	proto.RegisterType((*FooMsg)(nil), "wrapper.examples.FooMsg")
	proto.RegisterType((*BarMsg)(nil), "wrapper.examples.BarMsg")
	proto.RegisterType((*Foo)(nil), "wrapper.examples.foo")
	proto.RegisterType((*Bar)(nil), "wrapper.examples.bar")
	proto.RegisterType((*Baz)(nil), "wrapper.examples.baz")
}

func init() { proto.RegisterFile("examples/foo.proto", fileDescriptor_1ef2189187829863) }

var fileDescriptor_1ef2189187829863 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x2d, 0xd6, 0x4f, 0xcb, 0xcf, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x28, 0x2f, 0x4a, 0x2c, 0x28, 0x48, 0x2d, 0xd2, 0x83, 0xc9, 0x49, 0xf1, 0xc2, 0x44, 0xc0, 0x0a,
	0x94, 0xa6, 0x33, 0x72, 0xb1, 0xb9, 0xe5, 0xe7, 0xfb, 0x16, 0xa7, 0x0b, 0x69, 0x72, 0x31, 0xa7,
	0xe5, 0xe7, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xea, 0xa1, 0xeb, 0xd4, 0x4b, 0xcb,
	0xcf, 0xf7, 0x60, 0x08, 0x02, 0xa9, 0x01, 0x29, 0x4d, 0x4a, 0x2c, 0x92, 0x60, 0xc2, 0xa5, 0x34,
	0x29, 0xb1, 0x08, 0xa4, 0x34, 0x29, 0xb1, 0x08, 0xa2, 0xb4, 0x4a, 0x82, 0x19, 0xb7, 0xd2, 0x2a,
	0x88, 0xd2, 0x2a, 0x2b, 0xd6, 0x05, 0x8f, 0x3e, 0x49, 0x30, 0x3a, 0xb1, 0x72, 0x31, 0x17, 0x97,
	0xe6, 0x2a, 0x4d, 0x60, 0xe4, 0x62, 0x73, 0x4a, 0x2c, 0x1a, 0x0c, 0x2e, 0x83, 0x39, 0x49, 0x18,
	0xec, 0x0e, 0x21, 0x1e, 0x2e, 0xc6, 0x44, 0xb0, 0x63, 0x58, 0x82, 0x18, 0x13, 0x41, 0x82, 0x20,
	0xd3, 0xb0, 0x08, 0x56, 0xa1, 0x0a, 0x26, 0xb1, 0x81, 0x83, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0xa7, 0x97, 0xda, 0x3a, 0xa9, 0x01, 0x00, 0x00,
}
