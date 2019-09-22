package control

import (
	"instructions/base"
	"rtda"
)

type LOOPUPSWITCH struct {
	base.BranchInstruction

	defaultOffset int32
	pairs         int32
	matchPairs    []int32 // key1,value1,key2,value2
}

func (recv *LOOPUPSWITCH) FetchOperands(reader *base.ByteCodeReader) {

	// 有0-3个padding
	reader.SkipPadding()

	// 后面是一个 4-byte的偏移
	//three signed 32-bit values
	recv.defaultOffset = reader.ReadInt32()

	recv.pairs = reader.ReadInt32()
	recv.matchPairs = reader.ReadInt32s(uint(recv.pairs) * 2)

}

func (recv *LOOPUPSWITCH) Execute(frame *rtda.Frame) {

	index := frame.OpStack().PopInt()
	recv.Offset = recv.defaultOffset

	for i := 0; i < int(recv.pairs)*2; i += 2 {
		if recv.matchPairs[i] == index {
			recv.Offset = recv.matchPairs[i+1]
			break
		}
	}
	recv.GotoPC(frame)
}
