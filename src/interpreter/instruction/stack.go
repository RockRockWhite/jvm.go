package instruction

import "github.com/RockRockWhite/jvm.go/src/runtime_data"

type PopInst struct {
	NoOperandsInstruction
}

func (inst *PopInst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.Pop()
	return nil
}

type Pop2Inst struct {
	NoOperandsInstruction
}

func (inst *Pop2Inst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.Pop()
	frame.OperandStack.Pop()
	return nil
}
