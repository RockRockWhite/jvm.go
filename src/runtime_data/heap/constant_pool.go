package heap

type ConstantPool struct {
}

type Constant struct {
}

type SymbolRef struct {
	ConstantPool *ConstantPool
	ClassName    string
	Class        *Class
}
