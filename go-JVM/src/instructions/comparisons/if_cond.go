package comparisons

import (
	"instructions/base"
	"rtda"
)

// if 指令是 栈顶弹出一个int，和0比较，满足就跳转
//value = 0
type IFEQ struct {
	base.BranchInstruction
}

func (recv *IFEQ) Execute(frame *rtda.Frame) {

	value := frame.OpStack().PopInt()

	if value == 0 {
		recv.GotoPC(frame)
	}
}

//////////////////////////////////////////////////////

//value ！= 0
type IFNE struct {
	base.BranchInstruction
}

func (recv *IFNE) Execute(frame *rtda.Frame) {

	value := frame.OpStack().PopInt()

	if value != 0 {
		recv.GotoPC(frame)
	}
}

//////////////////////////////////////////////////////

//value < 0
type IFLT struct {
	base.BranchInstruction
}

func (recv *IFLT) Execute(frame *rtda.Frame) {

	value := frame.OpStack().PopInt()

	if value < 0 {
		recv.GotoPC(frame)
	}
}

//////////////////////////////////////////////////////
//value <= 0
type IFLE struct {
	base.BranchInstruction
}

func (recv *IFLE) Execute(frame *rtda.Frame) {

	value := frame.OpStack().PopInt()

	if value <= 0 {
		recv.GotoPC(frame)
	}
}

//////////////////////////////////////////////////////

//value > 0
type IFGT struct {
	base.BranchInstruction
}

func (recv *IFGT) Execute(frame *rtda.Frame) {

	value := frame.OpStack().PopInt()

	if value > 0 {
		recv.GotoPC(frame)
	}
}

//////////////////////////////////////////////////////
//value >= 0
type IFGE struct {
	base.BranchInstruction
}

func (recv *IFGE) Execute(frame *rtda.Frame) {

	value := frame.OpStack().PopInt()

	if value >= 0 {
		recv.GotoPC(frame)
	}
}
