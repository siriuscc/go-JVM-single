package heap

import (
	"classfile"
	"strings"
)

type Klass struct {
	sourceFile  string
	accessFlags *classfile.AccessFlags
	name string

	javaMirror   *OopDesc      // 指向java/lang/Class 对象
	constantPool *ConstantPool // 指向常量池

	superClassName string
	interfaceNames []string

	super       *Klass       // 指向 父类Klass
	classLoader *ClassLoader //

	interfaces []*Klass
	methods    []*Method
	fields     []*Field

	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initialized       bool // 是否已经初始化
}

func (recv *Klass) Init()                     { recv.initialized = true }
func (recv *Klass) IsInitialized() bool       { return recv.initialized }
func (recv *Klass) GetSourceFileName() string { return recv.sourceFile }
func (recv *Klass) IsPublic() bool            { return recv.accessFlags.IsPublic() }
func (recv *Klass) IsEnum() bool              { return recv.accessFlags.IsEnum() }
func (recv *Klass) IsAnnotation() bool        { return recv.accessFlags.IsAnnotation() }
func (recv *Klass) IsSynthetic() bool         { return recv.accessFlags.IsSynthetic() }
func (recv *Klass) IsAbstract() bool          { return recv.accessFlags.IsAbstract() }
func (recv *Klass) IsInterface() bool         { return recv.accessFlags.IsInterface() }
func (recv *Klass) IsSuper() bool             { return recv.accessFlags.IsSuper() }
func (recv *Klass) IsFinal() bool             { return recv.accessFlags.IsFinal() }
func (recv *Klass) GetJavaMirror() *OopDesc   { return recv.javaMirror }

func (recv *Klass) GetConstantPool() *ConstantPool { return recv.constantPool }
func (recv *Klass) GetStaticVars() Slots           { return recv.staticVars }
func (recv *Klass) GetSuper() *Klass               { return recv.super }
func (recv *Klass) GetClassLoader() *ClassLoader   { return recv.classLoader }
func (recv *Klass) GetClinitMethod() *Method       { return recv.GetStaticMethod("<clinit>", "()V") }

func (recv *Klass) isObjectClass() bool                  { return recv.name == "java/lang/Object" }
func (recv *Klass) isCloneableIFace() bool               { return recv.name == "java/lang/Cloneable" }
func (recv *Klass) isSerializableIFace() bool            { return recv.name == "java/io/Serializable" }
func (recv *Klass) IsSuperInterfaceOf(iface *Klass) bool { return iface.IsSubClassOf(recv) }

func (recv *Klass) GetName() string {
	if recv == nil {
		return ""
	}
	return recv.name
}

// 创建
func CreateClass(cf *classfile.ClassFile) *Klass {

	klass := &Klass{}
	klass.accessFlags = cf.GetAccessFlags()
	klass.name = cf.GetClassName()

	klass.superClassName = cf.GetSuperClassName()
	klass.interfaceNames = cf.GetInterfaces()

	klass.loadConstantPool(cf.GetConstantPool())
	klass.loadFields(cf.GetFields())
	klass.loadMethods(cf.GetMethodInfos())

	klass.sourceFile = getSourceFileName(cf)

	return klass
}

func (recv *Klass) isAccessibleTo(trigger *Klass) bool {
	return recv.IsPublic() || recv.GetPackageName() == trigger.GetPackageName()
}

func (recv *Klass) GetPackageName() string {

	if i := strings.LastIndex(recv.name, "/"); i >= 0 {
		return recv.name[:i]
	}
	return ""
}

func (recv *Klass) IsSubClassOf(super *Klass) bool {

	if recv.super == super {
		return true
	}
	// 间接父类
	return recv.super != nil && recv.super.IsSubClassOf(super)
}

func (recv *Klass) GetMainMethod() *Method {
	// ([Ljava/lang/String;)V
	return recv.GetStaticMethod("main", "([Ljava/lang/String;)V")

}
func (recv *Klass) GetInstanceMethod(name string, desc string) *Method {
	return recv.getMethod(name, desc, false)
}

func (recv *Klass) GetStaticMethod(name string, descriptor string) *Method {

	return recv.getMethod(name, descriptor, true)
}

func (recv *Klass) getMethod(name string, descriptor string, isStatic bool) *Method {

	for _, method := range recv.methods {
		if method.name == name && method.descriptor == descriptor {
			if !isStatic || method.accessFlags.IsStatic() {
				return method
			}
		}
	}
	return nil
}

// recv 是 interface 的实现类?
// S 可能实现了 接口T，或者T的子接口 T'
func (recv *Klass) IsImplementsOf(interfaceKlass *Klass) bool {

	for _, iface := range recv.interfaces {
		if iface == interfaceKlass || iface.IsSubInterfaceOf(interfaceKlass) {
			return true
		}
	}

	return recv.super != nil && recv.super.IsImplementsOf(interfaceKlass)
}

