package constants

import (
	"instructions/base"
	"rtda"
)

// BIPUSH 是 byte 扩展为 int 后入 opStack
// SIPUSH 是 short 扩展为 int 后入 opStack

type BIPUSH struct {
	val uint8 // push byte
}

func (recv *BIPUSH) FetchOperands(reader *base.ByteCodeReader) {

	// 获取一个byte，扩展为int，然后入栈
	recv.val = reader.ReadUint8()
}

func (recv *BIPUSH) Execute(frame *rtda.Frame) {

	frame.OpStack().PushInt(int32(recv.val))
}

type SIPUSH struct {
	val uint16 //push short
}

func (recv *SIPUSH) FetchOperands(reader *base.ByteCodeReader) {
	recv.val = reader.ReadUint16()
}

func (recv *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OpStack().PushInt(int32(recv.val))
}
