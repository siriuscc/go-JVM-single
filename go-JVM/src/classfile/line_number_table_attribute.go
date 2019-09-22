package classfile

import "fmt"

type LineNumberEntry struct {
	startPC    uint16 // 开始的指令位置
	lineNumber uint16 // 源代码中对应的行数
}

type LineNumberTableAttribute struct {
	lineNumberTableLength uint16
	lineNumberTable       []*LineNumberEntry
}

func (recv *LineNumberEntry) read(reader *ClassReader) {
	recv.startPC = reader.readUint16()
	recv.lineNumber = reader.readUint16()
}
func (recv *LineNumberEntry) ToString() string {

	s := fmt.Sprintf("  LineNumberEntry:startPC=%d,lineNumber=%d", recv.startPC, recv.lineNumber)
	return s
}

func (recv *LineNumberEntry) GetLineNumber() uint16 {
	return recv.lineNumber
}

func (recv *LineNumberEntry) GetStartPC() uint16 {
	return recv.startPC
}

func (recv *LineNumberTableAttribute) GetLineNumberTable() []*LineNumberEntry {
	return recv.lineNumberTable
}

func (recv *LineNumberTableAttribute) read(reader *ClassReader) {
	recv.lineNumberTableLength = reader.readUint16()
	recv.lineNumberTable = make([]*LineNumberEntry, recv.lineNumberTableLength)
	for i := range recv.lineNumberTable {

		recv.lineNumberTable[i] = &LineNumberEntry{}
		recv.lineNumberTable[i].read(reader)
	}
}

func (recv *LineNumberTableAttribute) ToString() string {

	s := fmt.Sprintf("  LineNumberTableAttribute: lineNumberTableLength=%d\n", recv.lineNumberTableLength)

	for _, i := range recv.lineNumberTable {
		s += i.ToString() + "\n"
	}

	return s
}
