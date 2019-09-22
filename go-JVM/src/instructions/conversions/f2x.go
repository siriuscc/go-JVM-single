package conversions

import (
	"instructions/base"
	"rtda"
)

type F2I struct {
	base.NoOperandsInstruction
}

func (recv *F2I) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopFloat()
	stack.PushInt(int32(value))
}

///////////////////////////////////////////

type F2L struct {
	base.NoOperandsInstruction
}

func (recv *F2L) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopFloat()
	stack.PushLong(int64(value))
}

///////////////////////////////////////////

type F2D struct {
	base.NoOperandsInstruction
}

func (recv *F2D) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopFloat()
	stack.PushDouble(float64(value))
}
