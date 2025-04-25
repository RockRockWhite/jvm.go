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

	r.byte_reader.errors = append(r.byte_reader.errors, err)

	return AttributeSourceFile{
		SourceFile: source_file_str,
	}
}

func (r *AttributeReader) readAttributeConstantValue() AttributeInfo {
	length := r.byte_reader.readUInt32()

	if length != 2 {
		r.byte_reader.errors = append(r.byte_reader.errors, fmt.Errorf("invalid length for SourceFile attribute"))
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

	r.byte_reader.errors = append(r.byte_reader.errors, err)
	return AttributeConstantValue{
		ConstantValue: value,
		ConstantType:  constant_type,
	}
}

// func (cf *ClassFile) readExpectionTableEntry() (ExpectionTableEntry, error) {
// 	return ExpectionTableEntry{}, nil
// }

// func (cf *ClassFile) readAttributeCode(constant_pool ConstantPoolInfo) (AttributeInfo, error) {
// 	// read length, unused
// 	_, err := cf.readUInt32()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// read max_stack
// 	max_stack, err := cf.readUInt16()
// 	if err != nil {
// 		return nil, err
// 	}

// 	max_locals, err := cf.readUInt16()
// 	if err != nil {
// 		return nil, err
// 	}

// 	code_len, err := cf.readUInt32()
// 	if err != nil {
// 		return nil, err
// 	}
// 	code := make([]byte, code_len)

// 	if _, err := io.ReadFull(cf.reader, code); err != nil {
// 		return nil, err
// 	}

// 	exception_table_len, err := cf.readUInt16()
// 	if err != nil {
// 		return nil, err
// 	}

// 	exception_table := make([]ExpectionTableEntry, 0, exception_table_len)
// 	for i := 0; i != int(exception_table_len); i++ {
// 		expection_table_entry, err := cf.readExpectionTableEntry()
// 		if err != nil {
// 			return nil, nil
// 		}
// 		exception_table = append(exception_table, expection_table_entry)
// 	}

// 	attributes, err := cf.readAttributes(constant_pool)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return AttributeCode{
// 		MaxStack:       max_stack,
// 		MaxLocals:      max_locals,
// 		Code:           code,
// 		ExpectionTable: exception_table,
// 		Attributes:     attributes,
// 	}, nil
// }

// func (cf *ClassFile) readAttributeException(constant_pool ConstantPoolInfo) (AttributeInfo, error) {
// 	// read length, unused
// 	_, err := cf.readUInt32()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// read exceptions
// 	exceptions_len, err := cf.readUInt16()
// 	if err != nil {
// 		return nil, err
// 	}
// 	exceptions := make([]string, 0, exceptions_len)

// 	for i := 0; i != int(exceptions_len); i++ {

// 		expection_name_index, err := cf.readUInt16()
// 		if err != nil {
// 			return nil, err
// 		}

// 		constant_class, ok := constant_pool.Entries[expection_name_index].(ConstantClass)
// 		if !ok {
// 			return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for Exception attribute")
// 		}

// 		exceptions = append(exceptions, constant_class.GetClassName(constant_pool))
// 	}

// 	return AttributeExceptions{
// 		Exceptions: exceptions,
// 	}, nil
// }

// func (cf *ClassFile) readAttributeLineNumber() (AttributeInfo, error) {
// 	// read length, unused
// 	_, err := cf.readUInt32()
// 	if err != nil {
// 		return nil, err
// 	}

// 	line_number_len, err := cf.readUInt16()
// 	if err != nil {
// 		return nil, err
// 	}

// 	line_numbers := make([]LineNumberTableEntry, 0, line_number_len)
// 	for i := 0; i != int(line_number_len); i += 1 {
// 		start_pc, err := cf.readUInt16()
// 		if err != nil {
// 			return nil, err
// 		}

// 		line_number, err := cf.readUInt16()
// 		if err != nil {
// 			return nil, err
// 		}

// 		line_numbers = append(line_numbers, LineNumberTableEntry{
// 			StartPC:    start_pc,
// 			LineNumber: line_number,
// 		})
// 	}

// 	return AttributeLineNumber{
// 		LineNumberTable: line_numbers,
// 	}, nil
// }

// func (cf *ClassFile) readAttributeLocalVariable(constant_pool ConstantPoolInfo) (AttributeInfo, error) {
// 	// read length, unused
// 	_, err := cf.readUInt32()
// 	if err != nil {
// 		return nil, err
// 	}

// 	local_variable_len, err := cf.readUInt16()
// 	if err != nil {
// 		return nil, err
// 	}

// 	local_variables := make([]LocalVariableTableEntry, 0, local_variable_len)
// 	for i := 0; i != int(local_variable_len); i += 1 {
// 		start_pc, err := cf.readUInt16()
// 		if err != nil {
// 			return nil, err
// 		}

// 		length, err := cf.readUInt16()
// 		if err != nil {
// 			return nil, err
// 		}

// 		name_index, err := cf.readUInt16()
// 		if err != nil {
// 			return nil, err
// 		}

// 		name_constant_utf8_ptr, ok := constant_pool.Entries[name_index].(*ConstantUtf8)
// 		if !ok {
// 			return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for LocalVariable attribute")
// 		}

// 		descriptor_index, err := cf.readUInt16()
// 		if err != nil {
// 			return nil, err
// 		}

// 		descriptor_constant_utf8_ptr, ok := constant_pool.Entries[descriptor_index].(*ConstantUtf8)
// 		if !ok {
// 			return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for LocalVariable attribute")
// 		}

// 		index, err := cf.readUInt16()
// 		if err != nil {
// 			return nil, err
// 		}

// 		local_variables = append(local_variables, LocalVariableTableEntry{
// 			StartPC:    start_pc,
// 			Length:     length,
// 			Name:       name_constant_utf8_ptr.Str,
// 			Descriptor: descriptor_constant_utf8_ptr.Str,
// 			Index:      index,
// 		})
// 	}

// 	return AttributeLocalVariable{
// 		LocalVariableTable: local_variables,
// 	}, nil
// }

// func (cf *ClassFile) readAttributeUnknown() (AttributeInfo, error) {
// 	// read length, unused
// 	length, err := cf.readUInt32()
// 	if err != nil {
// 		return nil, err
// 	}

// 	// read attribute data
// 	data := make([]byte, length)
// 	io.ReadFull(cf.reader, data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return AttributeUnknown{
// 		Data: data,
// 	}, nil
// }

// func (cf *ClassFile) readAttributeInfo(constant_pool ConstantPoolInfo) (AttributeInfo, error) {

// 	// read attribute name
// 	name_index, err := cf.readUInt16()
// 	if err != nil {
// 		return nil, err
// 	}

// 	constant_utf8_ptr, ok := constant_pool.Entries[name_index].(ConstantUtf8)
// 	if !ok {
// 		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for attribute name")
// 	}
// 	attribute_type := AttributeType(constant_utf8_ptr.Str)
// 	switch attribute_type {
// 	case ATTRIBUTE_Code:
// 		return cf.readAttributeCode(constant_pool)
// 	case ATTRIBUTE_ConstantValue:
// 		return cf.readAttributeConstantValue(constant_pool)
// 	case ATTRIBUTE_Deprecated:
// 		return cf.readAttributeDeprecated()
// 	case ATTRIBUTE_Exceptions:
// 		return cf.readAttributeException(constant_pool)
// 	case ATTRIBUTE_LineNumberTable:
// 		return cf.readAttributeLineNumber()
// 	case ATTRIBUTE_LocalVariableTable:
// 		return cf.readAttributeLocalVariable(constant_pool)
// 	case ATTRIBUTE_SourceFile:
// 		return cf.readAttributeSourceFile(constant_pool)
// 	case ATTRIBUTE_Synthetic:
// 		return cf.readAttributeSynthetic()
// 	default:
// 		return cf.readAttributeUnknown()
// 	}
// }

func (r *AttributeReader) ReadAttributeName() *AttributeReader {
	name_index := r.byte_reader.readUInt16()

	name_str, err := r.constant_pool.GetUtf8(name_index)
	r.byte_reader.errors = append(r.byte_reader.errors, err)

	r.attribute_type = AttributeType(name_str)

	return r
}

func (r *AttributeReader) ReadAttributeInfo() *AttributeReader {
	switch r.attribute_type {
	case ATTRIBUTE_Code:
	case ATTRIBUTE_ConstantValue:
		r.attribute_info = r.readAttributeConstantValue()
	case ATTRIBUTE_Deprecated:
		r.attribute_info = r.readAttributeDeprecated()
	case ATTRIBUTE_Exceptions:
		r.attribute_info = r.readAttributeSynthetic()
	case ATTRIBUTE_LineNumberTable:
	case ATTRIBUTE_LocalVariableTable:
	case ATTRIBUTE_SourceFile:
		r.attribute_info = r.readAttributeSourceFile()
	case ATTRIBUTE_Synthetic:
	default:
	}

	return r

	// switch attribute_type {
	// case ATTRIBUTE_Code:
	// 	return cf.readAttributeCode(constant_pool)
	// case ATTRIBUTE_ConstantValue:
	// 	return cf.readAttributeConstantValue(constant_pool)
	// case ATTRIBUTE_Deprecated:
	// 	return cf.readAttributeDeprecated()
	// case ATTRIBUTE_Exceptions:
	// 	return cf.readAttributeException(constant_pool)
	// case ATTRIBUTE_LineNumberTable:
	// 	return cf.readAttributeLineNumber()
	// case ATTRIBUTE_LocalVariableTable:
	// 	return cf.readAttributeLocalVariable(constant_pool)
	// case ATTRIBUTE_SourceFile:
	// 	return cf.readAttributeSourceFile(constant_pool)
	// case ATTRIBUTE_Synthetic:
	// 	return cf.readAttributeSynthetic()
	// default:
	// 	return cf.readAttributeUnknown()
	// }
}

func (r *AttributeReader) Build() (AttributeInfo, []error) {
	return r.attribute_info, r.byte_reader.errors
}

func NewAttributeReader(reader io.Reader, constant_pool ConstantPoolInfo) AttributeReader {
	return AttributeReader{
		byte_reader:   NewByteReader(reader),
		constant_pool: constant_pool,
	}
}
