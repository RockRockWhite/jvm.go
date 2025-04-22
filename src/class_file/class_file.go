package class_file

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type ClassFile struct {
	reader io.Reader
}

func NewClassFile(data []byte) ClassFile {
	return ClassFile{
		reader: bytes.NewReader(data),
	}
}

func (cf *ClassFile) readUInt8() (uint8, error) {
	var res uint8
	err := binary.Read(cf.reader, binary.BigEndian, &res)

	return res, err
}

func (cf *ClassFile) readUInt16() (uint16, error) {
	var res uint16
	err := binary.Read(cf.reader, binary.BigEndian, &res)

	return res, err
}

func (cf *ClassFile) readUInt32() (uint32, error) {
	var res uint32
	err := binary.Read(cf.reader, binary.BigEndian, &res)

	return res, err
}

func (cf *ClassFile) readUInt64() (uint64, error) {
	var res uint64
	err := binary.Read(cf.reader, binary.BigEndian, &res)

	return res, err
}

func (cf *ClassFile) readInt16() (int16, error) {
	var res int16
	err := binary.Read(cf.reader, binary.BigEndian, &res)

	return res, err
}

func (cf *ClassFile) readInt32() (int32, error) {
	var res int32
	err := binary.Read(cf.reader, binary.BigEndian, &res)

	return res, err
}

func (cf *ClassFile) readInt64() (int64, error) {
	var res int64
	err := binary.Read(cf.reader, binary.BigEndian, &res)

	return res, err
}

func (cf *ClassFile) readFloat32() (float32, error) {
	var res float32
	err := binary.Read(cf.reader, binary.BigEndian, &res)

	return res, err
}

func (cf *ClassFile) readFloat64() (float64, error) {
	var res float64
	err := binary.Read(cf.reader, binary.BigEndian, &res)

	return res, err
}

func (cf *ClassFile) readBytes(size uint16) ([]byte, error) {
	bytes := make([]byte, size)
	_, err := io.ReadFull(cf.reader, bytes)

	return bytes, err
}

func (cf *ClassFile) checkMagic() (bool, error) {
	magic, err := cf.readUInt32()
	if err != nil {
		return false, err
	}

	// magic shall be 0xcafebabe
	return (magic == 0xcafebabe), nil
}

func (cf *ClassFile) readAndCheckVersion() (bool, uint16, uint16, error) {
	minor_version, err := cf.readUInt16()
	if err != nil {
		return false, 0, 0, err
	}

	major_version, err := cf.readUInt16()
	if err != nil {
		return false, 0, 0, err
	}

	// check version
	if (major_version >= 45) && (major_version <= 52) {
		return true, minor_version, major_version, nil
	} else {
		return false, minor_version, major_version, nil
	}
}

func (cf *ClassFile) readAccessFlags() (uint16, error) {
	access_flags, err := cf.readUInt16()

	return access_flags, err
}

func (cf *ClassFile) readIndex() (uint16, error) {
	index, err := cf.readUInt16()
	return index, err
}

func (cf *ClassFile) readIndexTable() ([]uint16, error) {
	index_count, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	indexes := make([]uint16, 0, index_count)

	for i := 0; i != int(index_count); i++ {
		index, err := cf.readUInt16()
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, index)
	}
	return indexes, err
}

func (cf *ClassFile) readConstantInt() (ConstantInfo, error) {
	value, err := cf.readInt32()
	if err != nil {
		return nil, err
	}
	return ConstantInt{
		Value: value,
	}, nil
}

func (cf *ClassFile) readConstantLong() (ConstantInfo, error) {
	value, err := cf.readInt64()
	if err != nil {
		return nil, err
	}
	return ConstantLong{
		Value: value,
	}, nil
}

func (cf *ClassFile) readConstantFloat() (ConstantInfo, error) {
	value, err := cf.readFloat32()
	if err != nil {
		return nil, err
	}
	return ConstantFloat{
		Value: value,
	}, nil
}

