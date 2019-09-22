package classfile

import "fmt"

type CodeAttribute struct {
	//attributeNameIndex   uint16
	//attributeLength      uint32
	maxStack   uint16
	maxLocals  uint16
	codeLength uint32
	code       []uint8

	exceptionTableLength uint16
	exceptionTable       []*ExceptionInfo

	attributesCount uint16
	attributes      []AttributeInfo
}

func (recv *CodeAttribute) GetExceptionTable() []*ExceptionInfo {
	return recv.exceptionTable
}

func (recv *CodeAttribute) GetMaxStack() uint16 {
	return recv.maxStack
}
func (recv *CodeAttribute) GetMaxLocals() uint16 {
	return recv.maxLocals
}
func (recv *CodeAttribute) GetCode() []uint8 {
	return recv.code
}

type ExceptionInfo struct {
	//cp ConstantPool
	startPC   uint16
	endPC     uint16
	handlerPC uint16
	catchType uint16
}

func (recv *ExceptionInfo) GetStartPC() uint16 {
	return recv.startPC
}

func (recv *ExceptionInfo) GetEndPC() uint16 {
	return recv.endPC
}

func (recv *ExceptionInfo) GetHandlerPC() uint16 {
	return recv.handlerPC
}

func (recv *ExceptionInfo) GetCatchType() uint16 {
	return recv.catchType
}

func (recv *ExceptionInfo) read(reader *ClassReader) {
	//recv.cp=reader.cp
	recv.startPC = reader.readUint16()
	recv.endPC = reader.readUint16()
	recv.handlerPC = reader.readUint16()
	recv.catchType = reader.readUint16()
}

func (recv *CodeAttribute) read(reader *ClassReader) {

	recv.maxStack = reader.readUint16()
	recv.maxLocals = reader.readUint16()
	recv.codeLength = reader.readUint32()

	recv.code = make([]uint8, recv.codeLength)
	for i := range recv.code {
		recv.code[i] = reader.readUint8()
	}
	recv.exceptionTableLength = reader.readUint16()
	excpTable := make([]*ExceptionInfo, recv.exceptionTableLength)
	for i := range excpTable {
		excpTable[i] = &ExceptionInfo{}
		excpTable[i].read(reader)
	}
	recv.exceptionTable = excpTable
	recv.attributes, recv.attributesCount = parseAttributeInfos(reader)
}

func (recv *ExceptionInfo) ToString() string {

	s := fmt.Sprintf("Exception:{")
	s += fmt.Sprintf("  startPC=%d\n", recv.startPC)
	s += fmt.Sprintf("  endPC=%d\n", recv.endPC)
	s += fmt.Sprintf("  handlerPC=%d\n", recv.handlerPC)
	s += fmt.Sprintf("  catchType=%d\n", recv.catchType)
	s += "}"
	return s
}

func (recv *CodeAttribute) ToString() string {
	s := "Code:{\n"
	s += fmt.Sprintf("    maxStack=\t%d\n", recv.maxStack)
	s += fmt.Sprintf("    maxLocals=\t%d\n", recv.maxLocals)
	s += fmt.Sprintf("    codeLength=\t%d\n", recv.codeLength)

	for _, i := range recv.code {
		s += fmt.Sprintf("%x ", i)
	}
	s += "\n"
	s += fmt.Sprintf("    exceptionTableLength=\t%d\n", recv.exceptionTableLength)

	for _, i := range recv.exceptionTable {
		s += "  " + i.ToString()
	}

	s += fmt.Sprintf("    attributesCount=%d\n", recv.attributesCount)

	for _, i := range recv.attributes {
		s += "  " + i.ToString()
	}
	s += "}"

	return s
}
