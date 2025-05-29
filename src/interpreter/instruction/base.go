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

type Index8Instruction struct {
	Index uint8
}

func (inst *Index8Instruction) FetchOperands(reader *ByteCodeReader) error {
	return nil
}

type Index16Instruction struct {
	Index uint16
}

func (inst *Index16Instruction) FetchOperands(reader *ByteCodeReader) error {
	return nil
}