func (cf *ClassFile) readConstantDouble() (ConstantInfo, error) {
	value, err := cf.readFloat64()
	if err != nil {
		return nil, err
	}
	return ConstantDouble{
		Value: value,
	}, nil
}

func (cf *ClassFile) readConstantUtf8() (ConstantInfo, error) {
	length, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	var bytes []byte
	bytes, err = cf.readBytes(length)
	if err != nil {
		return nil, err
	}
	// convert bytes to string
	return ConstantUtf8{
		Str: string(bytes),
	}, nil
}

func (cf *ClassFile) readConstantString() (ConstantInfo, error) {
	index, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	// convert bytes to string
	return ConstantString{
		EntryIndex: index,
	}, nil
}

func (cf *ClassFile) readConstantClass() (ConstantInfo, error) {
	index, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	// convert bytes to string
	return ConstantClass{
		EntryIndex: index,
	}, nil
}

func (cf *ClassFile) readConstantNameAndType() (ConstantInfo, error) {
	name_index, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	descriptor_index, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	return ConstantNameAndType{
		NameEntryIndex:       name_index,
		DescriptorEntryIndex: descriptor_index,
	}, nil
}

func (cf *ClassFile) readConstantMemberRef() (ConstantMemberRef, error) {
	class_index, err := cf.readUInt16()
	if err != nil {
		return ConstantMemberRef{}, err
	}

	name_and_type_index, err := cf.readUInt16()
	if err != nil {
		return ConstantMemberRef{}, err
	}

	// convert bytes to string
	return ConstantMemberRef{
		ClassEntryIndex:       class_index,
		NameAndTypeEntryIndex: name_and_type_index,
	}, nil
}

func (cf *ClassFile) readConstantInfo() (ConstantInfo, error) {
	// read constant type tag
	tag, err := cf.readUInt8()
	if err != nil {
		return nil, err
	}

	constant_type := ConstantType(tag)
	switch constant_type {
	case CONSTANT_Class:
		return cf.readConstantClass()
	case CONSTANT_Fieldref:
		if constant_member_ref, err := cf.readConstantMemberRef(); err != nil {
			return nil, err
		} else {
			return NewConstantFieldRef(constant_member_ref), nil
		}
	case CONSTANT_Methodref:
		if constant_member_ref, err := cf.readConstantMemberRef(); err != nil {
			return nil, err
		} else {
			return NewConstantMethodRef(constant_member_ref), nil
		}
	case CONSTANT_InferfaceMethodref:
		if constant_member_ref, err := cf.readConstantMemberRef(); err != nil {
			return nil, err
		} else {
			return NewConstantInterfaceMethodRef(constant_member_ref), nil
		}
	case CONSTANT_String:
		return cf.readConstantString()
	case CONSTANT_Integer:
		return cf.readConstantInt()
	case CONSTANT_Float:
		return cf.readConstantFloat()
	case CONSTANT_Long:
		return cf.readConstantLong()
	case CONSTANT_Double:
		return cf.readConstantDouble()
	case CONSTANT_NameAndType:
		return cf.readConstantNameAndType()
	case CONSTANT_Utf8:
		return cf.readConstantUtf8()
	case CONSTANT_MethodHandle:
	case CONSTANT_MethodType:
	case CONSTANT_InvokeDynamic:
	default:
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type %d", tag)
	}

	return nil, nil
}

func (cf *ClassFile) readConstantPoolInfo() (ConstantPoolInfo, error) {

	// read constant pool count
	cnt, err := cf.readUInt16()
	if err != nil {
		return ConstantPoolInfo{}, err
	}

	entries := make([]ConstantInfo, 0, cnt-1)
	entries = append(entries, nil) // index 0 is not used

	// only avaliable till cnt-1
	// [1, cnt)
	for i := 1; i != int(cnt); i++ {
		constant_info, err := cf.readConstantInfo()
		if err != nil {
			return ConstantPoolInfo{}, err
		}

		entries = append(entries, constant_info)

		// long and double take up two slots
		// append one empty slot for convenience
		switch constant_info.(type) {
		case ConstantLong, ConstantDouble:
			entries = append(entries, nil)
			i++
		}
	}

	return ConstantPoolInfo{
		Entries: entries,
	}, nil
}

