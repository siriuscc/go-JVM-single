package stores

import (
	"instructions/base"
	"rtda"
)

func storeIntFromIndex(frame *rtda.Frame, index uint16) {
	val := frame.OpStack().PopInt()
	frame.LocalVars().SetInt(uint16(index), val)
}

//Store int into local variable
type ISTORE struct {
	base.Index8Instruction
}

// 把栈顶元素存储到 index
func (recv *ISTORE) Execute(frame *rtda.Frame) {

	storeIntFromIndex(frame, recv.Index)
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (recv *ISTORE_0) Execute(frame *rtda.Frame) {

	storeIntFromIndex(frame, 0)
}

/////////////////////////////////////////////////////////////
type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (recv *ISTORE_1) Execute(frame *rtda.Frame) {

	storeIntFromIndex(frame, 1)
}

/////////////////////////////////////////////////////////////

type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (recv *ISTORE_2) Execute(frame *rtda.Frame) {

	storeIntFromIndex(frame, 2)
}

/////////////////////////////////////////////////////////////////
type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (recv *ISTORE_3) Execute(frame *rtda.Frame) {

	storeIntFromIndex(frame, 3)
}
