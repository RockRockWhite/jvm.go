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

type IStoreNInst struct {
	idx uint16
	NoOperandsInstruction
}

func (inst *IStoreNInst) Execute(frame *runtime_data.Frame) error {
	return (&IStoreInstBase{}).execute(frame, inst.idx)
}
