package math

import (
	"instructions/base"
	"rtda"
)

type INEG struct {
	base.NoOperandsInstruction
}

func (recv *INEG) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	val := opStack.PopInt()
	opStack.PushInt(-val)
}

/////////////////////////////////////////////////////////

type LNEG struct {
	base.NoOperandsInstruction
}

func (recv *LNEG) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	val := opStack.PopLong()
	opStack.PushLong(-val)
}

/////////////////////////////////////////////////////////
type FNEG struct {
	base.NoOperandsInstruction
}

func (recv *FNEG) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	val := opStack.PopFloat()
	opStack.PushFloat(-val)
}

/////////////////////////////////////////////////////////
type DNEG struct {
	base.NoOperandsInstruction
}

func (recv *DNEG) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	val := opStack.PopDouble()
	opStack.PushDouble(-val)
}
