package base

import (
	"rtda"
)

type Instruction interface {
	// 提取操作数
	FetchOperands(reader *ByteCodeReader)
	// 执行指令
	Execute(frame *rtda.Frame)
}

type WideInstruction interface {
	Instruction
	MakeWide()
}

// 空操作
type NoOperandsInstruction struct{}

func (recv *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {

}

func (recv *NoOperandsInstruction) Execute(frame *rtda.Frame) {

}

//////////////////////////////////////////////////////
type Index8Instruction struct {
	Index uint16
	Wide  bool
	NoOperandsInstruction
}

func (recv *Index8Instruction) FetchOperands(reader *ByteCodeReader) {

	if recv.Wide {
		recv.Index = reader.ReadUint16()
	} else {
		recv.Index = uint16(reader.ReadUint8())
	}
}

func (recv *Index8Instruction) MakeWide() {
	recv.Wide = true
}

//////////////////////////////////////////////////////

type Index16Instruction struct {
	Index uint16
	NoOperandsInstruction
}

func (recv *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	recv.Index = reader.ReadUint16()
}

/////////////////////////////////////////////////////

type BranchInstruction struct {
	Offset int32 // code length 不超过2^16,所以跳转范围 [-2^16,2^16]
}

func (recv *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	recv.Offset = int32(reader.ReadInt16())
}

func (recv *BranchInstruction) GotoPC(frame *rtda.Frame) {

	//方法入口pc
	pc := frame.GetThread().GetPC()
	nextPC := int32(pc) + recv.Offset
	frame.SetNextPC(uint16(nextPC))

}
