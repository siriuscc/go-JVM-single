package stores

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// Store into reference array
type AASTORE struct{ base.NoOperandsInstruction }

type BASTORE struct{ base.NoOperandsInstruction }
type CASTORE struct{ base.NoOperandsInstruction }
type DASTORE struct{ base.NoOperandsInstruction }
type FASTORE struct{ base.NoOperandsInstruction }
type IASTORE struct{ base.NoOperandsInstruction }
type LASTORE struct{ base.NoOperandsInstruction }
type SASTORE struct{ base.NoOperandsInstruction }

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

func (recv *AASTORE) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value := opStack.PopRef()
	index := opStack.PopInt()
	arrayRef := opStack.PopRef()

	checkNotNil(arrayRef)
	arr := arrayRef.Refs()

	checkIndex(len(arr), index)
	arr[index] = value
}

func (recv *LASTORE) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value := opStack.PopLong()
	index := opStack.PopInt()
	arrayRef := opStack.PopRef()

	checkNotNil(arrayRef)
	arr := arrayRef.Longs()

	checkIndex(len(arr), index)
	arr[index] = value
}

func (recv *FASTORE) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value := opStack.PopFloat()
	index := opStack.PopInt()
	arrayRef := opStack.PopRef()

	checkNotNil(arrayRef)
	arr := arrayRef.Floats()

	checkIndex(len(arr), index)
	arr[index] = value
}

func (recv *CASTORE) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value := opStack.PopInt()
	index := opStack.PopInt()
	arrayRef := opStack.PopRef()

	checkNotNil(arrayRef)
	arr := arrayRef.Chars()

	checkIndex(len(arr), index)
	arr[index] = uint16(value)
}

func (recv *BASTORE) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value := opStack.PopInt()
	index := opStack.PopInt()
	arrayRef := opStack.PopRef()

	checkNotNil(arrayRef)
	arr := arrayRef.Bytes()

	checkIndex(len(arr), index)
	arr[index] = int8(value)
}

func (recv *DASTORE) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value := opStack.PopDouble()
	index := opStack.PopInt()
	arrayRef := opStack.PopRef()

	checkNotNil(arrayRef)
	arr := arrayRef.Doubles()

	checkIndex(len(arr), index)
	arr[index] = value
}

func (recv *IASTORE) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value := opStack.PopInt()
	index := opStack.PopInt()
	arrayRef := opStack.PopRef()

	checkNotNil(arrayRef)
	arr := arrayRef.Ints()

	checkIndex(len(arr), index)
	arr[index] = value
}

func (recv *SASTORE) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value := opStack.PopInt()
	index := opStack.PopInt()
	arrayRef := opStack.PopRef()

	checkNotNil(arrayRef)
	arr := arrayRef.Shorts()

	checkIndex(len(arr), index)
	arr[index] = int16(value)
}
