package class_file

import (
	"fmt"
	"io"
)

type ConstantPoolReader struct {
	byte_reader       ByteReader
	constant_pool_cnt uint16
	constant_pool     ConstantPoolInfo
}

func (r *ConstantPoolReader) ReadCount() *ConstantPoolReader {
	r.constant_pool_cnt = r.byte_reader.readUInt16()
	return r
}

func (r *ConstantPoolReader) readConstantType() ConstantType {
	tag := r.byte_reader.readUInt8()
	constant_type := ConstantType(tag)
	return constant_type
}

func (r *ConstantPoolReader) readConstantInt() ConstantInfo {
	value := r.byte_reader.readInt32()

	return ConstantInt{
		Value: value,
	}
}

func (r *ConstantPoolReader) readConstantLong() ConstantInfo {
	value := r.byte_reader.readInt64()

	return ConstantLong{
		Value: value,
	}
}

func (r *ConstantPoolReader) readConstantFloat() ConstantInfo {
	value := r.byte_reader.readFloat32()

	return ConstantFloat{
		Value: value,
	}
}

func (r *ConstantPoolReader) readConstantDouble() ConstantInfo {
	value := r.byte_reader.readFloat64()

	return ConstantDouble{
		Value: value,
	}
}

func (r *ConstantPoolReader) readConstantUtf8() ConstantInfo {
	size := r.byte_reader.readUInt16()
	bytes := r.byte_reader.readBytes(uint64(size))

	// convert bytes to string
	return ConstantUtf8{
		Str: string(bytes),
	}
}

func (r *ConstantPoolReader) readConstantString() ConstantInfo {
	index := r.byte_reader.readUInt16()

	return ConstantString{
		EntryIndex: index,
	}
}

func (r *ConstantPoolReader) readConstantClass() ConstantInfo {
	index := r.byte_reader.readUInt16()

	return ConstantClass{
		EntryIndex: index,
	}
}

func (r *ConstantPoolReader) readConstantNameAndType() ConstantInfo {
	name_index := r.byte_reader.readUInt16()
	descriptor_index := r.byte_reader.readUInt16()

	return ConstantNameAndType{
		NameEntryIndex:       name_index,
		DescriptorEntryIndex: descriptor_index,
	}
}

func (r *ConstantPoolReader) readConstantMemberRef() ConstantMemberRef {
	class_index := r.byte_reader.readUInt16()
	name_and_type_index := r.byte_reader.readUInt16()

	return ConstantMemberRef{
		ClassEntryIndex:       class_index,
		NameAndTypeEntryIndex: name_and_type_index,
	}
}

func (r *ConstantPoolReader) readConstantInfo() ConstantInfo {
	constant_type := r.readConstantType()
	switch constant_type {
	case CONSTANT_Class:
		return r.readConstantClass()
	case CONSTANT_Fieldref:
		return NewConstantFieldRef(r.readConstantMemberRef())
	case CONSTANT_Methodref:
		return NewConstantMethodRef(r.readConstantMemberRef())
	case CONSTANT_InferfaceMethodref:
		return NewConstantInterfaceMethodRef(r.readConstantMemberRef())
	case CONSTANT_String:
		return r.readConstantString()
	case CONSTANT_Integer:
		return r.readConstantInt()
	case CONSTANT_Float:
		return r.readConstantFloat()
	case CONSTANT_Long:
		return r.readConstantLong()
	case CONSTANT_Double:
		return r.readConstantDouble()
	case CONSTANT_NameAndType:
		return r.readConstantNameAndType()
	case CONSTANT_Utf8:
		return r.readConstantUtf8()
	case CONSTANT_MethodHandle:
		fallthrough
	case CONSTANT_MethodType:
		fallthrough
	case CONSTANT_InvokeDynamic:
		fallthrough
	default:
		r.byte_reader.errors = append(r.byte_reader.errors, fmt.Errorf("java.lang.ClassFormatError: invalid constant type %d", constant_type))
		return nil
	}
}

func (r *ConstantPoolReader) ReadConstantPoolInfos() *ConstantPoolReader {
	r.constant_pool_cnt = r.byte_reader.readUInt16()

	entries := make([]ConstantInfo, 0, r.constant_pool_cnt-1)

	// index 0 is not used
	entries = append(entries, nil)

	// only avaliable till cnt-1
	// [1, cnt)
	for i := 1; i != int(r.constant_pool_cnt); i++ {
		constant_info := r.readConstantInfo()
		entries = append(entries, &constant_info)

		// long and double take up two slots
		// append one empty slot for convenience
		switch constant_info.(type) {
		case ConstantLong, ConstantDouble:
			entries = append(entries, nil)
			i++
		}
	}

	r.constant_pool.Entries = entries
	return r
}

func (r *ConstantPoolReader) BuildConstantPool() (ConstantPoolInfo, []error) {
	return ConstantPoolInfo{}, r.byte_reader.errors
}

func NewConstantPoolReader(reader io.Reader) ConstantPoolReader {
	return ConstantPoolReader{
		byte_reader: NewByteReader(reader),
	}
}
