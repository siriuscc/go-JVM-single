package classfile

import "fmt"

type FieldInfo struct {
	cp              ConstantPool
	accessFlags     *AccessFlags
	nameIndex       uint16
	descriptorIndex uint16
	attributesCount uint16
	attributes      []AttributeInfo //attributes_count
}

func (recv *FieldInfo) read(reader *ClassReader) {

	recv.cp = reader.cp
	recv.accessFlags = &AccessFlags{accessType: ACCESS_FIELD}
	recv.accessFlags.read(reader)

	recv.nameIndex = reader.readUint16()
	recv.descriptorIndex = reader.readUint16()
	//recv.attributesCount = reader.readUint16()

	recv.attributes, recv.attributesCount = parseAttributeInfos(reader)
}

func (recv *FieldInfo) GetAccessFlags() *AccessFlags {

	return recv.accessFlags
}

func (recv *FieldInfo) GetName() string {
	return recv.cp.GetUtf8(recv.nameIndex)
}

func (recv *FieldInfo) GetDescriptor() string {
	return recv.cp.GetUtf8(recv.descriptorIndex)
}

func parseFields(reader *ClassReader) ([]FieldInfo, uint16) {

	fieldsCount := reader.readUint16()

	infos := make([]FieldInfo, fieldsCount)
	for i := range infos {
		infos[i].read(reader)
	}

	return infos, fieldsCount
}

func (recv *FieldInfo) GetConstantValueAttribute() *ConstantValueAttribute {

	for _, attrInfo := range recv.attributes {

		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (recv *FieldInfo) ToString() string {

	msg := "FieldInfo: "
	msg += recv.accessFlags.ToString() + "\n"

	msg += fmt.Sprintf("  nameIndex:#%d", recv.nameIndex)
	msg += fmt.Sprintf("  descriptorIndex:#%d", recv.descriptorIndex)
	msg += fmt.Sprintf("  attributesCount:#%d", recv.attributesCount)
	msg += "  attributes:"

	for _, i := range recv.attributes {
		msg += i.ToString()
	}

	return msg
}
