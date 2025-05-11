package runtime_data

type LocalVariableSlot struct {
	Data any
}

type LocalVariables struct {
	Slots []LocalVariableSlot
}

type OperandStack struct {
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
