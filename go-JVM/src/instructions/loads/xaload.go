package loads

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// Load reference from array
type AALOAD struct{ base.NoOperandsInstruction }

// Load byte or boolean from array
type BALOAD struct{ base.NoOperandsInstruction }
type CALOAD struct{ base.NoOperandsInstruction }
type DALOAD struct{ base.NoOperandsInstruction }
type FALOAD struct{ base.NoOperandsInstruction }
type IALOAD struct{ base.NoOperandsInstruction }
type LALOAD struct{ base.NoOperandsInstruction }
type SALOAD struct{ base.NoOperandsInstruction }

func checkNotNil(ref *heap.OopDesc) {

	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(length int, index int32) {

	if index < 0 || index >= int32(length) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

func (recv *AALOAD) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	index := opStack.PopInt()
	arrRef := opStack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	opStack.PushRef(refs[index])
}

func (recv *BALOAD) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	index := opStack.PopInt()
	arrRef := opStack.PopRef()
	checkNotNil(arrRef)
	arr := arrRef.Bytes()
	checkIndex(len(arr), index)
	opStack.PushInt(int32(arr[index])) // 扩展为 有符号的
}

func (recv *CALOAD) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	index := opStack.PopInt()
	arrRef := opStack.PopRef()
	checkNotNil(arrRef)
	arr := arrRef.Chars()
	checkIndex(len(arr), index)
	opStack.PushInt(int32(arr[index])) // 扩展为 有符号的
}

func (recv *IALOAD) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	index := opStack.PopInt()
	arrRef := opStack.PopRef()
	checkNotNil(arrRef)
	arr := arrRef.Ints()
	checkIndex(len(arr), index)
	opStack.PushInt(arr[index]) // 扩展为 有符号的
}

func (recv *DALOAD) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	index := opStack.PopInt()
	arrRef := opStack.PopRef()
	checkNotNil(arrRef)
	arr := arrRef.Doubles()
	checkIndex(len(arr), index)
	opStack.PushDouble(arr[index]) // 扩展为 有符号的
}

func (recv *FALOAD) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	index := opStack.PopInt()
	arrRef := opStack.PopRef()
	checkNotNil(arrRef)
	arr := arrRef.Floats()
	checkIndex(len(arr), index)
	opStack.PushFloat(arr[index]) // 扩展为 有符号的
}

func (recv *SALOAD) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	index := opStack.PopInt()
	arrRef := opStack.PopRef()
	checkNotNil(arrRef)
	arr := arrRef.Shorts()
	checkIndex(len(arr), index)
	opStack.PushInt(int32(arr[index])) // 扩展为 有符号的
}

func (recv *LALOAD) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	index := opStack.PopInt()
	arrRef := opStack.PopRef()
	checkNotNil(arrRef)
	arr := arrRef.Longs()
	checkIndex(len(arr), index)
	opStack.PushLong(arr[index]) // 扩展为 有符号的
}
