package stores

import (
	"instructions/base"
	"rtda"
)

func storeRefFromIndex(frame *rtda.Frame, index uint16) {
	val := frame.OpStack().PopRef()
	frame.LocalVars().SetRef(uint16(index), val)
}

//Store int into local variable
type ASTORE struct {
	base.Index8Instruction
}

func (recv *ASTORE) Execute(frame *rtda.Frame) {
	storeRefFromIndex(frame, recv.Index)
}

type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (recv *ASTORE_0) Execute(frame *rtda.Frame) {
	storeRefFromIndex(frame, 0)
}

/////////////////////////////////////////////////////////////
type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (recv *ASTORE_1) Execute(frame *rtda.Frame) {
	storeRefFromIndex(frame, 1)
}

/////////////////////////////////////////////////////////////

type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (recv *ASTORE_2) Execute(frame *rtda.Frame) {
	storeRefFromIndex(frame, 2)
}

/////////////////////////////////////////////////////////////////
type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (recv *ASTORE_3) Execute(frame *rtda.Frame) {
	storeRefFromIndex(frame, 3)
}
