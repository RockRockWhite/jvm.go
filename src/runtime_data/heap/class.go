package heap

import "github.com/RockRockWhite/jvm.go/src/runtime_data"

type Field struct {
}

type Method struct {
}

type Class struct {
	AccessFlags      uint16
	Name             string
	SuperClassName   string
	InterfaceNames   []string
	ConstantPool     *ConstantPool
	Fields           []*Field
	Methods          []*Method
	ClassLoader      *ClassLoader
	SuperClass       *Class
	Interfaces       *Class
	InstantSlotCount uint
	StaticSlotCount  uint
	StaticVars       []runtime_data.VariableSlot
}
