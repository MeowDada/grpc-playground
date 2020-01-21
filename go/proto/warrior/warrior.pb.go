// Code generated by protoc-gen-go. DO NOT EDIT.
// source: warrior.proto

// Explicitly declaring the generating package name.

package warrior

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

type Warrior_WeaponType int32

const (
	Warrior_KNIFE Warrior_WeaponType = 0
	Warrior_SWORD Warrior_WeaponType = 1
	Warrior_AXE   Warrior_WeaponType = 2
)

var Warrior_WeaponType_name = map[int32]string{
	0: "KNIFE",
	1: "SWORD",
	2: "AXE",
}

var Warrior_WeaponType_value = map[string]int32{
	"KNIFE": 0,
	"SWORD": 1,
	"AXE":   2,
}

func (x Warrior_WeaponType) String() string {
	return proto.EnumName(Warrior_WeaponType_name, int32(x))
}

func (Warrior_WeaponType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f32876558ada2f2b, []int{0, 0}
}

type Warrior struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Keyword repeated means dynamic array size.
	Weapons              []*Warrior_Weapon `protobuf:"bytes,2,rep,name=weapons,proto3" json:"weapons,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Warrior) Reset()         { *m = Warrior{} }
func (m *Warrior) String() string { return proto.CompactTextString(m) }
func (*Warrior) ProtoMessage()    {}
func (*Warrior) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32876558ada2f2b, []int{0}
}

func (m *Warrior) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Warrior.Unmarshal(m, b)
}
func (m *Warrior) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Warrior.Marshal(b, m, deterministic)
}
func (m *Warrior) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Warrior.Merge(m, src)
}
func (m *Warrior) XXX_Size() int {
	return xxx_messageInfo_Warrior.Size(m)
}
func (m *Warrior) XXX_DiscardUnknown() {
	xxx_messageInfo_Warrior.DiscardUnknown(m)
}

var xxx_messageInfo_Warrior proto.InternalMessageInfo

func (m *Warrior) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Warrior) GetWeapons() []*Warrior_Weapon {
	if m != nil {
		return m.Weapons
	}
	return nil
}

type Warrior_Weapon struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 Warrior_WeaponType `protobuf:"varint,2,opt,name=type,proto3,enum=warrior.Warrior_WeaponType" json:"type,omitempty"`
	Damage               int32              `protobuf:"varint,3,opt,name=damage,proto3" json:"damage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Warrior_Weapon) Reset()         { *m = Warrior_Weapon{} }
func (m *Warrior_Weapon) String() string { return proto.CompactTextString(m) }
func (*Warrior_Weapon) ProtoMessage()    {}
func (*Warrior_Weapon) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32876558ada2f2b, []int{0, 0}
}

func (m *Warrior_Weapon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Warrior_Weapon.Unmarshal(m, b)
}
func (m *Warrior_Weapon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Warrior_Weapon.Marshal(b, m, deterministic)
}
func (m *Warrior_Weapon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Warrior_Weapon.Merge(m, src)
}
func (m *Warrior_Weapon) XXX_Size() int {
	return xxx_messageInfo_Warrior_Weapon.Size(m)
}
func (m *Warrior_Weapon) XXX_DiscardUnknown() {
	xxx_messageInfo_Warrior_Weapon.DiscardUnknown(m)
}

var xxx_messageInfo_Warrior_Weapon proto.InternalMessageInfo

func (m *Warrior_Weapon) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Warrior_Weapon) GetType() Warrior_WeaponType {
	if m != nil {
		return m.Type
	}
	return Warrior_KNIFE
}

func (m *Warrior_Weapon) GetDamage() int32 {
	if m != nil {
		return m.Damage
	}
	return 0
}

type Army struct {
	Warriors             []*Warrior `protobuf:"bytes,1,rep,name=warriors,proto3" json:"warriors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Army) Reset()         { *m = Army{} }
func (m *Army) String() string { return proto.CompactTextString(m) }
func (*Army) ProtoMessage()    {}
func (*Army) Descriptor() ([]byte, []int) {
	return fileDescriptor_f32876558ada2f2b, []int{1}
}

func (m *Army) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Army.Unmarshal(m, b)
}
func (m *Army) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Army.Marshal(b, m, deterministic)
}
func (m *Army) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Army.Merge(m, src)
}
func (m *Army) XXX_Size() int {
	return xxx_messageInfo_Army.Size(m)
}
func (m *Army) XXX_DiscardUnknown() {
	xxx_messageInfo_Army.DiscardUnknown(m)
}

var xxx_messageInfo_Army proto.InternalMessageInfo

func (m *Army) GetWarriors() []*Warrior {
	if m != nil {
		return m.Warriors
	}
	return nil
}

func init() {
	proto.RegisterEnum("warrior.Warrior_WeaponType", Warrior_WeaponType_name, Warrior_WeaponType_value)
	proto.RegisterType((*Warrior)(nil), "warrior.Warrior")
	proto.RegisterType((*Warrior_Weapon)(nil), "warrior.Warrior.Weapon")
	proto.RegisterType((*Army)(nil), "warrior.Army")
}

func init() { proto.RegisterFile("warrior.proto", fileDescriptor_f32876558ada2f2b) }

var fileDescriptor_f32876558ada2f2b = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4f, 0x2c, 0x2a,
	0xca, 0xcc, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0x9e, 0x30,
	0x72, 0xb1, 0x87, 0x43, 0xd8, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x90, 0x21, 0x17, 0x7b, 0x79, 0x6a, 0x62, 0x41, 0x7e, 0x5e,
	0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0xb8, 0x1e, 0xcc, 0xa4, 0x70, 0x18, 0x0d, 0x96,
	0x0f, 0x82, 0xa9, 0x93, 0x4a, 0xe5, 0x62, 0x83, 0x08, 0x61, 0x35, 0x50, 0x9f, 0x8b, 0xa5, 0xa4,
	0xb2, 0x20, 0x55, 0x82, 0x49, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x1a, 0x87, 0x69, 0x21, 0x95, 0x05,
	0xa9, 0x41, 0x60, 0x85, 0x42, 0x62, 0x5c, 0x6c, 0x29, 0x89, 0xb9, 0x89, 0xe9, 0xa9, 0x12, 0xcc,
	0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x50, 0x9e, 0x92, 0x36, 0x17, 0x17, 0x42, 0xad, 0x10, 0x27, 0x17,
	0xab, 0xb7, 0x9f, 0xa7, 0x9b, 0xab, 0x00, 0x03, 0x88, 0x19, 0x1c, 0xee, 0x1f, 0xe4, 0x22, 0xc0,
	0x28, 0xc4, 0xce, 0xc5, 0xec, 0x18, 0xe1, 0x2a, 0xc0, 0xa4, 0x64, 0xc2, 0xc5, 0xe2, 0x58, 0x94,
	0x5b, 0x29, 0xa4, 0xc3, 0xc5, 0x01, 0xb5, 0xb0, 0x58, 0x82, 0x11, 0xec, 0x1f, 0x01, 0x74, 0x17,
	0x04, 0xc1, 0x55, 0x24, 0xb1, 0x81, 0x03, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x9a,
	0xd8, 0x89, 0x3d, 0x01, 0x00, 0x00,
}