package comparisons

import (
	"instructions/base"
	"rtda"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (recv *IF_ACMPEQ) Execute(frame *rtda.Frame) {

	v2, v1 := frame.OpStack().PopRef2()

	if v1 == v2 {
		recv.GotoPC(frame)
	}
}

///////////////////////////////////////////////////////////////////////

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (recv *IF_ACMPNE) Execute(frame *rtda.Frame) {

	v2, v1 := frame.OpStack().PopRef2()

	if v1 != v2 {
		recv.GotoPC(frame)
	}
}
