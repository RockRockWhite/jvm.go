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
	Maigc           uint32
	MinorVersion    uint16
	MajorVersion    uint16
	ConstantPool    []ConstantPoolInfo
	AccessFlags     uint16
	ThisClassIndex  uint16
	SuperClassIndex uint16
	InterfaceIndexs []InterfaceInfo
	Fields          []FieldInfo
	Methods         []MethodInfo
	Attrubutes      []AttributeInfo
}
