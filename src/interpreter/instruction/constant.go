package instruction

import (
	"github.com/RockRockWhite/jvm.go/src/runtime_data"
)

type Nop struct {
	NoOperandsInstruction
}

func (inst *Nop) Execute(frame *runtime_data.Frame) error {
	// NOP does nothing, so we simply return nil.
	return nil
}



