package classfile

import (
	"errors"
	"fmt"
)

type ClassFile struct {
	magic uint32 // 魔数，标志这个文件是class文件，0xCAFEBABE.

	minorVersion      uint16       // 次版本号
	majorVersion      uint16       // 主版本号
	constantPoolCount uint16       // 常量池大小，[1,constantPoolCount]
	constantPool      ConstantPool // 常量池

	accessFlags AccessFlags // 访问标识
	thisClass   uint16      // 本类信息
	superClass  uint16      // 父类信息

	interfacesCount uint16   //
	interfaces      []uint16 // 每一个都是一个CONSTANT_Class_info

	fieldsCount uint16
	fields      []FieldInfo // 属性信息

	methodsCount uint16
	methods      []MethodInfo // 方法信息

	attributesCount uint16
	attributes      []AttributeInfo // attr，大杂烩，存放各种信息，比如代码
}

// 负责整体解析
func ParseClassFile(bytes []byte) (*ClassFile, error) {

	if len(bytes) < 1 {
		return nil, errors.New("bytes no len")
	}
	reader := &ClassReader{data: bytes}
	classFile := &ClassFile{}

	classFile.readMagic(reader)
	classFile.readVersion(reader)
	classFile.constantPool, classFile.constantPoolCount = parseConstantPool(reader)

	//logger.Msg(classFile.constantPool.ToString())

	classFile.readAccessFlags(reader)
	classFile.readThisSuperClass(reader)
	classFile.readInterfaces(reader)

	classFile.fields, classFile.fieldsCount = parseFields(reader)
	classFile.methods, classFile.methodsCount = parseMethods(reader)
	classFile.attributes, classFile.attributesCount = parseAttributeInfos(reader)

	// 解析完reader 应该为空
	//log.Println("reader status:", reader.ToString())

	return classFile, nil
}

func (recv *ClassFile) readMagic(reader *ClassReader) {

	recv.magic = reader.readUint32()

	// 检查
	if recv.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatEror:magic")
	}
}

func (recv *ClassFile) readVersion(reader *ClassReader) {

	recv.minorVersion = reader.readUint16()
	recv.majorVersion = reader.readUint16()

	// 检查, 45.0-45.3, 45.65535
	if recv.majorVersion == 45 {
		return
	} else if recv.majorVersion <= 52 && recv.minorVersion == 0 {
		return
	}

	panic("java.lang.UnsupportedClassVersionError!")
}

func (recv *ClassFile) readAccessFlags(reader *ClassReader) {

	recv.accessFlags = AccessFlags{accessType: ACCESS_CLASS}
	recv.accessFlags.read(reader)
}

func (recv *ClassFile) readThisSuperClass(reader *ClassReader) {
	recv.thisClass = reader.readUint16()
	recv.superClass = reader.readUint16()
}

func (recv *ClassFile) readInterfaces(reader *ClassReader) {
	recv.interfaces = reader.readUint16s()
	recv.interfacesCount = uint16(len(recv.interfaces))
}

func (recv *ClassFile) GetMainMethod() *MethodInfo {

	for _, method := range recv.methods {

		name := method.GetMethodName()

		if "main" == name {
			return &method
		}
	}

	return nil
}

func (recv *ClassFile) ToString() string {

	msg := "ClassFile:\n"
	msg += fmt.Sprintf("  minor version: %d \n", recv.minorVersion)
	msg += fmt.Sprintf("  major version: %d \n", recv.majorVersion)
	msg += recv.accessFlags.ToString() + "\n"
	msg += recv.constantPool.ToString() + "\n"
	msg += fmt.Sprintf("  interfacesCount:%d: ", recv.interfacesCount)
	for _, i := range recv.interfaces {
		msg += fmt.Sprintf("#%d ", i)
	}
	msg += "\n"

	msg += fmt.Sprintf("  fieldsCount:%d: ", recv.fieldsCount)
	for _, i := range recv.fields {
		msg += i.ToString()
	}
	msg += "\n"

	msg += fmt.Sprintf("  methodsCount:%d: \n", recv.methodsCount)
	for _, i := range recv.methods {
		msg += "  " + i.ToString()
	}
	msg += "\n"

	return msg
}

func (recv *ClassFile) GetAccessFlags() *AccessFlags {
	return &recv.accessFlags
}

func (recv *ClassFile) GetClassName() string {

	classInfo := recv.constantPool.GetConstantClassInfo(recv.thisClass)
	return classInfo.GetName()
}

func (recv *ClassFile) GetConstantPool() ConstantPool {
	return recv.constantPool
}

func (recv *ClassFile) GetMethodInfos() []MethodInfo {
	return recv.methods
}

func (recv *ClassFile) GetFields() []FieldInfo {
	return recv.fields
}

func (recv *ClassFile) GetSuperClassName() string {

	info := recv.constantPool.GetConstantClassInfo(recv.superClass)

	if info != nil {
		return info.GetName()

	}
	return "" // Object类或者 接口
}

func (recv *ClassFile) GetInterfaces() []string {

	interfaceNames := make([]string, len(recv.interfaces))
	for i, index := range recv.interfaces {

		interfaceNames[i] = recv.constantPool.GetUtf8(recv.constantPool[index].(*ConstantClassInfo).nameIndex)
	}
	return interfaceNames
}

func (recv *ClassFile) GetSourceFileAttribute() *SourceFileAttribute {

	for i := range recv.attributes {
		if item, ok := recv.attributes[i].(*SourceFileAttribute); ok {
			return item
		}
	}

	return nil
}
