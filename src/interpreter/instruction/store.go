package instruction

import "github.com/RockRockWhite/jvm.go/src/runtime_data"

type IStoreInstBase struct {
}

func (inst *IStoreInstBase) execute(frame *runtime_data.Frame, idx uint16) error {
	local_variable := frame.OperandStack.PopInt()

	frame.LocalVariables.SetInt(idx, local_variable)
	return nil
}

type IStoreInst struct {
	IStoreInstBase
	ByteOprandInstruction
}

func (inst *IStoreInst) Execute(frame *runtime_data.Frame) error {
	idx := uint16(inst.Operand)
	return inst.execute(frame, idx)
}

type IStoreNInst struct {
	idx uint16
	IStoreInstBase
	NoOperandsInstruction
}

func (inst *IStoreNInst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, inst.idx)
}

type LStoreInstBase struct {
}

func (inst *LStoreInstBase) execute(frame *runtime_data.Frame, idx uint16) error {
	local_variable := frame.OperandStack.PopLong()

	frame.LocalVariables.SetLong(idx, local_variable)
	return nil
}

type LStoreInst struct {
	LStoreInstBase
	ByteOprandInstruction
}

func (inst *LStoreInst) Execute(frame *runtime_data.Frame) error {
	idx := uint16(inst.Operand)
	return inst.execute(frame, idx)
}

type LStoreNInst struct {
	idx uint16
	LStoreInstBase
	NoOperandsInstruction
}

func (inst *LStoreNInst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, inst.idx)
}

type FStoreInstBase struct {
}

func (inst *FStoreInstBase) execute(frame *runtime_data.Frame, idx uint16) error {
	local_variable := frame.OperandStack.PopFloat()

	frame.LocalVariables.SetFloat(idx, local_variable)
	return nil
}

type FStoreInst struct {
	FStoreInstBase
	ByteOprandInstruction
}

func (inst *FStoreInst) Execute(frame *runtime_data.Frame) error {
	idx := uint16(inst.Operand)
	return inst.execute(frame, idx)
}

type FStoreNInst struct {
	idx uint16
	FStoreInstBase
	NoOperandsInstruction
}

func (inst *FStoreNInst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, inst.idx)
}

type DStoreInstBase struct {
}

func (inst *DStoreInstBase) execute(frame *runtime_data.Frame, idx uint16) error {
	local_variable := frame.OperandStack.PopDouble()

	frame.LocalVariables.SetDouble(idx, local_variable)
	return nil
}

type DStoreInst struct {
	DStoreInstBase
	ByteOprandInstruction
}

func (inst *DStoreInst) Execute(frame *runtime_data.Frame) error {
	idx := uint16(inst.Operand)
	return inst.execute(frame, idx)
}

type DStoreNInst struct {
	idx uint16
	DStoreInstBase
	NoOperandsInstruction
}

func (inst *DStoreNInst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, inst.idx)
}

type AStoreInstBase struct {
}

func (inst *AStoreInstBase) execute(frame *runtime_data.Frame, idx uint16) error {
	local_variable := frame.OperandStack.PopAddress()

	frame.LocalVariables.SetAddress(idx, local_variable)
	return nil
}

type AStoreInst struct {
	AStoreInstBase
	ByteOprandInstruction
}

func (inst *AStoreInst) Execute(frame *runtime_data.Frame) error {
	idx := uint16(inst.Operand)
	return inst.execute(frame, idx)
}

type AStoreNInst struct {
	idx uint16
	AStoreInstBase
	NoOperandsInstruction
}

func (inst *AStoreNInst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, inst.idx)
}
