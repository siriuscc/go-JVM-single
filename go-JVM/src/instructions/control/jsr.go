package control

import (
	"instructions/base"
	"rtda"
)

// 后面跟着一个 uint16
type JSR struct {
	returnAddress int16
}

func (recv *JSR) FetchOperands(reader *base.ByteCodeReader) {
	recv.returnAddress = reader.ReadInt16()
}

func (recv *JSR) Execute(frame *rtda.Frame) {

	//TODO
	//frame.OpStack().PushRef((*rtda.Object)recv.returnAddress)

}
