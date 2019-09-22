package loads

import (
	"instructions/base"
	"rtda"
)

func loadDoubleFromIndex(frame *rtda.Frame, index uint16) {
	value := frame.LocalVars().GetDouble(index)
	frame.OpStack().PushDouble(value)
}

type DLOAD struct {
	base.Index8Instruction
}

func (recv *DLOAD) Execute(frame *rtda.Frame) {
	loadDoubleFromIndex(frame, recv.Index)
}

/////////////////////////////////////////////////////////////

type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (recv *DLOAD_0) Execute(frame *rtda.Frame) {
	loadDoubleFromIndex(frame, 0)
}

////////////////////////////////////////////////////////////

type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (recv *DLOAD_1) Execute(frame *rtda.Frame) {
	loadDoubleFromIndex(frame, 1)
}

///////////////////////////////////////////////////////////
type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (recv *DLOAD_2) Execute(frame *rtda.Frame) {
	loadDoubleFromIndex(frame, 2)
}

///////////////////////////////////////////////////////////
type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (recv *DLOAD_3) Execute(frame *rtda.Frame) {
	loadDoubleFromIndex(frame, 3)
}
