package class_file

import (
	"fmt"
	"io"
)

type AttributeReader struct {
	byte_reader    ByteReader
	constant_pool  ConstantPoolInfo
	attribute_type AttributeType
	attribute_info AttributeInfo
}

func (r *AttributeReader) readAttributeDeprecated() AttributeInfo {
	length := r.byte_reader.readUInt32()

	if length != 0 {
		r.byte_reader.errors = append(r.byte_reader.errors, fmt.Errorf("invalid length for Deprecated attribute"))
		return nil
	}

	return AttributeDeprecated{}
}

func (r *AttributeReader) readAttributeSynthetic() AttributeInfo {
	length := r.byte_reader.readUInt32()

	if length != 0 {
		r.byte_reader.errors = append(r.byte_reader.errors, fmt.Errorf("invalid length for Synthetic attribute"))
		return nil
	}

	return AttributeSynthetic{}
}

func (r *AttributeReader) readAttributeSourceFile() AttributeInfo {
	length := r.byte_reader.readUInt32()

	if length != 2 {
		r.byte_reader.errors = append(r.byte_reader.errors, fmt.Errorf("invalid length for SourceFile attribute"))
		return nil
	}

	source_file_index := r.byte_reader.readUInt16()
	source_file_str, err := r.constant_pool.GetUtf8(source_file_index)

	if err != nil {
		r.byte_reader.errors = append(r.byte_reader.errors, err)
	}

	return AttributeSourceFile{
		SourceFile: source_file_str,
	}
}

func (r *AttributeReader) readAttributeConstantValue() AttributeInfo {
	length := r.byte_reader.readUInt32()

	if length != 2 {
		r.byte_reader.errors = append(r.byte_reader.errors, fmt.Errorf("invalid length for ConstantValue attribute"))
		return nil
	}

	constant_value_index := r.byte_reader.readUInt16()

	var value any
	var err error
	constant_type := r.constant_pool.GetType(constant_value_index)

	switch constant_type {
	case CONSTANT_Integer:
		value, err = r.constant_pool.GetInt(constant_value_index)
	case CONSTANT_Float:
		value, err = r.constant_pool.GetFloat(constant_value_index)
	case CONSTANT_Long:
		value, err = r.constant_pool.GetLong(constant_value_index)
	case CONSTANT_Double:
		value, err = r.constant_pool.GetDouble(constant_value_index)
	case CONSTANT_String:
		value, err = r.constant_pool.GetString(constant_value_index)
	default:
		value = nil
		err = fmt.Errorf("invalid constant type for ConstantValue attribute")
	}

	if err != nil {
		r.byte_reader.errors = append(r.byte_reader.errors, err)
	}
	return AttributeConstantValue{
		ConstantValue: value,
		ConstantType:  constant_type,
	}
}

func (r *AttributeReader) readExpectionTableEntry() ExpectionTableEntry {
	start_pc := r.byte_reader.readUInt16()
	end_pc := r.byte_reader.readUInt16()
	handler_pc := r.byte_reader.readUInt16()
	catch_type := r.byte_reader.readUInt16()

	return ExpectionTableEntry{
		StartPC:   start_pc,
		EndPC:     end_pc,
		HandlerPC: handler_pc,
		CatchType: catch_type,
	}
}

func (r *AttributeReader) readAttributeCode() AttributeInfo {
	// read length, unused
	_ = r.byte_reader.readUInt32()

	max_stack := r.byte_reader.readUInt16()
	max_locals := r.byte_reader.readUInt16()

	code_len := r.byte_reader.readUInt32()
	code := r.byte_reader.readBytes(uint64(code_len))

	exception_table_len := r.byte_reader.readUInt16()
	exception_table := make([]ExpectionTableEntry, 0, exception_table_len)
	for i := 0; i != int(exception_table_len); i++ {
		exception_table = append(exception_table, r.readExpectionTableEntry())
	}

	attributes_count := r.byte_reader.readUInt16()
	attributes := make([]AttributeInfo, 0, attributes_count)
	for i := 0; i != int(attributes_count); i++ {
		inner_attribute_reader := NewAttributeReader(r.byte_reader.reader, r.constant_pool)
		inner_attribute, err := inner_attribute_reader.
			ReadAttributeName().
			ReadAttributeInfo().
			Build()

		if err != nil {
			r.byte_reader.errors = append(r.byte_reader.errors, err)
		}

		attributes = append(attributes, inner_attribute)
	}

	return AttributeCode{
		MaxStack:       max_stack,
		MaxLocals:      max_locals,
		Code:           code,
		ExpectionTable: exception_table,
		Attributes:     attributes,
	}
}

