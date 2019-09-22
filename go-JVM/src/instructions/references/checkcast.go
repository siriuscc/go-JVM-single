package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// 检查是否能进行转换
//    (A)a
//    后面跟一个 u16 index， 指向 run-time cp 中的一个ClassRef
//    如果 能转换，栈保持不变，
//    如果不能转换，抛出异常
type CHECKCAST struct {
	base.Index16Instruction
}

func (recv *CHECKCAST) Execute(frame *rtda.Frame) {

	currentMethod := frame.GetMethod()
	currentClass := currentMethod.GetOwner()
	cp := currentClass.GetConstantPool()

	opStack := frame.OpStack()
	instance := opStack.TopRef()

	if instance == nil {
		return
	}

	classRef := cp.GetConstant(uint(recv.Index)).(*heap.ClassRef)
	klass := classRef.ResolvedClass()
	// todo
	if !instance.IsInstanceOf(klass) {

		panic("java.lang.ClassCastException")
	}
}
