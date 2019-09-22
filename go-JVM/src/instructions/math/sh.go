package math

import (
	"instructions/base"
	"rtda"
)

//移位运算

//算术左移
type ISHL struct {
	base.NoOperandsInstruction
}

func (recv *ISHL) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	v2 := stack.PopInt()
	v1 := stack.PopInt()
	//where s is the value of the low 5 bits of value2.
	s := uint32(v2) & 0x1f
	stack.PushInt(v1 << s)
}

/////////////////////////////////////////////////////////

type ISHR struct {
	base.NoOperandsInstruction
}

func (recv *ISHR) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	stack.PushInt(v1 >> s)
}

/////////////////////////////////////////////////////////

//算术左移
type LSHL struct {
	base.NoOperandsInstruction
}

func (recv *LSHL) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	v2 := stack.PopInt()
	v1 := stack.PopLong()
	//where s is the value of the low 5 bits of value2.
	s := uint64(v2) & 0x3f
	stack.PushLong(v1 << s)
}

/////////////////////////////////////////////////////////

type LSHR struct {
	base.NoOperandsInstruction
}

func (recv *LSHR) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f
	stack.PushLong(v1 >> s)
}

/////////////////////////////////////////////////////////

//算术左移
type IUSHR struct {
	base.NoOperandsInstruction
}

func (recv *IUSHR) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	stack.PushInt(int32(uint32(v1) >> s))
}

/////////////////////////////////////////////////////////

type LUSHR struct {
	base.NoOperandsInstruction
}

func (recv *LUSHR) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint64(v2) & 0x3f
	stack.PushLong(int64(uint64(v1) >> s))
}
