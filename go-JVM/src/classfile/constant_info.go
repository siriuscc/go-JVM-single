package classfile

import (
	"fmt"
)

type ConstantInfo interface {
	read(reader *ClassReader)
}

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16 // utf8
}

func (recv *ConstantClassInfo) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.nameIndex = reader.readUint16()
}

func (recv *ConstantClassInfo) GetName() string {
	return recv.cp.GetUtf8(recv.nameIndex)
}

func (recv *ConstantClassInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t#%d", "Class", recv.nameIndex)
}

/////////////////////////////////////////////////////////////////////

type ConstantFieldRefInfo struct {
	cp               ConstantPool
	classIndex       uint16 // 指向拥有者的名字
	nameAndTypeIndex uint16 //
}

func (recv *ConstantFieldRefInfo) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.classIndex = reader.readUint16()
	recv.nameAndTypeIndex = reader.readUint16()
}

func (recv *ConstantFieldRefInfo) GetClassName() string {

	return recv.cp.GetConstantClassInfo(recv.classIndex).GetName()
}

func (recv *ConstantFieldRefInfo) GetNameAndType() *ConstantNameAndTypeInfo {

	return recv.cp[recv.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
}

func (recv *ConstantFieldRefInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t#%d,#%d", "FieldRef", recv.classIndex, recv.nameAndTypeIndex)
}

///////////////////////////////////////////////////////////////////////

type ConstantInvokeDynamicInfo struct {
	cp ConstantPool

	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (recv *ConstantInvokeDynamicInfo) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.bootstrapMethodAttrIndex = reader.readUint16()
	recv.nameAndTypeIndex = reader.readUint16()
}

func (recv *ConstantInvokeDynamicInfo) GetNameAndType() *ConstantNameAndTypeInfo {
	return recv.cp.GetNameAndType(recv.nameAndTypeIndex)
}

func (recv *ConstantInvokeDynamicInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t#%d,#%d", "MethodType", recv.bootstrapMethodAttrIndex, recv.nameAndTypeIndex)
}

///////////////////////////////////////////////////////////////

type ConstantNameAndTypeInfo struct {
	cp              ConstantPool
	nameIndex       uint16
	descriptorIndex uint16
}

func (recv *ConstantNameAndTypeInfo) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.nameIndex = reader.readUint16()
	recv.descriptorIndex = reader.readUint16()
}

func (recv *ConstantNameAndTypeInfo) GetName() string {
	return recv.cp.GetUtf8(recv.nameIndex)
}

func (recv *ConstantNameAndTypeInfo) GetDescriptor() string {
	return recv.cp.GetUtf8(recv.descriptorIndex)
}

func (recv *ConstantNameAndTypeInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t#%d,#%d", "NameAndType", recv.nameIndex, recv.descriptorIndex)
}
