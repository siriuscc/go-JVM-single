package conversions

import (
	"instructions/base"
	"rtda"
)

type L2I struct {
	base.NoOperandsInstruction
}

func (recv *L2I) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopLong()
	stack.PushInt(int32(value))
}

////////////////////////////////////////////////////

type L2F struct {
	base.NoOperandsInstruction
}

func (recv *L2F) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopLong()
	stack.PushFloat(float32(value))
}

////////////////////////////////////////////////////

type L2D struct {
	base.NoOperandsInstruction
}

func (recv *L2D) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopLong()
	stack.PushDouble(float64(value))
}
