package references

import (
	"instructions/base"
	"rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (recv *ARRAY_LENGTH) Execute(frame *rtda.Frame) {

	opStack := frame.OpStack()
	arrRef := opStack.PopRef()

	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	opStack.PushInt(arrLen)
}
