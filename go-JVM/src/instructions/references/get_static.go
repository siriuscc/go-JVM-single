package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// 得到类的静态变量的 值，push到opStack
type GETSTATIC struct {
	base.Index16Instruction
}

func (recv *GETSTATIC) Execute(frame *rtda.Frame) {

	cp := frame.GetMethod().GetOwner().GetConstantPool()
	fieldRef := cp.GetConstant(uint(recv.Index)).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	ownerClass := field.GetOwner()

	if !ownerClass.IsInitialized() {
		frame.RevertNextPC()
		base.InitClass(frame.GetThread(), ownerClass)
		return
	}

	if !field.GetAccessFlags().IsStatic() {
		panic("java.lang.IncompatibleChangeError")
	}

	descriptor := field.GetDescriptor()
	slotId := field.GetSlotId()
	slots := ownerClass.GetStaticVars()
	stack := frame.OpStack()
	switch descriptor[0] {
	case 'Z', 'S', 'B', 'C', 'I':
		stack.PushInt(slots.GetInt(slotId))
		break
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
		break
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
		break
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
		break
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}
