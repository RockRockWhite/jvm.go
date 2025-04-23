package class_file

import (
	"fmt"
	"io"
)

type ClassReader struct {
	byte_reader   ByteReader
	magic         uint32
	minor_version uint16
	major_version uint16
	constant_pool ConstantPoolInfo
}

func (r *ClassReader) ReadMagic() *ClassReader {
	r.magic = r.byte_reader.readUInt32()
	if r.magic != 0xcafebabe {
		r.byte_reader.errors = append(r.byte_reader.errors, fmt.Errorf("java.lang.ClassFormatError: invalid magic number"))
	}

	return r
}

func (r *ClassReader) ReadVersion() *ClassReader {
	r.minor_version = r.byte_reader.readUInt16()
	r.major_version = r.byte_reader.readUInt16()

	compatiable := (r.major_version >= 45) && (r.major_version <= 52)

	if !compatiable {
		r.byte_reader.errors = append(r.byte_reader.errors, fmt.Errorf("java.lang.UnsupportedClassVersionError"))
	}

	return r
}

func (r *ClassReader) ReadConstantPool() *ClassReader {
	constant_pool_reader := NewConstantPoolReader(r.byte_reader.reader)

	constant_pool, errors := constant_pool_reader.ReadCount().ReadConstantPoolInfos().BuildConstantPool()

	r.constant_pool = constant_pool
	r.byte_reader.errors = append(r.byte_reader.errors, errors...)

	return r
}

func (r *ClassReader) BuildClass() (ClassInfo, error) {
	return ClassInfo{}, nil
}

func NewClassReader(reader io.Reader) ClassReader {
	return ClassReader{
		byte_reader: NewByteReader(reader),
	}
}

func ReadClassInfo(reader io.Reader) (ClassInfo, error) {
	class_reader := NewClassReader(reader)
	return class_reader.ReadMagic().ReadVersion().BuildClass()
}
