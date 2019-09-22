package stack

import (
	"instructions/base"
	"rtda"
)

// 复制栈顶的单个变量
type DUP struct {
	base.NoOperandsInstruction
}

func (recv *DUP) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	slot := stack.PopSlot()
	frame.OpStack().PushSlot(slot)
	frame.OpStack().PushSlot(slot)
}

//
type DUP_X1 struct {
	base.NoOperandsInstruction
}

func (recv *DUP_X1) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	slot1 := opStack.PopSlot()
	slot2 := opStack.PopSlot()

	opStack.PushSlot(slot1)
	opStack.PushSlot(slot2)
	opStack.PushSlot(slot1)
}

type DUP_X2 struct {
	base.NoOperandsInstruction
}

func (recv *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	frame.OpStack().PushSlot(slot1)
	frame.OpStack().PushSlot(slot3)
	frame.OpStack().PushSlot(slot2)
	frame.OpStack().PushSlot(slot1)
}

type DUP2 struct {
	base.NoOperandsInstruction
}

func (recv *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OpStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

func (recv *DUP2_X1) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (recv *DUP2_X2) Execute(frame *rtda.Frame) {

	stack := frame.OpStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()

	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
