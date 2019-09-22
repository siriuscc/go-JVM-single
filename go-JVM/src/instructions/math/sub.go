package math

import (
	"instructions/base"
	"rtda"
)

//减法

type ISUB struct {
	base.NoOperandsInstruction
}

func (recv *ISUB) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value2 := opStack.PopInt()
	value1 := opStack.PopInt()

	opStack.PushInt(value1 - value2)
}

////////////////////////////////////////////////
type FSUB struct {
	base.NoOperandsInstruction
}

func (recv *FSUB) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value2 := opStack.PopFloat()
	value1 := opStack.PopFloat()

	opStack.PushFloat(value1 - value2)
}

////////////////////////////////////////////////

type DSUB struct {
	base.NoOperandsInstruction
}

func (recv *DSUB) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value2 := opStack.PopDouble()
	value1 := opStack.PopDouble()

	opStack.PushDouble(value1 - value2)
}

////////////////////////////////////////////////

type LSUB struct {
	base.NoOperandsInstruction
}

func (recv *LSUB) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	value2 := opStack.PopLong()
	value1 := opStack.PopLong()

	opStack.PushLong(value1 - value2)
}
