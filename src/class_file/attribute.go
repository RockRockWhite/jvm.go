package class_file

type AttributeInfo interface {
}

type AttributeType string

const (
	ATTRIBUTE_Code               AttributeType = "Code"
	ATTRIBUTE_ConstantValue      AttributeType = "ConstantValue"
	ATTRIBUTE_Deprecated         AttributeType = "Deprecated"
	ATTRIBUTE_Exceptions         AttributeType = "Exceptions"
	ATTRIBUTE_LineNumberTable    AttributeType = "LineNumberTable"
	ATTRIBUTE_LocalVariableTable AttributeType = "LocalVariableTable"
	ATTRIBUTE_SourceFile         AttributeType = "SourceFile"
	ATTRIBUTE_Synthetic          AttributeType = "Synthetic"
)

type AttributeDeprecated struct {
}

type AttributeSynthetic struct {
}

type AttributeSourceFile struct {
	SourceFile string
}

type AttributeConstantValue struct {
	ConstantValue any
}

type ExpectionTableEntry struct {
	StartPC   uint16
	EndPC     uint16
	HandlerPC uint16
	CatchType uint16
}

type AttributeCode struct {
	MaxStack       uint16
	MaxLocals      uint16
	Code           []byte
	ExpectionTable []ExpectionTableEntry
	Attributes     []AttributeInfo
}

type AttributeExceptions struct {
	Exceptions []string
}

type LineNumberTableEntry struct {
	StartPC    uint16
	LineNumber uint16
}

type AttributeLineNumber struct {
	LineNumberTable []LineNumberTableEntry
}
