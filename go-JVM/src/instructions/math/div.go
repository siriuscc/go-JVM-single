package math

import (
	"instructions/base"
	"rtda"
)

type IDIV struct {
	base.NoOperandsInstruction
}

func (recv *IDIV) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()

	val2 := opStack.PopInt()
	val1 := opStack.PopInt()

	if val2 == 0 {
		base.ThrowRuntimeException(frame, "java/lang/ArithmeticException", " / by zero")
		return
	}

	// 溢出不会抛异常
	opStack.PushInt(val1 / val2)
}

/////////////////////////////////////////////////////////

type LDIV struct {
	base.NoOperandsInstruction
}

func (recv *LDIV) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()

	val2 := opStack.PopLong()
	val1 := opStack.PopLong()
	// 溢出不会抛异常
	opStack.PushLong(val1 / val2)
}

/////////////////////////////////////////////////////////

type FDIV struct {
	base.NoOperandsInstruction
}

func (recv *FDIV) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()

	val2 := opStack.PopFloat()
	val1 := opStack.PopFloat()
	// 溢出不会抛异常
	opStack.PushFloat(val1 / val2)
}

/////////////////////////////////////////////////////////

type DDIV struct {
	base.NoOperandsInstruction
}

func (recv *DDIV) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()

	val2 := opStack.PopDouble()
	val1 := opStack.PopDouble()
	// 溢出不会抛异常
	opStack.PushDouble(val1 / val2)
}
