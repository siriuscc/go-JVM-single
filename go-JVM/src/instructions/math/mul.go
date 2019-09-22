package math

import (
	"instructions/base"
	"rtda"
)

type IMUL struct {
	base.NoOperandsInstruction
}

func (recv *IMUL) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()

	val1 := opStack.PopInt()
	val2 := opStack.PopInt()
	// 溢出不会抛异常
	opStack.PushInt(val1 * val2)
}

/////////////////////////////////////////////////////////
type LMUL struct {
	base.NoOperandsInstruction
}

func (recv *LMUL) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()

	val1 := opStack.PopLong()
	val2 := opStack.PopLong()
	// 溢出不会抛异常
	opStack.PushLong(val1 * val2)
}

/////////////////////////////////////////////////////////
type FMUL struct {
	base.NoOperandsInstruction
}

func (recv *FMUL) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()

	val1 := opStack.PopFloat()
	val2 := opStack.PopFloat()
	// 溢出不会抛异常
	opStack.PushFloat(val1 * val2)
}

/////////////////////////////////////////////////////////

type DMUL struct {
	base.NoOperandsInstruction
}

func (recv *DMUL) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()

	val1 := opStack.PopDouble()
	val2 := opStack.PopDouble()
	// 溢出不会抛异常
	opStack.PushDouble(val1 * val2)
}

/////////////////////////////////////////////////////////
