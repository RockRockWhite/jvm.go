package runtime_data

import "fmt"

type VariableSlot struct {
	Data any
}

type LocalVariables struct {
	Slots []VariableSlot
}

func (vars *LocalVariables) PushInt(data int32) {
	vars.Slots = append(vars.Slots, VariableSlot{
		Data: &data,
	})
}

func (vars *LocalVariables) GetInt(index uint16) int32 {
	data_ptr, ok := vars.Slots[index].Data.(*int32)
	if !ok {
		// just panic
		panic(fmt.Sprintf("Local variable at index %d is not an int", index))
	}

	return *data_ptr
}

func (vars *LocalVariables) SetInt(index uint16, data int32) {
	if index >= uint16(len(vars.Slots)) {
		// just panic
		panic(fmt.Sprintf("Local variable index %d out of bounds", index))
	}

	data_ptr, ok := vars.Slots[index].Data.(*int32)
	if !ok {
		// just panic
		panic(fmt.Sprintf("Local variable at index %d is not an int", index))
	}

	*data_ptr = data
}


func (vars *LocalVariables) PushLong(data int64) {
	vars.Slots = append(vars.Slots, VariableSlot{
		Data: &data,
	})

	// long values take two slots, so we append a nil slot
	vars.Slots = append(vars.Slots, VariableSlot{
		Data: nil,
	})
}

func (vars *LocalVariables) GetLong(index uint16) int64 {
	data_ptr, ok := vars.Slots[index].Data.(*int64)
	if !ok {
		// just panic
		panic(fmt.Sprintf("Local variable at index %d is not an int", index))
	}

	return *data_ptr
}

func (vars *LocalVariables) SetLong(index uint16, data int64) {
	if index >= uint16(len(vars.Slots)) {
		// just panic
		panic(fmt.Sprintf("Local variable index %d out of bounds", index))
	}

	data_ptr, ok := vars.Slots[index].Data.(*int64)
	if !ok {
		// just panic
		panic(fmt.Sprintf("Local variable at index %d is not an int", index))
	}

	*data_ptr = data
}
