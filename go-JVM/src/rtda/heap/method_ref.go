package heap

import "classfile"

type MethodRef struct {
	SymRef

	name       string // 属性名称		field.name
	descriptor string // 属性描述符	java/lang/Object

	method *Method
}

func (recv *MethodRef) init(cp *ConstantPool, fileCP classfile.ConstantPool, info classfile.ConstantInfo) *MethodRef {

	methodInfo := info.(*classfile.ConstantMethodRefInfo)

	recv.cp = cp
	recv.className = methodInfo.GetClassName()
	nameAndType := methodInfo.GetNameAndType()

	recv.descriptor = nameAndType.GetDescriptor()
	recv.name = nameAndType.GetName()

	return recv
}

func (recv *MethodRef) GetName() string {

	return recv.name
}

func (recv *MethodRef) GetDescriptor() string {

	return recv.descriptor
}

func (recv *MethodRef) ResolveMethod() *Method {

	if recv.method == nil {
		recv.resolveMethod()
	}

	return recv.method
}

func (recv *MethodRef) resolveMethod() {

	d := recv.cp.klass
	c := recv.ResolvedClass() // d 触发了c

	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(c, recv.name, recv.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	recv.method = method
}

// 先从类继承结构查找
// 再从接口中查找
func lookupMethod(klass *Klass, name string, descriptor string) *Method {

	method := LookupMethodInExtendsTree(klass, name, descriptor)

	if method == nil {
		method = lookupMethodInInterfaces(klass, name, descriptor)
	}

	return method
}

func lookupMethodInInterfaces(klass *Klass, name string, descriptor string) *Method {

	if klass == nil {
		return nil
	}

	for _, method := range klass.methods {
		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}

	for _, iface := range klass.interfaces {

		method := lookupMethodInInterfaces(iface, name, descriptor)
		if method != nil {
			return method
		}
	}

	return nil
}

// 在继承树上查找方法
func LookupMethodInExtendsTree(klass *Klass, name string, descriptor string) *Method {

	if klass == nil {
		return nil
	}

	// 本类中查找
	for _, method := range klass.methods {

		if method.name == name && method.descriptor == descriptor {
			return method
		}
	}
	return LookupMethodInExtendsTree(klass.super, name, descriptor)
}
