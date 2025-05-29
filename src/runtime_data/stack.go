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

func NewOperandStack(size uint) OperandStack {
	return OperandStack{
		Size:  size,
		Slots: make([]VariableSlot, 0, size),
	}
}

func (stack *OperandStack) Push(data any) {
	if uint(len(stack.Slots)) >= stack.Size {
		// just panic
		panic("Operand stack overflow!")
	}

	stack.Slots = append(stack.Slots, VariableSlot{
		Data: data,
	})
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
