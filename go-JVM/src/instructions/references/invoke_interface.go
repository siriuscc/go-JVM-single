package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// invoke interface method
type INVOKE_INTERFACE struct {
	index uint16 // 指向运行时常量池的 interfaceMethod
	count uint8  // 有n个参数
	zero  uint8
}

func (recv *INVOKE_INTERFACE) FetchOperands(reader *base.ByteCodeReader) {

	recv.index = reader.ReadUint16()
	recv.count = reader.ReadUint8()
	recv.zero = reader.ReadUint8()

}

func (recv *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {

	cp := frame.GetMethod().GetOwner().GetConstantPool()

	invokeMethodRef := cp.GetConstant(uint(recv.index)).(*heap.InterfaceMethodRef)
	invokeMethod := invokeMethodRef.ResolvedInterfaceMethod()

	if invokeMethod.GetAccessFlags().IsStatic() || invokeMethod.GetAccessFlags().IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	slotCount := invokeMethod.ArgSlotCount()

	this := frame.OpStack().GetTopRef(slotCount - 1)
	if this == nil {
		panic("java.lang.NullPointerException")
	}

	if !this.GetKlass().IsImplementsOf(invokeMethodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	realMethod := heap.LookupMethodInExtendsTree(this.GetKlass(), invokeMethodRef.GetName(), invokeMethodRef.GetDescriptor())

	if realMethod == nil || realMethod.GetAccessFlags().IsAbstract() {
		panic("java.lang.IllegalAccessError")
	}

	base.InvokeAMethod(frame, realMethod)
}
