package classfile

import "fmt"

type ConstantValueAttribute struct {
	constantValueIndex uint16 // long,int...string
}

func (recv *ConstantValueAttribute) GetConstantValueIndex() uint16 {
	return recv.constantValueIndex
}

func (recv *ConstantValueAttribute) read(reader *ClassReader) {
	recv.constantValueIndex = reader.readUint16()
}
func (recv *ConstantValueAttribute) ToString() string {

	return fmt.Sprintf("  ConstantValueAttribute:\n \t\t #%d", recv.constantValueIndex)
}
