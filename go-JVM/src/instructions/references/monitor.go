package references

import (
	"instructions/base"
	"rtda"
)

type MONITOR_ENTER struct {
	base.NoOperandsInstruction
}

func (recv *MONITOR_ENTER) Execute(frame *rtda.Frame) {

	ref := frame.OpStack().PopRef()

	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

type MONITOR_EXIT struct {
	base.NoOperandsInstruction
}

func (recv *MONITOR_EXIT) Execute(frame *rtda.Frame) {
	ref := frame.OpStack().PopRef()

	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
