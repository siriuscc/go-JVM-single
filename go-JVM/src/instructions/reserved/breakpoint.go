package reserved

import (
	"instructions/base"
	"rtda"
)

type BREAKPOINT struct {
	returnAddress int32
}

func (recv *BREAKPOINT) FetchOperands(reader *base.ByteCodeReader) {
	panic("java.unsupport.opcode:BREAKPOINT")

}

func (recv *BREAKPOINT) Execute(frame *rtda.Frame) {

	//TODO break point
	panic("java.unsupport.opcode:BREAKPOINT")

}
