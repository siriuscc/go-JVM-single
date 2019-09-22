package classfile

import "fmt"

type ExceptionsAttribute struct {
	numberOfExceptions  uint16
	exceptionIndexTable []uint16
}

func (recv *ExceptionsAttribute) read(reader *ClassReader) {
	recv.exceptionIndexTable = reader.readUint16s()
	recv.numberOfExceptions = uint16(len(recv.exceptionIndexTable))
}

func (recv *ExceptionsAttribute) ToString() string {

	s := fmt.Sprintf("  ExceptionsAttribute:\n \t\t %d", recv.numberOfExceptions)
	for _, i := range recv.exceptionIndexTable {
		s += fmt.Sprintf("\t\t#%d \n", i)
	}
	return s
}
