package math

import (
	"instructions/base"
	"math"
	"rtda"
)

// 求余
type IREM struct {
	base.NoOperandsInstruction
}

func (recv *IREM) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	value2 := stack.PopInt()
	value1 := stack.PopInt()

	if value2 == 0 {
		panic("java.ArithmeticException: / by zero")
	}

	stack.PushInt(value1 % value2)
}

////////////////////////////////////////////////////////////////
type LREM struct {
	base.NoOperandsInstruction
}

func (recv *LREM) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	value2 := stack.PopLong()
	value1 := stack.PopLong()

	if value2 == 0 {
		panic("java.ArithmeticException: / by zero")
	}

	stack.PushLong(value1 % value2)
}

////////////////////////////////////////////////////////////////
// 求余
type DREM struct {
	base.NoOperandsInstruction
}

func (recv *DREM) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	value2 := stack.PopDouble()
	value1 := stack.PopDouble()
	stack.PushDouble(math.Mod(value1, value2))
}

////////////////////////////////////////////////////////////////
type FREM struct {
	base.NoOperandsInstruction
}

func (recv *FREM) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	value2 := stack.PopFloat()
	value1 := stack.PopFloat()
	stack.PushFloat(float32(math.Mod(float64(value1), float64(value2))))
}

////////////////////////////////////////////////////////////////
