// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gomatcha.io/matcha/pb/view/ios/stackview.proto

package ios

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import matcha "gomatcha.io/matcha/pb"
import matcha_text "gomatcha.io/matcha/pb/text"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type StackChildView struct {
	ScreenId int64 `protobuf:"varint,3,opt,name=screenId" json:"screenId,omitempty"`
}

func (m *StackChildView) Reset()                    { *m = StackChildView{} }
func (m *StackChildView) String() string            { return proto.CompactTextString(m) }
func (*StackChildView) ProtoMessage()               {}
func (*StackChildView) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *StackChildView) GetScreenId() int64 {
	if m != nil {
		return m.ScreenId
	}
	return 0
}

type StackView struct {
	Children       []*StackChildView      `protobuf:"bytes,1,rep,name=children" json:"children,omitempty"`
	TitleTextStyle *matcha_text.TextStyle `protobuf:"bytes,2,opt,name=titleTextStyle" json:"titleTextStyle,omitempty"`
	BackTextStyle  *matcha_text.TextStyle `protobuf:"bytes,3,opt,name=backTextStyle" json:"backTextStyle,omitempty"`
	BarColor       *matcha.Color          `protobuf:"bytes,4,opt,name=barColor" json:"barColor,omitempty"`
}

func (m *StackView) Reset()                    { *m = StackView{} }
func (m *StackView) String() string            { return proto.CompactTextString(m) }
func (*StackView) ProtoMessage()               {}
func (*StackView) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *StackView) GetChildren() []*StackChildView {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *StackView) GetTitleTextStyle() *matcha_text.TextStyle {
	if m != nil {
		return m.TitleTextStyle
	}
	return nil
}

func (m *StackView) GetBackTextStyle() *matcha_text.TextStyle {
	if m != nil {
		return m.BackTextStyle
	}
	return nil
}

func (m *StackView) GetBarColor() *matcha.Color {
	if m != nil {
		return m.BarColor
	}
	return nil
}

