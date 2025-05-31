package runtime_data

import "fmt"

type VariableSlot struct {
	Data any
}

type LocalVariables struct {
	Slots []VariableSlot
}

func (vars *LocalVariables) PushInt(data int32) {
	vars.Slots = append(vars.Slots, VariableSlot{
		Data: &data,
	})
}

func (vars *LocalVariables) GetInt(index uint16) int32 {
	data_ptr, ok := vars.Slots[index].Data.(*int32)
	if !ok {
		// just panic
		panic(fmt.Sprintf("Local variable at index %d is not an int", index))
	}

	return *data_ptr
}

func (vars *LocalVariables) SetInt(index uint16, data int32) {
	if index >= uint16(len(vars.Slots)) {
		// just panic
		panic(fmt.Sprintf("Local variable index %d out of bounds", index))
	}

	data_ptr, ok := vars.Slots[index].Data.(*int32)
	if !ok {
		// just panic
		panic(fmt.Sprintf("Local variable at index %d is not an int", index))
	}

	*data_ptr = data
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
