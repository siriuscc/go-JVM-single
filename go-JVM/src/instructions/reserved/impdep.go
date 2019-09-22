package reserved

import (
	"instructions/base"
	"rtda"
)

type IMPDEP1 struct {
}

func (recv *IMPDEP1) FetchOperands(reader *base.ByteCodeReader) {
	panic("java.unsupport.opcode:IMPDEP1")

}

func (recv *IMPDEP1) Execute(frame *rtda.Frame) {

	//TODO impdep1
	panic("java.unsupport.opcode:IMPDEP1")

}

type IMPDEP2 struct {
}

func (recv *IMPDEP2) FetchOperands(reader *base.ByteCodeReader) {
	panic("java.unsupport.opcode:IMPDEP2")

}

func (recv *IMPDEP2) Execute(frame *rtda.Frame) {

	//TODO impdep2
	panic("java.unsupport.opcode:IMPDEP2")

}
