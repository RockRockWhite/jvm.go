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

	return res, nil
}
