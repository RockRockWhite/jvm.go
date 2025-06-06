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

type DupInst struct {
	NoOperandsInstruction
}

func (inst *DupInst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.Dup()
	return nil
}

type DupX1Inst struct {
	NoOperandsInstruction
}

func (inst *DupX1Inst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.DupX1()
	return nil
}

type DupX2Inst struct {
	NoOperandsInstruction
}

func (inst *DupX2Inst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.DupX2()
	return nil
}

type Dup2Inst struct {
	NoOperandsInstruction
}

func (inst *Dup2Inst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.Dup2()
	return nil
}

type Dup2X1Inst struct {
	NoOperandsInstruction
}

func (inst *Dup2X1Inst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.Dup2X1()
	return nil
}

type Dup2X2Inst struct {
	NoOperandsInstruction
}

func (inst *Dup2X2Inst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.Dup2X2()
	return nil
}

type SwapInst struct {
	NoOperandsInstruction
}

func (inst *SwapInst) Execute(frame *runtime_data.Frame) error {
	frame.OperandStack.Swap()
	return nil
}
