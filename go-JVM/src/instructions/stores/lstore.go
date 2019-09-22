package stores

import (
	"instructions/base"
	"rtda"
)

func storeLongFromIndex(frame *rtda.Frame, index uint16) {
	val := frame.OpStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

//Store int into local variable
type LSTORE struct {
	base.Index8Instruction
}

func (recv *LSTORE) Execute(frame *rtda.Frame) {

	storeLongFromIndex(frame, recv.Index)
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (recv *LSTORE_0) Execute(frame *rtda.Frame) {

	storeLongFromIndex(frame, 0)
}

/////////////////////////////////////////////////////////////
type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (recv *LSTORE_1) Execute(frame *rtda.Frame) {

	storeLongFromIndex(frame, 1)
}

/////////////////////////////////////////////////////////////

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (recv *LSTORE_2) Execute(frame *rtda.Frame) {

	storeLongFromIndex(frame, 2)
}

/////////////////////////////////////////////////////////////////
type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (recv *LSTORE_3) Execute(frame *rtda.Frame) {

	storeLongFromIndex(frame, 3)
}
