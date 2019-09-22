package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// 给实例变量赋值
// 	后面接一个运行时常量池索引，
// 	从操作数栈中依次pop出 变量值，对象引用
type PUTFIELD struct {
	base.Index16Instruction
}

func (recv *PUTFIELD) Execute(frame *rtda.Frame) {

	currentMethod := frame.GetMethod()
	ownerClass := currentMethod.GetOwner()
	cp := ownerClass.GetConstantPool()

	fieldRef := cp.GetConstant(uint(recv.Index)).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	slotId := field.GetSlotId()

	opStack := frame.OpStack()
	// 权限校验
	if field.GetAccessFlags().IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	if field.GetAccessFlags().IsFinal() {
		// 所属类不是当前类
		if ownerClass != field.GetOwner() || currentMethod.GetName() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.GetDescriptor()

	switch descriptor[0] {
	case 'Z', 'S', 'C', 'I', 'B':
		val := opStack.PopInt()
		ref := opStack.PopRef()

		if ref == nil {

			panic("java.lang.NullPointerException")
		}
		ref.GetFields().SetInt(slotId, val)
	case 'J':
		val := opStack.PopLong()
		ref := opStack.PopRef()

		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.GetFields().SetLong(slotId, val)
		break
	case 'D':
		val := opStack.PopDouble()
		ref := opStack.PopRef()

		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.GetFields().SetDouble(slotId, val)
		break
	case 'F':
		val := opStack.PopFloat()
		ref := opStack.PopRef()

		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.GetFields().SetFloat(slotId, val)
		break
	case 'L', '[':
		val := opStack.PopRef()
		ref := opStack.PopRef()

		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.GetFields().SetRef(slotId, val)
	}

}
