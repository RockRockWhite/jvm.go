package heap

import (
	"github.com/RockRockWhite/jvm.go/src/class_file"
	"github.com/RockRockWhite/jvm.go/src/runtime_data"
)

type Field struct {
}

type Method struct {
}

type AccessFlags uint16

const (
	AccessFlagPublic       AccessFlags = 1
	AccessFlagPrivate      AccessFlags = 1 << 1
	AccessFlagProtected    AccessFlags = 1 << 2
	AccessFlagStatic       AccessFlags = 1 << 3
	AccessFlagFinal        AccessFlags = 1 << 4
	AccessFlagSynchronized AccessFlags = 1 << 5
	AccessFlagVolatile     AccessFlags = 1 << 6
	AccessFlagBridge       AccessFlags = 1 << 6
	AccessFlagTransient    AccessFlags = 1 << 7
	AccessFlagVarargs      AccessFlags = 1 << 7
	AccessFlagNative       AccessFlags = 1 << 8
	AccessFlagInterface    AccessFlags = 1 << 9
	AccessFlagAbstract     AccessFlags = 1 << 10
	AccessFlagStrict       AccessFlags = 1 << 11
	AccessFlagSynthetic    AccessFlags = 1 << 12
	AccessFlagAnnotation   AccessFlags = 1 << 13
	AccessFlagEnum         AccessFlags = 1 << 14
)

type Class struct {
	AccessFlags      AccessFlags
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

func NewClass(cf *class_file.ClassFile) *Class {
	return &Class{
		AccessFlags:      AccessFlags(cf.AccessFlags),
		Name:             cf.ThisClass,
		SuperClassName:   cf.SuperClass,
		InterfaceNames:   cf.Interfaces,
		ConstantPool:     &ConstantPool{},
		Fields:           []*Field{},
		Methods:          []*Method{},
		ClassLoader:      &ClassLoader{},
		SuperClass:       &Class{},
		Interfaces:       &Class{},
		InstantSlotCount: 0,
		StaticSlotCount:  0,
		StaticVars:       []runtime_data.VariableSlot{},
	}
}

type ClassMember struct {
	AccessFlags AccessFlags
	Name        string
	Descriptor  string
	MemberClass *Class
}
