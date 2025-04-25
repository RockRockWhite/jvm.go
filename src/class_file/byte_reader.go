package class_file

import (
	"encoding/binary"
	"io"
)

type ByteReader struct {
	errors []error
	reader io.Reader
}

func (r *ByteReader) readUInt8() uint8 {
	var res uint8
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteReader) readUInt16() uint16 {
	var res uint16
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteReader) readUInt32() uint32 {
	var res uint32
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteReader) readUInt64() uint64 {
	var res uint64
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteReader) readInt8() int8 {
	var res int8
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteReader) readInt16() int16 {
	var res int16
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteReader) readInt32() int32 {
	var res int32
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteReader) readInt64() int64 {
	var res int64
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteReader) readFloat32() float32 {
	var res float32
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteReader) readFloat64() float64 {
	var res float64
	err := binary.Read(r.reader, binary.BigEndian, &res)
	r.errors = append(r.errors, err)

	return res
}

func (r *ByteReader) readBytes(size uint64) []byte {
	bytes := make([]byte, size)
	_, err := io.ReadFull(r.reader, bytes)
	r.errors = append(r.errors, err)

	return bytes
}

func NewByteReader(reader io.Reader) ByteReader {
	return ByteReader{
		errors: make([]error, 0),
		reader: reader,
	}
}
