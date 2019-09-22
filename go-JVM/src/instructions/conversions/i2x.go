package conversions

import (
	"instructions/base"
	"rtda"
)

type I2L struct {
	base.NoOperandsInstruction
}

func (recv *I2L) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopInt()
	stack.PushLong(int64(value))
}

/////////////////////////////////////////////////////////

type I2F struct {
	base.NoOperandsInstruction
}

func (recv *I2F) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopInt()
	stack.PushFloat(float32(value))
}

/////////////////////////////////////////////////////////

type I2D struct {
	base.NoOperandsInstruction
}

func (recv *I2D) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopInt()
	stack.PushDouble(float64(value))
}

/////////////////////////////////////////////////////////

type I2B struct {
	base.NoOperandsInstruction
}

func (recv *I2B) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopInt()
	stack.PushInt(int32(uint8(value)))
}

/////////////////////////////////////////////////////////

type I2C struct {
	base.NoOperandsInstruction
}

func (recv *I2C) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopInt()
	stack.PushInt(int32(uint16(value)))
}

/////////////////////////////////////////////////////////

type I2S struct {
	base.NoOperandsInstruction
}

func (recv *I2S) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	value := stack.PopInt()
	stack.PushInt(int32(int16(value)))
}
