package instruction

import (
	"github.com/RockRockWhite/jvm.go/src/runtime_data"
)

type NopInst struct {
	NoOperandsInstruction
}

func (inst *NopInst) Execute(frame *runtime_data.Frame) error {
	// NOP does nothing, so we simply return nil.
	return nil
}


type IConstInst struct {
	NoOperandsInstruction
}

func (inst *IConstInst) PushValue(frame *runtime_data.Frame, value int32) error {
	// Push the integer value onto the operand stack.
	frame.OperandStack.Slots[frame.OperandStack.Size].Data = value
	frame.OperandStack.Size++
	return nil
}
