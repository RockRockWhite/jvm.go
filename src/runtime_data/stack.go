package runtime_data

type VariableSlot struct {
	Data any
}

type LocalVariables struct {
	Slots []VariableSlot
}

type OperandStack struct {
	Size  uint
	Slots []VariableSlot
}

type Frame struct {
	LocalVariables LocalVariables
	OperandStack   OperandStack
}

type Stack struct {
	MaxSize int
	Size    int
	Frames  []Frame
}
