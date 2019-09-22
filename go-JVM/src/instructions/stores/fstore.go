package stores

import (
	"instructions/base"
	"rtda"
)

func storeFloatFromIndex(frame *rtda.Frame, index uint16) {
	val := frame.OpStack().PopFloat()
	frame.LocalVars().SetFloat(uint16(index), val)
}

//Store int into local variable
type FSTORE struct {
	base.Index8Instruction
}

func (recv *FSTORE) Execute(frame *rtda.Frame) {
	storeFloatFromIndex(frame, recv.Index)
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (recv *FSTORE_0) Execute(frame *rtda.Frame) {
	storeFloatFromIndex(frame, 0)
}

/////////////////////////////////////////////////////////////
type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (recv *FSTORE_1) Execute(frame *rtda.Frame) {
	storeFloatFromIndex(frame, 1)
}

/////////////////////////////////////////////////////////////

type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (recv *FSTORE_2) Execute(frame *rtda.Frame) {
	storeFloatFromIndex(frame, 2)
}

/////////////////////////////////////////////////////////////////
type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (recv *FSTORE_3) Execute(frame *rtda.Frame) {
	storeFloatFromIndex(frame, 3)
}
