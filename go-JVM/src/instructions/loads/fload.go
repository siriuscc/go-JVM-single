package loads

import (
	"instructions/base"
	"rtda"
)

func loadFloatFromIndex(frame *rtda.Frame, index uint16) {
	value := frame.LocalVars().GetFloat(index)
	frame.OpStack().PushFloat(value)
}

type FLOAD struct {
	base.Index8Instruction
}

func (recv *FLOAD) Execute(frame *rtda.Frame) {
	loadFloatFromIndex(frame, recv.Index)
}

/////////////////////////////////////////////////////////////

type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (recv *FLOAD_0) Execute(frame *rtda.Frame) {
	loadFloatFromIndex(frame, 0)
}

////////////////////////////////////////////////////////////

type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (recv *FLOAD_1) Execute(frame *rtda.Frame) {
	loadFloatFromIndex(frame, 1)
}

///////////////////////////////////////////////////////////
type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (recv *FLOAD_2) Execute(frame *rtda.Frame) {
	loadFloatFromIndex(frame, 2)
}

///////////////////////////////////////////////////////////
type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (recv *FLOAD_3) Execute(frame *rtda.Frame) {
	loadFloatFromIndex(frame, 3)
}
