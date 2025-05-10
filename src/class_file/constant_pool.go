package class_file

import "fmt"

type ConstantType uint8

const (
	CONSTANT_Unexpected         ConstantType = 0
	CONSTANT_Class              ConstantType = 7
	CONSTANT_Fieldref           ConstantType = 9
	CONSTANT_Methodref          ConstantType = 10
	CONSTANT_InferfaceMethodref ConstantType = 11
	CONSTANT_String             ConstantType = 8
	CONSTANT_Integer            ConstantType = 3
	CONSTANT_Float              ConstantType = 4
	CONSTANT_Long               ConstantType = 5
	CONSTANT_Double             ConstantType = 6
	CONSTANT_NameAndType        ConstantType = 12
	CONSTANT_Utf8               ConstantType = 1
	CONSTANT_MethodHandle       ConstantType = 15
	CONSTANT_MethodType         ConstantType = 16
	CONSTANT_InvokeDynamic      ConstantType = 18
)

type ConstantInfo interface {
}

type ConstantInt struct {
	Value int32
}

type ConstantFloat struct {
	Value float32
}

type ConstantLong struct {
	Value int64
}

type ConstantDouble struct {
	Value float64
}

type ConstantUtf8 struct {
	Str string
}

type ConstantString struct {
	EntryIndex uint16
}

func (cs *ConstantString) GetString(constant_pool ConstantPoolInfo) string {
	constant_utf8, _ := constant_pool.Entries[cs.EntryIndex].(ConstantUtf8)
	return constant_utf8.Str
}

type ConstantClass struct {
	EntryIndex uint16
}

// will be deprecated
func (cc *ConstantClass) GetClassName(constant_pool ConstantPoolInfo) string {
	constant_utf8, _ := constant_pool.Entries[cc.EntryIndex].(ConstantUtf8)
	return constant_utf8.Str
}

type ConstantNameAndType struct {
	NameEntryIndex       uint16
	DescriptorEntryIndex uint16
}

func (c *ConstantNameAndType) GetName(constant_pool ConstantPoolInfo) string {
	name_utf8, _ := constant_pool.Entries[c.NameEntryIndex].(ConstantUtf8)

	return name_utf8.Str
}

func (c *ConstantNameAndType) GetDescriptor(constant_pool ConstantPoolInfo) string {
	descriptor_utf8, _ := constant_pool.Entries[c.DescriptorEntryIndex].(ConstantUtf8)

	return descriptor_utf8.Str
}

type ConstantMemberRef struct {
	ClassEntryIndex       uint16
	NameAndTypeEntryIndex uint16
}

type ConstantFieldRef struct {
	MemberRef ConstantMemberRef
}

func NewConstantFieldRef(constant_member_ref ConstantMemberRef) ConstantFieldRef {
	return ConstantFieldRef{
		MemberRef: constant_member_ref,
	}
}

type ConstantMethodRef struct {
	MemberRef ConstantMemberRef
}

func NewConstantMethodRef(constant_member_ref ConstantMemberRef) ConstantMethodRef {
	return ConstantMethodRef{
		MemberRef: constant_member_ref,
	}
}

type ConstantInterfaceMethodRef struct {
	MemberRef ConstantMemberRef
}

func NewConstantInterfaceMethodRef(constant_member_ref ConstantMemberRef) ConstantInterfaceMethodRef {
	return ConstantInterfaceMethodRef{
		MemberRef: constant_member_ref,
	}
}

type ConstantPoolInfo struct {
	Entries []ConstantInfo
}

func (cp *ConstantPoolInfo) GetType(idx uint16) ConstantType {
	switch cp.Entries[idx].(type) {
	case ConstantClass:
		return CONSTANT_Class
	case ConstantFieldRef:
		return CONSTANT_Fieldref
	case ConstantMethodRef:
		return CONSTANT_Methodref
	case ConstantInterfaceMethodRef:
		return CONSTANT_InferfaceMethodref
	case ConstantString:
		return CONSTANT_String
	case ConstantInt:
		return CONSTANT_Integer
	case ConstantFloat:
		return CONSTANT_Float
	case ConstantLong:
		return CONSTANT_Long
	case ConstantDouble:
		return CONSTANT_Double
	case ConstantNameAndType:
		return CONSTANT_NameAndType
	case ConstantUtf8:
		return CONSTANT_Utf8
	default:
		return CONSTANT_Unexpected
	}
}

func (cp *ConstantPoolInfo) GetUtf8(idx uint16) (string, error) {
	constant_utf8, ok := cp.Entries[idx].(ConstantUtf8)
	if !ok {
		return "", fmt.Errorf("index %d in is not utf8", idx)
	}

	return constant_utf8.Str, nil
}

func (cp *ConstantPoolInfo) GetString(idx uint16) (string, error) {
	constant_str, ok := cp.Entries[idx].(ConstantString)
	if !ok {
		return "", fmt.Errorf("index %d in is not utf8", idx)
	}

	return cp.GetUtf8(constant_str.EntryIndex)
}

func (cp *ConstantPoolInfo) GetInt(idx uint16) (int32, error) {
	constant_value, ok := cp.Entries[idx].(ConstantInt)
	if !ok {
		return 0, fmt.Errorf("index %d in is not int", idx)
	}

	return constant_value.Value, nil
}

func (cp *ConstantPoolInfo) GetLong(idx uint16) (int64, error) {
	constant_value, ok := cp.Entries[idx].(ConstantLong)
	if !ok {
		return 0, fmt.Errorf("index %d in is not long", idx)
	}

	return constant_value.Value, nil
}

func (cp *ConstantPoolInfo) GetFloat(idx uint16) (float32, error) {
	constant_value, ok := cp.Entries[idx].(ConstantFloat)
	if !ok {
		return 0, fmt.Errorf("index %d in is not long", idx)
	}

	return constant_value.Value, nil
}

func (cp *ConstantPoolInfo) GetDouble(idx uint16) (float64, error) {
	constant_value, ok := cp.Entries[idx].(ConstantDouble)
	if !ok {
		return 0, fmt.Errorf("index %d in is not long", idx)
	}

	return constant_value.Value, nil
}

func (cp *ConstantPoolInfo) GetClass(idx uint16) (string, error) {
	constant_class, ok := cp.Entries[idx].(ConstantClass)
	if !ok {
		return "", fmt.Errorf("index %d in is not class", idx)
	}

	class_name, err := cp.GetUtf8(constant_class.EntryIndex)
	return class_name, err
}
