package heap

import (
	"classfile"
	"strings"
)

// [I   [I
// [Ljava/lang/String	[
// 把参数描述符 拆解为 一个，剩余段；
func splitOneParam(descriptor string) (string, string) {

	if len(descriptor) < 1 {
		return "", ""
	}

	switch descriptor[0] {
	case 'B', 'C', 'D', 'F', 'I', 'J', 'S', 'Z':
		return string(descriptor[0]), descriptor[1:]
	case 'L':

		end := strings.Index(descriptor, ";")
		return descriptor[0:end], descriptor[end+1:]

	case '[':
		item, last := splitOneParam(descriptor[1:])
		return "[" + item, last
	}
	return "", ""
}

var DriverMethod = &Method{
	ClassMember: ClassMember{
		accessFlags: classfile.CreateAccessFlag(classfile.ACCESS_METHOD, classfile.ACC_PUBLIC),
		name:        "<DriverReturn>",
		owner:       _DriverKlass,
	},
	code: []byte{0xb1}, //return
}

var AThrowMethod = &Method{
	ClassMember: ClassMember{
		accessFlags: classfile.CreateAccessFlag(classfile.ACCESS_METHOD, classfile.ACC_PUBLIC),
		name:        "<DriverAThrow>",
		owner:       _DriverKlass,
	},
	code: []byte{0xBF, 0xb1}, //athrow return
}
