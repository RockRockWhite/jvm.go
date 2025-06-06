package runtime_data

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

func (stack *OperandStack) push(slot VariableSlot) {
	if uint(len(stack.Slots)) >= stack.Size {
		// just panic
		panic("Operand stack overflow!")
	}

	stack.Slots = append(stack.Slots, slot)
}

func (stack *OperandStack) top() *VariableSlot {
	if uint(len(stack.Slots)) == 0 {
		// just panic
		panic("Operand stack overflow!")
	}

	top := stack.Slots[len(stack.Slots)-1]
	return &top
}

func (stack *OperandStack) Pop() *VariableSlot {
	top := stack.top()
	stack.Slots = stack.Slots[:len(stack.Slots)-1]

	return top
}

func (stack *OperandStack) PushInt(data int32) {
	stack.push(VariableSlot{
		Data: &data,
	})
}

func (stack *OperandStack) PopInt() int32 {
	top := stack.Pop()
	data_ptr, ok := top.Data.(*int32)
	if !ok {
		// just panic
		panic("Top of operand stack is not an int")
	}

	return *data_ptr
}

func (stack *OperandStack) PushLong(data int64) {
	stack.push(VariableSlot{
		Data: &data,
	})
}