// 是子接口？
func (recv *Klass) IsSubInterfaceOf(interfaceKlass *Klass) bool {

	if !recv.IsInterface() || !interfaceKlass.IsInterface() {
		return false
	}
	// interface 多继承
	for _, i := range recv.interfaces {
		if i == interfaceKlass {
			return true
		}
	}
	return false
}

func (recv *Klass) ComponentClass() *Klass {

	if recv.name[0] == '[' {
		// [I I
		// [[I [I
		// [Ljava/lang.Object;
		componentTypeName := recv.name[1:]
		name := convertTypeDescToClassName(componentTypeName)

		return recv.classLoader.LoadClass(name)
	}

	panic("Not array:" + recv.name)
}

func (recv *Klass) IsAssignableFrom(klass *Klass) bool {

	return klass.isAccessibleTo(recv)
}

func (recv *Klass) IsAssignableTo(klass *Klass) bool {

	s, t := recv, klass
	if s == t {
		return true
	}
	if t == nil {
		return false
	}
	if !s.IsArray() {
		// 普通的class
		if !s.IsInterface() {
			if t.IsInterface() {
				return s.IsImplementsOf(t)
			} else { // t is class
				return s.IsSubClassOf(t)
			}
		} else { // 接口
			if !t.IsInterface() { //Class
				return t.isObjectClass()
			} else {
				return t.IsSuperInterfaceOf(s)
			}
		}
	} else { // 数组类
		if !t.IsArray() {
			if !t.IsInterface() { // class
				return t.isObjectClass()
			} else { // interface
				return t.isCloneableIFace() || t.isSerializableIFace()
			}
		} else {

			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.IsAssignableTo(sc)
		}
	}
}

func (recv *Klass) getField(name string, desc string, isStatic bool) *Field {

	for _, field := range recv.fields {
		if isStatic == field.accessFlags.IsStatic() && field.name == name && field.descriptor == desc {
			return field
		}
	}

	if !isStatic && recv.super != nil {
		// 不是静态属性，去父类查找
		return recv.super.getField(name, desc, false)
	}

	return nil
}

func (recv *Klass) GetStaticFieldRef(name string, desc string) *OopDesc {
	field := recv.getField(name, desc, true)
	return recv.staticVars.GetRef(field.slotId)
}

func (recv *Klass) SetStaticFieldRef(name string, desc string, ref *OopDesc) {

	field := recv.getField(name, desc, true)
	recv.staticVars.SetRef(field.slotId, ref)
}

func (recv *Klass) ArrayClass() *Klass {

	return recv.classLoader.LoadClass(getArrayClassName(recv.name))
}

func (recv *Klass) GetJavaName() string {

	return strings.Replace(recv.name, "/", ".", -1)
}

func (recv *Klass) IsPrimitive() bool {

	_, ok := primitiveTypes[recv.name]
	return ok
}

// 没有父类，返回0
// 只有一级父类，返回1
func (recv *Klass) GetExtendsLayerCount() int {

	count := 0

	for p := recv.super; p != nil; p = p.super {
		count++
	}
	return count
}

func (recv *Klass) GetInstanceField(name string, desc string) *Field {

	return recv.getField(name, desc, false)
}

func (recv *Klass) GetAccessFlags() uint16 {

	return recv.accessFlags.GetFlags()
}

func (recv *Klass) GetDescName() string {

	if recv.IsPrimitive() {
		return recv.name
	}
	if !recv.IsArray() {
		return "L" + recv.name
	}

	return recv.name
}

func (recv *Klass) GetInitMethod(desc string) *Method {

	return recv.GetInstanceMethod("<init>", desc)
}

func (recv *Klass) GetMethods(publicOnly bool) []*Method {

	methods := make([]*Method, 0)

	for _, method := range recv.methods {
		if !publicOnly || method.accessFlags.IsPublic() {
			methods = append(methods, method)
		}
	}

	return methods
}

func (recv *Klass) GetFields(publicOnly bool) []*Field {

	fields := make([]*Field, 0)
	for _, field := range recv.fields {
		if !publicOnly || field.accessFlags.IsPublic() {
			fields = append(fields, field)
		}
	}
	return fields
}

func (recv *Klass) CreateObject() *OopDesc {

	return &OopDesc{
		oopType:  recv,
		fields:   createSlots(recv.instanceSlotCount),
		hashCode: GetNotZeroRand(),
	}
}

// 获取所有的init函数
func (recv *Klass) GetConstructors(publicOnly bool) []*Method {

	methods := make([]*Method, 0)

	for _, method := range recv.methods {
		if method.name == "<init>" && (!publicOnly || method.accessFlags.IsPublic()) {
			methods = append(methods, method)
		}
	}

	return methods
}
