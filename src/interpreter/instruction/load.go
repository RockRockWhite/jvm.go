package instruction

import "github.com/RockRockWhite/jvm.go/src/runtime_data"

type ILoadInstBase struct {
}

func (inst *ILoadInstBase) execute(frame *runtime_data.Frame, idx uint16) error {
	local_variable := frame.LocalVariables.GetInt(idx)

	frame.OperandStack.PushInt(local_variable)
	return nil
}

type ILoadInst struct {
	ByteOprandInstruction
	ILoadInstBase
}

func (inst *ILoadInst) Execute(frame *runtime_data.Frame) error {
	idx := uint16(inst.Operand)
	return inst.execute(frame, idx)
}

// == ILoad0Inst, ILoad1Inst, ILoad2Inst, ILoad3Inst ==
type ILoadNInst struct {
	idx uint16
	NoOperandsInstruction
	ILoadInstBase
}

func (inst *ILoadNInst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, inst.idx)
}

type LLoadInstBase struct {
}

func (inst *LLoadInstBase) execute(frame *runtime_data.Frame, idx uint16) error {
	local_variable := frame.LocalVariables.GetLong(idx)

	frame.OperandStack.PushLong(local_variable)
	return nil
}

type LLoadInst struct {
	ByteOprandInstruction
	LLoadInstBase
}

func (inst *LLoadInst) Execute(frame *runtime_data.Frame) error {
	idx := uint16(inst.Operand)
	return inst.execute(frame, idx)
}

// == LLoad0Inst, LLoad1Inst, LLoad2Inst, LLoad3Inst ==
type LLoadNInst struct {
	idx uint16
	NoOperandsInstruction
	LLoadInstBase
}

func (inst *LLoadNInst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, inst.idx)
}

type FLoadInstBase struct {
}

func (inst *FLoadInstBase) execute(frame *runtime_data.Frame, idx uint16) error {
	local_variable := frame.LocalVariables.GetFloat(idx)

	frame.OperandStack.PushFloat(local_variable)
	return nil
}

type FLoadInst struct {
	FLoadInstBase
	ByteOprandInstruction
}

func (inst *FLoadInst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, uint16(inst.Operand))
}

type FLoadNInst struct {
	idx uint16
	FLoadInstBase
	ByteOprandInstruction
}

func (inst *FLoadNInst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, inst.idx)
}
