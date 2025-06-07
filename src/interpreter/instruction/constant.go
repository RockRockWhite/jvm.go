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
	value int32
	NoOperandsInstruction
}

func (inst *IConstInst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.PushInt(inst.value)

	return nil
}

type LConstInst struct {
	value int64
	NoOperandsInstruction
}

func (inst *LConstInst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.PushLong(inst.value)

	return nil
}

type FConstInst struct {
	value float32
	NoOperandsInstruction
}

func (inst *FConstInst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.PushFloat(inst.value)

	return nil
}

type DConstInst struct {
	value float64
	NoOperandsInstruction
}

func (inst *DConstInst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.PushDouble(inst.value)

	return nil
}

type AConstInst struct {
	value *runtime_data.Object
	NoOperandsInstruction
}

func (inst *AConstInst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.PushAddress(inst.value)

	return nil
}

type BiPushInst struct {
	ByteOprandInstruction
}

func (inst *BiPushInst) Execute(frame *runtime_data.Frame) error {
	// Sign-extend the byte operand to
	val := int32(inst.Operand)
	frame.OperandStack.PushInt(val)

	return nil
}

type SiPushInst struct {
	DByteOprandInstruction
}

func (inst *SiPushInst) Execute(frame *runtime_data.Frame) error {
	// Sign-extend the short operand to 32 bits.
	val := int32(inst.Index)
	frame.OperandStack.PushInt(val)

	return nil
}
