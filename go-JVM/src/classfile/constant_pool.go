package classfile

import (
	"fmt"
	"logger"
)

type ConstantPool []ConstantInfo

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

func newCpInfo(reader *ClassReader) ConstantInfo {

	tag := reader.readUint8()
	switch tag {

	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	}
	return nil
}

func parseConstantPool(reader *ClassReader) ([]ConstantInfo, uint16) {

	constantPoolCount := reader.readUint16()
	constantPool := make([]ConstantInfo, constantPoolCount)
	//logger.Println("constantPoolCount:",constantPoolCount)
	reader.cp = constantPool
	// [1,constantPoolCount)
	for i := 1; i < int(constantPoolCount); i++ {

		constantPool[i] = newCpInfo(reader)
		//of := reflect.TypeOf(constantPool[i])
		//fmt.Printf("%d--------->%s \n",i,of)

		constantPool[i].read(reader)
		switch constantPool[i].(type) {
		case *ConstantDoubleInfo:
			i++
		case *ConstantLongInfo:
			i++
		}
	}

	return constantPool, constantPoolCount
}

// index 指向 UTF8_string
func (recv ConstantPool) GetUtf8(index uint16) string {

	if index < 1 || int(index) > len(recv) {
		return ""
	}

	utf8Info, ok := recv[index].(*ConstantUtf8Info)

	if !ok {
		panic(fmt.Sprintf("error:%s", recv[index]))
	}

	//log.Printf("    @GetUtf8.cp[%d]= %s :\n", index, cp[index])

	return utf8Info.val
}

func (recv ConstantPool) GetNameAndType(index uint16) *ConstantNameAndTypeInfo {

	return recv[index].(*ConstantNameAndTypeInfo)
}

func (recv ConstantPool) ToString() string {

	var s = "Constant pool:\n"

	for i := 1; i < len(recv); i++ {
		s += fmt.Sprintf("  %d\t:%s\n", i, recv[i].(logger.Debuggable).ToString())
		//s = s + cp[i].(logger.Debuggable).ToString() + "\n"
	}
	return s
}

func (recv ConstantPool) GetConstantClassInfo(index uint16) *ConstantClassInfo {

	if index < 1 {
		return nil
	}

	return recv[index].(*ConstantClassInfo)
}
