package class_file

type InterfaceInfo struct {
	NameIndex uint16
}

type MemberInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Attrubutes      []AttributeInfo
}

func (m *MemberInfo) GetName() string {
	return ""
}

func (m *MemberInfo) GetDescriptor() string {
	return ""
}

type FieldInfo MemberInfo
type MethodInfo MemberInfo

type AttributeInfo struct {
}

type ClassInfo struct {
	Maigc           uint32
	MinorVersion    uint16
	MajorVersion    uint16
	ConstantPool    ConstantPoolInfo
	AccessFlags     uint16
	ThisClassIndex  uint16
	SuperClassIndex uint16
	Interfaces      []InterfaceInfo
	Fields          []FieldInfo
	Methods         []MethodInfo
	Attrubutes      []AttributeInfo
}
