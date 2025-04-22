package class_file

type ConstantType uint8

const (
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
