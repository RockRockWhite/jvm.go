package instruction

import "github.com/RockRockWhite/jvm.go/src/runtime_data"

type ILoadInstBase struct {
	ByteOprandInstruction
}

func (inst *ILoadInstBase) Execute(frame *runtime_data.Frame) error {
	idx := int(inst.Operand)
	

}