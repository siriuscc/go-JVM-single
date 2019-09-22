package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

// Create new multidimensional array
//
type MULTI_ANEW_ARRAY struct {
	index      uint16 // 指向一个ClassRef,对应多维数组类
	dimensions uint8  // 数组维度
}

// 从opStack 中弹出n个数，分别表示每个维度的数组长度

func (recv *MULTI_ANEW_ARRAY) FetchOperands(reader *base.ByteCodeReader) {

	recv.index = reader.ReadUint16()
	recv.dimensions = reader.ReadUint8()
}

func (recv *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	cp := frame.GetMethod().GetOwner().GetConstantPool()

	classRef := cp.GetConstant(uint(recv.index)).(*heap.ClassRef)
	klass := classRef.ResolvedClass()

	counts := popAndCheckDimensions(opStack, recv.dimensions)

	obj := klass.CreateMultiArrayObject(counts)

	opStack.PushRef(obj)
}

func popAndCheckDimensions(opStack *rtda.OperandStack, count uint8) []int32 {

	dis := make([]int32, count)

	for i := int(count) - 1; i >= 0; i-- {
		dis[i] = opStack.PopInt()

		if dis[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return dis
}
