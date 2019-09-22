package math

import (
	"instructions/base"
	"rtda"
)

type IOR struct {
	base.NoOperandsInstruction
}

func (recv *IOR) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	v2 := stack.PopInt()
	v1 := stack.PopInt()

	stack.PushInt(v2 | v1)
}

///////////////////////////////////////////////////////

type LOR struct {
	base.NoOperandsInstruction
}

func (recv *LOR) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	v2 := stack.PopLong()
	v1 := stack.PopLong()

	stack.PushLong(v2 | v1)
}

///////////////////////////////////////////////////////

type IXOR struct {
	base.NoOperandsInstruction
}

func (recv *IXOR) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	v2 := stack.PopInt()
	v1 := stack.PopInt()

	stack.PushInt(v2 ^ v1)
}

///////////////////////////////////////////////////////
type LXOR struct {
	base.NoOperandsInstruction
}

func (recv *LXOR) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	v2 := stack.PopLong()
	v1 := stack.PopLong()

	stack.PushLong(v2 ^ v1)
}
