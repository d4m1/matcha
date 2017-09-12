// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/view/scrollview.proto

package view

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import matcha_layout "gomatcha.io/matcha/pb/layout"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ScrollView struct {
	ScrollEnabled                  bool `protobuf:"varint,1,opt,name=scrollEnabled" json:"scrollEnabled,omitempty"`
	Horizontal                     bool `protobuf:"varint,4,opt,name=horizontal" json:"horizontal,omitempty"`
	Vertical                       bool `protobuf:"varint,5,opt,name=vertical" json:"vertical,omitempty"`
	ShowsHorizontalScrollIndicator bool `protobuf:"varint,2,opt,name=showsHorizontalScrollIndicator" json:"showsHorizontalScrollIndicator,omitempty"`
	ShowsVerticalScrollIndicator   bool `protobuf:"varint,3,opt,name=showsVerticalScrollIndicator" json:"showsVerticalScrollIndicator,omitempty"`
}

func (m *ScrollView) Reset()                    { *m = ScrollView{} }
func (m *ScrollView) String() string            { return proto.CompactTextString(m) }
func (*ScrollView) ProtoMessage()               {}
func (*ScrollView) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *ScrollView) GetScrollEnabled() bool {
	if m != nil {
		return m.ScrollEnabled
	}
	return false
}

func (m *ScrollView) GetHorizontal() bool {
	if m != nil {
		return m.Horizontal
	}
	return false
}

func (m *ScrollView) GetVertical() bool {
	if m != nil {
		return m.Vertical
	}
	return false
}

func (m *ScrollView) GetShowsHorizontalScrollIndicator() bool {
	if m != nil {
		return m.ShowsHorizontalScrollIndicator
	}
	return false
}

func (m *ScrollView) GetShowsVerticalScrollIndicator() bool {
	if m != nil {
		return m.ShowsVerticalScrollIndicator
	}
	return false
}

type ScrollEvent struct {
	ContentOffset *matcha_layout.Point `protobuf:"bytes,1,opt,name=contentOffset" json:"contentOffset,omitempty"`
}

func (m *ScrollEvent) Reset()                    { *m = ScrollEvent{} }
func (m *ScrollEvent) String() string            { return proto.CompactTextString(m) }
func (*ScrollEvent) ProtoMessage()               {}
func (*ScrollEvent) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ScrollEvent) GetContentOffset() *matcha_layout.Point {
	if m != nil {
		return m.ContentOffset
	}
	return nil
}

func init() {
	proto.RegisterType((*ScrollView)(nil), "matcha.view.ScrollView")
	proto.RegisterType((*ScrollEvent)(nil), "matcha.view.ScrollEvent")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/pb/view/scrollview.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 284 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x69, 0xad, 0x52, 0x26, 0xed, 0x65, 0xf1, 0x10, 0x82, 0x14, 0x29, 0x1e, 0x14, 0x21,
	0x01, 0xbd, 0x79, 0x11, 0x02, 0x15, 0x7b, 0x10, 0x43, 0x85, 0x1e, 0xbc, 0x6d, 0xd2, 0xad, 0x59,
	0x58, 0x77, 0x42, 0x32, 0x26, 0xe8, 0xcf, 0xf1, 0x67, 0x7a, 0x92, 0x4c, 0x96, 0x68, 0x45, 0xf4,
	0x34, 0xb3, 0x6f, 0xbf, 0x79, 0xcb, 0xbe, 0x81, 0xf3, 0x27, 0x7c, 0x96, 0x94, 0xe5, 0x32, 0xd4,
	0x18, 0x75, 0x5d, 0x54, 0xa4, 0x51, 0xad, 0x55, 0x13, 0x55, 0x59, 0x89, 0xc6, 0xb4, 0x6d, 0x58,
	0x94, 0x48, 0x28, 0x3c, 0x87, 0xb6, 0x52, 0x70, 0xf6, 0xfb, 0xa4, 0x91, 0xaf, 0xf8, 0x42, 0xae,
	0x74, 0x73, 0xf3, 0x8f, 0x01, 0xc0, 0x03, 0x9b, 0xad, 0xb5, 0x6a, 0xc4, 0x09, 0x4c, 0x3b, 0xeb,
	0x85, 0x95, 0xa9, 0x51, 0x1b, 0x7f, 0x70, 0x3c, 0x38, 0x1d, 0xaf, 0x76, 0x45, 0x31, 0x03, 0xc8,
	0xb1, 0xd4, 0x6f, 0x68, 0x49, 0x1a, 0x7f, 0xc4, 0xc8, 0x37, 0x45, 0x04, 0x30, 0xae, 0x55, 0x49,
	0x3a, 0x93, 0xc6, 0xdf, 0xe7, 0xdb, 0xfe, 0x2c, 0x6e, 0x60, 0x56, 0xe5, 0xd8, 0x54, 0xb7, 0x3d,
	0xde, 0x3d, 0xbf, 0xb4, 0x1b, 0x9d, 0x49, 0xc2, 0xd2, 0x1f, 0xf2, 0xc4, 0x3f, 0x94, 0x88, 0xe1,
	0x88, 0x89, 0xb5, 0x33, 0xfe, 0xe9, 0xb2, 0xc7, 0x2e, 0x7f, 0x32, 0xf3, 0x25, 0x78, 0x9d, 0xb4,
	0xa8, 0x95, 0x25, 0x71, 0x05, 0xd3, 0x0c, 0x2d, 0x29, 0x4b, 0xf7, 0xdb, 0x6d, 0xa5, 0x88, 0x3f,
	0xef, 0x5d, 0x1c, 0x86, 0x2e, 0x4c, 0x17, 0x5c, 0x82, 0xda, 0xd2, 0x6a, 0x17, 0x8d, 0xaf, 0x21,
	0xd0, 0x18, 0xf6, 0xb9, 0xbb, 0x52, 0xa4, 0xbc, 0x90, 0x78, 0x92, 0xa4, 0x5f, 0x21, 0x3f, 0x8e,
	0x5a, 0xed, 0x7d, 0x38, 0xb9, 0x63, 0xa8, 0x95, 0x92, 0x38, 0x3d, 0xe0, 0x7d, 0x5c, 0x7e, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x73, 0x0b, 0xa6, 0x51, 0xf6, 0x01, 0x00, 0x00,
}
