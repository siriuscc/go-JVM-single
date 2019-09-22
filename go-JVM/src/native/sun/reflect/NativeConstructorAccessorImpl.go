package reflect

import (
	"instructions/base"
	"native"
	"rtda"
	"rtda/heap"
)

func init() {

	const class = "sun/reflect/NativeConstructorAccessorImpl"
	native.Register(class, "newInstance0", "(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;", newInstance0)
}

//private static native Object newInstance0(Constructor<?> constructor, Object[]  params) throws InstantiationException, IllegalArgumentException, InvocationTargetException;
//   例：获取 String 的一个构造函数， 参数为一个 String
//	  String.class.getConstructor(String.class);
// (Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;
func newInstance0(frame *rtda.Frame) {

	locals := frame.LocalVars()
	constructor := locals.GetRef(0) //  构造函数的描述
	params := locals.GetRef(1)      // 调用构造函数的参数

	// 根据构造器得到对应的类
	declaringKlass := constructor.GetObjectField("clazz", "Ljava/lang/Class;").GetMetaData()
	parameterTypes := constructor.GetObjectField("parameterTypes", "[Ljava/lang/Class;")

	desc := convertToDesc(parameterTypes)
	desc = "(" + desc + ")V"
	initMethod := declaringKlass.GetInitMethod(desc)

	driverFrame := frame.GetThread().CreateDriverFrame(int(initMethod.GetMaxLocals()))
	opStack := driverFrame.OpStack()

	// 手动new一个实例
	obj := declaringKlass.CreateObject()
	// 根据构造器得到初始化方法的描述符
	// 根据描述符得到对应的初始化方法
	// 写入参数
	opStack.PushRef(obj)
	if params != nil {
		for _, o := range params.GetRefTable() {
			opStack.PushRef(o)
		}
	}
	// 手动调用初始化方法
	base.InvokeAMethod(driverFrame, initMethod)

	frame.OpStack().PushRef(obj)
}

func convertToDesc(paramsType *heap.OopDesc) string {

	//private Class<?>[]          parameterTypes;
	// vtable里，每一个应该都是oop才对，可以得到对应的metaData
	// oop 数组
	params := paramsType.GetRefTable()

	desc := ""

	for _, typeOop := range params {
		desc += typeOop.GetMetaData().GetDescName()
	}

	return desc
}
