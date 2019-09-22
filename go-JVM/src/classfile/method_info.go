package classfile

import (
	"fmt"
)

type MethodInfo struct {
	cp ConstantPool

	/**
	 *  public,private..static,final,native
	 */
	accessFlags *AccessFlags

	/**
	 * 指向常量池的一个CONSTANT_Utf8_info
	 *  表示一个 <init>或者<cinit>或者 普通方法
	 */
	nameIndex uint16
	/**
	 * CONSTANT_Utf8_info ，描述符
	 *     例如：Object m(int i, double d, Thread t) {...}
	 *        方法描述符 (IDLjava/lang/Thread;)Ljava/lang/Object;
	 */
	descriptorIndex uint16 //
	attributesCount uint16
	attributes      []AttributeInfo
}

func (recv *MethodInfo) read(reader *ClassReader) {
	recv.cp = reader.cp
	recv.accessFlags = &AccessFlags{accessType: ACCESS_METHOD}
	recv.accessFlags.read(reader)
	recv.nameIndex = reader.readUint16()
	recv.descriptorIndex = reader.readUint16()

	//logger.Println("Method Name:" ,reader.cp.GetUtf8(recv.nameIndex))

	recv.attributes, recv.attributesCount = parseAttributeInfos(reader)
}

func parseMethods(reader *ClassReader) ([]MethodInfo, uint16) {

	methodsCount := reader.readUint16()
	methodInfos := make([]MethodInfo, methodsCount)

	for i := range methodInfos {
		methodInfos[i].read(reader)
	}
	return methodInfos, methodsCount
}

func (recv *MethodInfo) ToString() string {

	msg := "Method:{\n"

	msg += recv.accessFlags.ToString() + "\n"

	msg += fmt.Sprintf("    nameIndex:\t#%d\n", recv.nameIndex)
	msg += fmt.Sprintf("    descriptorIndex:\t#%d\n", recv.descriptorIndex)
	msg += fmt.Sprintf("    attributesCount:\t#%d\n", recv.attributesCount)

	for _, i := range recv.attributes {
		msg += "    " + i.ToString()
	}

	msg += "}\n"

	return msg
}

func (recv *MethodInfo) GetCodeAttribute() *CodeAttribute {
	for _, i := range recv.attributes {
		attr, ok := i.(*CodeAttribute)
		if ok {
			return attr
		}
	}
	return nil
}

func (recv *MethodInfo) GetMethodName() string {

	return recv.cp.GetUtf8(recv.nameIndex)
}

func (recv *MethodInfo) GetAccessFlags() *AccessFlags {

	return recv.accessFlags
}

func (recv *MethodInfo) GetDescriptor() string {
	return recv.cp.GetUtf8(recv.descriptorIndex)
}

func (recv *MethodInfo) GetLineNumberTableAttr() *LineNumberTableAttribute {

	codeAttribute := recv.GetCodeAttribute()

	if codeAttribute == nil {
		return nil
	}

	for _, attr := range codeAttribute.attributes {
		if attr, ok := attr.(*LineNumberTableAttribute); ok {
			return attr
		}
	}
	return nil
}

func (recv *MethodInfo) GetExceptionTableAttr() []*ExceptionInfo {

	codeAttribute := recv.GetCodeAttribute()

	if codeAttribute == nil {
		return nil
	}

	return codeAttribute.GetExceptionTable()
}

func (recv *MethodInfo) GetSignaturesAttr() *SignatureAttribute {

	for _, i := range recv.attributes {
		attr, ok := i.(*SignatureAttribute)
		if ok {
			return attr
		}
	}
	return nil
}

// VisibleAnnotationsAttribute 不需要解析，直接返回
func (recv *MethodInfo) RuntimeVisibleAnnotationsAttributeData() []byte {
	return recv.getUnparsedAttributeData("RuntimeVisibleAnnotations")
}
func (recv *MethodInfo) RuntimeVisibleParameterAnnotationsAttributeData() []byte {
	return recv.getUnparsedAttributeData("RuntimeVisibleParameterAnnotationsAttribute")
}
func (recv *MethodInfo) AnnotationDefaultAttributeData() []byte {
	return recv.getUnparsedAttributeData("AnnotationDefault")
}

func (recv *MethodInfo) getUnparsedAttributeData(name string) []byte {
	for _, attrInfo := range recv.attributes {
		switch attrInfo.(type) {
		case *DefaultAttribute:
			unparsedAttr := attrInfo.(*DefaultAttribute)
			if unparsedAttr.name == name {
				return unparsedAttr.info
			}
		}
	}
	return nil
}