func (cf *ClassFile) readAttributeDeprecated() (AttributeInfo, error) {
	length, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}

	if length != 0 {
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid length for Deprecated attribute")
	}

	return AttributeDeprecated{}, nil
}

func (cf *ClassFile) readAttributeSynthetic() (AttributeInfo, error) {
	length, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}

	if length != 0 {
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid length for Synthetic attribute")
	}

	return AttributeSynthetic{}, nil
}

func (cf *ClassFile) readAttributeSourceFile(constant_pool ConstantPoolInfo) (AttributeInfo, error) {
	length, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}

	if length != 2 {
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid length for SourceFile attribute")
	}

	source_file_index, err := cf.readUInt16()

	if err != nil {
		return nil, err
	}

	constant_string_ptr, ok := constant_pool.Entries[source_file_index].(ConstantUtf8)
	if !ok {
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for SourceFile attribute")
	}

	return AttributeSourceFile{
		SourceFile: constant_string_ptr.Str,
	}, nil
}

func (cf *ClassFile) readAttributeConstantValue(constant_pool ConstantPoolInfo) (AttributeInfo, error) {
	length, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}

	if length != 2 {
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid length for ConstantValue attribute")
	}

	constant_value_index, err := cf.readUInt16()

	if err != nil {
		return nil, err
	}

	switch constant_pool.Entries[constant_value_index].(type) {
	case ConstantInt:
		return AttributeConstantValue{
			ConstantValue: constant_pool.Entries[constant_value_index].(ConstantInt).Value,
		}, nil
	case ConstantFloat:
		return AttributeConstantValue{
			ConstantValue: constant_pool.Entries[constant_value_index].(ConstantFloat).Value,
		}, nil
	case ConstantLong:
		return AttributeConstantValue{
			ConstantValue: constant_pool.Entries[constant_value_index].(ConstantLong).Value,
		}, nil
	case ConstantDouble:
		return AttributeConstantValue{
			ConstantValue: constant_pool.Entries[constant_value_index].(ConstantDouble).Value,
		}, nil
	case ConstantString:
		return AttributeConstantValue{
			ConstantValue: constant_pool.Entries[constant_value_index].(ConstantString).EntryIndex,
		}, nil
	default:
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for ConstantValue attribute")
	}

}

func (cf *ClassFile) readExpectionTableEntry() (ExpectionTableEntry, error) {
	return ExpectionTableEntry{}, nil
}

func (cf *ClassFile) readAttributeCode(constant_pool ConstantPoolInfo) (AttributeInfo, error) {
	// read length, unused
	_, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}

	// read max_stack
	max_stack, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	max_locals, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	code_len, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}
	code := make([]byte, code_len)

	if _, err := io.ReadFull(cf.reader, code); err != nil {
		return nil, err
	}

	exception_table_len, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	exception_table := make([]ExpectionTableEntry, 0, exception_table_len)
	for i := 0; i != int(exception_table_len); i++ {
		expection_table_entry, err := cf.readExpectionTableEntry()
		if err != nil {
			return nil, nil
		}
		exception_table = append(exception_table, expection_table_entry)
	}

	attributes, err := cf.readAttributes(constant_pool)
	if err != nil {
		return nil, err
	}

	return AttributeCode{
		MaxStack:       max_stack,
		MaxLocals:      max_locals,
		Code:           code,
		ExpectionTable: exception_table,
		Attributes:     attributes,
	}, nil
}

func (cf *ClassFile) readAttributeException(constant_pool ConstantPoolInfo) (AttributeInfo, error) {
	// read length, unused
	_, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}

	// read exceptions
	exceptions_len, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}
	exceptions := make([]string, 0, exceptions_len)

	for i := 0; i != int(exceptions_len); i++ {

		expection_name_index, err := cf.readUInt16()
		if err != nil {
			return nil, err
		}

		constant_class, ok := constant_pool.Entries[expection_name_index].(ConstantClass)
		if !ok {
			return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for Exception attribute")
		}

		exceptions = append(exceptions, constant_class.GetClassName(constant_pool))
	}

	return AttributeExceptions{
		Exceptions: exceptions,
	}, nil
}

