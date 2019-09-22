package stack

import (
	"instructions/base"
	"rtda"
)

// 交换两个slot的位置
type SWAP struct {
	base.NoOperandsInstruction
}

func (recv *SWAP) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()

	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}
