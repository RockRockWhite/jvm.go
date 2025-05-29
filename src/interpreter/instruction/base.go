package instruction

import (
	"io"

	"github.com/RockRockWhite/jvm.go/src/runtime_data"
)

type ByteCodeReader struct {
	Reader *io.Reader
}

type Instruction interface {
	FetchOperands(reader *ByteCodeReader) error
	Execute(frame *runtime_data.Frame) error
}

type NoOperandsInstruction struct {
}

func (inst *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) error {
	// Do nothing, as this instruction does not require any operands.
	return nil
}

type BranchInstruction struct {
	Offset int
}

func (inst *BranchInstruction) FetchOperands(reader *ByteCodeReader) error {
	return nil
}

type ByteOprandInstruction struct {
	Operand uint8
}

func (inst *ByteOprandInstruction) FetchOperands(reader *ByteCodeReader) error {
	return nil
}

type DByteOprandInstruction struct {
	Index uint16
}

func (inst *DByteOprandInstruction) FetchOperands(reader *ByteCodeReader) error {
	return nil
}