func (cf *ClassFile) readAttributeLineNumber() (AttributeInfo, error) {
	// read length, unused
	_, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}

	line_number_len, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	line_numbers := make([]LineNumberTableEntry, 0, line_number_len)
	for i := 0; i != int(line_number_len); i += 1 {
		start_pc, err := cf.readUInt16()
		if err != nil {
			return nil, err
		}

		line_number, err := cf.readUInt16()
		if err != nil {
			return nil, err
		}

		line_numbers = append(line_numbers, LineNumberTableEntry{
			StartPC:    start_pc,
			LineNumber: line_number,
		})
	}

	return AttributeLineNumber{
		LineNumberTable: line_numbers,
	}, nil
}

func (cf *ClassFile) readAttributeLocalVariable(constant_pool ConstantPoolInfo) (AttributeInfo, error) {
	// read length, unused
	_, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}

	local_variable_len, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	local_variables := make([]LocalVariableTableEntry, 0, local_variable_len)
	for i := 0; i != int(local_variable_len); i += 1 {
		start_pc, err := cf.readUInt16()
		if err != nil {
			return nil, err
		}

		length, err := cf.readUInt16()
		if err != nil {
			return nil, err
		}

		name_index, err := cf.readUInt16()
		if err != nil {
			return nil, err
		}

		name_constant_utf8_ptr, ok := constant_pool.Entries[name_index].(*ConstantUtf8)
		if !ok {
			return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for LocalVariable attribute")
		}

		descriptor_index, err := cf.readUInt16()
		if err != nil {
			return nil, err
		}

		descriptor_constant_utf8_ptr, ok := constant_pool.Entries[descriptor_index].(*ConstantUtf8)
		if !ok {
			return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for LocalVariable attribute")
		}

		index, err := cf.readUInt16()
		if err != nil {
			return nil, err
		}

		local_variables = append(local_variables, LocalVariableTableEntry{
			StartPC:    start_pc,
			Length:     length,
			Name:       name_constant_utf8_ptr.Str,
			Descriptor: descriptor_constant_utf8_ptr.Str,
			Index:      index,
		})
	}

	return AttributeLocalVariable{
		LocalVariableTable: local_variables,
	}, nil
}

func (cf *ClassFile) readAttributeUnknown() (AttributeInfo, error) {
	// read length, unused
	length, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}

	// read attribute data
	data := make([]byte, length)
	io.ReadFull(cf.reader, data)
	if err != nil {
		return nil, err
	}

	return AttributeUnknown{
		Data: data,
	}, nil
}

func (cf *ClassFile) readAttributeInfo(constant_pool ConstantPoolInfo) (AttributeInfo, error) {

	// read attribute name
	name_index, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	constant_utf8_ptr, ok := constant_pool.Entries[name_index].(ConstantUtf8)
	if !ok {
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for attribute name")
	}
	attribute_type := AttributeType(constant_utf8_ptr.Str)
	switch attribute_type {
	case ATTRIBUTE_Code:
		return cf.readAttributeCode(constant_pool)
	case ATTRIBUTE_ConstantValue:
		return cf.readAttributeConstantValue(constant_pool)
	case ATTRIBUTE_Deprecated:
		return cf.readAttributeDeprecated()
	case ATTRIBUTE_Exceptions:
		return cf.readAttributeException(constant_pool)
	case ATTRIBUTE_LineNumberTable:
		return cf.readAttributeLineNumber()
	case ATTRIBUTE_LocalVariableTable:
		return cf.readAttributeLocalVariable(constant_pool)
	case ATTRIBUTE_SourceFile:
		return cf.readAttributeSourceFile(constant_pool)
	case ATTRIBUTE_Synthetic:
		return cf.readAttributeSynthetic()
	default:
		return cf.readAttributeUnknown()
	}
}

