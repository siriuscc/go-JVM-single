package loads

import (
	"instructions/base"
	"rtda"
)

// 从局部变量表加载到操作数栈
func loadIntFromIndex(frame *rtda.Frame, index uint16) {
	value := frame.LocalVars().GetInt(index)
	frame.OpStack().PushInt(value)
}

type ILOAD struct {
	base.Index8Instruction
}

func (recv *ILOAD) Execute(frame *rtda.Frame) {
	loadIntFromIndex(frame, recv.Index)
}

/////////////////////////////////////////////////////////////

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (recv *ILOAD_0) Execute(frame *rtda.Frame) {
	loadIntFromIndex(frame, 0)
}

////////////////////////////////////////////////////////////

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (recv *ILOAD_1) Execute(frame *rtda.Frame) {
	loadIntFromIndex(frame, 1)
}

///////////////////////////////////////////////////////////
type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (recv *ILOAD_2) Execute(frame *rtda.Frame) {
	loadIntFromIndex(frame, 2)
}

///////////////////////////////////////////////////////////
type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (recv *ILOAD_3) Execute(frame *rtda.Frame) {
	loadIntFromIndex(frame, 3)
}
