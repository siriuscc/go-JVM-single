package base

import (
	"rtda"
	"rtda/heap"
)

func ThrowRuntimeException(frame *rtda.Frame, className string, msg string) {

	classLoader := frame.GetMethod().GetOwner().GetClassLoader()
	exKlass := classLoader.LoadClass(className)
	object := exKlass.CreateObject()

	initMethod := exKlass.GetInitMethod("(Ljava/lang/String;)V")
	msgObj := heap.JString(classLoader, msg)

	throwFrame := frame.GetThread().CreateThrowFrame(3)
	opStack := throwFrame.OpStack()
	opStack.PushRef(object) // this
	opStack.PushRef(object) // this
	opStack.PushRef(msgObj) // msg

	InvokeAMethod(throwFrame, initMethod)
}
