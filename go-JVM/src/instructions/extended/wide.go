package extended

import (
	"instructions/base"
	"instructions/loads"
	"instructions/math"
	"instructions/stores"
	"rtda"
)

// 改变其他指令的行为
type WIDE struct {
	wideInst base.WideInstruction
}

func (recv *WIDE) FetchOperands(reader *base.ByteCodeReader) {

	opcode := reader.ReadUint8()

	//TODO 这里创建了一把,会创建所有的指令实现，消耗很大，后面优化一下
	//interpreter:=&instructions.Interpreter{}

	switch opcode {

	case 0x15: //iload
		recv.wideInst = &loads.ILOAD{}
	case 0x16: //lload
		recv.wideInst = &loads.LLOAD{}
	case 0x17: //fload
		recv.wideInst = &loads.FLOAD{}
	case 0x18: //dload
		recv.wideInst = &loads.DLOAD{}
	case 0x19: //aload
		recv.wideInst = &loads.ALOAD{}
	case 0x36: //istore
		recv.wideInst = &stores.ISTORE{}
	case 0x37: //lstore
		recv.wideInst = &stores.LSTORE{}
	case 0x38: //fstore
		recv.wideInst = &stores.FSTORE{}
	case 0x39: //dstore
		recv.wideInst = &stores.DSTORE{}
	case 0x3a: //astore
		recv.wideInst = &stores.ASTORE{}
	case 0x84: //iinc
		recv.wideInst = &math.IINC{}

	case 0xa9: //ret
		panic("Unsupported opcode 0xa9!")
	}

	recv.wideInst.MakeWide()
	recv.wideInst.FetchOperands(reader)
}

func (recv *WIDE) Execute(frame *rtda.Frame) {

	recv.wideInst.Execute(frame)
}
