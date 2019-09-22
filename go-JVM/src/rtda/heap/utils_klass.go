package heap

import (
	"classfile"
	"math/rand"
)

func GetNotZeroRand() int32 {
	hashCode := int32(0)

	for hashCode == 0 {
		hashCode = rand.Int31()
	}
	return hashCode
}

func getSourceFileName(cf *classfile.ClassFile) string {

	if sourceFileAttr := cf.GetSourceFileAttribute(); sourceFileAttr != nil {
		return sourceFileAttr.GetSourceFileName()
	}

	return "Unknown"
}
func CreateByteArray(loader *ClassLoader, bytes []int8) *OopDesc {
	return &OopDesc{oopType: loader.LoadClass("[B"), vtable: bytes, metaData: nil}
}

var _DriverKlass = &Klass{
	name: "<JVMDriver>",
}

// 不应该有 基本类型的
var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

// java/lang/Object  => [Ljava/lang/Object;
// [I	=>[[I
func getArrayClassName(name string) string {

	if name[0] == '[' {
		return name
	}

	if d, ok := primitiveTypes[name]; ok {
		panic("Error type" + d)
		//return d
	}

	//[Ljava/lang.Object
	return "[L" + name + ";"
}

// [Ljava/lang/Object;  name:[Ljava/lang/Object;
// [I			[I
// Ljava/lang/String	java/lang/String
// I					int
// 转换为类名
func convertTypeDescToClassName(name string) string {

	if name[0] == '[' {
		return name
	}

	if name[0] == 'L' {
		return name[1 : len(name)-1]
	}

	for className, d := range primitiveTypes {
		if d == name {
			return className
		}
	}
	panic("Invalid descriptor:" + name)
}
