package class_file

import (
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

func (r *ConstantPoolReader) readConstantInfo() ConstantInfo {
	// // read constant type tag
	// tag, err := cf.readUInt8()
	// if err != nil {
	// 	return nil, err
	// }

	// constant_type := ConstantType(tag)
	// switch constant_type {
	// case CONSTANT_Class:
	// 	return cf.readConstantClass()
	// case CONSTANT_Fieldref:
	// 	if constant_member_ref, err := cf.readConstantMemberRef(); err != nil {
	// 		return nil, err
	// 	} else {
	// 		return NewConstantFieldRef(constant_member_ref), nil
	// 	}
	// case CONSTANT_Methodref:
	// 	if constant_member_ref, err := cf.readConstantMemberRef(); err != nil {
	// 		return nil, err
	// 	} else {
	// 		return NewConstantMethodRef(constant_member_ref), nil
	// 	}
	// case CONSTANT_InferfaceMethodref:
	// 	if constant_member_ref, err := cf.readConstantMemberRef(); err != nil {
	// 		return nil, err
	// 	} else {
	// 		return NewConstantInterfaceMethodRef(constant_member_ref), nil
	// 	}
	// case CONSTANT_String:
	// 	return cf.readConstantString()
	// case CONSTANT_Integer:
	// 	return cf.readConstantInt()
	// case CONSTANT_Float:
	// 	return cf.readConstantFloat()
	// case CONSTANT_Long:
	// 	return cf.readConstantLong()
	// case CONSTANT_Double:
	// 	return cf.readConstantDouble()
	// case CONSTANT_NameAndType:
	// 	return cf.readConstantNameAndType()
	// case CONSTANT_Utf8:
	// 	return cf.readConstantUtf8()
	// case CONSTANT_MethodHandle:
	// case CONSTANT_MethodType:
	// case CONSTANT_InvokeDynamic:
	// default:
	// 	return nil, fmt.Errorf("java.lang.ClassFormatError: invalid constant type %d", tag)
	// }

	return nil
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
		entries = append(entries, constant_info)

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