func (r *AttributeReader) readAttributeException() AttributeInfo {
	// read length, unused
	_ = r.byte_reader.readUInt32()

	exceptions_len := r.byte_reader.readUInt16()
	exceptions := make([]string, 0, exceptions_len)

	for i := 0; i != int(exceptions_len); i++ {

		expection_name_index := r.byte_reader.readUInt16()
		expection_name, err := r.constant_pool.GetClass(expection_name_index)
		if err != nil {
			r.byte_reader.errors = append(r.byte_reader.errors, err)
		}

		exceptions = append(exceptions, expection_name)
	}

	return AttributeExceptions{
		Exceptions: exceptions,
	}
}

func (r *AttributeReader) readAttributeLineNumber() AttributeInfo {
	// read length, unused
	_ = r.byte_reader.readUInt32()

	line_number_len := r.byte_reader.readUInt16()
	line_numbers := make([]LineNumberTableEntry, 0, line_number_len)

	for i := 0; i != int(line_number_len); i += 1 {
		start_pc := r.byte_reader.readUInt16()
		line_number := r.byte_reader.readUInt16()

		line_numbers = append(line_numbers, LineNumberTableEntry{
			StartPC:    start_pc,
			LineNumber: line_number,
		})
	}

	return AttributeLineNumber{
		LineNumberTable: line_numbers,
	}
}

func (r *AttributeReader) readAttributeLocalVariable() AttributeInfo {
	// read length, unused
	_ = r.byte_reader.readUInt32()

	local_variable_len := r.byte_reader.readUInt16()
	local_variables := make([]LocalVariableTableEntry, 0, local_variable_len)

	for i := 0; i != int(local_variable_len); i += 1 {
		start_pc := r.byte_reader.readUInt16()
		length := r.byte_reader.readUInt16()

		name_index := r.byte_reader.readUInt16()
		name, err := r.constant_pool.GetUtf8(name_index)
		if err != nil {
			r.byte_reader.errors = append(r.byte_reader.errors, err)
		}

		descriptor_index := r.byte_reader.readUInt16()
		descriptor, err := r.constant_pool.GetUtf8(descriptor_index)
		if err != nil {
			r.byte_reader.errors = append(r.byte_reader.errors, err)
		}

		index := r.byte_reader.readUInt16()

		local_variables = append(local_variables, LocalVariableTableEntry{
			StartPC:    start_pc,
			Length:     length,
			Name:       name,
			Descriptor: descriptor,
			Index:      index,
		})
	}

	return AttributeLocalVariable{
		LocalVariableTable: local_variables,
	}
}

func (r *AttributeReader) readAttributeUnknown() AttributeInfo {
	length := r.byte_reader.readUInt32()

	// just read and store all the bytes
	data := r.byte_reader.readBytes(uint64(length))

	return AttributeUnknown{
		Data: data,
	}
}

func (r *AttributeReader) ReadAttributeName() *AttributeReader {
	name_index := r.byte_reader.readUInt16()

	name_str, err := r.constant_pool.GetUtf8(name_index)

	if err != nil {
		r.byte_reader.errors = append(r.byte_reader.errors, err)
	}

	r.attribute_type = AttributeType(name_str)

	return r
}

func (r *AttributeReader) ReadAttributeInfo() *AttributeReader {
	switch r.attribute_type {
	case ATTRIBUTE_Code:
		r.attribute_info = r.readAttributeCode()
	case ATTRIBUTE_ConstantValue:
		r.attribute_info = r.readAttributeConstantValue()
	case ATTRIBUTE_Deprecated:
		r.attribute_info = r.readAttributeDeprecated()
	case ATTRIBUTE_Exceptions:
		r.attribute_info = r.readAttributeException()
	case ATTRIBUTE_LineNumberTable:
		r.attribute_info = r.readAttributeLineNumber()
	case ATTRIBUTE_LocalVariableTable:
		r.attribute_info = r.readAttributeLocalVariable()
	case ATTRIBUTE_SourceFile:
		r.attribute_info = r.readAttributeSourceFile()
	case ATTRIBUTE_Synthetic:
		r.attribute_info = r.readAttributeSynthetic()
	default:
		r.attribute_info = r.readAttributeUnknown()
	}

	return r
}

func (r *AttributeReader) Build() (AttributeInfo, error) {
	if len(r.byte_reader.errors) != 0 {
		// convet all errors to a single error
		error_msg := ""
		for _, err := range r.byte_reader.errors {
			error_msg += fmt.Sprintf("%v; ", err)
		}
		return nil, fmt.Errorf("{ %v }", error_msg)
	}

	return r.attribute_info, nil
}

func NewAttributeReader(reader io.Reader, constant_pool ConstantPoolInfo) AttributeReader {
	return AttributeReader{
		byte_reader:   NewByteReader(reader),
		constant_pool: constant_pool,
	}
}
