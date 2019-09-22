package stores

import (
	"instructions/base"
	"rtda"
)

func storeDoubleFromIndex(frame *rtda.Frame, index uint16) {
	val := frame.OpStack().PopDouble()
	frame.LocalVars().SetDouble(uint16(index), val)
}

//Store int into local variable
type DSTORE struct {
	base.Index8Instruction
}

func (recv *DSTORE) Execute(frame *rtda.Frame) {
	storeDoubleFromIndex(frame, recv.Index)
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (recv *DSTORE_0) Execute(frame *rtda.Frame) {
	storeDoubleFromIndex(frame, 0)
}

/////////////////////////////////////////////////////////////
type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (recv *DSTORE_1) Execute(frame *rtda.Frame) {
	storeDoubleFromIndex(frame, 1)
}

/////////////////////////////////////////////////////////////

type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (recv *DSTORE_2) Execute(frame *rtda.Frame) {
	storeDoubleFromIndex(frame, 2)
}

/////////////////////////////////////////////////////////////////
type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (recv *DSTORE_3) Execute(frame *rtda.Frame) {
	storeDoubleFromIndex(frame, 3)
}
