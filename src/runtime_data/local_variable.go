package runtime_data

import "fmt"

type VariableSlot struct {
	Data any
}

type LocalVariables struct {
	Slots []VariableSlot
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

	vars.Slots[index].Data = &data
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

	vars.Slots[index].Data = &data
}

func (vars *LocalVariables) GetFloat(index uint16) float32 {
	data_ptr, ok := vars.Slots[index].Data.(*float32)
	if !ok {
		// just panic
		panic(fmt.Sprintf("Local variable at index %d is not an int", index))
	}

	return *data_ptr
}

func (vars *LocalVariables) SetFloat(index uint16, data float32) {
	if index >= uint16(len(vars.Slots)) {
		// just panic
		panic(fmt.Sprintf("Local variable index %d out of bounds", index))
	}

	vars.Slots[index].Data = &data
}

func (vars *LocalVariables) GetDouble(index uint16) float64 {
	data_ptr, ok := vars.Slots[index].Data.(*float64)
	if !ok {
		// just panic
		panic(fmt.Sprintf("Local variable at index %d is not an int", index))
	}

	return *data_ptr
}

func (vars *LocalVariables) SetDouble(index uint16, data float64) {
	if index >= uint16(len(vars.Slots)) {
		// just panic
		panic(fmt.Sprintf("Local variable index %d out of bounds", index))
	}

	vars.Slots[index].Data = &data
}

func (vars *LocalVariables) GetAddress(index uint16) *Object {
	data_ptr, ok := vars.Slots[index].Data.(*Object)
	if !ok {
		// just panic
		panic(fmt.Sprintf("Local variable at index %d is not an object", index))
	}

	return data_ptr
}

func (vars *LocalVariables) SetAddress(index uint16, data *Object) {
	if index >= uint16(len(vars.Slots)) {
		// just panic
		panic(fmt.Sprintf("Local variable index %d out of bounds", index))
	}

	vars.Slots[index].Data = data
}
