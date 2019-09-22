package references

import (
	"instructions/base"
	"rtda"
	"rtda/heap"
)

const (
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

// atype 要创建的哪种类型的指针
// count, 数组长度，从opStack 中弹出
type NEW_ARRAY struct {
	atype uint8
}

func (recv *NEW_ARRAY) FetchOperands(reader *base.ByteCodeReader) {
	recv.atype = reader.ReadUint8()
}

func (recv *NEW_ARRAY) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	count := opStack.PopInt()

	// java数组长为0是可以的。
	if count < 0 {
		panic("java.lang.NegativeArrayException")
	}

	classLoader := frame.GetMethod().GetOwner().GetClassLoader()
	arrClass := loadDetailArray(classLoader, recv.atype)

	arr := arrClass.CreateBaseArrayClass(count)
	opStack.PushRef(arr)

}

func loadDetailArray(classLoader *heap.ClassLoader, atype uint8) *heap.Klass {

	switch atype {

	case AT_BOOLEAN:
		return classLoader.LoadClass("[Z")
	case AT_CHAR:
		return classLoader.LoadClass("[C")
	case AT_FLOAT:
		return classLoader.LoadClass("[F")
	case AT_DOUBLE:
		return classLoader.LoadClass("[D")
	case AT_BYTE:
		return classLoader.LoadClass("[B")
	case AT_SHORT:
		return classLoader.LoadClass("[S")
	case AT_INT:
		return classLoader.LoadClass("[I")
	case AT_LONG:
		return classLoader.LoadClass("[J")
	default:
		panic("Invalid atype!")

	}

}
