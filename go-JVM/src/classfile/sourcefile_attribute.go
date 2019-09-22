package classfile

import "fmt"

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (recv *SourceFileAttribute) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.sourceFileIndex = reader.readUint16()
}

func (recv *SourceFileAttribute) ToString() string {

	return fmt.Sprintf("SourceFile #%d", recv.sourceFileIndex)
}

func (recv *SourceFileAttribute) GetSourceFileName() string {
	return recv.cp.GetUtf8(recv.sourceFileIndex)
}
