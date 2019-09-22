package loads

import (
	"instructions/base"
	"rtda"
)

func loadLongFromIndex(frame *rtda.Frame, index uint16) {
	value := frame.LocalVars().GetLong(index)
	frame.OpStack().PushLong(value)
}

type LLOAD struct {
	base.Index8Instruction
}

func (recv *LLOAD) Execute(frame *rtda.Frame) {
	loadLongFromIndex(frame, recv.Index)
}

/////////////////////////////////////////////////////////////

type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (recv *LLOAD_0) Execute(frame *rtda.Frame) {
	loadLongFromIndex(frame, 0)
}

////////////////////////////////////////////////////////////

type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (recv *LLOAD_1) Execute(frame *rtda.Frame) {
	loadLongFromIndex(frame, 1)
}

///////////////////////////////////////////////////////////
type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (recv *LLOAD_2) Execute(frame *rtda.Frame) {
	loadLongFromIndex(frame, 2)
}

///////////////////////////////////////////////////////////
type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (recv *LLOAD_3) Execute(frame *rtda.Frame) {
	loadLongFromIndex(frame, 3)
}
