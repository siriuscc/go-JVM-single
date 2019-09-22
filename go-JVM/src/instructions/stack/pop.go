package stack

import (
	"instructions/base"
	"rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

func (recv *POP) Execute(frame *rtda.Frame) {
	frame.OpStack().PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (recv *POP2) Execute(frame *rtda.Frame) {
	frame.OpStack().PopSlot()
	frame.OpStack().PopSlot()
}
