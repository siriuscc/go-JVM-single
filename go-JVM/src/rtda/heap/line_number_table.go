package heap

import (
	"classfile"
)

type LineNumberTable []*LineNumber

type LineNumber struct {
	startPC    uint16 // 开始的指令位置
	lineNumber uint16 // 源代码中对应的行数
}

func (recv *LineNumber) LoadData(lineNumber uint16, startPC uint16) {

	recv.lineNumber = lineNumber
	recv.startPC = startPC
}

// 根据PC获得对应的lineNumber
func (recv LineNumberTable) GetLineNumber(PC uint) int {

	for i := len(recv) - 1; i >= 0; i-- {

		if PC > uint(recv[i].startPC) {
			return int(recv[i].lineNumber)
		}
	}
	return -1
}

func (recv *Method) loadLineNumbers(linesAttribute *classfile.LineNumberTableAttribute) {

	if linesAttribute == nil {
		recv.lineNumbers = nil
		return
	}

	linesTable := linesAttribute.GetLineNumberTable()
	lines := make([]*LineNumber, len(linesTable))

	for i := range linesTable {
		lines[i] = &LineNumber{
			lineNumber: linesTable[i].GetLineNumber(),
			startPC:    linesTable[i].GetStartPC(),
		}
	}

	recv.lineNumbers = lines
}
