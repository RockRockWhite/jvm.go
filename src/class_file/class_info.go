package class_file

type ConstantPoolInfo struct {
}

type InterfaceInfo struct {
}

type FieldInfo struct {
}

type MethodInfo struct {
}

type AttributeInfo struct {
}

type ClassInfo struct {
	Maigc        uint32
	MinorVersion uint16
	MajorVersion uint16
	ConstantPool []ConstantPoolInfo
	AccessFlags  uint16
	ThisClass    uint16
	SuperClass   uint16
	Interfaces   []InterfaceInfo
	Fields       []FieldInfo
	Methods      []MethodInfo
	Attrubutes   []AttributeInfo
}
