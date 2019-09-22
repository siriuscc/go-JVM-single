package loads

import (
	"instructions/base"
	"rtda"
)

func loadRefFromIndex(frame *rtda.Frame, index uint16) {
	value := frame.LocalVars().GetRef(index)
	frame.OpStack().PushRef(value)
}

type ALOAD struct {
	base.Index8Instruction
}

func (recv *ALOAD) Execute(frame *rtda.Frame) {
	loadRefFromIndex(frame, recv.Index)
}

/////////////////////////////////////////////////////////////

type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (recv *ALOAD_0) Execute(frame *rtda.Frame) {
	loadRefFromIndex(frame, 0)
}

////////////////////////////////////////////////////////////

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (recv *ALOAD_1) Execute(frame *rtda.Frame) {
	loadRefFromIndex(frame, 1)
}

///////////////////////////////////////////////////////////
type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (recv *ALOAD_2) Execute(frame *rtda.Frame) {
	loadRefFromIndex(frame, 2)
}

///////////////////////////////////////////////////////////
type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (recv *ALOAD_3) Execute(frame *rtda.Frame) {
	loadRefFromIndex(frame, 3)
}
