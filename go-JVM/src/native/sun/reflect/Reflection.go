package reflect

import (
	"native"
	"rtda"
)

func init() {

	const class = "sun/reflect/Reflection"
	native.Register(class, "getCallerClass", "()Ljava/lang/Class;", reflection_getCallerClass)
	native.Register(class, "getClassAccessFlags", "(Ljava/lang/Class;)I", reflection_getClassAccessFlags)

}

//public static native Class<?> getCallerClass();
func reflection_getCallerClass(frame *rtda.Frame) {
	// top0 is sun/reflect/Reflection
	// top1 is the caller of getCallerClass()
	// top2 is the caller of method

	callerFrame := frame.GetThread().GetFrames()[2]
	callerMethod := callerFrame.GetMethod()
	mirror := callerMethod.GetOwner().GetJavaMirror()
	frame.OpStack().PushRef(mirror)
}

//public static native int getClassAccessFlags(Class<?> classOop);
func reflection_getClassAccessFlags(frame *rtda.Frame) {

	locals := frame.LocalVars()

	classOop := locals.GetRef(0)
	metaData := classOop.GetMetaData()
	flags := metaData.GetAccessFlags()
	frame.OpStack().PushInt(int32(flags))
}
