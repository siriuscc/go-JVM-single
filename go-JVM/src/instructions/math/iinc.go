package math

import (
	"instructions/base"
	"rtda"
)

// 局部变量表中index位置的变量加Const
type IINC struct {
	Index uint16 // 一般为uint8 ,uint16是为了兼容wide
	Const int16
	Wide  bool
}

func (recv *IINC) FetchOperands(reader *base.ByteCodeReader) {

	if recv.Wide {
		//logger.Println("innc wide")
		recv.Index = reader.ReadUint16()
		recv.Const = reader.ReadInt16()
	} else {
		//logger.Printf("innc no wide:%d,%d",recv.Index,recv.Const)
		recv.Index = uint16(reader.ReadUint8())
		recv.Const = int16(reader.ReadInt8())
	}
}

func (recv *IINC) Execute(frame *rtda.Frame) {

	value := frame.LocalVars().GetInt(recv.Index)
	result := value + int32(recv.Const)
	frame.LocalVars().SetInt(recv.Index, result)
}

func (recv *IINC) MakeWide() {
	recv.Wide = true
}
