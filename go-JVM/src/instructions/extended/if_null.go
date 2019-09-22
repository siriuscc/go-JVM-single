package extended

import (
	"instructions/base"
	"rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

func (recv *IFNULL) Execute(frame *rtda.Frame) {

	ref := frame.OpStack().PopRef()

	if ref == nil {
		recv.GotoPC(frame)
	}
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (recv *IFNONNULL) Execute(frame *rtda.Frame) {

	ref := frame.OpStack().PopRef()

	if ref != nil {
		recv.GotoPC(frame)
	}
}
