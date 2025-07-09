package class_file

import (
	"fmt"
	"io"
)

type ClassFileReader struct {
	byte_reader   ByteReader
	magic         uint32
	minor_version uint16
	major_version uint16
	constant_pool ConstantPoolInfo
	access_flags  uint16
	this_class    string
	super_class   string
	interfaces    []string
	fields        []FieldInfo
	methods       []MethodInfo
	attributes    []AttributeInfo
}

func (r *ClassFileReader) ReadMagic() *ClassFileReader {
	r.magic = r.byte_reader.readUInt32()
	if r.magic != 0xcafebabe {
		r.byte_reader.errors = append(r.byte_reader.errors, fmt.Errorf("java.lang.ClassFormatError: invalid magic number"))
	}

	return r
}

func (r *ClassFileReader) ReadVersion() *ClassFileReader {
	r.minor_version = r.byte_reader.readUInt16()
	r.major_version = r.byte_reader.readUInt16()

	compatiable := (r.major_version >= 45) && (r.major_version <= 52)

	if !compatiable {
		r.byte_reader.errors = append(r.byte_reader.errors, fmt.Errorf("java.lang.UnsupportedClassVersionError"))
	}

	return r
}

func (r *ClassFileReader) ReadConstantPool() *ClassFileReader {
	constant_pool_reader := NewConstantPoolReader(r.byte_reader.reader)
	constant_pool, err := constant_pool_reader.
		ReadConstantPoolInfos().
		BuildConstantPool()

	r.constant_pool = constant_pool

	if err != nil {
		r.byte_reader.errors = append(r.byte_reader.errors, err)
	}

	return r
}

func (r *ClassFileReader) ReadAccessFlags() *ClassFileReader {
	r.access_flags = r.byte_reader.readUInt16()
	return r
}

func (r *ClassFileReader) ReadThisClass() *ClassFileReader {
	this_class_idx := r.byte_reader.readUInt16()
	class_name, err := r.constant_pool.GetClass(this_class_idx)

	if err != nil {
		r.byte_reader.errors = append(r.byte_reader.errors, err)
	}

	r.this_class = class_name
	return r
}

func (r *ClassFileReader) ReadSuperClass() *ClassFileReader {
	this_class_idx := r.byte_reader.readUInt16()
	class_name, err := r.constant_pool.GetClass(this_class_idx)
	if err != nil {
		r.byte_reader.errors = append(r.byte_reader.errors, err)
	}
	r.super_class = class_name
	return r
}

func (r *ClassFileReader) ReadInterfaces() *ClassFileReader {
	interfaces_count := r.byte_reader.readUInt16()
	interfaces := make([]string, 0, interfaces_count)

	for i := 0; i != int(interfaces_count); i++ {
		interface_idx := r.byte_reader.readUInt16()
		interface_name, err := r.constant_pool.GetClass(interface_idx)
		if err != nil {
			r.byte_reader.errors = append(r.byte_reader.errors, err)
		}

		interfaces = append(interfaces, interface_name)
	}
	r.interfaces = interfaces

	return r
}

func (r *ClassFileReader) readAttributes() []AttributeInfo {
	attributes_count := r.byte_reader.readUInt16()
	attributes := make([]AttributeInfo, 0, attributes_count)

	for i := 0; i != int(attributes_count); i++ {

		attribute_reader := NewAttributeReader(r.byte_reader.reader, r.constant_pool)
		attribute, err := attribute_reader.
			ReadAttributeName().
			ReadAttributeInfo().
			Build()

		if err != nil {
			r.byte_reader.errors = append(r.byte_reader.errors, err)
		}
		attributes = append(attributes, attribute)
	}

	return attributes
}

func (r *ClassFileReader) readMember() MemberInfo {

	access_flags := r.byte_reader.readUInt16()

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

	attributes := r.readAttributes()

	return MemberInfo{
		AccessFlags: access_flags,
		Name:        name,
		Descriptor:  descriptor,
		Attrubutes:  attributes,
	}
}

func (r *ClassFileReader) ReadFields() *ClassFileReader {
	fields_count := r.byte_reader.readUInt16()
	fields := make([]FieldInfo, 0, fields_count)

	for i := 0; i != int(fields_count); i++ {
		field := r.readMember()
		fields = append(fields, FieldInfo(field))
	}
	r.fields = fields

	return r
}

func (r *ClassFileReader) ReadMethods() *ClassFileReader {
	methods_count := r.byte_reader.readUInt16()
	methods := make([]MethodInfo, 0, methods_count)

	for i := 0; i != int(methods_count); i++ {
		method := r.readMember()
		methods = append(methods, MethodInfo(method))
	}
	r.methods = methods

	return r
}

func (r *ClassFileReader) ReadAttributes() *ClassFileReader {
	r.attributes = r.readAttributes()

	return r
}

func (r *ClassFileReader) BuildClass() (ClassInfo, error) {
	if len(r.byte_reader.errors) != 0 {
		// convet all errors to a single error
		error_msg := ""
		for _, err := range r.byte_reader.errors {
			error_msg += fmt.Sprintf("%v; ", err)
		}
		return ClassInfo{}, fmt.Errorf("java.lang.ClassFormatError: { %v }", error_msg)
	}

	return ClassInfo{
		Maigc:        r.magic,
		MinorVersion: r.minor_version,
		MajorVersion: r.major_version,
		ConstantPool: r.constant_pool,
		AccessFlags:  r.access_flags,
		ThisClass:    r.this_class,
		SuperClass:   r.super_class,
		Interfaces:   r.interfaces,
		Fields:       r.fields,
		Methods:      r.methods,
		Attributes:   r.attributes,
	}, nil
}

func NewClassReader(reader io.Reader) *ClassFileReader {
	return &ClassFileReader{
		byte_reader: NewByteReader(reader),
	}
}

func ReadClassInfo(reader io.Reader) (ClassInfo, error) {
	class_reader := NewClassReader(reader)
	return class_reader.
		ReadMagic().
		ReadVersion().
		ReadConstantPool().
		ReadAccessFlags().
		ReadThisClass().
		ReadSuperClass().
		ReadInterfaces().
		ReadFields().
		ReadMethods().
		ReadAttributes().
		BuildClass()
}
