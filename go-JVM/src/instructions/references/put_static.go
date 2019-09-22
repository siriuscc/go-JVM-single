package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// 后面接一个u16的索引，可以对应 运行时常量池的一个FieldRef
// 从opStack.pop 出 赋给静态变量的值
// 给类的某个静态变量赋值
type PUTSTATIC struct {
	base.Index16Instruction
}

func (recv *PUTSTATIC) Execute(frame *rtda.Frame) {

	currentMethod := frame.GetMethod()
	currentClass := currentMethod.GetOwner()
	rtCP := currentClass.GetConstantPool()

	fieldRef := rtCP.GetConstant(uint(recv.Index)).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	klass := field.GetOwner()
	if !klass.IsInitialized() {
		frame.RevertNextPC()
		base.InitClass(frame.GetThread(), klass)
		return
	}
	// 不是静态字段
	if !field.GetAccessFlags().IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.GetAccessFlags().IsFinal() {
		if currentClass != klass || currentMethod.GetName() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.GetDescriptor()
	slotId := field.GetSlotId()

	slots := klass.GetStaticVars()
	stack := frame.OpStack()

	switch descriptor[0] {
	case 'Z', 'B', 'S', 'C', 'I':
		slots.SetInt(slotId, stack.PopInt())
		break
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
		break
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
		break
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
		break
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
		break
	}
}
