package classfile

import "fmt"

type LocalVariableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16 // 变量名，指向一个 CONSTANT_Utf8_info
	descriptorIndex uint16 // 描述符，CONSTANT_Utf8_info
	index           uint16 // 此变量必须在本地变量表的索引 为index
}
type LocalVariableTableAttribute struct {
	localVariableTableLength uint16
	localVariableTable       []LocalVariableEntry
}

func (recv *LocalVariableEntry) ToString() string {
	s := "  LocalVariable:{\n"
	s += fmt.Sprintf("    startPC=%d\n", recv.startPc)
	s += fmt.Sprintf("    length=%d\n", recv.length)
	s += fmt.Sprintf("    nameIndex=%d\n", recv.nameIndex)
	s += fmt.Sprintf("    descriptorIndex=%d\n", recv.descriptorIndex)
	s += fmt.Sprintf("    index=%d\n", recv.index)

	s += "  }\n"
	return s
}

func (recv *LocalVariableEntry) read(reader *ClassReader) {
	recv.startPc = reader.readUint16()
	recv.length = reader.readUint16()
	recv.nameIndex = reader.readUint16()
	recv.descriptorIndex = reader.readUint16()
	recv.index = reader.readUint16()
}

func (recv *LocalVariableTableAttribute) ToString() string {

	s := "LocalVariableTable:{\n"
	s += fmt.Sprintf("  localVariableTableLength=%d\n", recv.localVariableTableLength)

	for _, i := range recv.localVariableTable {
		s += i.ToString()
	}
	s += "}\n"
	return s
}

func (recv *LocalVariableTableAttribute) read(reader *ClassReader) {

	recv.localVariableTableLength = reader.readUint16()
	recv.localVariableTable = make([]LocalVariableEntry, recv.localVariableTableLength)

	for _, i := range recv.localVariableTable {
		i.read(reader)
	}
}
