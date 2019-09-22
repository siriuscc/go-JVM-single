package math

import (
	"instructions/base"
	"rtda"
)

type IAND struct {
	base.NoOperandsInstruction
}

func (recv *IAND) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	v2 := stack.PopInt()
	v1 := stack.PopInt()

	stack.PushInt(v2 & v1)
}

//////////////////////////////////////////////////

type LAND struct {
	base.NoOperandsInstruction
}

func (recv *LAND) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	v2 := stack.PopLong()
	v1 := stack.PopLong()

	stack.PushLong(v2 & v1)
}
