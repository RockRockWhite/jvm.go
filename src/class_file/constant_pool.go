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
	value int32
}

type ConstantFloat struct {
	value float32
}

type ConstantLong struct {
	value int64
}

type ConstantDouble struct {
	value float64
}

type ConstantUtf8 struct {
	str string
}

type ConstantString struct {
	EntryIndex uint16
}

func (cs *ConstantString) GetString(constant_pool ConstantPoolInfo) string {
	constant_utf8, _ := constant_pool.Entries[cs.EntryIndex].(*ConstantUtf8)
	return constant_utf8.str
}

type ConstantClass struct {
	EntryIndex uint16
}

func (cc *ConstantClass) GetClassName(constant_pool ConstantPoolInfo) string {
	constant_utf8, _ := constant_pool.Entries[cc.EntryIndex].(*ConstantUtf8)
	return constant_utf8.str
}

type ConstantNameAndType struct {
	NameEntryIndex       uint16
	DescriptorEntryIndex uint16
}

func (c *ConstantNameAndType) GetName(constant_pool ConstantPoolInfo) string {
	name_utf8, _ := constant_pool.Entries[c.NameEntryIndex].(*ConstantUtf8)

	return name_utf8.str
}

func (c *ConstantNameAndType) GetDescriptor(constant_pool ConstantPoolInfo) string {
	descriptor_utf8, _ := constant_pool.Entries[c.DescriptorEntryIndex].(*ConstantUtf8)

	return descriptor_utf8.str
}

type ConstantPoolInfo struct {
	Entries []ConstantInfo
}
