package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// a instanceof A 判断a是不是A的实例
//    后面跟着一个u16 的 index指向一个ClassRef
//    操作数栈中 pop 一个ref
//	    如果ref==nil，pushInt(0)
type INSTANCE_OF struct {
	base.Index16Instruction
}

func (recv *INSTANCE_OF) Execute(frame *rtda.Frame) {

	currentMethod := frame.GetMethod()
	currentClass := currentMethod.GetOwner()
	cp := currentClass.GetConstantPool()

	opStack := frame.OpStack()
	instance := opStack.PopRef()

	if instance == nil {
		opStack.PushInt(0)
		return
	}

	classRef := cp.GetConstant(uint(recv.Index)).(*heap.ClassRef)
	klass := classRef.ResolvedClass()
	if instance.IsInstanceOf(klass) {
		opStack.PushInt(1)
	} else {
		opStack.PushInt(0)
	}

}
