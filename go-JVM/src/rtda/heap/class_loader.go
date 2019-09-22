package heap

import (
	"classfile"
	"classpath"
	"log"
)

type ClassLoader struct {
	logFlag  bool
	cPath    *classpath.Classpath
	klassMap map[string]*Klass // <类的完全限定名，Klass>
}

func CreateClassLoader(cPath *classpath.Classpath, printLog bool) *ClassLoader {

	classLoader := &ClassLoader{
		cPath:    cPath,
		klassMap: make(map[string]*Klass),
		logFlag:  printLog,
	}

	classLoader.loadBasicClasses()
	classLoader.loadPrimitiveClasses()

	return classLoader
}

// 装载入口
func (recv *ClassLoader) LoadClass(name string) *Klass {

	if klass, ok := recv.klassMap[name]; ok {
		return klass
	}

	// 构建klass 结构
	var klass *Klass
	if name[0] == '[' {
		klass = recv.loadArrayClass(name)
	} else {
		klass = recv.loadNonArrayClass(name)
	}

	// 构建对应的oop结构
	if classKlass, ok := recv.klassMap["java/lang/Class"]; ok {
		klass.javaMirror = classKlass.CreateObject()
		klass.javaMirror.metaData = klass
	}

	return klass
}

// 根据全限定名称 name 加载非数组类
func (recv *ClassLoader) loadNonArrayClass(fullName string) *Klass {

	bytes, entry := recv.readClass(fullName)
	klass := recv.defineClass(bytes)
	link(klass)

	if recv.logFlag {
		log.Printf("[Loaded (%s) from (%s)]\n", fullName, entry)
	}
	return klass
}

// 从 数据源中读取出 二进制流
func (recv *ClassLoader) readClass(fullName string) ([]byte, classpath.Entry) {

	bytes, entry, err := recv.cPath.ReadClass(fullName)
	if err == nil {
		return bytes, entry
	}
	return nil, nil
}

// 将二进制流 转换为 Klass 对象
func (recv *ClassLoader) defineClass(bytes []byte) *Klass {

	klass := parseKlass(bytes)
	hackClass(klass)
	klass.classLoader = recv

	recv.resolveSuperKlass(klass)

	recv.resolveInterfaces(klass)
	recv.klassMap[klass.name] = klass
	return klass
}

func parseKlass(bytes []byte) *Klass {

	cFile, e := classfile.ParseClassFile(bytes)
	if e != nil {
		panic("java.lang.ClassFormatError")
	}
	return CreateClass(cFile)
}

func (recv *ClassLoader) resolveSuperKlass(klass *Klass) {

	if klass.name != "java/lang/Object" {
		klass.super = klass.classLoader.LoadClass(klass.superClassName)
	}
}

func (recv *ClassLoader) resolveInterfaces(klass *Klass) {

	if len(klass.interfaceNames) < 1 {
		return
	}

	interfaces := make([]*Klass, len(klass.interfaceNames))
	for i, name := range klass.interfaceNames {
		interfaces[i] = klass.classLoader.LoadClass(name)
	}
	klass.interfaces = interfaces
}

func (recv *ClassLoader) loadArrayClass(name string) *Klass {

	arrKlass := &Klass{
		accessFlags: classfile.CreateAccessFlag(classfile.ACCESS_CLASS, classfile.ACC_PUBLIC),
		name:        name,
		classLoader: recv,
		initialized: true,
		super:       recv.LoadClass("java/lang/Object"),
		interfaces: []*Klass{
			recv.LoadClass("java/lang/Cloneable"),
			recv.LoadClass("java/io/Serializable"),
		},
	}

	recv.klassMap[name] = arrKlass

	return arrKlass
}

func (recv *ClassLoader) loadBasicClasses() {

	classClass := recv.LoadClass("java/lang/Class")
	for _, class := range recv.klassMap {
		if class.javaMirror == nil {
			// 创建 类对应的 oop结构
			class.javaMirror = classClass.CreateObject()
			class.javaMirror.metaData = class
		}
	}
	// 这里会导致 Class 也有oop结构
}

func (recv *ClassLoader) loadPrimitiveClasses() {

	for className := range primitiveTypes {
		recv.loadPrimitiveClass(className)
	}
}

func (recv *ClassLoader) loadPrimitiveClass(className string) {
	class := &Klass{
		accessFlags: classfile.CreateAccessFlag(classfile.ACCESS_CLASS, classfile.ACC_PUBLIC),
		name:        className,
		classLoader: recv,
		initialized: true,
	}

	class.javaMirror = recv.klassMap["java/lang/Class"].CreateObject()
	class.javaMirror.metaData = class
	recv.klassMap[className] = class
}

func link(klass *Klass) {

	verify(klass)
	prepare(klass)
}

func verify(klass *Klass) {
	//todo  something for valifity
}

func prepare(klass *Klass) {

	// 统计一下所有需要的槽位
	calcInstantFieldSlotIds(klass)
	calcStaticFieldSlotIds(klass)
	// 分配槽位
	allocAndInitStaticVars(klass)

}

func allocAndInitStaticVars(klass *Klass) {

	klass.staticVars = createSlots(klass.staticSlotCount)

	for _, field := range klass.fields {
		// 常量池缓存
		if field.accessFlags.IsStatic() && field.accessFlags.IsFinal() {
			initStaticFinalVar(klass, field)
		}
	}
}

// final static ,
func initStaticFinalVar(klass *Klass, field *Field) {

	vars := klass.staticVars
	cp := klass.constantPool

	cpIndex := field.constValueIndex

	// 常量池，索引从 1 开始
	if cpIndex > 0 {

		switch field.descriptor {

		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(field.slotId, val)
		case "J": //long
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(field.slotId, val)
		case "D": //Double
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(field.slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(klass.GetClassLoader(), goStr)
			vars.SetRef(field.slotId, jStr)
		}

	}

}

func calcStaticFieldSlotIds(klass *Klass) {

	slotId := uint(0)
	if klass.super != nil {
		slotId = klass.super.staticSlotCount
	}

	for _, field := range klass.fields {

		if field.accessFlags.IsStatic() {
			field.slotId = slotId
			slotId++

			if field.isLongOrDouble() {
				slotId++
			}
		}
	}

	klass.staticSlotCount = slotId
}

func calcInstantFieldSlotIds(klass *Klass) {

	slotId := uint(0)
	if klass.super != nil {
		slotId = klass.super.instanceSlotCount
	}

	for _, field := range klass.fields {

		if !field.accessFlags.IsStatic() {
			field.slotId = slotId
			slotId++

			if field.isLongOrDouble() {
				slotId++
			}
		}
	}

	klass.instanceSlotCount = slotId
}

// todo 不让它继续 正常的load，要不然会native爆炸的。
func hackClass(class *Klass) {
	if class.name == "java/lang/ClassLoader" {
		loadLibrary := class.GetStaticMethod("loadLibrary", "(Ljava/lang/Class;Ljava/lang/String;Z)V")
		loadLibrary.code = []byte{0xb1} // return void
	}
}
