package classfile

import (
	"logger"
	"reflect"
)

type AttributeInfo interface {
	read(reader *ClassReader)
	logger.Debuggable
}

type TagAttribute struct {
	attributeLength uint32
}

func (recv *TagAttribute) read(read *ClassReader) {
	if recv.attributeLength != 0 {
		panic("java.lang.Format:Error SyntheticAttribute")
	}
}

func (recv *TagAttribute) ToString() string {

	return reflect.TypeOf(recv).String()
}

type SyntheticAttribute struct {
	TagAttribute
}

type DeprecatedAttribute struct {
	TagAttribute
}

func newAttributeInfo(reader *ClassReader) AttributeInfo {

	attributeNameIndex := reader.readUint16()
	attributeLength := reader.readUint32() // attributeLength
	attrName := reader.cp.GetUtf8(attributeNameIndex)

	//logger.Println("    newAttributeInfo.attrName:", attrName)

	switch attrName {
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Code":
		return &CodeAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{TagAttribute{attributeLength}}
	case "Deprecated":
		return &DeprecatedAttribute{TagAttribute{attributeLength}}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{}
	case "Signature":
		return &SignatureAttribute{}

		//case "BootstrapMethods":
		//case "InnerClasses":
		//case "EnclosingMethod":
		//case "RuntimeVisibleAnnotations":
		//case "RuntimeInvisibleAnnotations":
		//case "RuntimeVisibleParameterAnnotations":
		//case "RuntimeInvisibleParameterAnnotations":
		//case "RuntimeVisibleTypeAnnotations":
		//case "RuntimeInvisibleTypeAnnotations":
		//case "AnnotationDefault":
		//case "MethodParameters":
	}
	return &DefaultAttribute{name: attrName, attributeLength: attributeLength}

}

func parseAttributeInfos(reader *ClassReader) ([]AttributeInfo, uint16) {

	attributesCount := reader.readUint16()
	infos := make([]AttributeInfo, attributesCount)

	//logger.Println("attributesCount:", attributesCount)

	for i := range infos {
		infos[i] = newAttributeInfo(reader)
		infos[i].read(reader)
	}

	return infos, attributesCount
}
