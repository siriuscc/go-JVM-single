package misc

import (
	"instructions/base"
	"native"
	"rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize",
		"()V", initialize)
}

//private static native void initialize();
// 调用 initializeSystemClass()
func initialize(frame *rtda.Frame) {

	classLoader := frame.GetMethod().GetOwner().GetClassLoader()
	SysClass := classLoader.LoadClass("java/lang/System")
	initSysClassMethod := SysClass.GetStaticMethod("initializeSystemClass", "()V")

	base.InvokeAMethod(frame, initSysClassMethod)
}
