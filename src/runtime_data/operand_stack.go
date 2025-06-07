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
	// Long occupies two slots in the operand stack
	stack.push(VariableSlot{
		Data: nil,
	})
}

func (stack *OperandStack) PopLong() int64 {
	top := stack.Pop()
	data_ptr, ok := top.Data.(*int64)
	if !ok {
		// just panic
		panic("Top of operand stack is not an long")
	}

	return *data_ptr
}

func (stack *OperandStack) PushFloat(data float32) {
	stack.push(VariableSlot{
		Data: &data,
	})
}

func (stack *OperandStack) PopFloat() float32 {
	top := stack.Pop()
	data_ptr, ok := top.Data.(*float32)
	if !ok {
		// just panic
		panic("Top of operand stack is not a float")
	}

	return *data_ptr
}

func (stack *OperandStack) PushDouble(data float64) {
	stack.push(VariableSlot{
		Data: &data,
	})
	// Double occupies two slots in the operand stack
	stack.push(VariableSlot{
		Data: nil,
	})
}

func (stack *OperandStack) PopDouble() float64 {
	top := stack.Pop()
	data_ptr, ok := top.Data.(*float64)
	if !ok {
		// just panic
		panic("Top of operand stack is not a double")
	}

	return *data_ptr
}

func (stack *OperandStack) PushAddress(data *Object) {
	stack.push(VariableSlot{
		Data: data,
	})
}

func (stack *OperandStack) PopAddress() *Object {
	top := stack.Pop()
	data_ptr, ok := top.Data.(*Object)
	if !ok {
		// just panic
		panic("Top of operand stack is not an address")
	}

	return data_ptr
}

func (stack *OperandStack) Dup() {
	top := stack.Slots[len(stack.Slots)-1]
	stack.Slots = append(stack.Slots, top)
}

func (stack *OperandStack) DupX1() {
	top := stack.Slots[len(stack.Slots)-1]
	rear := stack.Slots[len(stack.Slots)-2:]

	stack.Slots = append(stack.Slots[:len(stack.Slots)-2], top)
	stack.Slots = append(stack.Slots, rear...)
}

func (stack *OperandStack) DupX2() {
	top := stack.Slots[len(stack.Slots)-1]
	rear := stack.Slots[len(stack.Slots)-3:]

	stack.Slots = append(stack.Slots[:len(stack.Slots)-3], top)
	stack.Slots = append(stack.Slots, rear...)
}

func (stack *OperandStack) Dup2() {
	top1 := stack.Slots[len(stack.Slots)-1]
	top2 := stack.Slots[len(stack.Slots)-2]

	stack.Slots = append(stack.Slots, top2, top1)
}

func (stack *OperandStack) Dup2X1() {
	top1 := stack.Slots[len(stack.Slots)-1]
	top2 := stack.Slots[len(stack.Slots)-2]
	rear := stack.Slots[len(stack.Slots)-3:]

	stack.Slots = append(stack.Slots[:len(stack.Slots)-3], top2, top1)
	stack.Slots = append(stack.Slots, rear...)
}

func (stack *OperandStack) Dup2X2() {
	top1 := stack.Slots[len(stack.Slots)-1]
	top2 := stack.Slots[len(stack.Slots)-2]
	rear := stack.Slots[len(stack.Slots)-4:]

	stack.Slots = append(stack.Slots[:len(stack.Slots)-4], top2, top1)
	stack.Slots = append(stack.Slots, rear...)
}

func (stack *OperandStack) Swap() {
	if len(stack.Slots) < 2 {
		// just panic
		panic("Not enough elements in operand stack to swap")
	}

	top1 := stack.Slots[len(stack.Slots)-1]
	top2 := stack.Slots[len(stack.Slots)-2]

	stack.Slots[len(stack.Slots)-1] = top2
	stack.Slots[len(stack.Slots)-2] = top1
}
