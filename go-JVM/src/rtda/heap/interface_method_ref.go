package heap

import "classfile"

type InterfaceMethodRef struct {
	//cp             *ConstantPool
	//ownerClassName string
	//owner          *Klass // 属性属于谁

	SymRef

	name       string // 属性名称		field.name
	descriptor string // 属性描述符	java/lang/Object

	method *Method
}

func (recv *InterfaceMethodRef) init(cp *ConstantPool, fileCP classfile.ConstantPool, info classfile.ConstantInfo) *InterfaceMethodRef {

	methodInfo := info.(*classfile.ConstantInterfaceMethodRefInfo)
	recv.cp = cp
	recv.className = methodInfo.GetClassName()
	nameAndType := methodInfo.GetNameAndType()

	recv.descriptor = nameAndType.GetDescriptor()
	recv.name = nameAndType.GetName()

	return recv
}

func (recv *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if recv.method == nil {
		recv.resolvedInterfaceMethodRef()
	}
	return recv.method
}

func (recv *InterfaceMethodRef) resolvedInterfaceMethodRef() {
	d := recv.cp.klass        // 触发者
	c := recv.ResolvedClass() // 拥有者

	if !c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethodInInterfaces(c, recv.name, recv.descriptor)

	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	recv.method = method

}
func (recv *InterfaceMethodRef) GetName() string {

	return recv.name
}

func (recv *InterfaceMethodRef) GetDescriptor() string {

	return recv.descriptor
}
