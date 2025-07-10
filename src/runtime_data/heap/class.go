package heap

type Field struct {
}

type Method struct {
}

type Class struct {
	AccessFlags    uint16
	Name           string
	SuperClassName string
	InterfaceNames []string
	ConstantPool   *ConstantPool
	Fields         []*Field
	Methods        []*Method
	ClassLoader    *ClassLoader
	SuperClass     *Class
	Interfaces     *Class
}
