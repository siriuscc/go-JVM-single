package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// 创建 引用类型数组
// 	后接 index u16 指向 rtCP的 一个class，array，interface
//		从opStack中pop一个 int count,
type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (recv *ANEW_ARRAY) Execute(frame *rtda.Frame) {

	// get count
	opStack := frame.OpStack()
	count := opStack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	// get class
	cp := frame.GetMethod().GetOwner().GetConstantPool()
	classRef := cp.GetConstant(uint(recv.Index)).(*heap.ClassRef)
	klass := classRef.ResolvedClass()

	// return [class
	arrRef := klass.CreateArrayObject(count)
	opStack.PushRef(arrRef)
}
