package heap

type ClassPath struct {
}

type ClassLoader struct {
	classPath ClassPath
	classMap  map[string]*Class
}

func (cl *ClassLoader) loadNonArrayClass(name string) *Class {
	return nil
}

func (cl *ClassLoader) LoadClass(name string) *Class {
	// if a class is already loaded, return it
	if class, ok := cl.classMap[name]; ok {
		return class
	}

	return cl.loadNonArrayClass(name)
}
