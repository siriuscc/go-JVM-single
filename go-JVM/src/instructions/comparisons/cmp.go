package comparisons

import (
	"instructions/base"
	"rtda"
)

// 判断是否相等
type LCMP struct {
	base.NoOperandsInstruction
}

func (recv *LCMP) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	v2 := opStack.PopLong()
	v1 := opStack.PopLong()

	if v1 > v2 {
		opStack.PushInt(1)
	} else if v1 == v2 {
		opStack.PushInt(0)
	} else {
		opStack.PushInt(-1)
	}
}

/////////////////////////////////////////////////
func _fcmp(frame *rtda.Frame, gFlag bool) {

	opStack := frame.OpStack()
	v2 := opStack.PopFloat()
	v1 := opStack.PopFloat()

	if v1 > v2 {
		opStack.PushInt(1)
	} else if v1 == v2 {
		opStack.PushInt(0)
	} else if v1 < v2 {
		opStack.PushInt(-1)
	} else if gFlag {
		opStack.PushInt(1)
	} else {
		opStack.PushInt(-1)
	}
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (recv *FCMPL) Execute(frame *rtda.Frame) {

	// 两个变量至少有一个是NaN时， fcmpg 返回1，fcmpl返回-1
	_fcmp(frame, false)
}

type FCMPG struct {
	base.NoOperandsInstruction
}

func (recv *FCMPG) Execute(frame *rtda.Frame) {

	// 两个变量至少有一个是NaN时， fcmpg 返回1，fcmpl返回-1
	_fcmp(frame, true)
}

///////////////////////////////////////////////////////////////

func _dcmp(frame *rtda.Frame, gFlag bool) {

	opStack := frame.OpStack()
	v2, v1 := frame.OpStack().PopDouble2()

	if v1 > v2 {
		opStack.PushInt(1)
	} else if v1 == v2 {
		opStack.PushInt(0)
	} else if v1 < v2 {
		opStack.PushInt(-1)
	} else if gFlag {
		opStack.PushInt(1)
	} else {
		opStack.PushInt(-1)
	}
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (recv *DCMPL) Execute(frame *rtda.Frame) {

	// 两个变量至少有一个是NaN时， fcmpg 返回1，fcmpl返回-1
	_dcmp(frame, false)
}

type DCMPG struct {
	base.NoOperandsInstruction
}

func (recv *DCMPG) Execute(frame *rtda.Frame) {

	// 两个变量至少有一个是NaN时， fcmpg 返回1，fcmpl返回-1
	_dcmp(frame, true)
}
