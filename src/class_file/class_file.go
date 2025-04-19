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

func (cf *ClassFile) readBytes(size uint32) ([]byte, error) {
	bytes := make([]byte, size)
	cf.reader.Read(bytes)
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

func (cf *ClassFile) readThisClassIndex() (uint16, error) {
	index, err := cf.readUInt16()
	return index, err
}

func (cf *ClassFile) readSuperClassIndex() (uint16, error) {
	index, err := cf.readUInt16()
	return index, err
}

func (cf *ClassFile) readConstantInt() (ConstantInfo, error) {
	value, err := cf.readInt32()
	if err != nil {
		return nil, err
	}
	return ConstantInt{
		value: value,
	}, nil
}

func (cf *ClassFile) readConstantLong() (ConstantInfo, error) {
	value, err := cf.readInt64()
	if err != nil {
		return nil, err
	}
	return ConstantLong{
		value: value,
	}, nil
}

func (cf *ClassFile) readConstantFloat() (ConstantInfo, error) {
	value, err := cf.readFloat32()
	if err != nil {
		return nil, err
	}
	return ConstantFloat{
		value: value,
	}, nil
}

func (cf *ClassFile) readConstantDouble() (ConstantInfo, error) {
	value, err := cf.readFloat64()
	if err != nil {
		return nil, err
	}
	return ConstantDouble{
		value: value,
	}, nil
}

func (cf *ClassFile) readConstantUtf8() (ConstantInfo, error) {
	len, err := cf.readUInt32()
	if err != nil {
		return nil, err
	}

	var bytes []byte
	bytes, err = cf.readBytes(len)
	if err != nil {
		return nil, err
	}

	// convert bytes to string
	return ConstantUtf8{
		str: string(bytes),
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
	tag, err := cf.readUInt16()
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

	// only avaliable till cnt-1
	for i := 0; i != int(cnt)-1; i++ {
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

	if access_flags, err := cf.readAccessFlags(); err != nil {
		return ClassInfo{}, err
	} else {
		res.AccessFlags = access_flags
	}

	if this_class, err := cf.readThisClassIndex(); err != nil {
		return ClassInfo{}, err
	} else {
		res.ThisClassIndex = this_class
	}

	if super_class, err := cf.readSuperClassIndex(); err != nil {
		return ClassInfo{}, err
	} else {
		res.SuperClassIndex = super_class
	}

	return res, nil
}
