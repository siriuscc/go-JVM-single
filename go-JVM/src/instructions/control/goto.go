package control

import (
	"instructions/base"
	"rtda"
)

// 无条件跳转
type GOTO struct {
	base.BranchInstruction
}

func (recv *GOTO) Execute(frame *rtda.Frame) {
	recv.GotoPC(frame)
}
