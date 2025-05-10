package class_file

type InterfaceInfo struct {
	NameIndex uint16
}

type MemberInfo struct {
	AccessFlags uint16
	Name        string
	Descriptor  string
	Attrubutes  []AttributeInfo
}

func (m *MemberInfo) GetName() string {
	return ""
}

func (m *MemberInfo) GetDescriptor() string {
	return ""
}

type FieldInfo MemberInfo
type MethodInfo MemberInfo

type ClassInfo struct {
	Maigc        uint32
	MinorVersion uint16
	MajorVersion uint16
	ConstantPool ConstantPoolInfo
	AccessFlags  uint16
	ThisClass    string
	SuperClass   string
	Interfaces   []string
	Fields       []FieldInfo
	Methods      []MethodInfo
	Attributes   []AttributeInfo
}
