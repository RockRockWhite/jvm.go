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

type IConstInst0 struct {
	NoOperandsInstruction
}

func (inst *IConstInst0) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.Push(int32(0))

	return nil
}

type ConstInst struct {
	value any
	NoOperandsInstruction
}

func (inst *ConstInst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.Push(inst.value)

	return nil
}

type BiPushInst struct {
	ByteOprandInstruction
}

func (inst *BiPushInst) Execute(frame *runtime_data.Frame) error {
	// Sign-extend the byte operand to
	val := int32(inst.Operand)
	frame.OperandStack.Push(val)

	return nil
}

type SiPushInst struct {
	DByteOprandInstruction
}

func (inst *SiPushInst) Execute(frame *runtime_data.Frame) error {
	// Sign-extend the short operand to 32 bits.
	val := int32(inst.Index)
	frame.OperandStack.Push(val)
	
	return nil
}
