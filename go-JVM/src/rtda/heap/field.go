package heap

import "classfile"

type Field struct {
	ClassMember
	constValueIndex uint
	slotId          uint
}

func (recv *Field) isLongOrDouble() bool {

	return recv.descriptor == "J" || recv.descriptor == "D"
}

func (recv *Field) GetSlotId() uint {
	return recv.slotId
}

// 在 类构造完成后
func (recv *Klass) loadFields(fieldInfos []classfile.FieldInfo) []*Field {

	fields := make([]*Field, len(fieldInfos))
	for i, info := range fieldInfos {

		fields[i] = &Field{}
		fields[i].owner = recv

		fields[i].accessFlags = info.GetAccessFlags()
		fields[i].name = info.GetName()
		fields[i].descriptor = info.GetDescriptor()
		// 如果属性是常量，会有一个Constant_value_attribute，其中index 指向 常量池
		if info.GetConstantValueAttribute() != nil {
			fields[i].constValueIndex = uint(info.GetConstantValueAttribute().GetConstantValueIndex())
		}
	}
	recv.fields = fields
	return fields
}

func (recv *Field) GetType() *OopDesc {

	typeKlass := recv.owner.GetClassLoader().LoadClass(convertTypeDescToClassName(recv.descriptor))
	// 把desc转为类名，装载类
	return typeKlass.GetJavaMirror()
}
