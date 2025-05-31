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
	ByteOprandInstruction
}

func (inst *IStoreInst) Execute(frame *runtime_data.Frame) error {
	idx := uint16(inst.Operand)
	return (&IStoreInstBase{}).execute(frame, idx)
}

type IStore0Inst struct {
	NoOperandsInstruction
}

func (inst *IStore0Inst) Execute(frame *runtime_data.Frame) error {
	return (&IStoreInstBase{}).execute(frame, 0)
}

type IStore1Inst struct {
	NoOperandsInstruction
}

func (inst *IStore1Inst) Execute(frame *runtime_data.Frame) error {
	return (&IStoreInstBase{}).execute(frame, 1)
}

type IStore2Inst struct {
	NoOperandsInstruction
}

func (inst *IStore2Inst) Execute(frame *runtime_data.Frame) error {
	return (&IStoreInstBase{}).execute(frame, 2)
}

type IStore3Inst struct {
	NoOperandsInstruction
}

func (inst *IStore3Inst) Execute(frame *runtime_data.Frame) error {
	return (&IStoreInstBase{}).execute(frame, 3)
}
