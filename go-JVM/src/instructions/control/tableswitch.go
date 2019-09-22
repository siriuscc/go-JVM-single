package control

import (
	"instructions/base"
	"rtda"
)

type TABLESWITCH struct {
	base.BranchInstruction

	defaultOffset int32
	low           int32
	high          int32
	jumps         []int32
}

func (recv *TABLESWITCH) FetchOperands(reader *base.ByteCodeReader) {

	// 有0-3个padding
	reader.SkipPadding()

	// 后面是一个 4-byte的偏移
	//three signed 32-bit values
	recv.defaultOffset = reader.ReadInt32()

	recv.low = reader.ReadInt32()
	recv.high = reader.ReadInt32()

	recv.jumps = reader.ReadInt32s(uint(recv.high - recv.low + 1))

}

func (recv *TABLESWITCH) Execute(frame *rtda.Frame) {

	index := frame.OpStack().PopInt()

	if index < recv.low || index > recv.high {
		recv.Offset = recv.defaultOffset
	} else {
		recv.Offset = recv.jumps[index-recv.low]
	}
	recv.GotoPC(frame)
}
