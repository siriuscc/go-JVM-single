package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// 根据 u16 的索引，在运行时常量池找到对应的ClassRef符号引用
// 然后加载对应的类，用对应的类创建实例
// 将实例ref push到栈顶
type NEW struct {
	base.Index16Instruction
}

func (recv *NEW) Execute(frame *rtda.Frame) {

	rtCP := frame.GetMethod().GetOwner().GetConstantPool()

	classRef := rtCP.GetConstant(uint(recv.Index)).(*heap.ClassRef)

	klass := classRef.ResolvedClass()
	if !klass.IsInitialized() {
		frame.RevertNextPC()
		base.InitClass(frame.GetThread(), klass)
		return
	}

	if klass.IsInterface() || klass.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := klass.CreateObject()
	frame.OpStack().PushRef(ref)
}
