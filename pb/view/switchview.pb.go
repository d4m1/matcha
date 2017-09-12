// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/view/switchview.proto

package view

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SwitchView struct {
	Value   bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Enabled bool `protobuf:"varint,2,opt,name=enabled" json:"enabled,omitempty"`
}

func (m *SwitchView) Reset()                    { *m = SwitchView{} }
func (m *SwitchView) String() string            { return proto.CompactTextString(m) }
func (*SwitchView) ProtoMessage()               {}
func (*SwitchView) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *SwitchView) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func (m *SwitchView) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type SwitchEvent struct {
	Value bool `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *SwitchEvent) Reset()                    { *m = SwitchEvent{} }
func (m *SwitchEvent) String() string            { return proto.CompactTextString(m) }
func (*SwitchEvent) ProtoMessage()               {}
func (*SwitchEvent) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *SwitchEvent) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

func init() {
	proto.RegisterType((*SwitchView)(nil), "matcha.view.SwitchView")
	proto.RegisterType((*SwitchEvent)(nil), "matcha.view.SwitchEvent")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/pb/view/switchview.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4e, 0xcf, 0xcf, 0x4d,
	0x2c, 0x49, 0xce, 0x48, 0xd4, 0xcb, 0xcc, 0xd7, 0x87, 0xb0, 0xf4, 0x0b, 0x92, 0xf4, 0xcb, 0x32,
	0x53, 0xcb, 0xf5, 0x8b, 0xcb, 0x33, 0x4b, 0x92, 0x33, 0x40, 0x4c, 0xbd, 0x82, 0xa2, 0xfc, 0x92,
	0x7c, 0x21, 0x6e, 0xa8, 0x52, 0x90, 0x90, 0x92, 0x0d, 0x17, 0x57, 0x30, 0x58, 0x41, 0x58, 0x66,
	0x6a, 0xb9, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x69, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x47, 0x10, 0x84, 0x23, 0x24, 0xc1, 0xc5, 0x9e, 0x9a, 0x97, 0x98, 0x94, 0x93, 0x9a, 0x22, 0xc1,
	0x04, 0x16, 0x87, 0x71, 0x95, 0x94, 0xb9, 0xb8, 0x21, 0xba, 0x5d, 0xcb, 0x52, 0xf3, 0x4a, 0xb0,
	0x6b, 0x77, 0xb2, 0xe7, 0x92, 0xca, 0xcc, 0xd7, 0x83, 0xbb, 0x10, 0x4a, 0x15, 0x24, 0x81, 0x1d,
	0xe0, 0xc4, 0x13, 0x90, 0x84, 0x70, 0x40, 0x14, 0x0b, 0x48, 0x6c, 0x11, 0x13, 0x8f, 0x2f, 0x58,
	0x11, 0x48, 0x28, 0x20, 0x29, 0x89, 0x0d, 0xec, 0x6e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf7, 0xb8, 0x11, 0x99, 0xe6, 0x00, 0x00, 0x00,
}
