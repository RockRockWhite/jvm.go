package class_file

import (
	"encoding/binary"
	"io"
)

type ByteFileReader struct {
	errors []error
	reader io.Reader
}

func (r *ByteFileReader) readUInt8() uint8 {
	var res uint8
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteFileReader) readUInt16() uint16 {
	var res uint16
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteFileReader) readUInt32() uint32 {
	var res uint32
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteFileReader) readUInt64() uint64 {
	var res uint64
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteFileReader) readInt8() int8 {
	var res int8
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteFileReader) readInt16() int16 {
	var res int16
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteFileReader) readInt32() int32 {
	var res int32
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteFileReader) readInt64() uint64 {
	var res uint64
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteFileReader) readFloat32() float32 {
	var res float32
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteFileReader) readFloat64() float64 {
	var res float64
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteFileReader) readBytes(size uint64) []byte {
	bytes := make([]byte, size)
	_, err := io.ReadFull(r.reader, bytes)
	r.errors = append(r.errors, err)

	return bytes
}
