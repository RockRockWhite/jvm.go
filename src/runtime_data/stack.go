package runtime_data

type Frame struct {
	LocalVariables LocalVariables
	OperandStack   OperandStack
}

type Stack struct {
	MaxSize int
	Size    int
	Frames  []Frame
}
