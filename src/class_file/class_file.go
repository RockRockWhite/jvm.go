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

	constant_string_ptr, ok := constant_pool.Entries[source_file_index].(*ConstantUtf8)
	if !ok {
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for SourceFile attribute")
	}

	return AttributeSourceFile{
		SourceFile: constant_string_ptr.Str,
	}, nil
}

func (cf *ClassFile) readAttributeInfo(constant_pool ConstantPoolInfo) (AttributeInfo, error) {

	// read attribute name
	name_index, err := cf.readUInt16()
	if err != nil {
		return nil, err
	}

	constant_utf8_ptr, ok := constant_pool.Entries[name_index].(*ConstantUtf8)
	if !ok {
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type for attribute name")
	}

	attribute_type := AttributeType(constant_utf8_ptr.Str)
	switch attribute_type {
	// case ATTRIBUTE_Code:
	// case ATTRIBUTE_ConstantValue:
	case ATTRIBUTE_Deprecated:
		return cf.readAttributeDeprecated()
	// case ATTRIBUTE_Exceptions:
	// case ATTRIBUTE_LineNumberTable:
	// case ATTRIBUTE_LocalVariableTable:
	case ATTRIBUTE_SourceFile:
		return cf.readAttributeSourceFile(constant_pool)
	case ATTRIBUTE_Synthetic:
		return cf.readAttributeSynthetic()
	default:
		return nil, fmt.Errorf("java.lang.ClassFormatError: invalid attribute type %s", attribute_type)
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

	if field_indexes, err := cf.readIndexTable(); err != nil {
		return ClassInfo{}, err
	} else {
		fields := make([]FieldInfo, 0, len(field_indexes))
		for _, index := range field_indexes {
			fields = append(fields, FieldInfo{
				NameIndex: index,
			})
		}

		res.Fields = fields
	}

	if method_indexes, err := cf.readIndexTable(); err != nil {
		return ClassInfo{}, err
	} else {
		methods := make([]MethodInfo, 0, len(method_indexes))
		for _, index := range method_indexes {
			methods = append(methods, MethodInfo{
				NameIndex: index,
			})
		}

		res.Methods = methods
	}

	return res, nil
}
