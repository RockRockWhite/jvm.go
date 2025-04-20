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
