package heap

import "github.com/RockRockWhite/jvm.go/src/runtime_data"

type Field struct {
}

type Method struct {
}

type AccessFlag uint16

const (
	AccessFlagPublic       AccessFlag = 1
	AccessFlagPrivate      AccessFlag = 1 << 1
	AccessFlagProtected    AccessFlag = 1 << 2
	AccessFlagStatic       AccessFlag = 1 << 3
	AccessFlagFinal        AccessFlag = 1 << 4
	AccessFlagSynchronized AccessFlag = 1 << 5
	AccessFlagVolatile     AccessFlag = 1 << 6
	AccessFlagBridge       AccessFlag = 1 << 6
	AccessFlagTransient    AccessFlag = 1 << 7
	AccessFlagVarargs      AccessFlag = 1 << 7
	AccessFlagNative       AccessFlag = 1 << 8
	AccessFlagInterface    AccessFlag = 1 << 9
	AccessFlagAbstract     AccessFlag = 1 << 10
	AccessFlagStrict       AccessFlag = 1 << 11
	AccessFlagSynthetic    AccessFlag = 1 << 12
	AccessFlagAnnotation   AccessFlag = 1 << 13
	AccessFlagEnum         AccessFlag = 1 << 14
)

type Class struct {
	AccessFlags      AccessFlag
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
