package math

import (
	"instructions/base"
	"rtda"
)

type IADD struct {
	base.NoOperandsInstruction
}

func (recv *IADD) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	val1 := stack.PopInt()
	val2 := stack.PopInt()
	// 溢出也不会抛异常

	stack.PushInt(val1 + val2)
}

////////////////////////////////////////////////////

type DADD struct {
	base.NoOperandsInstruction
}

func (recv *DADD) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	val1 := stack.PopDouble()
	val2 := stack.PopDouble()
	// 溢出也不会抛异常

	stack.PushDouble(val1 + val2)
}

////////////////////////////////////////////////////

type LADD struct {
	base.NoOperandsInstruction
}

func (recv *LADD) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	val1 := stack.PopLong()
	val2 := stack.PopLong()
	// 溢出也不会抛异常

	stack.PushLong(val1 + val2)
}

////////////////////////////////////////////////////

type FADD struct {
	base.NoOperandsInstruction
}

func (recv *FADD) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	val1 := stack.PopFloat()
	val2 := stack.PopFloat()
	// 溢出也不会抛异常

	stack.PushFloat(val1 + val2)
}

////////////////////////////////////////////////////
