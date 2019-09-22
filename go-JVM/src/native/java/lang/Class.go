package lang

import (
	"instructions/base"
	"native"
	"rtda"
	"rtda/heap"
	"strings"
)

func init() {
	const class = "java/lang/Class"
	native.Register(class, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(class, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(class, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	native.Register(class, "forName0", "(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", forName0)
	native.Register(class, "getModifiers", "()I", getModifiers)
	native.Register(class, "isPrimitive", "()Z", isPrimitive)
	native.Register(class, "getSuperclass", "()Ljava/lang/Class;", getSuperclass)
	native.Register(class, "isArray", "()B", isArray)
	native.Register(class, "isAssignableFrom", "(Ljava/lang/Class;)Z", isAssignableFrom)
	native.Register(class, "isInterface", "()Z", isInterface)

	native.Register(class, "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;", getDeclaredConstructors0)
	native.Register(class, "getDeclaredMethods0", "(Z)[Ljava/lang/reflect/Method;", getDeclaredMethods0)
	native.Register(class, "getDeclaredFields0", "(Z)[Ljava/lang/reflect/Field;", getDeclaredFields0)

}

//     public native boolean isInterface();
func isInterface(frame *rtda.Frame) {

	locals := frame.LocalVars()
	this := locals.GetThis()

	isIface := this.GetKlass().IsInterface()

	frame.OpStack().PushBool(isIface)
}

// ()B
// public native boolean isArray();
func isArray(frame *rtda.Frame) {
	locals := frame.LocalVars()
	opStack := frame.OpStack()
	this := locals.GetThis()

	if this.GetKlass().IsArray() {
		opStack.PushBool(true)
		return
	}
	opStack.PushBool(false)
}

// 得到访问权限描述符
// return accessFlags
// public native int getModifiers();
// getModifiers   ()I
func getModifiers(frame *rtda.Frame) {

	locals := frame.LocalVars()
	this := locals.GetRef(0)
	flags := this.GetKlass().GetAccessFlags()
	frame.OpStack().PushInt(int32(flags))
}

// 获取父类
//public native Class<? super T> getSuperclass();
// ()Ljava/lang/Class;
func getSuperclass(frame *rtda.Frame) {

	locals := frame.LocalVars()
	opStack := frame.OpStack()
	this := locals.GetThis()

	super := this.GetMetaData().GetSuper()

	if super == nil {
		opStack.PushRef(nil)
		return
	}
	opStack.PushRef(super.GetJavaMirror())
}

//public native boolean isPrimitive();
// ()B
func isPrimitive(frame *rtda.Frame) {
	locals := frame.LocalVars()
	opStack := frame.OpStack()
	this := locals.GetThis()

	if this.GetKlass().IsPrimitive() {
		opStack.PushBool(true)
	} else {
		opStack.PushBool(false)
	}
}

//根据给定的 权限定名 name ，用loader加载，返回一个 类oop
//private static native Class<?> forName0(String name, boolean initialize,
//    ClassLoader loader,
//    Class<?> caller)
//    throws ClassNotFoundException;
func forName0(frame *rtda.Frame) {

	opStack := frame.OpStack()
	locals := frame.LocalVars()

	name := locals.GetRef(0)
	initialize := locals.GetBoolean(1)
	//loader := locals.GetRef(3)
	//caller := locals.GetRef(4)

	className := heap.GoString(name)
	className = strings.Replace(className, ".", "/", -1)
	klass := frame.GetMethod().GetOwner().GetClassLoader().LoadClass(className)

	oop := klass.GetJavaMirror()
	if initialize && !klass.IsInitialized() {

		// pc 回退，插入 一个 类初始化的 frame
		thread := frame.GetThread()
		frame.SetNextPC(thread.GetPC())
		// init class
		base.InitClass(thread, klass)
	} else {
		opStack.PushRef(oop)
	}

}

//    private native String getName0();
func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.GetMetaData()

	name := class.GetJavaName()
	nameObj := heap.JString(class.GetClassLoader(), name)

	frame.OpStack().PushRef(nameObj)
}

// 获取基本类型,静态方法
// static native Class<?> getPrimitiveClass(String name);
//		 Return the Virtual Machine's Class object for the named primitive type.
func getPrimitiveClass(frame *rtda.Frame) {

	nameStringObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameStringObj) // 字符串

	loader := frame.GetMethod().GetOwner().GetClassLoader()
	mirror := loader.LoadClass(name).GetJavaMirror() // 对应的oop

	frame.OpStack().PushRef(mirror)
}

//private static native boolean desiredAssertionStatus0(Class<?> clazz);
func desiredAssertionStatus0(frame *rtda.Frame) {
	// 断言 全部默认false，不做逻辑处理
	frame.OpStack().PushInt(0)
}

//public native boolean isAssignableFrom(Class<?> cls);
//(Ljava/lang/Class;)Z isAssignableFrom
func isAssignableFrom(frame *rtda.Frame) {

	locals := frame.LocalVars()

	this := locals.GetThis()
	clsOop := locals.GetRef(1)

	assignableTo := this.GetKlass().IsAssignableTo(clsOop.GetKlass())
	frame.OpStack().PushBool(assignableTo)
}
