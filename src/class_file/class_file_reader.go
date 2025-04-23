package class_file

type ClassFileReader struct {
	byte_reader ByteFileReader
}

func (r *ClassFileReader) BuildClass() (ClassInfo, error) {
	return ClassInfo{}, nil
}
