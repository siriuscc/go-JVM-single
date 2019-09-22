package constants

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// 从运行时常量池加载 常量值，并将其推入opStack.
// index 指向一个 int,float, ref,string字面量，或者符号引用（for class,method,method handle）
// LDC 和LDC_W 的区别仅仅是后者的index是u16而已，他们都只能操作一个槽位
// LDC2_W 能操作两个槽位，也就是double和long
type LDC struct{ base.Index8Instruction }
type LDC_W struct{ base.Index16Instruction }
type LDC2_W struct{ base.Index16Instruction }

func ldc(frame *rtda.Frame, index uint) {

	opStack := frame.OpStack()

	currentMethod := frame.GetMethod()
	currentClass := currentMethod.GetOwner()
	cp := currentClass.GetConstantPool()

	constant := cp.GetConstant(index)

	switch constant.(type) {
	case int32:
		opStack.PushInt(constant.(int32))
	case float32:
		opStack.PushFloat(constant.(float32))
	case string:
		strRef := heap.JString(currentClass.GetClassLoader(), constant.(string))
		opStack.PushRef(strRef)
	case *heap.ClassRef:
		klassRef := constant.(*heap.ClassRef)
		oop := klassRef.ResolvedClass().GetJavaMirror()
		opStack.PushRef(oop)
	default:
		panic("todo : string, classRef LDC")
	}

}
func (recv *LDC) Execute(frame *rtda.Frame) {

	ldc(frame, uint(recv.Index))
}

func (recv *LDC_W) Execute(frame *rtda.Frame) {

	ldc(frame, uint(recv.Index))
}

func (recv *LDC2_W) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()

	currentMethod := frame.GetMethod()
	currentClass := currentMethod.GetOwner()
	cp := currentClass.GetConstantPool()

	constant := cp.GetConstant(uint(recv.Index))

	switch constant.(type) {
	case int64:
		opStack.PushLong(constant.(int64))
	case float64:
		opStack.PushDouble(constant.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}

}
