package extended

import (
	"instructions/base"
	"rtda"
)

type JSR_W struct {
	returnAddress int32
}

func (recv *JSR_W) FetchOperands(reader *base.ByteCodeReader) {
	recv.returnAddress = reader.ReadInt32()
}

func (recv *JSR_W) Execute(frame *rtda.Frame) {

	//TODO
	panic("java.unsupport.opcode:JSR_W")

}
