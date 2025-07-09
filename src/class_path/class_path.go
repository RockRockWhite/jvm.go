package class_path

import (
	"bytes"
	"fmt"
	"os"

	"github.com/RockRockWhite/jvm.go/src/class_file"
)

type ClassPath struct {
	defaultClassPath string
}

func NewClassPath(default_patch string) *ClassPath {
	return &ClassPath{
		defaultClassPath: default_patch,
	}
}

func (cp *ClassPath) GetClassFileReader(class_name string) (*class_file.ClassFileReader, error) {
	found_class_path := fmt.Sprintf("%s/%s", cp.defaultClassPath, class_name)

	data, err := os.ReadFile(found_class_path)
	if err != nil {
		return nil, err
	}

	return class_file.NewClassReader(bytes.NewReader(data)), nil
}
