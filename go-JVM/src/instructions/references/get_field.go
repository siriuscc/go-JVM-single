package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// 得到实例变量
// 	后面接一个运行时常量池索引 index
// 	从操作数栈中依次pop出 变量值 objRef
type GETFIELD struct {
	base.Index16Instruction
}

func (recv *GETFIELD) Execute(frame *rtda.Frame) {

	currentMethod := frame.GetMethod()
	currentClass := currentMethod.GetOwner()
	cp := currentClass.GetConstantPool()

	fieldRef := cp.GetConstant(uint(recv.Index)).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	// 权限校验
	if field.GetAccessFlags().IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//if field.GetAccessFlags().IsFinal() {
	//	// 所属类不是当前类
	//	if currentClass != field.GetOwner() || currentMethod.GetName() != "<init>" {
	//		panic("java.lang.IllegalAccessError")
	//	}
	//}

	slotId := field.GetSlotId()

	opStack := frame.OpStack()
	descriptor := field.GetDescriptor()

	ref := opStack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	fieldSlots := ref.GetFields()

	switch descriptor[0] {
	case 'Z', 'S', 'C', 'I', 'B':
		opStack.PushInt(fieldSlots.GetInt(slotId))

	case 'J':
		opStack.PushLong(fieldSlots.GetLong(slotId))

	case 'D':
		opStack.PushDouble(fieldSlots.GetDouble(slotId))

	case 'F':
		opStack.PushFloat(fieldSlots.GetFloat(slotId))
		break
	case 'L', '[':
		opStack.PushRef(fieldSlots.GetRef(slotId))

	}

}
