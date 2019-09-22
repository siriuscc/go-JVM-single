package classfile

import "reflect"

type SignatureAttribute struct {
	cp             ConstantPool
	signatureIndex uint16
}

func (recv *SignatureAttribute) read(reader *ClassReader) {

	recv.cp = reader.cp
	recv.signatureIndex = reader.readUint16()
}

func (recv *SignatureAttribute) ToString() string {
	return reflect.TypeOf(recv).String()
}

func (recv *SignatureAttribute) GetSignature() string {

	if recv == nil {
		return ""
	}
	return recv.cp.GetUtf8(recv.signatureIndex)
}
