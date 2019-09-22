package classfile

import "fmt"

/////////////////////////////////////////////////////////////////////

type ConstantMethodRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (recv *ConstantMethodRefInfo) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.classIndex = reader.readUint16()
	recv.nameAndTypeIndex = reader.readUint16()
}

func (recv *ConstantMethodRefInfo) GetClassName() string {

	return recv.cp.GetConstantClassInfo(recv.classIndex).GetName()
}

func (recv *ConstantMethodRefInfo) GetNameAndType() *ConstantNameAndTypeInfo {

	return recv.cp[recv.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
}

func (recv *ConstantMethodRefInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t#%d,#%d", "MethodRef", recv.classIndex, recv.nameAndTypeIndex)
}

type ConstantMethodHandleInfo struct {
	cp             ConstantPool
	referenceKind  uint8
	referenceIndex uint16
}

func (recv *ConstantMethodHandleInfo) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.referenceKind = reader.readUint8()
	recv.referenceIndex = reader.readUint16()
}
func (recv *ConstantMethodHandleInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t%d,#%d", "MethodHandle", recv.referenceKind, recv.referenceIndex)
}

type ConstantMethodTypeInfo struct {
	cp              ConstantPool
	descriptorIndex uint16
}

func (recv *ConstantMethodTypeInfo) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.descriptorIndex = reader.readUint16()
}
func (recv *ConstantMethodTypeInfo) ToString() string {
	return fmt.Sprintf("\t%s:\t\t#%d", "MethodType", recv.descriptorIndex)
}

/////////////////////////////////////////////////////////////////////

type ConstantInterfaceMethodRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (recv *ConstantInterfaceMethodRefInfo) GetClassName() string {

	return recv.cp.GetConstantClassInfo(recv.classIndex).GetName()
}

func (recv *ConstantInterfaceMethodRefInfo) GetNameAndType() *ConstantNameAndTypeInfo {
	return recv.cp[recv.nameAndTypeIndex].(*ConstantNameAndTypeInfo)
}

func (recv *ConstantInterfaceMethodRefInfo) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.classIndex = reader.readUint16()
	recv.nameAndTypeIndex = reader.readUint16()
}

func (recv *ConstantInterfaceMethodRefInfo) ToString() string {
	return "    InterfaceMethodRef:        #" + string(recv.classIndex) + "#" + string(recv.nameAndTypeIndex)
	return fmt.Sprintf("\t%s:\t\t#%d,#%d", "InterfaceMethodRef", recv.classIndex, recv.nameAndTypeIndex)
}

/////////////////////////////////////////////////////////////////////
