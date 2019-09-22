package heap

import "classfile"

type FieldRef struct {
	SymRef

	//className string		// 属性属于谁
	//oopType          *Klass // 属性属于谁
	name       string // 属性名称		field.name
	descriptor string // 属性描述符	java/lang/Object
	field      *Field
}

func (recv *FieldRef) ResolvedField() *Field {

	if recv.field == nil {
		recv.resolveFieldRef()
	}

	return recv.field
}

func (recv *FieldRef) resolveFieldRef() {

	d := recv.cp.klass
	c := recv.ResolvedClass()
	field := lookupField(c, recv.name, recv.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}

	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	recv.field = field
}

func lookupField(klass *Klass, name string, descriptor string) *Field {

	if klass == nil {
		return nil
	}

	// 在本类中查找
	for _, field := range klass.fields {

		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}

	// 在接口中查找

	for _, ifield := range klass.interfaces {
		if field := lookupField(ifield, name, descriptor); field != nil {
			return field
		}
	}

	return lookupField(klass.super, name, descriptor)
}

func (recv *FieldRef) init(cp *ConstantPool, fileCP classfile.ConstantPool, info classfile.ConstantInfo) *FieldRef {

	fieldInfo := info.(*classfile.ConstantFieldRefInfo)

	recv.cp = cp
	recv.className = fieldInfo.GetClassName()
	nameAndType := fieldInfo.GetNameAndType()

	recv.descriptor = nameAndType.GetDescriptor()
	recv.name = nameAndType.GetName()

	return recv
}
