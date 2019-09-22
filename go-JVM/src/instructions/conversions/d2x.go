package conversions

import (
	"instructions/base"
	"rtda"
)

type D2I struct {
	base.NoOperandsInstruction
}

func (recv *D2I) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopDouble()
	stack.PushInt(int32(value))
}

///////////////////////////////////////////////////

type D2F struct {
	base.NoOperandsInstruction
}

func (recv *D2F) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopDouble()
	stack.PushFloat(float32(value))
}

///////////////////////////////////////////////////

type D2L struct {
	base.NoOperandsInstruction
}

func (recv *D2L) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopDouble()
	stack.PushLong(int64(value))
}
