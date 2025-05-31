package instruction

import "github.com/RockRockWhite/jvm.go/src/runtime_data"

type ILoadInstBase struct {
}

func (inst *ILoadInstBase) execute(frame *runtime_data.Frame, idx uint16) error {
	local_variable := frame.LocalVariables.GetInt(idx)

	frame.OperandStack.Push(local_variable)
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

type ILoad0Inst struct {
	NoOperandsInstruction
	ILoadInstBase
}

func (inst *ILoad0Inst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, 0)
}

type ILoad1Inst struct {
	NoOperandsInstruction
	ILoadInstBase
}

func (inst *ILoad1Inst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, 1)
}

type ILoad2Inst struct {
	NoOperandsInstruction
	ILoadInstBase
}

func (inst *ILoad2Inst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, 2)
}

type ILoad3Inst struct {
	NoOperandsInstruction
	ILoadInstBase
}

func (inst *ILoad3Inst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, 3)
}

type LLoadInstBase struct {
}

func (inst *LLoadInstBase) execute(frame *runtime_data.Frame, idx uint16) error {
	local_variable := frame.LocalVariables.GetLong(idx)

	frame.OperandStack.Push(local_variable)
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

type LLoad0Inst struct {
	NoOperandsInstruction
	LLoadInstBase
}

func (inst *LLoad0Inst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, 0)
}

type LLoad1Inst struct {
	NoOperandsInstruction
	LLoadInstBase
}

func (inst *LLoad1Inst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, 1)
}

type LLoad2Inst struct {
	NoOperandsInstruction
	LLoadInstBase
}

func (inst *ILoad2Inst) LLoad2Inst(frame *runtime_data.Frame) error {
	return inst.execute(frame, 2)
}

type LLoad3Inst struct {
	NoOperandsInstruction
	LLoadInstBase
}

func (inst *LLoad3Inst) Execute(frame *runtime_data.Frame) error {
	return inst.execute(frame, 3)
}