type StackBar struct {
	Title                 string `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	BackButtonHidden      bool   `protobuf:"varint,7,opt,name=backButtonHidden" json:"backButtonHidden,omitempty"`
	CustomBackButtonTitle bool   `protobuf:"varint,2,opt,name=customBackButtonTitle" json:"customBackButtonTitle,omitempty"`
	BackButtonTitle       string `protobuf:"bytes,3,opt,name=backButtonTitle" json:"backButtonTitle,omitempty"`
	HasTitleView          bool   `protobuf:"varint,4,opt,name=hasTitleView" json:"hasTitleView,omitempty"`
	RightViewCount        int64  `protobuf:"varint,5,opt,name=rightViewCount" json:"rightViewCount,omitempty"`
	LeftViewCount         int64  `protobuf:"varint,6,opt,name=leftViewCount" json:"leftViewCount,omitempty"`
}

func (m *StackBar) Reset()                    { *m = StackBar{} }
func (m *StackBar) String() string            { return proto.CompactTextString(m) }
func (*StackBar) ProtoMessage()               {}
func (*StackBar) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *StackBar) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *StackBar) GetBackButtonHidden() bool {
	if m != nil {
		return m.BackButtonHidden
	}
	return false
}

func (m *StackBar) GetCustomBackButtonTitle() bool {
	if m != nil {
		return m.CustomBackButtonTitle
	}
	return false
}

func (m *StackBar) GetBackButtonTitle() string {
	if m != nil {
		return m.BackButtonTitle
	}
	return ""
}

func (m *StackBar) GetHasTitleView() bool {
	if m != nil {
		return m.HasTitleView
	}
	return false
}

func (m *StackBar) GetRightViewCount() int64 {
	if m != nil {
		return m.RightViewCount
	}
	return 0
}

func (m *StackBar) GetLeftViewCount() int64 {
	if m != nil {
		return m.LeftViewCount
	}
	return 0
}

type StackEvent struct {
	Id []int64 `protobuf:"varint,1,rep,packed,name=id" json:"id,omitempty"`
}

func (m *StackEvent) Reset()                    { *m = StackEvent{} }
func (m *StackEvent) String() string            { return proto.CompactTextString(m) }
func (*StackEvent) ProtoMessage()               {}
func (*StackEvent) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *StackEvent) GetId() []int64 {
	if m != nil {
		return m.Id
	}
	return nil
}

func init() {
	proto.RegisterType((*StackChildView)(nil), "matcha.view.ios.StackChildView")
	proto.RegisterType((*StackView)(nil), "matcha.view.ios.StackView")
	proto.RegisterType((*StackBar)(nil), "matcha.view.ios.StackBar")
	proto.RegisterType((*StackEvent)(nil), "matcha.view.ios.StackEvent")
}

func init() { proto.RegisterFile("gomatcha.io/matcha/pb/view/ios/stackview.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 424 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0xcf, 0x93, 0x40,
	0x10, 0xc6, 0x03, 0xf8, 0xbe, 0xd2, 0xc1, 0x52, 0xb3, 0x51, 0x43, 0x1a, 0xa3, 0x48, 0xd4, 0xa0,
	0x31, 0x90, 0x54, 0x6f, 0x1a, 0x63, 0x68, 0x4c, 0xf4, 0x60, 0x6c, 0xb6, 0x8d, 0x07, 0x6f, 0x0b,
	0xac, 0x65, 0x53, 0xca, 0x36, 0xb0, 0xfd, 0xe3, 0x97, 0xf1, 0xe0, 0x37, 0xf4, 0x1b, 0x18, 0x86,
	0x96, 0x0a, 0x36, 0xef, 0x05, 0x76, 0x9e, 0xf9, 0x3d, 0xc3, 0xec, 0x30, 0x10, 0x2c, 0xe5, 0x9a,
	0xa9, 0x24, 0x63, 0x81, 0x90, 0x61, 0x73, 0x0a, 0x37, 0x71, 0xb8, 0x13, 0x7c, 0x1f, 0x0a, 0x59,
	0x85, 0x95, 0x62, 0xc9, 0xaa, 0x8e, 0x82, 0x4d, 0x29, 0x95, 0x24, 0xa3, 0x23, 0x8d, 0x92, 0x90,
	0xd5, 0xf8, 0xc9, 0xe5, 0x02, 0x62, 0xcd, 0x96, 0xbc, 0xf1, 0x8c, 0x9f, 0x5d, 0x46, 0x14, 0x3f,
	0x28, 0x7c, 0x34, 0x98, 0xf7, 0x0a, 0xec, 0x79, 0xfd, 0xb5, 0x69, 0x26, 0xf2, 0xf4, 0x9b, 0xe0,
	0x7b, 0x32, 0x06, 0xb3, 0x4a, 0x4a, 0xce, 0x8b, 0xcf, 0xa9, 0x63, 0xb8, 0x9a, 0x6f, 0xd0, 0x36,
	0xf6, 0xfe, 0x68, 0x30, 0x40, 0x1c, 0xc9, 0xb7, 0x60, 0x26, 0xb5, 0xad, 0xe4, 0x85, 0xa3, 0xb9,
	0x86, 0x6f, 0x4d, 0x1e, 0x07, 0xbd, 0x4e, 0x83, 0x6e, 0x71, 0xda, 0x1a, 0xc8, 0x7b, 0xb0, 0x95,
	0x50, 0x39, 0x5f, 0xf0, 0x83, 0x9a, 0xab, 0x9f, 0x39, 0x77, 0x74, 0x57, 0xf3, 0xad, 0xc9, 0x83,
	0x53, 0x09, 0x6c, 0xb2, 0xcd, 0xd2, 0x1e, 0x4d, 0xde, 0xc1, 0x30, 0x66, 0xc9, 0xea, 0x6c, 0x37,
	0x6e, 0xb4, 0x77, 0x61, 0xf2, 0x02, 0xcc, 0x98, 0x95, 0x53, 0x99, 0xcb, 0xd2, 0xb9, 0x85, 0xc6,
	0xe1, 0xc9, 0x88, 0x22, 0x6d, 0xd3, 0xde, 0x2f, 0x1d, 0x4c, 0xbc, 0x45, 0xc4, 0x4a, 0x72, 0x0f,
	0xae, 0xb0, 0x0f, 0x47, 0x73, 0x35, 0x7f, 0x40, 0x9b, 0x80, 0xbc, 0x84, 0xbb, 0x75, 0xf9, 0x68,
	0xab, 0x94, 0x2c, 0x3e, 0x89, 0x34, 0xe5, 0x85, 0x73, 0xdb, 0xd5, 0x7c, 0x93, 0xfe, 0xa7, 0x93,
	0x37, 0x70, 0x3f, 0xd9, 0x56, 0x4a, 0xae, 0xa3, 0x36, 0xb3, 0xc0, 0x8a, 0x3a, 0x1a, 0x2e, 0x27,
	0x89, 0x0f, 0xa3, 0xb8, 0xc7, 0x1b, 0xd8, 0x41, 0x5f, 0x26, 0x1e, 0xdc, 0xc9, 0x58, 0x85, 0xe7,
	0x7a, 0xe2, 0x78, 0x3b, 0x93, 0x76, 0x34, 0xf2, 0x1c, 0xec, 0x52, 0x2c, 0x33, 0x55, 0x07, 0x53,
	0xb9, 0x2d, 0x94, 0x73, 0x85, 0x3f, 0xba, 0xa7, 0x92, 0xa7, 0x30, 0xcc, 0xf9, 0x8f, 0x7f, 0xb0,
	0x6b, 0xc4, 0xba, 0xa2, 0xf7, 0x10, 0x00, 0xe7, 0xf3, 0x71, 0xc7, 0x0b, 0x45, 0x6c, 0xd0, 0x45,
	0x8a, 0xeb, 0x60, 0x50, 0x5d, 0xa4, 0xd1, 0x07, 0x78, 0x24, 0xe4, 0x79, 0xe1, 0x8f, 0xaf, 0x4d,
	0xdc, 0xae, 0x48, 0x64, 0xcd, 0xe2, 0x76, 0xa7, 0xbe, 0x1b, 0x42, 0x56, 0xbf, 0x75, 0xeb, 0x0b,
	0x62, 0xe2, 0xeb, 0x7c, 0x16, 0xc5, 0xd7, 0xb8, 0xa9, 0xaf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x44, 0x4a, 0x74, 0x58, 0x36, 0x03, 0x00, 0x00,
}
