package heap

import "classfile"

type ClassRef struct {
	SymRef
}

func (recv *ClassRef) init(cp *ConstantPool, fileCP classfile.ConstantPool, info classfile.ConstantInfo) *ClassRef {

	classInfo := info.(*classfile.ConstantClassInfo)
	recv.className = classInfo.GetName()
	recv.cp = cp
	return recv
}
