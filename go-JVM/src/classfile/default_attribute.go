package classfile

import "fmt"

type DefaultAttribute struct {
	name            string
	attributeLength uint32
	info            []uint8 // 一些附带信息，比如 ConstantValue，Code
}

func (recv *DefaultAttribute) read(reader *ClassReader) {

	recv.info = make([]uint8, recv.attributeLength)

	for i := range recv.info {
		recv.info[i] = reader.readUint8()
	}
}

func (recv *DefaultAttribute) ToString() string {
	return fmt.Sprintf("DefaultAttribute: len:%d", recv.attributeLength)
}
