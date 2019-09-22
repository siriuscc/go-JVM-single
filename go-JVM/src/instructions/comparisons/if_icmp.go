package comparisons

import (
	"instructions/base"
	"rtda"
)

//v1==v2
type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (recv *IF_ICMPEQ) Execute(frame *rtda.Frame) {

	v2, v1 := frame.OpStack().PopInt2()

	if v1 == v2 {
		recv.GotoPC(frame)
	}
}

//////////////////////////////////////////////////////
//v1!=v2
type IF_ICMPNE struct {
	base.BranchInstruction
}

func (recv *IF_ICMPNE) Execute(frame *rtda.Frame) {

	v2, v1 := frame.OpStack().PopInt2()

	if v1 != v2 {
		recv.GotoPC(frame)
	}
}

//////////////////////////////////////////////////////

//v1!=v2
type IF_ICMPLT struct {
	base.BranchInstruction
}

func (recv *IF_ICMPLT) Execute(frame *rtda.Frame) {

	v2, v1 := frame.OpStack().PopInt2()

	if v1 < v2 {
		recv.GotoPC(frame)
	}
}

//////////////////////////////////////////////////////

//>=
type IF_ICMPGE struct {
	base.BranchInstruction
}

func (recv *IF_ICMPGE) Execute(frame *rtda.Frame) {

	v2, v1 := frame.OpStack().PopInt2()

	if v1 >= v2 {
		recv.GotoPC(frame)
	}
}

//////////////////////////////////////////////////////
type IF_ICMPGT struct {
	base.BranchInstruction
}

func (recv *IF_ICMPGT) Execute(frame *rtda.Frame) {

	v2, v1 := frame.OpStack().PopInt2()

	if v1 > v2 {
		recv.GotoPC(frame)
	}
}

//////////////////////////////////////////////////////

type IF_ICMPLE struct {
	base.BranchInstruction
}

func (recv *IF_ICMPLE) Execute(frame *rtda.Frame) {

	v2, v1 := frame.OpStack().PopInt2()

	if v1 <= v2 {
		recv.GotoPC(frame)
	}
}
