package heap

import "github.com/RockRockWhite/jvm.go/src/class_path"

type ClassLoader struct {
	classPath *class_path.ClassPath
	classMap  map[string]*Class
}

func (cl *ClassLoader) loadNonArrayClass(name string) (*Class, error) {
	// get class file reader from class path
	class_file_reader, err := cl.classPath.GetClassFileReader(name)
	if err != nil {
		return nil, err
	}

	// read class file
	class_file, err := class_file_reader.ReadClassFile()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (cl *ClassLoader) LoadClass(name string) (*Class, error) {
	// if a class is already loaded, return it
	if class, ok := cl.classMap[name]; ok {
		return class, nil
	}

	return cl.loadNonArrayClass(name)
}
