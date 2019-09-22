package extended

import (
	"instructions/base"
	"rtda"
)

// 无条件跳转
type GOTO_W struct {
	base.BranchInstruction
}

func (recv *GOTO_W) FetchOperands(reader *base.ByteCodeReader) {
	recv.Offset = int32(reader.ReadInt32())
}

func (recv *GOTO_W) Execute(frame *rtda.Frame) {
	recv.GotoPC(frame)
}