func (cf *ClassFile) readAttributes(constant_pool ConstantPoolInfo) ([]AttributeInfo, error) {
	cnt, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	attributes := make([]AttributeInfo, 0, cnt)

	for i := 0; i != int(cnt); i++ {
		if attribute, err := cf.readAttributeInfo(constant_pool); err != nil {
			return nil, err
		} else {
			attributes = append(attributes, attribute)
		}
	}

	return attributes, nil
}

func (cf *ClassFile) readMember(constant_pool ConstantPoolInfo) (MemberInfo, error) {
	access_flags, err := cf.readAccessFlags()
	if err != nil {
		return MemberInfo{}, err
	}

	name_index, err := cf.readIndex()
	if err != nil {
		return MemberInfo{}, err
	}

	descriptor_index, err := cf.readIndex()
	if err != nil {
		return MemberInfo{}, err
	}

	attributes, err := cf.readAttributes(constant_pool)

	if err != nil {
		return MemberInfo{}, err
	}

	return MemberInfo{
		AccessFlags:     access_flags,
		NameIndex:       name_index,
		DescriptorIndex: descriptor_index,
		Attrubutes:      attributes,
	}, nil
}

func (cf *ClassFile) IntoClassInfo() (ClassInfo, error) {
	var res ClassInfo

	if ok, err := cf.checkMagic(); err != nil {
		return ClassInfo{}, err
	} else if !ok {
		return ClassInfo{}, fmt.Errorf("java.lang.ClassFormatError: invalid magic number")
	}

	if ok, minor_version, major_version, err := cf.readAndCheckVersion(); err != nil {
		return ClassInfo{}, err
	} else if !ok {
		return ClassInfo{}, fmt.Errorf("java.lang.UnsupportedClassVersionError")
	} else {
		res.MinorVersion = minor_version
		res.MajorVersion = major_version
	}

	if constant_pool, err := cf.readConstantPoolInfo(); err != nil {
		return ClassInfo{}, err
	} else {
		res.ConstantPool = constant_pool
	}

	if access_flags, err := cf.readAccessFlags(); err != nil {
		return ClassInfo{}, err
	} else {
		res.AccessFlags = access_flags
	}

	if this_class, err := cf.readIndex(); err != nil {
		return ClassInfo{}, err
	} else {
		res.ThisClassIndex = this_class
	}

	if super_class, err := cf.readIndex(); err != nil {
		return ClassInfo{}, err
	} else {
		res.SuperClassIndex = super_class
	}

	if interface_indexes, err := cf.readIndexTable(); err != nil {
		return ClassInfo{}, err
	} else {
		interfaces := make([]InterfaceInfo, 0, len(interface_indexes))
		for _, index := range interface_indexes {
			interfaces = append(interfaces, InterfaceInfo{
				NameIndex: index,
			})
		}

		res.Interfaces = interfaces
	}

	if fields_len, err := cf.readUInt16(); err != nil {
		return ClassInfo{}, err
	} else {
		fields := make([]FieldInfo, 0, fields_len)
		for i := 0; i != int(fields_len); i++ {
			member_info, err := cf.readMember(res.ConstantPool)
			if err != nil {
				return ClassInfo{}, err
			}

			fields = append(fields, FieldInfo(member_info))
		}

		res.Fields = fields
	}

	if methods_len, err := cf.readUInt16(); err != nil {
		return ClassInfo{}, err
	} else {
		methods := make([]MethodInfo, 0, methods_len)
		for i := 0; i != int(methods_len); i++ {
			member_info, err := cf.readMember(res.ConstantPool)
			if err != nil {
				return ClassInfo{}, err
			}

			methods = append(methods, MethodInfo(member_info))
		}

		res.Methods = methods
	}

	attributes, err := cf.readAttributes(res.ConstantPool)
	if err != nil {
		return ClassInfo{}, err
	}
	res.Attributes = attributes

	return res, nil
}
